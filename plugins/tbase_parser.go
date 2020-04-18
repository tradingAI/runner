package plugins

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

func (p *TbasePlugin) ParseBar(encode string) (currentStep uint32, totalSteps uint32, err error) {
	values := strings.Split(encode, p.Sep)
	if cap(values) != 2 {
		errMsg := fmt.Sprintf("TbasePlugin ParseBar input invalid encode: %s, expected: currentStep, totalSteps", encode)
		err = errors.New(errMsg)
		glog.Error(err)
		return
	}
	currentStepI64, err := strconv.ParseInt(strings.TrimSpace(values[0]), 10, 32)
	if err != nil {
		glog.Error(err)
		return
	}
	currentStep = uint32(currentStepI64)

	totalStepsI64, err := strconv.ParseInt(strings.TrimSpace(values[1]), 10, 32)
	if err != nil {
		glog.Error(err)
		return
	}
	totalSteps = uint32(totalStepsI64)
	return
}

func (p *TbasePlugin) ParseEval(encode string) (out *pb.JobOutput, err error) {
	return
}

func (p *TbasePlugin) ParseInfer(encode string) (out *pb.JobOutput, err error) {
	return
}
