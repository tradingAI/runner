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
		filePath := path.Join(r.Conf.EvalDir, id, "eval.txt")
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			glog.Error(err)
			return err
		}
		evalOutput, err := p.ParseEval(string(content))
		job.Output = evalOutput
		return err
	}
	return
}

func (r *Runner) UpdateInferOutput(job *pb.Job, p plugins.Plugin) (err error) {
	if job.Type == pb.JobType_INFER {
		id := strconv.FormatUint(job.Id, 10)
		filePath := path.Join(r.Conf.InferDir, id, "infer.txt")
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			glog.Error(err)
			return err
		}
		inferOutput, err := p.ParseInfer(string(content))
		job.Output = inferOutput
		return err
	}
	return
}

func (r *Runner) UpdateBar(job *pb.Job, p plugins.Plugin) (err error) {
	id := strconv.FormatUint(job.Id, 10)
	filePath := path.Join(r.Conf.JobLogDir, id)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		glog.Error(err)
		return
	}
	currentStep, totalSteps, err := p.ParseBar(string(content))
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
	err = r.UpdateEvalOutput(job, p)
	if err != nil {
		glog.Error(err)
		return
	}
	err = r.UpdateInferOutput(job, p)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

// upload files: model, tensorboard
// Update job.output clean dirs and files: model dir, job log, tensorboard dir, evals, infers, shells, progress_bars
func (r *Runner) UpdateJob(job *pb.Job) (err error) {
	if job.Type == pb.JobType_TRAIN {
		err = r.uploadTrainModel(job)
		if err != nil {
			glog.Error(err)
			job.Status = pb.JobStatus_FAILED
			return
		}
		err = r.uploadTensorboard(job)
		if err != nil {
			glog.Error(err)
			job.Status = pb.JobStatus_FAILED
			return
		}
	}
	err = r.UpdateOutput(job)
	if err != nil {
		glog.Error(err)
		job.Status = pb.JobStatus_FAILED
		return
	}
	r.Clean(strconv.FormatUint(job.Id, 10))
	job.Status = pb.JobStatus_SUCCESSED
	return
}
