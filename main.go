package main

import (
	"fmt"
	"github.com/j-white/underling/underlinglib"
	"io/ioutil"
	"time"
)

var stop = make(chan bool)

func sendHearbeat(conf underlinglib.UnderlingConfig) {
	skc := underlinglib.SinkClient{Config: conf}
	fmt.Println("Sending heartbeat.")

	identity := underlinglib.MinionIdentityDTO{
		Id:       conf.Minion.Id,
		Location: conf.Minion.Location,
	}
	identityXml, _ := underlinglib.MarshalToXml(identity)
	skc.Send("OpenNMS.Sink.Heartbeat", identityXml)
}

func main() {
	yamlConfig, err := ioutil.ReadFile("underling.yaml")
	if err != nil {
		fmt.Println("failed to read underling.yaml", err.Error())
		return
	}
	conf, err := underlinglib.GetConfig(yamlConfig)
	if err != nil {
		fmt.Println("failed to parse underling.yaml", err.Error())
		return
	}

	sc := underlinglib.StompClient{Config: conf}
	sc.RegisterModule(underlinglib.SNMPRpcModule{})
	sc.RegisterModule(underlinglib.DetectorRpcModule{})
	stop := sc.Start()

	heartbeat := time.NewTicker(time.Second * 30).C
	sendHearbeat(conf)
	for {
		select {
		case <-heartbeat:
			sendHearbeat(conf)
		case <-stop:
			fmt.Println("Done")
			return
		}
	}
}
