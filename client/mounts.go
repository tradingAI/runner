package client

import (
	"path"

	"docker.io/go-docker/api/types/mount"
	"github.com/tradingAI/runner/plugins"
)

func (c *Client) getTbaseMounts(id string) (mounts []mount.Mount) {
	mounts = []mount.Mount{
		c.getTbaseShellMount(id),
		c.getTbaseModelMount(id),
		c.getTbaseProgressBarMount(id),
		c.getTbaseTensorboardMount(id),
		c.getTbaseInferMount(id),
		c.getTbaseEvalMount(id),
	}
	return mounts
}

func (c *Client) getTbaseShellMount(id string) (m mount.Mount) {
	mustCreateDir(c.Conf.JobShellDir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: c.Conf.JobShellDir,
		Target: TARGET_SHELL_DIR,
	}
}

func (c *Client) getTbaseModelMount(id string) (m mount.Mount) {
	jobDir := path.Join(c.Conf.ModelDir, id)
	mustCreateDir(jobDir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: jobDir,
		Target: plugins.DEFAULT_TBASE_MODEL_DIR,
	}
}

func (c *Client) getTbaseProgressBarMount(id string) (m mount.Mount) {
	mustCreateDir(c.Conf.ProgressBarDir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: c.Conf.ProgressBarDir,
		Target: plugins.DEFAULT_TBASE_PROGRESS_BAR_DIR,
	}
}

func (c *Client) getTbaseTensorboardMount(id string) (m mount.Mount) {
	jobDir := path.Join(c.Conf.TensorboardDir, id)
	mustCreateDir(jobDir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: jobDir,
		Target: plugins.DEFAULT_TBASE_TENSORBOARD_DIR,
	}
}

func (c *Client) getTbaseInferMount(id string) (m mount.Mount) {
	mustCreateDir(c.Conf.InferDir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: c.Conf.InferDir,
		Target: plugins.DEFAULT_TBASE_INFER_DIR,
	}
}

func (c *Client) getTbaseEvalMount(id string) (m mount.Mount) {
	mustCreateDir(c.Conf.EvalDir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: c.Conf.EvalDir,
		Target: plugins.DEFAULT_TBASE_EVAL_DIR,
	}
}
