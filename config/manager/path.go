package manager

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/spf13/viper"
)

var (
	pathDescriptorHelper = stringConfigDescriptor{
		ShortFlag:        "p",
		LongFlag:         "path",
		ShortDescription: "path for %s",
		LongDescription:  "path for %s",
		Default:          "",
	}
)

func (this *ConfigurationManager) ConfigurePath(filename string) string {
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	this.configureStringFlag(stringConfigDescriptor{
		ShortFlag: pathDescriptorHelper.ShortFlag,
		LongFlag: pathDescriptorHelper.LongFlag,
		ShortDescription: fmt.Sprintf(pathDescriptorHelper.ShortDescription, filename),
		LongDescription: fmt.Sprintf(pathDescriptorHelper.LongDescription, filename),
		Default: filepath.Join(currentDir, filename),
	})

	return viper.GetString(pathDescriptorHelper.LongFlag)
}

func (this *ConfigurationManager) GetPath() string {
	return this.v.GetString(pathDescriptorHelper.LongFlag)
}
