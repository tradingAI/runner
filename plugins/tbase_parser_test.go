package plugins

import (
	"testing"

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
