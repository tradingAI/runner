package runner

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"

	"path"
)

const UPLOAD_CONTENT_TYPE = "application/zip"

func (r *Runner) uploadTrainModel(job *pb.Job) (err error) {
	glog.Infof("Runner uplading train model job.id=%d", job.Id)
	id := strconv.FormatUint(job.Id, 10)
	dir := path.Join(r.Conf.ModelDir, id)
	// achive
	modelPath, err := Archive(dir)
	if err != nil {
		glog.Error(err)
		return err
	}
	defer DeleteFile(modelPath)
	// upload
	trainInput := job.GetInput().GetTrainInput()
	if trainInput == nil {
		errMsg := "trainInput is nil, JobInput invalid!"
		glog.Error("errMsg")
		return errors.New(errMsg)
	}
	bucket := trainInput.GetBucket()
	objName := path.Join(trainInput.GetModelFileDir(), filepath.Base(modelPath))
	err = r.Minio.MinioUpload(bucket, modelPath, objName, UPLOAD_CONTENT_TYPE)
	if err != nil {
		glog.Error(err)
		return err
	}
	return
}

func (r *Runner) uploadTensorboard(job *pb.Job) (err error) {
	glog.Infof("Runner uplading tensorboard event job.id=%d", job.Id)
	id := strconv.FormatUint(job.Id, 10)
	dir := path.Join(r.Conf.TensorboardDir, id)
	// achive
	tbPath, err := Archive(dir)
	if err != nil {
		glog.Error(err)
		return err
	}
	defer DeleteFile(tbPath)
	// upload
	trainInput := job.GetInput().GetTrainInput()
	if trainInput == nil {
		err = fmt.Errorf("runner job request_id: %s trainInput is nil, JobInput invalid!", job.RequestId)
		glog.Error(err)
		return err
	}
	bucket := trainInput.GetBucket()
	objName := path.Join(trainInput.GetTensorboardFileDir(), filepath.Base(tbPath))
	err = r.Minio.MinioUpload(bucket, tbPath, objName, UPLOAD_CONTENT_TYPE)
	if err != nil {
		glog.Error(err)
		return err
	}
	return
}
