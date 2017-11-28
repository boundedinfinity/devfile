package cmd

import (
	"github.com/spf13/cobra"
	"github.com/boundedinfinity/devfile/file/brewfile"
	"github.com/boundedinfinity/devfile/config/manager"
)

var brewfileCmd = &cobra.Command{
	Use:   "brewfile",
	Short: "Brewfile management",
	Long:  `Brewfile management
Simple management of the contents of the Brewfile package management system.
You can read more about the project here: https://github.com/Homebrew/homebrew-bundle.
	`,
	Run:   brewfileRun,
}

var brewfileConfiguration *manager.ConfigurationManager

func init() {
	RootCmd.AddCommand(brewfileCmd)

	cm, err := manager.NewConfigurationManager(
		manager.Logger(rootLogger),
		manager.FlagSet(brewfileCmd.PersistentFlags()),
	)

	if err != nil {
		rootLogger.Fatal(err)
		return
	}

	brewfileConfiguration = cm
	brewfileConfiguration.ConfigureFormat()
	brewfileConfiguration.ConfigureDebug()
	brewfileConfiguration.ConfigureClean()
	brewfileConfiguration.ConfigurePath("Brewfile")
}

func brewfileRun(cmd *cobra.Command, args []string) {
    if err := brewfileConfiguration.ValidateFormat(); err != nil {
        rootLogger.Fatal(err)
        return
    }

	p, err := brewfile.NewBrewFileProcessor(
		brewfile.Logger(rootLogger),
		brewfile.Path(brewfileConfiguration.GetPath()),
		brewfile.OutputFormat(brewfileConfiguration.GetFormat()),
		brewfile.Debug(brewfileConfiguration.GetDebug()),
		brewfile.Clean(brewfileConfiguration.GetClean()),
	)

	if err != nil {
		rootLogger.Fatal(err)
		return
	}

	if err := p.Execute(); err != nil {
		rootLogger.Fatal(err)
	}
}

