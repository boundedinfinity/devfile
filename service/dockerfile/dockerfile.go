package dockerfile

import (
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/file"
)

type DockerfileService struct {
	cm *manager.ConfigurationManager
	fs *file.FileService
}

func NewDockerfileService(cm *manager.ConfigurationManager) *DockerfileService {
	return &DockerfileService{
		cm: cm,
		fs: file.NewFileService(cm),
	}
}

func (this *DockerfileService) Create() error {
	return this.fs.CreateProjectFile("Dockerfile", "Dockerfile")
}
