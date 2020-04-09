package client

import (
	"io/ioutil"
	"log"
	"path"
	"strconv"

	"github.com/golang/glog"
)

func (c *Client) refreshBars() (err error) {
	for _, container := range c.Containers {
		err = c.refreshBar(container)
		if err != nil {
			glog.Error(err)
			return
		}
	}
	return
}

func (c *Client) refreshBar(ct Container) (err error) {
	id := strconv.FormatUint(ct.Job.Id, 10)
	barPath := path.Join(c.Conf.ProgressBarDir, id)
	content, err := ioutil.ReadFile(barPath)
	if err != nil {
		log.Fatal(err)
	}
	encode := string(content)
	currentStep, totalSteps, err := ct.Plugin.ParseBar(encode)
	if err != nil {
		log.Fatal(err)
	}
	ct.Job.CurrentStep = currentStep
	ct.Job.TotalSteps = totalSteps
	return
}
