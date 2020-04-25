package runner

import (
	"io/ioutil"
	"path"
	"strconv"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

func (r *Runner) refreshBars() (err error) {
	for _, container := range r.Containers {
		if container.Job.Status == pb.JobStatus_SUCCESSED {
			continue
		}
		id := strconv.FormatUint(container.Job.Id, 10)
		barPath := path.Join(r.Conf.ProgressBarDir, id)
		err = container.refreshBar(barPath)
		if err != nil {
			glog.Error(err)
			return
		}
	}
	return
}

func (ct *Container) refreshBar(barPath string) (err error) {
	content, err := ioutil.ReadFile(barPath)
	if err != nil {
		glog.Error(err)
		return
	}
	encode := string(content)
	currentStep, totalSteps, err := ct.Plugin.ParseBar(encode)
	if err != nil {
		glog.Error(err)
		return
	}
	ct.Job.CurrentStep = currentStep
	ct.Job.TotalSteps = totalSteps
	return
}
