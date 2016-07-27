package underlinglib

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"strings"
)

type RPCModule interface {
	GetId() (id string)

	HandleRequest(requestBody string) (responseBody string)
}

type RPCManager interface {
	Start(stop chan bool)

	RegisterModule(module RPCModule)
}

type StompResponse struct {
	QueueName     string
	Body          string
	CorrelationID string
}

// these are the options work with ActiveMQ
var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
	stomp.ConnOpt.Login("guest", "guest"),
	stomp.ConnOpt.Host("/"),
}

type StompClient struct {
	Config  UnderlingConfig
	modules []RPCModule
}

func (sc *StompClient) RegisterModule(module RPCModule) {
	sc.modules = append(sc.modules, module)
}

func (sc *StompClient) Start() chan bool {
	stop := make(chan bool)

	incomingMessages, outgoingMessages := make(chan *stomp.Message), make(chan *StompResponse)

	for _, module := range sc.modules {
		go recvMessages(module.GetId(), sc.Config, incomingMessages, stop)
	}

	go handleMessages(sc, incomingMessages, outgoingMessages)
	go sendMessages(sc.Config, outgoingMessages)

	return stop
}

func recvMessages(moduleId string, conf UnderlingConfig, incomingMessages chan *stomp.Message, stop chan bool) {
	defer func() {
		stop <- true
	}()

	conn, err := stomp.Dial("tcp", conf.OpenNMS.Mq, options...)
	if err != nil {
		println("cannot connect to server", conf.OpenNMS.Mq, err.Error())
		return
	}
	println("successfully conected consumer to server!")

	queueName := "/queue/OpenNMS.RPC." + moduleId + "@" + conf.Minion.Location
	sub, err := conn.Subscribe(queueName, stomp.AckAuto)
	if err != nil {
		println("cannot subscribe to", queueName, err.Error())
		return
	}

	for {
		msg := <-sub.C
		println("got message")
		incomingMessages <- msg
	}
}

func handleMessages(sc *StompClient, incomingMessages chan *stomp.Message, outgoingMessages chan *StompResponse) {
	for {
		msg := <-incomingMessages
		msg.Header.Len()
		println("Received message with body:", string(msg.Body))
		if msg.Header == nil {
			println("Message has no headers.")
		} else {
			for i := 0; i < msg.Header.Len(); i++ {
				key, value := msg.Header.GetAt(i)
				fmt.Printf("Header %d: %s=%s\n", i, key, value)
			}

			requestBody := string(msg.Body)
			sourceQueueName := msg.Header.Get("JmsQueueName")

			for _, module := range sc.modules {
				if strings.HasPrefix(sourceQueueName, "OpenNMS.RPC."+module.GetId()+"@") {
					fmt.Printf("Handling request with %s module\n", module.GetId())
					responseBody := module.HandleRequest(requestBody)
					fmt.Printf("Generated response body %s\n", responseBody)
					res := StompResponse{
						QueueName:     msg.Header.Get("reply-to"),
						Body:          responseBody,
						CorrelationID: msg.Header.Get("correlation-id"),
					}
					outgoingMessages <- &res
				}
			}
		}
	}
}

func sendMessages(conf UnderlingConfig, outgoingMessages chan *StompResponse) {
	conn, err := stomp.Dial("tcp", conf.OpenNMS.Mq, options...)
	if err != nil {
		println("cannot connect to server", conf.OpenNMS.Mq, err.Error())
		return
	}
	println("successfully conected producer to server!")

	for {
		msg := <-outgoingMessages
		fmt.Printf("sending message to server on queue '%s' with correlationd-id %s: %s\n", msg.QueueName, msg.CorrelationID, msg.Body)
		err = conn.Send(msg.QueueName, "text/plain", []byte(msg.Body),
			stomp.SendOpt.Header("correlation-id", msg.CorrelationID))
		if err != nil {
			println("failed to send to server", err)
			return
		}
		println("succesfully sent message to server")
	}
}

type SNMPRpcModule struct {
}

func (snmp SNMPRpcModule) GetId() (id string) {
	return "SNMP"
}

func (snmp SNMPRpcModule) HandleRequest(requestBody string) (responseBody string) {
	request := SNMPRequestDTO{}
	UnmarshalFromXml(strings.NewReader(requestBody), &request)
	response := Exec(request)
	responseBody, _ = MarshalToXml(response)
	return responseBody
}

type DetectorRpcModule struct {
}

func (detect DetectorRpcModule) GetId() (id string) {
	return "Detect"
}

func (detect DetectorRpcModule) HandleRequest(requestBody string) (responseBody string) {
	request := DetectorRequestDTO{}
	UnmarshalFromXml(strings.NewReader(requestBody), &request)
	response := Detect(request)
	responseBody, _ = MarshalToXml(response)
	return responseBody
}
