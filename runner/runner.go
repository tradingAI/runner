package runner

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/google/uuid"
	minio "github.com/tradingAI/go/s3/minio"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"google.golang.org/grpc"
)

type Runner struct {
	Conf            *Conf
	Minio           *minio.Client
	ID              string
	Containers      map[uint64]Container // key: jobID, value: Container
	Machine         *Machine
	schedulerClient pb.SchedulerClient
	schedulerConn   *grpc.ClientConn
	Status          pb.RunnerStatus
}

func New(conf *Conf) (r *Runner, err error) {
	// make runner
	r = &Runner{
		Conf: conf,
		ID:   uuid.New().String(),
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
	r.RunJobs(resp.Jobs)
	return
}

// NOTE(wen): 一个job出现错误，不影响runner继续执行, 所以没有返回error
func (r *Runner) RunJobs(jobs []*pb.Job) {
	for _, job := range jobs {
		glog.Infof("Runner RunJobs: job.id=%d", job.Id)
		r.RunJob(job)
	}
}

func (r *Runner) RunJob(job *pb.Job) (err error) {
	glog.Infof("Runner RunJob: job=[%v]", job)
	switch job.Status {
	case pb.JobStatus_CREATED:
		glog.Infof("Runner RunJob: creating job.id=%d", job.Id)
		job.Status = pb.JobStatus_RUNNING
		err = r.CreateJob(job)
		if err != nil {
			glog.Error(err)
			job.Status = pb.JobStatus_FAILED
			return
		}
		job.Status = pb.JobStatus_SUCCESSED
		return
	case pb.JobStatus_CANCELLED:
		glog.Infof("Runner RunJob: stopping job.id=%d", job.Id)
		err = r.StopJob(job)
		if err != nil {
			glog.Error(err)
			job.Status = pb.JobStatus_FAILED
		}
		job.Status = pb.JobStatus_CANCELLED
		// TODO: 重新定义stop状态, stopped
		return
	default:
		err = errors.New(fmt.Sprintf("Runner RunJob error: invalid job status %v", job))
		glog.Error(err)
		job.Status = pb.JobStatus_FAILED
		return
	}

	return
}
