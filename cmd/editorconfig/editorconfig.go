package editorconfig

import (
	"log"

	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/brewfile"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "editorconfig",
	Short: ".editorconfig management",
	Long: `.editorconfig management
Simple management of the contents of a .editorconfig file.
You can read more about the markdown syntax here: http://editorconfig.org.
	`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create .editorconfig",
	Long:  `create .editorconfig`,
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
	createConfig.ConfigurePath(".editorconfig")

	rootCmd.AddCommand(createCmd)

	return rootCmd, nil
}

func createRun(cmd *cobra.Command, args []string) {
	service := brewfile.NewBrewfileService(createConfig)

	if err := service.Create(); err != nil {

	}
}
