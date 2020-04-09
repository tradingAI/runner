package runner

import (
	"errors"
	"os"
	"strconv"

	"github.com/golang/glog"
	minio "github.com/tradingAI/go/s3/minio"
	"github.com/tradingAI/runner/plugins"
)

type Conf struct {
	StorageDir       string
	Minio            minio.MinioConf
	HeartbeatSeconds int
	TushareToken     string
	DataRootDir      string
	JobLogDir        string
	JobShellDir      string
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
		TushareToken:     os.Getenv("TUSHARE_TOKEN"),
		DataRootDir:      plugins.ROOT_DATA_DIR,
		JobLogDir:        plugins.JOB_LOG_DIR,
		JobShellDir:      plugins.JOB_SHELL_DIR,
		ModelDir:         plugins.MODEL_DIR,
		ProgressBarDir:   plugins.PROGRESS_BAR_DIR,
		TensorboardDir:   plugins.TENSORBOARD_DIR,
		InferDir:         plugins.INFER_DIR,
		EvalDir:          plugins.EVAL_DIR,
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
