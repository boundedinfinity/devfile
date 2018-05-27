package brewfile

import (
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/common"
)

type BrewfileService struct {
	cm *manager.ConfigurationManager
}

func NewBrewfileService(cm *manager.ConfigurationManager) *BrewfileService {
	return &BrewfileService{
		cm: cm,
	}
}

func (this *BrewfileService) Create() error {
	return common.CreateFile(this.cm, "brewfile/Brewfile", "Brewfile")
}
