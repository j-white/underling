package underlinglib

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"strings"
	"time"
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

	// Start a receiver routine for each registered RPC module
	for _, module := range sc.modules {
		go recvMessages(module.GetId(), sc.Config, incomingMessages, stop)
	}

	// Start N number of handlers
	// TODO: These requests should be handled asynchronously instead
	for i := 0; i < 10; i++ {
		go handleMessages(sc, incomingMessages, outgoingMessages)
	}

	// Start a single sender
	go sendMessages(sc.Config, outgoingMessages)

	return stop
}

func recvMessages(moduleId string, conf UnderlingConfig, incomingMessages chan *stomp.Message, stop chan bool) {
	defer func() {
		stop <- true
	}()

	for {
		conn, err := stomp.Dial("tcp", conf.OpenNMS.Mq, options...)
		if err != nil {
			println("failed to connect to server", conf.OpenNMS.Mq, err.Error())
			println("sleeping for 5 seconds before trying again")
			time.Sleep(5 * time.Second)
			continue
		}
		println("successfully conected consumer to server!")

		queueName := "/queue/" + conf.OpenNMS.Id + "." + conf.Minion.Location + ".RPC." + moduleId
		sub, err := conn.Subscribe(queueName, stomp.AckAuto)
		if err != nil {
			println("cannot subscribe to", queueName, err.Error())
			return
		}

		for {
			msg := <-sub.C
			if msg == nil {
				println("got nil message. attempting to reconnect.")
				break
			}
			println("received message", msg)
			incomingMessages <- msg
		}
	}
}

func handleMessages(sc *StompClient, incomingMessages chan *stomp.Message, outgoingMessages chan *StompResponse) {
	for {
		msg := <-incomingMessages
		println("handling message", msg)
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

			matchedModules := 0
			for _, module := range sc.modules {
				if strings.HasSuffix(sourceQueueName, "RPC."+module.GetId()) {
					fmt.Printf("Handling request with %s module\n", module.GetId())
					responseBody := module.HandleRequest(requestBody)
					fmt.Printf("Generated response body %s\n", responseBody)
					res := StompResponse{
						QueueName:     msg.Header.Get("reply-to"),
						Body:          responseBody,
						CorrelationID: msg.Header.Get("correlation-id"),
					}
					outgoingMessages <- &res
					matchedModules += 1
				}
			}

			if matchedModules < 1 {
				println("No modules were matched for queue", sourceQueueName)
			}
		}
		println("done handling message", msg)
	}
}

func sendMessages(conf UnderlingConfig, outgoingMessages chan *StompResponse) {

	for {
		conn, err := stomp.Dial("tcp", conf.OpenNMS.Mq, options...)
		if err != nil {
			println("failed to connect to server", conf.OpenNMS.Mq, err.Error())
			println("sleeping for 5 seconds before trying again")
			time.Sleep(5 * time.Second)
			continue
		}
		println("successfully conected producer to server!")

		for {
			msg := <-outgoingMessages
			fmt.Printf("sending message to server on queue '%s' with correlationd-id %s: %s\n", msg.QueueName, msg.CorrelationID, msg.Body)
			err = conn.Send(msg.QueueName, "text/plain", []byte(msg.Body),
				stomp.SendOpt.Header("correlation-id", msg.CorrelationID))
			if err != nil {
				// TODO: The message will be dropped, should we re-attempt?
				println("failed to send to server", err)
				continue
			}
			println("succesfully sent message to server")
		}
	}
}
