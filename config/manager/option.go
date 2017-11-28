package manager

import (
	"log"
	"github.com/spf13/viper"
	flag "github.com/spf13/pflag"
)

type ConfigurationManagerOption func(*ConfigurationManager) error

func Logger(input *log.Logger) ConfigurationManagerOption {
	return func(this *ConfigurationManager) error {
		this.l = input
		return nil
	}
}

func Viper(input *viper.Viper) ConfigurationManagerOption {
	return func(this *ConfigurationManager) error {
		this.v = input
		return nil
	}
}

func Debug(input bool) ConfigurationManagerOption {
	return func(this *ConfigurationManager) error {
		this.debug = input
		return nil
	}
}

func FlagSet(input *flag.FlagSet) ConfigurationManagerOption {
	return func(this *ConfigurationManager) error {
		this.fs = input
		return nil
	}
}