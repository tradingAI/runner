package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/runner/runner"
)

const (
	trainJobId1 = uint64(1111)
	trainJobId2 = uint64(2222)
	evalJobId   = uint64(3333)
	inferJobId  = uint64(4444)
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")
	conf, _ := runner.LoadConf()
	jobRunner, _ := runner.New(conf)
	testCreateJobs(jobRunner)
}

func newCreateTrainJob(id uint64, runnerId string) (job *pb.Job) {
	job = runner.CreateTestTrainJob()
	job.Id = id
	job.RunnerId = runnerId
	job.Status = pb.JobStatus_CREATED
	return job
}

func newCreateEvalJob(id, modelJobId uint64, runnerId string) (job *pb.Job) {
	job = runner.CreateTestEvalJob()
	job.Id = id
	job.RunnerId = runnerId
	job.Status = pb.JobStatus_CREATED
	job.GetInput().GetEvalInput().GetModel().ObjName = fmt.Sprintf("/test_user/model/%d.zip", modelJobId)
	return job
}

func newCreateInferJob(id, modelJobId uint64, runnerId string) (job *pb.Job) {
	job = runner.CreateTestInferJob()
	job.Id = id
	job.RunnerId = runnerId
	job.Status = pb.JobStatus_CREATED
	job.GetInput().GetInferInput().GetModel().ObjName = fmt.Sprintf("/test_user/model/%d.zip", modelJobId)
	return job
}

func newStopJob(id uint64) (job *pb.Job) {
	return
}

func testCreateJobs(r *runner.Runner) {
	glog.Info("test create Jobs")
	var jobs []*pb.Job
	job1 := newCreateTrainJob(trainJobId1, r.ID)
	job2 := newCreateTrainJob(trainJobId2, r.ID)
	jobs = []*pb.Job{job1, job2}
	r.RunJobs(jobs)
	if job1.Status != pb.JobStatus_SUCCESSED {
		glog.Error("testCreateJobs job1 Failed")
	}
	if job2.Status != pb.JobStatus_SUCCESSED {
		glog.Error("testCreateJobs job2 Failed")
	}

	job3 := newCreateEvalJob(evalJobId, trainJobId1, r.ID)
	job4 := newCreateInferJob(inferJobId, trainJobId2, r.ID)
	jobs = []*pb.Job{job3, job4}
	r.RunJobs(jobs)
	if job3.Status != pb.JobStatus_SUCCESSED {
		glog.Error("testCreateJobs job3 Failed")
	}
	if job4.Status != pb.JobStatus_SUCCESSED {
		glog.Error("testCreateJobs job4 Failed")
	}
}
