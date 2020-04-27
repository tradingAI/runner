package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/runner/runner"
	"google.golang.org/grpc"
)

const (
	trainJobId1 = uint64(1111)
	trainJobId2 = uint64(2222)
	evalJobId   = uint64(3333)
	inferJobId  = uint64(4444)
)

type server struct {
	pb.SchedulerServer
	numHeartbeat int32
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

func (s *server) HeartBeat(ctx context.Context, req *pb.HeartBeatRequest) (resp *pb.HeartBeatResponse, err error) {
	runnerPb := req.GetRunner()
	if runnerPb == nil {
		err = errors.New("HeartBeat runnerPb is nil")
		glog.Error(err)
		return
	}
	glog.Infof("Received: %v %v", s.numHeartbeat, req.GetRunner())
	runnerId := runnerPb.GetId()

	if runnerId == "" {
		err = errors.New("HeartBeat runner id is empty")
		glog.Error(err)
		return
	}
	var newJobs []*pb.Job
	jobsPb := runnerPb.GetJobs()
	if len(jobsPb) < 1 {
		if s.numHeartbeat % 20 == 0 {
			job1 := newCreateTrainJob(trainJobId1, runnerId)
			job2 := newCreateTrainJob(trainJobId2, runnerId)
			newJobs = []*pb.Job{job1, job2}
		}
		if s.numHeartbeat % 4 == 3 {
			job1 := newCreateEvalJob(evalJobId, trainJobId1, runnerId)
			job2 := newCreateInferJob(inferJobId, trainJobId2, runnerId)
			newJobs = []*pb.Job{job1, job2}
		}
	}
	s.numHeartbeat += 1
	if  s.numHeartbeat >= 25 {
		return &pb.HeartBeatResponse{
			Ok:      true,
			Jobs:    newJobs,
			Destory: true,
		}, nil
	}
	return &pb.HeartBeatResponse{
		Ok:      true,
		Jobs:    newJobs,
		Destory: false,
	}, nil
}

func (s *server) CreateJob(ctx context.Context, req *pb.CreateJobRequest) (resp *pb.CreateJobResponse, err error) {
	return
}

func (s *server) StopJob(ctx context.Context, req *pb.StopJobRequest) (resp *pb.StopJobResponse, err error) {
	return
}

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")
	port := os.Getenv("SCHEDULER_PORT")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSchedulerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
