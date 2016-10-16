package underlinglib

import (
	"fmt"
	"time"
)

type Underling struct {
	stop chan bool
}

func sendHearbeat(conf UnderlingConfig) {
	skc := SinkClient{Config: conf}
	fmt.Println("Sending heartbeat.")

	identity := MinionIdentityDTO{
		Id:       conf.Minion.Id,
		Location: conf.Minion.Location,
	}
	identityXml, _ := MarshalToXml(identity)
	skc.Send(conf.OpenNMS.Id+".Sink.Heartbeat", identityXml)
}

func (underling *Underling) Start(conf UnderlingConfig) {
	sc := StompClient{Config: conf}
	sc.RegisterModule(SNMPRpcModule{})
	sc.RegisterModule(DetectorRpcModule{})
	stop := sc.Start()

	heartbeat := time.NewTicker(time.Second * 30).C
	sendHearbeat(conf)
	
	go func() {
		for {
			select {
			case <-heartbeat:
				sendHearbeat(conf)
			case <-stop:
				fmt.Println("Done")
				return
			}
		}
    }()
}

func (underling *Underling) Stop() {
	underling.stop <- true
}

func (underling *Underling) Wait() {
	for {
		select {
		case <-underling.stop:
			fmt.Println("Done")
			return
		}
	}
}
