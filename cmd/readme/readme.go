package readme

import (
	"log"

	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/readme"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "readme",
	Short: "readme.md management",
	Long: `readme.md management
Simple management of the contents of a readme.md file.
You can read more about the markdown syntax here: https://daringfireball.net/projects/markdown/syntax.
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
	createConfig.ConfigurePath("readme.md")

	rootCmd.AddCommand(createCmd)

	return rootCmd, nil
}

func createRun(cmd *cobra.Command, args []string) {
	service := readme.NewReadmeService(createConfig)

	if err := service.Create(); err != nil {

	}
}
