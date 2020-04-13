package runner

import (
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/runner/plugins"
)


func creatTestRunner() (r *Runner) {
    conf, _ := LoadConf()
	r, _ = New(conf)
    return
}

func createTestTrainJob() (job *pb.Job) {
	job = &pb.Job{
		Id:       uint64(123456789),
		RunnerId: "test_runner_id",
		Type:     pb.JobType_TRAIN,
		Input:    plugins.CreateDefaultTbaseTrainJobInput(),
	}
	return
}

func createTestEvalJob() (job *pb.Job) {
	job = &pb.Job{
		Id:       uint64(3),
		RunnerId: "test_runner_id",
		Type:     pb.JobType_EVALUATION,
		Input:    plugins.CreateDefaultTbaseEvalJobInput(),
	}
	return
}

func createTestInferJob() (job *pb.Job) {
	job = &pb.Job{
		Id:       uint64(4),
		RunnerId: "test_runner_id",
		Type:     pb.JobType_INFER,
		Input:    plugins.CreateDefaultTbaseInferJobInput(),
	}
	return
}


func createTestContainer() (container *Container) {
	job := createTestTrainJob()
	container = &Container{
		Name:    "123456789",
		ID:      "123456789",
		ShortID: "123",
		Job:     job,
		Plugin:  plugins.New(job),
	}
	return
}
