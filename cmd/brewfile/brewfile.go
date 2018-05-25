package brewfile

import (
	"log"

	"github.com/boundedinfinity/devfile/config/manager"
	"github.com/boundedinfinity/devfile/service/brewfile"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "brewfile",
	Short: "Brewfile management",
	Long: `Brewfile management
Simple management of the contents of the Brewfile package management system.
You can read more about the project here: https://github.com/Homebrew/homebrew-bundle.
	`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create Brewfile",
	Long:  `create Brewfile`,
	Run:   brewfileCreateRun,
}

func GetCommand(logger *log.Logger) (*cobra.Command, error) {

	createConfig, err := manager.NewConfigurationManager(
		manager.Logger(logger),
		manager.FlagSet(createCmd.PersistentFlags()),
	)

	if err != nil {
		return nil, err
	}

	createConfig.ConfigurePath("Brewfile")

	rootCmd.AddCommand(createCmd)

	return rootCmd, nil
}

func brewfileCreateRun(cmd *cobra.Command, args []string) {
	service := brewfile.NewBrewfileService()

	if err := service.Create(); err != nil {

	}
}

// var brewfileReadCmd = &cobra.Command{
//     Use:   "read",
//     Short: "read Brewfile",
//     Long:  `read Brewfile`,
//     Run:   brewfileReadRun,
// }

// var brewfileWriteCmd = &cobra.Command{
//     Use:   fmt.Sprintf("write <%s> <name> [<comment>]", ts("|")),
//     Short: "write or update Brewfile",
//     Long:  `write or update Brewfile`,
//     Run:   brewfileWriteRun,
//     Args:  brewfileWriteArgs,
// }

// func ts(sep string) string {
//     ss := make([]string, 0)
//     for _, s := range brewfile.BrewfileLineTypes {
//         ss = append(ss, string(s))
//     }
//     return strings.Join(ss, sep)
// }

var brewfileCreateConfiguration *manager.ConfigurationManager

// var brewfileReadConfiguration *manager.ConfigurationManager
// var brewfileWriteConfiguration *manager.ConfigurationManager

// func brewfileReadRun(cmd *cobra.Command, args []string) {
// 	if err := brewfileReadConfiguration.ValidateFormat(); err != nil {
// 		rootLogger.Fatal(err)
// 		return
// 	}

// 	p, err := brewfile.NewBrewFileProcessor(
// 		brewfile.Logger(rootLogger),
// 		brewfile.Path(brewfileReadConfiguration.GetPath()),
// 		brewfile.OutputFormat(brewfileReadConfiguration.GetFormat()),
// 		brewfile.Debug(brewfileReadConfiguration.GetDebug()),
// 	)

// 	if err != nil {
// 		rootLogger.Fatal(err)
// 		return
// 	}

// 	if err := p.Read(); err != nil {
// 		rootLogger.Fatal(err)
// 	}
// }

// func brewfileWriteRun(cmd *cobra.Command, args []string) {
// 	p, err := brewfile.NewBrewFileProcessor(
// 		brewfile.Logger(rootLogger),
// 		brewfile.Path(brewfileWriteConfiguration.GetPath()),
// 		brewfile.Debug(brewfileWriteConfiguration.GetDebug()),
// 	)

// 	if err != nil {
// 		rootLogger.Fatal(err)
// 		return
// 	}

// 	l := &brewfile.BrewFileLine{
// 		Name:  brewfile.String2BrewfileLineType(args[0]),
// 		Value: args[1],
// 	}

// 	if len(args) > 2 {
// 		l.Comment = args[2]
// 	}

// 	if err := p.Write(*l); err != nil {
// 		rootLogger.Fatal(err)
// 	}
// }

// func brewfileWriteArgs(cmd *cobra.Command, args []string) error {
// 	if len(args) < 2 {
// 		return errors.Errorf("missing args")
// 	}

// 	t := brewfile.String2BrewfileLineType(args[0])

// 	if t == brewfile.BrewfileLineType_Unknown {
// 		return errors.Errorf("invalid line type: %s.  must be one of: [%s]", ts(","))
// 	}

// 	return nil
// }
