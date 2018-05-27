package readme

import (
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/file"
)

type ReadmeService struct {
	cm *manager.ConfigurationManager
	fs *file.FileService
}

func NewReadmeService(cm *manager.ConfigurationManager) *ReadmeService {
	return &ReadmeService{
		cm: cm,
		fs: file.NewFileService(cm),
	}
}

func (this *ReadmeService) Create() error {
	return this.fs.CreateProjectFile("readme.md", "readme.md")
}
