package brewfile

import (
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/file"
)

type BrewfileService struct {
	cm *manager.ConfigurationManager
	fs *file.FileService
}

func NewBrewfileService(cm *manager.ConfigurationManager) *BrewfileService {
	return &BrewfileService{
		cm: cm,
		fs: file.NewFileService(cm),
	}
}

func (this *BrewfileService) Create() error {
	return this.fs.CreateProjectFile("Brewfile", "Brewfile")
}
