package underlinglib

import (
	"github.com/go-stomp/stomp"
)

type SinkClient struct {
	Config UnderlingConfig
}

func (sc *SinkClient) Send(queueName string, messageBody string) {
	conn, err := stomp.Dial("tcp", sc.Config.OpenNMS.Mq, stompClientOptions(sc.Config)...)
	if err != nil {
		println("failed to connect to server", sc.Config.OpenNMS.Mq, err.Error())
		return
	}
	println("successfully sink producer to server!")
	println("sending message to", queueName)
	conn.Send(queueName, "text/plain", []byte(messageBody))
	println("disconnecting...")
	conn.Disconnect()
}
