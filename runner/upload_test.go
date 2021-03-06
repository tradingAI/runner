package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTrainModel(t *testing.T) {
	r := creatTestRunner()
	job := CreateTestTrainJob()
	err := r.uploadTrainModel(job)
    assert.Nil(t, err)
}

func TestUploadTensorboard(t *testing.T) {
	r := creatTestRunner()
	job := CreateTestTrainJob()
	err := r.uploadTensorboard(job)
    assert.Nil(t, err)
}
