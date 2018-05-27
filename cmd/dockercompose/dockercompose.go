package dockercompose

import (
	"log"

	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/dockercompose"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dockercompose",
	Short: "docker-compose.yml management",
	Long: `docker-compose.yml management
Simple management of the contents of the docker-compose.yml package management system.
You can read more about the project here: https://docs.docker.com/compose/compose-file.
	`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create docker-compose.yml",
	Long:  `create docker-compose.yml`,
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
	createConfig.ConfigurePath("docker-compose.yml")

	rootCmd.AddCommand(createCmd)

	return rootCmd, nil
}

func createRun(cmd *cobra.Command, args []string) {
	service := dockercompose.NewDockerComposeService(createConfig)

	if err := service.Create(); err != nil {

	}
}
