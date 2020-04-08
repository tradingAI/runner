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
	dir := path.Join(plugins.MODEL_DIR, id)
	mustCreateDir(dir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: dir,
		Target: plugins.MODEL_DIR,
	}
}

func (c *Client) getTbaseProgressBarMount(id string) (m mount.Mount) {
	mustCreateDir(plugins.PROGRESS_BAR_DIR)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: plugins.PROGRESS_BAR_DIR,
		Target: plugins.PROGRESS_BAR_DIR,
	}
}

func (c *Client) getTbaseTensorboardMount(id string) (m mount.Mount) {
	dir := path.Join(c.Conf.TensorboardDir, id)
	mustCreateDir(dir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: dir,
		Target: plugins.TENSORBOARD_DIR,
	}
}

func (c *Client) getTbaseInferMount(id string) (m mount.Mount) {
	mustCreateDir(plugins.INFER_DIR)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: plugins.INFER_DIR,
		Target: plugins.INFER_DIR,
	}
}

func (c *Client) getTbaseEvalMount(id string) (m mount.Mount) {
	mustCreateDir(c.Conf.EvalDir)
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: plugins.EVAL_DIR,
		Target: plugins.EVAL_DIR,
	}
}
