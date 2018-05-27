package dockerignore

import (
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/file"
)

type DockerIgnoreService struct {
	cm *manager.ConfigurationManager
	fs *file.FileService
}

func NewDockerIgnoreService(cm *manager.ConfigurationManager) *DockerIgnoreService {
	return &DockerIgnoreService{
		cm: cm,
		fs: file.NewFileService(cm),
	}
}

func (this *DockerIgnoreService) Create() error {
	return this.fs.CreateProjectFile(".dockerignore", ".dockerignore")
}
