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
	r, err := runner.New(conf)
	if err != nil {
		glog.Fatal(err)
	}
	defer r.Free()

	// start runner
	err = r.StartOrDie()
	if err != nil {
		glog.Fatal(err)
	}
}
