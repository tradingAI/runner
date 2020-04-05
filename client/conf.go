package client

import (
	"os"
	"strconv"

	"github.com/golang/glog"
	minio "github.com/tradingAI/go/s3/minio"
)

type Conf struct {
	StorageDir       string
	Minio            minio.MinioConf
	HeartbeatSeconds int
	JobLogDir 		 string
}

// LoadConf load config from env
func LoadConf() (conf Conf, err error) {
	minioPort, err := strconv.Atoi(os.Getenv("TWEB_MINIO_PORT"))
	if err != nil {
		glog.Error(err)
		return
	}

	minioSecure, err := strconv.ParseBool(os.Getenv("TWEB_MINIO_SECURE"))
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
			AccessKey: os.Getenv("TWEB_MINIO_ACCESS_KEY"),
			SecretKey: os.Getenv("TWEB_MINIO_SECRET_KEY"),
			Host:      os.Getenv("TWEB_MINIO_HOST"),
			Port:      minioPort,
			Secure:    minioSecure,
		},
		HeartbeatSeconds: heartbeatSeconds,
		JobLogDir: os.Getenv("RUNNER_LOG_DIR"),
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
	return
}
