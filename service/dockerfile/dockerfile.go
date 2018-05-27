package dockerfile

import (
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/common"
)

type DockerfileService struct {
	cm *manager.ConfigurationManager
}

func NewDockerfileService(cm *manager.ConfigurationManager) *DockerfileService {
	return &DockerfileService{
		cm: cm,
	}
}

func (this *DockerfileService) Create() error {
	return common.CreateFile(this.cm, "dockerfile/Dockerfile", "Dockerfile")
}
