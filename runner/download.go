package runner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"path"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

func getFileExt(fileType string) (ext string, err error) {
	switch fileType {
	case "application/x-tar":
		ext = "tar"
		return
	case "application/zip":
		ext = "zip"
		return
	default:
		// https://github.com/tradingAI/proto/blob/e3a639a8f6599629ac7f49fc3c638e435e27dcec/common/model.proto#L23
		err = errors.New("file type invalid")
		return
	}
	return
}

func (r *Runner) getFileName(job *pb.Job) (fileName string, err error) {
	id := strconv.FormatUint(job.Id, 10)
	input := job.GetInput()
	var fileType string
	switch input.GetInput().(type) {
	case *pb.JobInput_EvalInput:
		fileType = input.GetEvalInput().GetModel().GetFileType()
		ext, err := getFileExt(fileType)
		if err != nil {
			glog.Error(err)
			return fileName, err
		}
		fileName = fmt.Sprintf("%s.%s", id, ext)
		return fileName, err
	case *pb.JobInput_InferInput:
		fileType = input.GetInferInput().GetModel().GetFileType()
		ext, err := getFileExt(fileType)
		if err != nil {
			glog.Error(err)
			return fileName, err
		}
		fileName = fmt.Sprintf("%s.%s", id, ext)
		return fileName, err
	default:
		err = errors.New("runner getFileName job.JobInput invalid")
		glog.Error(err)
		return fileName, err
	}
	return
}

func (r *Runner) getMinioInfo(job *pb.Job) (bucket, objName string, err error) {
	input := job.GetInput()
	switch input.GetInput().(type) {
	case *pb.JobInput_EvalInput:
		bucket = input.GetEvalInput().GetModel().GetBucket()
		objName = input.GetEvalInput().GetModel().GetObjName()
		return
	case *pb.JobInput_InferInput:
		bucket = input.GetInferInput().GetModel().GetBucket()
		objName = input.GetInferInput().GetModel().GetObjName()
		return
	default:
		err = errors.New("runner getBucket job.JobInput invalid")
		glog.Error(err)
		return
	}
	return
}

func (r *Runner) downloadAndUnarchiveModel(job *pb.Job) (modelDir string, err error) {
	fileName, err := r.getFileName(job)
	if err != nil {
		glog.Error(err)
		return
	}
	modelPath := path.Join(r.Conf.ModelDir, fileName)
	bucket, objName, err := r.getMinioInfo(job)
	if err != nil {
		glog.Error(err)
		return
	}
	// download
	err = r.Minio.MinioDownload(bucket, modelPath, objName)
	if err != nil {
		glog.Error(err)
		return
	}
	defer DeleteFile(modelPath)
	// unarchive model
	id := strconv.FormatUint(job.Id, 10)
	modelDir = path.Join(r.Conf.ModelDir, id)
	tmpDir := path.Join(r.Conf.ModelDir, id+"_tmp")
	err = Unarchive(modelPath, tmpDir)
	if err != nil {
		glog.Error(err)
		return
	}
	defer DeleteDirectory(tmpDir)
	OriginModelDir, err := getOriginModelDir(tmpDir)
	if err != nil {
		glog.Error(err)
		return
	}
	err = os.Rename(OriginModelDir, modelDir)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

// 从解压的目录中提取出源模型所在的目录
func getOriginModelDir(srcDir string) (dir string, err error) {
	dirs, err := ioutil.ReadDir(srcDir)
	if err != nil {
		glog.Error(err)
		return
	}
	if cap(dirs) != 1 {
		err = errors.New("runner getOriginModelDir: model file is invalid, only 1 directory is allowed")
		glog.Error(err)
		return
	}
	dir = path.Join(srcDir, dirs[0].Name())
	if !dirs[0].IsDir() {
		err = errors.New("runner getOriginModelDir: model file is invalid, only 1 directory is allowed")
		glog.Error(err)
		return
	}
	return
}
