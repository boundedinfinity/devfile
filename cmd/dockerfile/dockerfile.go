package dockerfile

import (
	"log"

	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/docker_compose"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Dockerfile management",
	Long: `Dockerfile management
Simple management of the contents of the Dockerfile package management system.
You can read more about the project here: https://docs.docker.com/engine/reference/builder/.
	`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create Dockerfile",
	Long:  `create Dockerfile`,
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
	createConfig.ConfigurePath("Dockerfile")

	rootCmd.AddCommand(createCmd)

	return rootCmd, nil
}

func createRun(cmd *cobra.Command, args []string) {
	service := docker_compose.NewDockerComposeService(createConfig)

	if err := service.Create(); err != nil {

	}
}
