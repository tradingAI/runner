package runner

import (
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/runner/plugins"
)

type Container struct {
	Name    string
	ID      string
	ShortID string
	Job     *pb.Job
	Plugin  plugins.Plugin
}
