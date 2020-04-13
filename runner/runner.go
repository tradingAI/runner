package runner

import (
	"time"

	"github.com/golang/glog"
	minio "github.com/tradingAI/go/s3/minio"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/runner/plugins"
)

type Runner struct {
	Conf       Conf
	Minio      *minio.Client
	ID         string
	Token      string
	Containers map[uint64]Container // key: jobID, value: Container
}

func New(conf Conf) (r *Runner, err error) {
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

	return
}

func (r *Runner) StartOrDie() (err error) {
	glog.Info("Starting runner")
	r.Heartbeat()
	r.Listen()
	// TODO: remove , 测试用
	r.Conf.HeartbeatSeconds = 1500
	d := time.Duration(int64(time.Second) * int64(r.Conf.HeartbeatSeconds))
	t := time.NewTicker(d)
	defer t.Stop()

	for {
		<-t.C
		r.Heartbeat()
		r.Listen()
		// TODO: remove 本地测试用
		return
	}
	return
}

func (r *Runner) Free() {
	return
}

func (r *Runner) Heartbeat() (err error) {
	glog.Infof("runner[%s] heartbeat", r.ID)
	r.refreshBars()
	// TODO: collect machine info call rpc hearbeat
	return
}

func (r *Runner) getCreateJobFromRedis() (job *pb.Job, err error) {
	// TODO
	job = &pb.Job{
		Id:       uint64(123456789),
		RunnerId: r.ID,
		Type:     pb.JobType_TRAIN,
		Input:    plugins.CreateDefaultTbaseTrainJobInput(),
	}
	return job, nil
}

func (r *Runner) getStopJobFromRedis() (job *pb.Job, err error) {
	// TODO
	// job = &pb.Job{
	// 	Id:       uint64(123456789),
	// 	RunnerId: r.ID,
	// 	Type:     pb.JobType_TRAIN,
	// }
	// return job, nil
	return nil, nil
}

func (r *Runner) Listen() (err error) {
	glog.Infof("runner[%s] listen job from redis", r.ID)
	// TODO: listen redis status and excute actions
	// create
	createJob, _ := r.getCreateJobFromRedis()
	if createJob != nil {
		go func(r *Runner) {
			err := r.CreateJob(createJob)
			if err != nil {
				glog.Error(err)
			}
		}(r)
	}
	// TODO: 删除sleep, 暂时用于本地测试用
	time.Sleep(15 * time.Second)
	// stop
	stopJob, _ := r.getStopJobFromRedis()
	if stopJob != nil {
		go func(r *Runner) {
			err := r.StopJob(stopJob.Id)
			if err != nil {
				glog.Error(err)
			}
		}(r)
	}
	return
}
