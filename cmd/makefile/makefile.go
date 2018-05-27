package makefile

import (
	"log"

	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/brewfile"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "makefile",
	Short: "Makefile management",
	Long: `Makefile management
Simple management of the contents of a Makefile file.
You can read more about the markdown syntax here: https://www.gnu.org/software/make/manual/make.html.
	`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create Makefile",
	Long:  `create Makefile`,
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
	createConfig.ConfigurePath("Makefile")

	rootCmd.AddCommand(createCmd)

	return rootCmd, nil
}

func createRun(cmd *cobra.Command, args []string) {
	service := brewfile.NewBrewfileService(createConfig)

	if err := service.Create(); err != nil {

	}
}
