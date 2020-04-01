package client

import (
	"os"
	"strconv"
	"time"

	"github.com/golang/glog"
	pg "github.com/tradingAI/go/db/postgres"
	minio "github.com/tradingAI/go/s3/minio"
)

type Conf struct {
	DB         pg.DBConf
	StorageDir string
	Minio      minio.MinioConf
}

// LoadConf load config from env
func LoadConf() (conf Conf, err error) {
	dbReconnectSec, err := strconv.Atoi(os.Getenv("TWEB_POSTGRES_RECONNECT_SEC"))
	if err != nil {
		glog.Error(err)
		return
	}

	dbPort, err := strconv.Atoi(os.Getenv("TWEB_POSTGRES_PORT"))
	if err != nil {
		glog.Error(err)
		return
	}

	minioPort, err := strconv.Atoi(os.Getenv("SCHEDULER_MINIO_PORT"))
	if err != nil {
		glog.Error(err)
		return
	}

	minioSecure, err := strconv.ParseBool(os.Getenv("TWEB_MINIO_SECURE"))
	if err != nil {
		glog.Error(err)
		return
	}
	conf = Conf{
		DB: pg.DBConf{
			Database:     os.Getenv("TWEB_POSTGRES_DB"),
			Username:     os.Getenv("TWEB_POSTGRES_USER"),
			Password:     os.Getenv("TWEB_POSTGRES_PASSWORD"),
			Port:         dbPort,
			Host:         os.Getenv("TWEB_POSTGRES_HOST"),
			ReconnectSec: time.Duration(dbReconnectSec) * time.Second,
		},
		StorageDir: os.Getenv("TWEB_STORAGE_DIR"),
		Minio: minio.MinioConf{
			AccessKey: os.Getenv("TWEB_MINIO_ACCESS_KEY"),
			SecretKey: os.Getenv("TWEB_MINIO_SECRET_KEY"),
			Host:      os.Getenv("TWEB_MINIO_HOST"),
			Port:      minioPort,
			Secure:    minioSecure,
		},
	}

	if err = conf.Validate(); err != nil {
		glog.Error(err)
		return
	}

	return
}
