package dockercompose

import (
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/file"
)

type DockerComposeService struct {
	cm *manager.ConfigurationManager
	fs *file.FileService
}

func NewDockerComposeService(cm *manager.ConfigurationManager) *DockerComposeService {
	return &DockerComposeService{
		cm: cm,
		fs: file.NewFileService(cm),
	}
}

func (this *DockerComposeService) Create() error {
	return this.fs.CreateProjectFile("docker-compose.yml", "docker-compose.yml")
}
