package plugins

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/glog"
	mpb "github.com/tradingAI/proto/gen/go/model"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

func (p *TbasePlugin) getBaseStrategy(name string) (s mpb.BaseStrategy) {
	switch name {
	case "buy&hold":
		return mpb.BaseStrategy_BUY_AND_HOLD
	default:
		return mpb.BaseStrategy_UNKNOWN
	}
	return
}

func (p *TbasePlugin) ParseBar(encode string) (currentStep uint32, totalSteps uint32, err error) {
	values := strings.Split(encode, p.Sep)
	if cap(values) != 2 {
		err = errors.New(fmt.Sprintf("TbasePlugin ParseBar input invalid encode: %s, expected: currentStep, totalSteps", encode))
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

func (p *TbasePlugin) ParseEval(encode string, jobId, modelId uint64) (out *pb.JobOutput, err error) {
	values := strings.Split(encode, p.Sep)
	if cap(values) != 8 {
		err = errors.New(fmt.Sprintf("TbasePlugin ParseEval input invalid encode: %s, expected 8 values", encode))
		glog.Error(err)
		return
	}

	absoluteReturn, err := strconv.ParseFloat(values[0], 32)
	if err != nil {
		glog.Error(err)
		return
	}

	annualizedReturn, err := strconv.ParseFloat(values[1], 32)
	if err != nil {
		glog.Error(err)
		return
	}

	maxDrawdown, err := strconv.ParseFloat(values[2], 32)
	if err != nil {
		glog.Error(err)
		return
	}

	sharpeRatio, err := strconv.ParseFloat(values[3], 32)
	if err != nil {
		glog.Error(err)
		return
	}

	absoluteExcessValue, err := strconv.ParseFloat(values[6], 32)
	if err != nil {
		glog.Error(err)
		return
	}

	annualizedExcessValue, err := strconv.ParseFloat(values[7], 32)
	if err != nil {
		glog.Error(err)
		return
	}

	excessReturn := &mpb.ExcessReturn{
		BaseCode:        values[4],
		Strategy:        p.getBaseStrategy(values[5]),
		AbsoluteValue:   float32(absoluteExcessValue),
		AnnualizedValue: float32(annualizedExcessValue),
	}

	evaluateMetricts := &mpb.EvaluateMetrics{
		AbsoluteReturn:   float32(absoluteReturn),
		AnnualizedReturn: float32(annualizedReturn),
		MaxDrawdown:      float32(maxDrawdown),
		SharpeRatio:      float32(sharpeRatio),
		ExcessReturn:     excessReturn,
	}

	tbaseEvaluateOutput := &mpb.TbaseEvaluateOutput{
		JobId:   jobId,
		ModelId: modelId,
		Metrics: evaluateMetricts,
	}

	out = &pb.JobOutput{
		Output: &pb.JobOutput_EvalOutput{tbaseEvaluateOutput},
	}

	return
}

func (p *TbasePlugin) getInferActionType(act string) (t mpb.ActionType, err error) {
	switch act {
	case "suggest_sell":
		t = mpb.ActionType_SUGGEST_SELL
		return
	case "suggest_buy":
		t = mpb.ActionType_SUGGEST_BUY
		return
	case "suggest_volume":
		t = mpb.ActionType_SUGGEST_VOLUME
		return
	case "buy_volume_pct":
		t = mpb.ActionType_BUY_VOLUME_PCT
		return
	case "sell_volume_pct":
		t = mpb.ActionType_SELL_VOLUME_PCT
		return
	default:
		err = errors.New(fmt.Sprintf("TbasePlugin getInferActionType: invalid act: %s", act))
		glog.Error(err)
		return
	}
	return
}

func (p *TbasePlugin) ParseInfer(encode, date string, jobId, modelId uint64) (out *pb.JobOutput, err error) {
	lines := strings.Split(encode, "\n")
	actions := []*mpb.TbaseAction{}
	for _, line := range lines {
		values := strings.Split(line, p.Sep)
		if cap(values) != 3 {
			err = errors.New(fmt.Sprintf("TbasePlugin ParseInfer: invalid act: %s", line))
			glog.Error(err)
			return
		}
		actType, err := p.getInferActionType(values[0])
		if err != nil {
			glog.Error(err)
			return nil, err
		}
		action := &mpb.TbaseAction{
			ActionType: actType,
			Value: values[1],
			Code: values[2],
		}
		actions = append(actions, action)
	}
	tbaseInferOutput := &mpb.TbaseInferOutput{
		JobId:   jobId,
		ModelId: modelId,
		Date: date,
		Actions: actions,
	}

	out = &pb.JobOutput{
		Output: &pb.JobOutput_InferOutput{tbaseInferOutput},
	}
	return
}
