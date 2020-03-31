package client

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	pg "github.com/tradingAI/go/db/postgres"
	minio2 "github.com/tradingAI/go/s3/minio"
)

type Client struct {
	Conf  Conf
	DB    *gorm.DB
	Minio *minio.Client
}

func New(conf Conf) (s *Client, err error) {
	// make server
	c = &Client{
		Conf: conf,
	}

	// Init db
	c.DB, err = pg.NewPostgreSQL(conf.DB)
	if err != nil {
		glog.Error(err)
		return
	}

	c.Minio, err = minio2.NewMinioClient(s.Conf.Minio)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func (c *Client) StartOrDie() (err error) {
	glog.Info("Hello world")
	return
}

func (c *Client) Free() {
	if err := c.DB.Close(); err != nil {
		glog.Warning(err)
	}
	return
}