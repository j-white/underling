package main

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"github.com/j-white/underling/underlinglib"
	"io/ioutil"
	"strings"
)

const defaultPort = ":61613"

var stop = make(chan bool)

type StompResponse struct {
	QueueName     string
	Body          string
	CorrelationID string
}

// these are the default options that work with RabbitMQ
var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
	stomp.ConnOpt.Login("guest", "guest"),
	stomp.ConnOpt.Host("/"),
}

func main() {
	yamlConfig, err := ioutil.ReadFile("underling.yaml")
	if err != nil {
		println("failed to read underling.yaml", err.Error())
		return
	}
	conf, err := underlinglib.GetConfig(yamlConfig)
	if err != nil {
		println("failed to parse underling.yaml", err.Error())
		return
	}

	incomingMessages, outgoingMessages := make(chan *stomp.Message), make(chan *StompResponse)
	go recvMessages(conf, incomingMessages)
	go handleMessages(incomingMessages, outgoingMessages)
	go sendMessages(conf, outgoingMessages)
	<-stop
}

func handleMessages(incomingMessages chan *stomp.Message, outgoingMessages chan *StompResponse) {
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

			request := underlinglib.SNMPRequestDTO{}
			underlinglib.UnmarshalFromXml(strings.NewReader(string(msg.Body)), &request)
			fmt.Println("request", request)
			response := underlinglib.Exec(request)
			fmt.Println("response", response)

			xmlResponse, _ := underlinglib.MarshalToXml(response)

			res := StompResponse{
				QueueName:     msg.Header.Get("reply-to"),
				Body:          xmlResponse,
				CorrelationID: msg.Header.Get("correlation-id"),
			}

			outgoingMessages <- &res
		}
	}
}

func recvMessages(conf underlinglib.UnderlingConfig, incomingMessages chan *stomp.Message) {
	defer func() {
		stop <- true
	}()

	conn, err := stomp.Dial("tcp", conf.OpenNMS.Mq, options...)
	if err != nil {
		println("cannot connect to server", conf.OpenNMS.Mq, err.Error())
		return
	}
	println("successfully conected consumer to server!")

	queueName := "/queue/OpenNMS.RPC.SNMP@" + conf.Minion.Location
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

func sendMessages(conf underlinglib.UnderlingConfig, outgoingMessages chan *StompResponse) {
	defer func() {
		stop <- true
	}()

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
