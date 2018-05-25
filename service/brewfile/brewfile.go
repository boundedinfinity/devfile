package brewfile

import (
	"io/ioutil"

	"github.com/boundedinfinity/devfile/config"
	"github.com/boundedinfinity/devfile/config/manager"
)

type BrewfileService struct {
	cm *manager.ConfigurationManager
}

func NewBrewfileService() *BrewfileService {
	return &BrewfileService{}
}

func (this *BrewfileService) Create() error {
	content, err := config.Packr().MustBytes("brewfile/Brewfile")

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(this.cm.GetPath(), content, 0755); err != nil {
		return err
	}

	return nil
}
