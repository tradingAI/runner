package runner

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang/glog"
	minio "github.com/tradingAI/go/s3/minio"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"google.golang.org/grpc"
)

type Runner struct {
	Conf            *Conf
	Minio           *minio.Client
	ID              string
	Token           string
	Containers      map[uint64]Container // key: jobID, value: Container
	Machine         *Machine
	schedulerClient pb.SchedulerClient
	schedulerConn   *grpc.ClientConn
	Status          pb.RunnerStatus
}

func New(conf *Conf) (r *Runner, err error) {
	// make runner
	r = &Runner{
		Conf:  conf,
		ID:    "test_runner_id", // TODO: use uuid
		Token: "test_token",     // TODO: use evn token
	}

	r.Minio, err = minio.NewMinioClient(r.Conf.Minio)
	if err != nil {
		glog.Error(err)
		return
	}

	r.Containers = make(map[uint64]Container)
	machine, err := NewMachine()
	if err != nil {
		glog.Error(err)
		return
	}
	r.Machine = machine
	err = r.Machine.Update()
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

func (r *Runner) StartOrDie() (err error) {
	glog.Info("Starting runner")
	r.Heartbeat()
	d := time.Duration(int64(time.Second) * int64(r.Conf.HeartbeatSeconds))
	t := time.NewTicker(d)
	defer t.Stop()

	for {
		<-t.C
		r.Heartbeat()
	}
	return
}

func (r *Runner) Free() {
	return
}

func (r *Runner) updateStatus() {
	if len(r.Containers) > 0 {
		r.Status = pb.RunnerStatus_BUSY
	} else {
		r.Status = pb.RunnerStatus_IDLE
	}
}

func (r *Runner) Heartbeat() (err error) {
	glog.Infof("runner[%s] heartbeat", r.ID)
	r.refreshBars()
	// update machine info
	err = r.Machine.Update()
	if err != nil {
		glog.Error(err)
		return
	}
	r.updateStatus()
	var jobs []*pb.Job
	for _, c := range r.Containers {
		jobs = append(jobs, c.Job)
	}
	pbRunner := &pb.Runner{
		Id:                 r.ID,
		Status:             r.Status,
		Jobs:               jobs,
		CpuCoreNum:         r.Machine.CPUNum,
		CpuUtilization:     r.Machine.CPUUtilization,
		GpuNum:             r.Machine.GPUNum,
		GpusIndex:          r.Machine.GPUsIndex,
		GpuUtilization:     r.Machine.GPUUtilization,
		Memory:             r.Machine.Memory,
		AvailableMemory:    r.Machine.AvailableMemory,
		GpuMemory:          r.Machine.GPUMemory,
		AvailableGpuMemory: r.Machine.AvailableGPUMemory,
		Token:              r.Token,
	}
	glog.Infof("runner pbRunner %v", pbRunner)
	glog.Infof("r.Machine.CPUUtilization: %.3f", r.Machine.CPUUtilization)
	req := &pb.HeartBeatRequest{
		Runner: pbRunner,
	}
	conn, err := grpc.Dial(r.Conf.SchedulerHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		glog.Error(err)
		return
	}
	defer conn.Close()
	schedulerClient := pb.NewSchedulerClient(conn)

	resp, err := schedulerClient.HeartBeat(context.Background(), req)
	if err != nil {
		glog.Error(err)
		return
	}

	if !resp.Ok {
		err = errors.New("Runner heartbeat: scheduler rpc server err response")
		glog.Error(err)
		return
	}
	if resp != nil {
		for _, job := range resp.Jobs {
			err = r.RunJob(job)
			if err != nil {
				glog.Error(err)
			}
		}
	}
	return
}

func (r *Runner) RunJob(job *pb.Job) (err error) {

	switch job.Status {
	case pb.JobStatus_CREATED:
		go func(r *Runner) {
			err := r.CreateJob(job)
			if err != nil {
				glog.Error(err)
				return
			}
		}(r)
	case pb.JobStatus_CANCELLED:
		go func(r *Runner) {
			err := r.StopJob(job.Id)
			if err != nil {
				glog.Error(err)
				return
			}
		}(r)
	default:
		err = errors.New(fmt.Sprintf("Runner RunJob error: invalid job status %v", job))
		glog.Error(err)
		return
	}

	return
}
