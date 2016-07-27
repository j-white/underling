package main

import (
	"fmt"
	"github.com/j-white/underling/underlinglib"
	"io/ioutil"
)

var stop = make(chan bool)

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
	<-stop
}
