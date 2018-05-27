package file

import (
	"io/ioutil"

	"github.com/boundedinfinity/devfile/config"
	"github.com/boundedinfinity/devfile/config/manager"
)

type FileService struct {
	cm *manager.ConfigurationManager
}

func NewFileService(cm *manager.ConfigurationManager) *FileService {
	return &FileService{
		cm: cm,
	}
}

func (this *FileService) CreateProjectFile(templatePath, filename string) error {
	content, err := config.Packr().MustBytes(templatePath)

	if err != nil {
		return err
	}

	p, err := this.cm.GetFilePath(filename)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(p, content, config.GetFileMode()); err != nil {
		return err
	}

	return nil
}
