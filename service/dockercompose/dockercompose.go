package dockercompose

import (
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/common"
)

type DockerComposeService struct {
	cm *manager.ConfigurationManager
}

func NewDockerComposeService(cm *manager.ConfigurationManager) *DockerComposeService {
	return &DockerComposeService{
		cm: cm,
	}
}

func (this *DockerComposeService) Create() error {
	return common.CreateFile(this.cm, "dockercompose/docker-compose.yml", "docker-compose.yml")
}
