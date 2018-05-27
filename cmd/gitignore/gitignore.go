package gitignore

import (
	"log"

	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/brewfile"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitignore",
	Short: ".gitignore management",
	Long: `.gitignore management
Simple management of the contents of a .gitignore file.
You can read more about the markdown syntax here: https://git-scm.com/docs/gitignore.
	`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create .gitignore",
	Long:  `create .gitignore`,
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
	createConfig.ConfigurePath(".gitignore")

	rootCmd.AddCommand(createCmd)

	return rootCmd, nil
}

func createRun(cmd *cobra.Command, args []string) {
	service := brewfile.NewBrewfileService(createConfig)

	if err := service.Create(); err != nil {

	}
}
