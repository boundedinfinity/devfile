package cmd

import (
	"encoding/xml"
	"fmt"

	"github.com/boundedinfinity/devfile/config"
	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Long:  `Version information`,
	Run:   versionRun,
}

var versionConfiguration *manager.ConfigurationManager

type versionJson struct {
	XMLName xml.Name `json:"-" yaml:"-" xml:"devfile"`
	Version string   `json:"version" yaml:"version" xml:"version"`
}

func init() {
	RootCmd.AddCommand(versionCmd)

	cm, err := manager.NewConfigurationManager(
		manager.Logger(rootLogger),
		manager.FlagSet(versionCmd.PersistentFlags()),
	)

	if err != nil {
		rootLogger.Fatal(err)
		return
	}

	versionConfiguration = cm
	versionConfiguration.ConfigureDebug()
	versionConfiguration.ConfigureFormat()
}

func versionRun(cmd *cobra.Command, args []string) {
	output := versionJson{
		Version: config.Version,
	}

	if versionConfiguration.GetFormat() != manager.OutputFormat_None {
		if err := versionConfiguration.PrintFormatString(&output); err != nil {
			panic(err)
		}
	} else if versionConfiguration.GetQuiet() {
		fmt.Print(config.Version)
	} else {
		fmt.Printf("Version: %s\n", config.Version)
	}
}
