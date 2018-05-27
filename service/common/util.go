package common

import (
	"io/ioutil"

	"github.com/boundedinfinity/devfile/config"
	"github.com/boundedinfinity/devfile/config/manager"
)

func CreateFile(cm *manager.ConfigurationManager, templatePath, filename string) error {
	content, err := config.Packr().MustBytes(templatePath)

	if err != nil {
		return err
	}

	p, err := cm.GetFilePath(filename)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(p, content, config.GetFileMode()); err != nil {
		return err
	}

	return nil
}
