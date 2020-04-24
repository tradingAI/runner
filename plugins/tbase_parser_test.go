package plugins

import (
	"testing"
	mpb "github.com/tradingAI/proto/gen/go/model"

	"github.com/stretchr/testify/assert"
)

func TestParseBar(t *testing.T) {
	p := NewTbasePlugin()
	actualCurrentStep, actualTotolStep, err := p.ParseBar("30, 100")
	assert.Equal(t, uint32(30), actualCurrentStep)
	assert.Equal(t, uint32(100), actualTotolStep)
	_, _, err = p.ParseBar("30, 100, ")
	assert.NotNil(t, err)
	_, _, err = p.ParseBar("30, 100, foo")
	assert.NotNil(t, err)
	_, _, err = p.ParseBar("30, foo")
	assert.NotNil(t, err)
	_, _, err = p.ParseBar("30")
	assert.NotNil(t, err)
}

func TestParseEval(t *testing.T){
	p := NewTbasePlugin()
	actual, err := p.ParseEval("0.994,-0.071,-0.148,-0.205,000001.SZ,buy&hold,0.953,-0.588", uint64(1), uint64(1))
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), actual.GetEvalOutput().JobId)
	assert.Equal(t, uint64(1), actual.GetEvalOutput().ModelId)
	assert.Equal(t, float32(0.994), actual.GetEvalOutput().Metrics.AbsoluteReturn)
	assert.Equal(t, float32(-0.071), actual.GetEvalOutput().Metrics.AnnualizedReturn)
	assert.Equal(t, float32(-0.148), actual.GetEvalOutput().Metrics.MaxDrawdown)
	assert.Equal(t, float32(-0.205), actual.GetEvalOutput().Metrics.SharpeRatio)
	assert.Equal(t, "000001.SZ", actual.GetEvalOutput().Metrics.ExcessReturn.BaseCode)
	assert.Equal(t, mpb.BaseStrategy_BUY_AND_HOLD, actual.GetEvalOutput().Metrics.ExcessReturn.Strategy)
	assert.Equal(t, float32(0.953), actual.GetEvalOutput().Metrics.ExcessReturn.AbsoluteValue)
	assert.Equal(t, float32(-0.588), actual.GetEvalOutput().Metrics.ExcessReturn.AnnualizedValue)
}
