package main

import (
	"fmt"
	"github.com/j-white/underling/underlinglib"
	"io/ioutil"
)

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

	underling := new(underlinglib.Underling)
	underling.Start(conf)
	underling.Wait()
}
