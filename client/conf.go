package client

import (
	"errors"
	"os"
	"strconv"

	minio "github.com/tradingAI/go/s3/minio"
)

type Conf struct {
	StorageDir       string
	Minio            minio.MinioConf
	HeartbeatSeconds int
	JobLogDir        string
	JobShellDir      string
	TushareToken     string
}

// LoadConf load config from env
func LoadConf() (conf Conf, err error) {
	minioPort, err := strconv.Atoi(os.Getenv("RUNNER_MINIO_PORT"))
	if err != nil {
		panic(err)
	}

	minioSecure, err := strconv.ParseBool(os.Getenv("RUNNER_MINIO_SECURE"))
	if err != nil {
		panic(err)
	}

	heartbeatSeconds, err := strconv.Atoi(os.Getenv("HEARTBEAT_SECONDS"))
	if err != nil {
		panic(err)
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
	}

	if err = conf.Validate(); err != nil {
		panic(err)
	}

	return
}

func (c *Conf) Validate() (err error) {
	if err = c.Minio.Validate(); err != nil {
		panic(err)
	}
	if c.TushareToken == "" {
		errMsg := "TushareToken is empty"
		err = errors.New(errMsg)
		panic(err)
	}
	return
}
