package runner

import (
	"io/ioutil"
	"path"
	"strconv"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/runner/plugins"
)

func (r *Runner) UpdateEvalOutput(job *pb.Job, p plugins.Plugin) (err error) {
	if job.Type == pb.JobType_EVALUATION || job.Type == pb.JobType_TRAIN {
		id := strconv.FormatUint(job.Id, 10)
		filePath := path.Join(r.Conf.EvalDir, id)
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			glog.Error(err)
			return err
		}
		evalOutput, err := p.ParseEval(string(content), job.Id, job.Id)
		job.Output = evalOutput
		return err
	}
	return
}

func (r *Runner) UpdateInferOutput(job *pb.Job, p plugins.Plugin) (err error) {
	if job.Type == pb.JobType_INFER {
		id := strconv.FormatUint(job.Id, 10)
		filePath := path.Join(r.Conf.InferDir, id)
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			glog.Error(err)
			return err
		}
		inferOutput, err := p.ParseInfer(string(content), job.GetInput().GetInferInput().Date, job.Id, job.Id)
		job.Output = inferOutput
		return err
	}
	return
}

func (r *Runner) UpdateBar(job *pb.Job, p plugins.Plugin) (err error) {
	if job.Type != pb.JobType_TRAIN {
		job.CurrentStep = uint32(0)
		job.TotalSteps = uint32(1)
		return nil
	}
	id := strconv.FormatUint(job.Id, 10)
	filePath := path.Join(r.Conf.ProgressBarDir, id)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		glog.Error(err)
		return
	}
	currentStep, totalSteps, err := p.ParseBar(string(content))
	if err != nil {
		glog.Error(err)
		return
	}
	job.CurrentStep = currentStep
	job.TotalSteps = totalSteps
	return
}

func (r *Runner) UpdateOutput(job *pb.Job) (err error) {
	p := plugins.New(job)
	err = r.UpdateBar(job, p)
	if err != nil {
		glog.Error(err)
		return
	}
	if job.Type == pb.JobType_EVALUATION ||  job.Type == pb.JobType_TRAIN {
		err = r.UpdateEvalOutput(job, p)
		if err != nil {
			glog.Error(err)
			return
		}
	}
	if job.Type == pb.JobType_INFER {
		err = r.UpdateInferOutput(job, p)
		if err != nil {
			glog.Error(err)
			return
		}
	}
	return
}

// upload files: model, tensorboard
// Update job.output clean dirs and files: model dir, job log, tensorboard dir, evals, infers, progress_bars
func (r *Runner) FinishedJob(job *pb.Job) (err error) {
	defer r.Clean(strconv.FormatUint(job.Id, 10))
	if job.Type == pb.JobType_TRAIN {
		err = r.uploadTrainModel(job)
		if err != nil {
			glog.Error(err)
			return
		}
		err = r.uploadTensorboard(job)
		if err != nil {
			glog.Error(err)
			return
		}
	}
	err = r.UpdateOutput(job)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}
