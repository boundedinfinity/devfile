package dockerignore

import (
	"log"

	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/brewfile"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dockerignore",
	Short: ".dockerignore management",
	Long: `.dockerignore management
Simple management of the contents of a .dockerignore file.
You can read more about the markdown syntax here: https://docs.docker.com/engine/reference/builder/#dockerignore-file.
	`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create readme.md",
	Long:  `create readme.md`,
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
	createConfig.ConfigurePath(".dockerignore")

	rootCmd.AddCommand(createCmd)

	return rootCmd, nil
}

func createRun(cmd *cobra.Command, args []string) {
	service := brewfile.NewBrewfileService(createConfig)

	if err := service.Create(); err != nil {

	}
}
