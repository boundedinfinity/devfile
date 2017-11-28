package manager

import (
	"log"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	flag "github.com/spf13/pflag"
	"github.com/boundedinfinity/devfile/config"

)

type ConfigurationManager struct {
	v     *viper.Viper
	l     *log.Logger
	fs    *flag.FlagSet
	debug bool
}

func NewConfigurationManager(options ...ConfigurationManagerOption) (*ConfigurationManager, error) {
	service := &ConfigurationManager{}

	for _, option := range options {
		if err := option(service); err != nil {
			return nil, err
		}
	}

	if service.fs == nil {
		return nil, errors.New("FlagSet cannot be nil")
	}

	if service.l == nil {
		service.l = config.CreateLogger()
	}

	if service.v == nil {
		service.v = viper.New()
	}

	return service, nil
}
