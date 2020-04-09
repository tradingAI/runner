package runner

import (
	"path"

	"docker.io/go-docker/api/types/mount"
	"github.com/tradingAI/runner/plugins"
)

func (r *Runner) getTbaseMounts(id string) (mounts []mount.Mount) {
	mounts = []mount.Mount{
		r.getTbaseShellMount(id),
		r.getTbaseModelMount(id),
		r.getTbaseProgressBarMount(id),
		r.getTbaseTensorboardMount(id),
		r.getTbaseInferMount(id),
		r.getTbaseEvalMount(id),
	}
	return mounts
}

func (r *Runner) getTbaseShellMount(id string) (m mount.Mount) {
	mustCreateDir(plugins.JOB_SHELL_DIR)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: plugins.JOB_SHELL_DIR,
		Target: plugins.JOB_SHELL_DIR,
	}
}

func (r *Runner) getTbaseModelMount(id string) (m mount.Mount) {
	dir := path.Join(plugins.MODEL_DIR, id)
	mustCreateDir(dir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: dir,
		Target: plugins.MODEL_DIR,
	}
}

func (r *Runner) getTbaseProgressBarMount(id string) (m mount.Mount) {
	mustCreateDir(plugins.PROGRESS_BAR_DIR)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: plugins.PROGRESS_BAR_DIR,
		Target: plugins.PROGRESS_BAR_DIR,
	}
}

func (r *Runner) getTbaseTensorboardMount(id string) (m mount.Mount) {
	dir := path.Join(r.Conf.TensorboardDir, id)
	mustCreateDir(dir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: dir,
		Target: plugins.TENSORBOARD_DIR,
	}
}

func (r *Runner) getTbaseInferMount(id string) (m mount.Mount) {
	mustCreateDir(plugins.INFER_DIR)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: plugins.INFER_DIR,
		Target: plugins.INFER_DIR,
	}
}

func (r *Runner) getTbaseEvalMount(id string) (m mount.Mount) {
	mustCreateDir(r.Conf.EvalDir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: plugins.EVAL_DIR,
		Target: plugins.EVAL_DIR,
	}
}
