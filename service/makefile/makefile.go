package makefile

import (
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/file"
)

type MakefileService struct {
	cm *manager.ConfigurationManager
	fs *file.FileService
}

func NewMakefileService(cm *manager.ConfigurationManager) *MakefileService {
	return &MakefileService{
		cm: cm,
		fs: file.NewFileService(cm),
	}
}

func (this *MakefileService) Create() error {
	return this.fs.CreateProjectFile("Makefile", "Makefile")
}
