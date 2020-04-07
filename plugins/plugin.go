package plugins

import (
	"github.com/minio/minio-go/v6"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

type Plugin struct{}

func (p *Plugin) GenerateCmds(input *pb.JobInput, minioClient *minio.Client) (cmds []string, err error) {
	return
}
