package client

import (
	"time"

	"github.com/golang/glog"
	"github.com/minio/minio-go/v6"
	minio2 "github.com/tradingAI/go/s3/minio"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

type Container struct{
    Name string
    ID string
    ShortID string
    Job *pb.Job
}

type Client struct {
	Conf  Conf
	Minio *minio.Client
	ID string
	// key: jobID, value: Container
	Containers map[uint64]Container
}

func New(conf Conf) (c *Client, err error) {
	// make client
	c = &Client{
		Conf: conf,
	}

	c.Minio, err = minio2.NewMinioClient(c.Conf.Minio)
	if err != nil {
		glog.Error(err)
		return
	}

	// TODO: use uuid
	c.ID = "test_runner_id"

	c.Containers = make(map[uint64]Container)

	return
}

func (c *Client) StartOrDie() (err error) {
	glog.Info("Starting runner")
	// c.Heartbeat()
	// c.Listen()
	d := time.Duration(int64(time.Second) * int64(c.Conf.HeartbeatSeconds))
	t := time.NewTicker(d)
    defer t.Stop()

    for {
            <- t.C
			c.Heartbeat()
			c.Listen()
    }
	return
}

func (c *Client) Free() {
	return
}

func (c *Client) Heartbeat() (err error) {
	glog.Infof("runner[%s] heartbeat", c.ID)
	// TODO: collect machine info call rpc hearbeat
	return
}

func (c *Client)  getCreateJobFromRedis()(job *pb.Job, err error){
	// TODO
	job = &pb.Job{
		Id: uint64(123456789),
		RunnerId: c.ID,
		Type: pb.JobType_TRAIN,
	}
	return job, nil
}

func (c *Client)  getStopJobFromRedis()(job *pb.Job, err error){
	// TODO
	job = &pb.Job{
		Id: uint64(123456789),
		RunnerId: c.ID,
		Type: pb.JobType_TRAIN,
	}
	return job, nil
}

func (c *Client) Listen() (err error) {
	glog.Infof("runner[%s] listen job from redis", c.ID)
	// TODO: listen redis status and excute actions
	// create
	createJob, _ := c.getCreateJobFromRedis()
	if createJob != nil {
		go func(c *Client) {
			c.CreateJob(createJob)
		}(c)
	}

	glog.Infof("runner[%s] sleepping 50 seconds", c.ID)
	time.Sleep(10 * time.Second)
	// stop
	stopJob, _ := c.getStopJobFromRedis()
	if stopJob != nil {
		go func(c *Client) {
			c.StopJob(stopJob.Id)
			c.RemoveContainer(stopJob.Id)
		}(c)
	}
	return
}
