package bumpversion

import (
	"log"

	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/brewfile"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bumpversion",
	Short: ".bumpversion.cfg management",
	Long: `.bumpversion.cfg management
Simple management of the contents of the .bumpversion.cfg version management system.
You can read more about the project here: https://github.com/peritus/bumpversion.
	`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create .bumpversion.cfg",
	Long:  `create .bumpversion.cfg`,
	Run:   createRun,
}
var createConfig *manager.ConfigurationManager

func GetCommand(logger *log.Logger) (*cobra.Command, error) {
	cm, err := manager.NewConfigurationManager(
		manager.Logger(logger),
		manager.FlagSet(createCmd.PersistentFlags()),
	)

	if err != nil {
		return nil, err
	}

	createConfig = cm
	createConfig.ConfigurePath(".bumpversion.cfg")

	rootCmd.AddCommand(createCmd)

	return rootCmd, nil
}

func createRun(cmd *cobra.Command, args []string) {
	service := brewfile.NewBrewfileService(createConfig)

	if err := service.Create(); err != nil {

	}
}
