package runner

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadAndUnarchiveModel(t *testing.T) {
	r := creatTestRunner()
	trainJob := createTestTrainJob()
	trainJob.Id = uint64(22222)
	modelDir, err := r.downloadAndUnarchiveModel(trainJob)
	assert.NotNil(t, err)
	assert.Equal(t, "", modelDir)
	DeleteFile(path.Join(r.Conf.ModelDir, "22222.zip"))
	// eval job
	// upload test model for test
	err = r.uploadTrainModel(trainJob)
	assert.Nil(t, err)
	evalJob := createTestEvalJob()
	modelDir, err = r.downloadAndUnarchiveModel(evalJob)
	assert.Nil(t, err)
	expectedDir := path.Join(r.Conf.ModelDir, "3")
	assert.Equal(t, expectedDir, modelDir)
	DeleteDirectory(modelDir)
	DeleteFile(path.Join(r.Conf.ModelDir, "3.zip"))
	// infer job
	inferJob := createTestInferJob()
	modelDir, err = r.downloadAndUnarchiveModel(inferJob)
	assert.Nil(t, err)
	expectedDir = path.Join(r.Conf.ModelDir, "4")
	assert.Equal(t, expectedDir, modelDir)
	DeleteDirectory(modelDir)
	DeleteFile(path.Join(r.Conf.ModelDir, "4.zip"))
}
