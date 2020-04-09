package client

import (
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/runner/plugins"
)


func creatTestClient() (c *Client) {
    conf, _ := LoadConf()
	c, _ = New(conf)
    return
}

func createTestJob() (job *pb.Job) {
	job = &pb.Job{
		Id:       uint64(123456789),
		RunnerId: "test_runner_id",
		Type:     pb.JobType_TRAIN,
		Input:    plugins.CreateDefaultTbaseTrainJobInput(),
	}
	return
}

func createTestContainer() (container *Container) {
	job := createTestJob()
	container = &Container{
		Name:    "123456789",
		ID:      "123456789",
		ShortID: "123",
		Job:     job,
		Plugin:  plugins.New(job),
	}
	return
}
