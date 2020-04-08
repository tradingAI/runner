package client

import (
	"errors"
	"os"
	"strconv"
	"path"

	"github.com/golang/glog"
	minio "github.com/tradingAI/go/s3/minio"
)

const DATA_ROOT = "/tmp/data"

type Conf struct {
	StorageDir       string
	Minio            minio.MinioConf
	HeartbeatSeconds int
	JobLogDir        string
	JobShellDir      string
	TushareToken     string
	ModelDir         string
	ProgressBarDir   string
	TensorboardDir   string
	InferDir         string
	EvalDir          string
}

// LoadConf load config from env
func LoadConf() (conf Conf, err error) {
	minioPort, err := strconv.Atoi(os.Getenv("RUNNER_MINIO_PORT"))
	if err != nil {
		glog.Error(err)
		return
	}

	minioSecure, err := strconv.ParseBool(os.Getenv("RUNNER_MINIO_SECURE"))
	if err != nil {
		glog.Error(err)
		return
	}

	heartbeatSeconds, err := strconv.Atoi(os.Getenv("HEARTBEAT_SECONDS"))
	if err != nil {
		glog.Error(err)
		return
	}

	conf = Conf{
		StorageDir: os.Getenv("TWEB_STORAGE_DIR"),
		Minio: minio.MinioConf{
			AccessKey: os.Getenv("RUNNER_MINIO_ACCESS_KEY"),
			SecretKey: os.Getenv("RUNNER_MINIO_SECRET_KEY"),
			Host:      os.Getenv("RUNNER_MINIO_HOST"),
			Port:      minioPort,
			Secure:    minioSecure,
		},
		HeartbeatSeconds: heartbeatSeconds,
		JobLogDir:        os.Getenv("JOB_LOG_DIR"),
		JobShellDir:      os.Getenv("JOB_SHELL_DIR"),
		TushareToken:     os.Getenv("TUSHARE_TOKEN"),
		ModelDir:         path.Join(DATA_ROOT, "model"),
		ProgressBarDir:   path.Join(DATA_ROOT, "progress_bar"),
		TensorboardDir:   path.Join(DATA_ROOT, "tensorboard"),
		InferDir:         path.Join(DATA_ROOT, "inferences"),
		EvalDir:          path.Join(DATA_ROOT, "evals"),
	}

	if err = conf.Validate(); err != nil {
		glog.Error(err)
		return
	}

	return
}

func (c *Conf) Validate() (err error) {
	if err = c.Minio.Validate(); err != nil {
		glog.Error(err)
		return
	}
	if c.TushareToken == "" {
		errMsg := "TushareToken is empty"
		err = errors.New(errMsg)
		glog.Error(err)
		return
	}
	return
}
