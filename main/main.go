package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/tradingAI/runner/runner"
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")

	run()
}

func run() {
	// load config
	conf, err := runner.LoadConf()
	if err != nil {
		glog.Fatal(err)
	}

	// new runner
	c, err := runner.New(conf)
	if err != nil {
		glog.Fatal(err)
	}
	defer c.Free()

	// start runner
	err = c.StartOrDie()
	if err != nil {
		glog.Fatal(err)
	}
}
