package runner

import (
	"path"

)

func (r *Runner) DeleteModel(id string) (err error) {
	dir := path.Join(r.Conf.ModelDir, id)
	return DeleteDirectory(dir)
}

func (r *Runner) DeleteTensorboard(id string) (err error) {
	dir := path.Join(r.Conf.TensorboardDir, id)
	return DeleteDirectory(dir)
}

func (r *Runner) DeleteLog(id string) (err error) {
	filePath := path.Join(r.Conf.JobLogDir, id)
	return DeleteFile(filePath)
}

func (r *Runner) DeleteEval(id string) (err error) {
	filePath := path.Join(r.Conf.EvalDir, id)
	return DeleteFile(filePath)
}

func (r *Runner) DeleteInfer(id string) (err error) {
	filePath := path.Join(r.Conf.InferDir, id)
	return DeleteFile(filePath)
}

func (r *Runner) DeleteShell(id string) (err error) {
	filePath := path.Join(r.Conf.JobShellDir, id)
	return DeleteFile(filePath)
}

func (r *Runner) DeleteProgressBar(id string) (err error) {
	filePath := path.Join(r.Conf.ProgressBarDir, id)
	return DeleteFile(filePath)
}

func (r *Runner) Clean(id string) {
	r.DeleteModel(id)
	r.DeleteTensorboard(id)
	r.DeleteLog(id)
	r.DeleteEval(id)
	r.DeleteInfer(id)
	r.DeleteShell(id)
	r.DeleteProgressBar(id)
}
