package main

import (
	"github.com/golang/glog"
  "github.com/tradingAI/runner/client"
)


func main() {
	runClient()
}

func runClient() {
	// load config
	conf, err := client.LoadConf()
	if err != nil {
		glog.Fatal(err)
	}

	// new client
	c, err := client.New(conf, frontend)
	if err != nil {
		glog.Fatal(err)
	}
	defer c.Free()

	// start client
	err = c.StartOrDie()
	if err != nil {
		glog.Fatal(err)
	}
}
