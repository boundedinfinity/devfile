package manager

import (
	"fmt"
	"os"
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
		ShortFlag:        pathDescriptorHelper.ShortFlag,
		LongFlag:         pathDescriptorHelper.LongFlag,
		ShortDescription: fmt.Sprintf(pathDescriptorHelper.ShortDescription, filename),
		LongDescription:  fmt.Sprintf(pathDescriptorHelper.LongDescription, filename),
		Default:          filepath.Join(currentDir, filename),
	})

	return viper.GetString(pathDescriptorHelper.LongFlag)
}

func (this *ConfigurationManager) GetPath() string {
	return this.v.GetString(pathDescriptorHelper.LongFlag)
}

func (this *ConfigurationManager) GetFilePath(filename string) (string, error) {
	info, err := os.Stat(this.GetPath())

	if err != nil {
		return "", err
	}

	var p string

	if info.IsDir() {
		p = filepath.Join(this.GetPath(), filename)
	} else {
		p = this.GetPath()
	}

	return p, nil
}
