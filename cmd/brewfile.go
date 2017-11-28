package cmd

import (
    "fmt"
    "strings"
    "github.com/spf13/cobra"
    "github.com/pkg/errors"
    "github.com/boundedinfinity/devfile/file/brewfile"
    "github.com/boundedinfinity/devfile/config/manager"
)

var brewfileCmd = &cobra.Command{
    Use:   "brewfile",
    Short: "Brewfile management",
    Long: `Brewfile management
Simple management of the contents of the Brewfile package management system.
You can read more about the project here: https://github.com/Homebrew/homebrew-bundle.
	`,
}

var brewfileReadCmd = &cobra.Command{
    Use:   "read",
    Short: "read Brewfile",
    Long:  `read Brewfile`,
    Run:   brewfileReadRun,
}

var brewfileWriteCmd = &cobra.Command{
    Use:   fmt.Sprintf("write <%s> <name> [<comment>]", ts("|")),
    Short: "write or update Brewfile",
    Long:  `write or update Brewfile`,
    Run:   brewfileWriteRun,
    Args:  brewfileWriteArgs,
}

func ts(sep string) string {
    ss := make([]string, 0)
    for _, s := range brewfile.BrewfileLineTypes {
        ss = append(ss, string(s))
    }
    return strings.Join(ss, sep)
}

var brewfileReadConfiguration *manager.ConfigurationManager
var brewfileWriteConfiguration *manager.ConfigurationManager

func init() {
    RootCmd.AddCommand(brewfileCmd)
    brewfileCmd.AddCommand(brewfileReadCmd)
    brewfileCmd.AddCommand(brewfileWriteCmd)

    rcm, err := manager.NewConfigurationManager(
        manager.Logger(rootLogger),
        manager.FlagSet(brewfileReadCmd.PersistentFlags()),
    )

    if err != nil {
        rootLogger.Fatal(err)
        return
    }

    wcm, err := manager.NewConfigurationManager(
        manager.Logger(rootLogger),
        manager.FlagSet(brewfileWriteCmd.PersistentFlags()),
    )

    if err != nil {
        rootLogger.Fatal(err)
        return
    }

    brewfileReadConfiguration = rcm
    brewfileReadConfiguration.ConfigureFormat()
    brewfileReadConfiguration.ConfigureDebug()
    brewfileReadConfiguration.ConfigurePath("Brewfile")

    brewfileWriteConfiguration = wcm
    brewfileWriteConfiguration.ConfigureDebug()
    brewfileWriteConfiguration.ConfigurePath("Brewfile")
}

func brewfileReadRun(cmd *cobra.Command, args []string) {
    if err := brewfileReadConfiguration.ValidateFormat(); err != nil {
        rootLogger.Fatal(err)
        return
    }

    p, err := brewfile.NewBrewFileProcessor(
        brewfile.Logger(rootLogger),
        brewfile.Path(brewfileReadConfiguration.GetPath()),
        brewfile.OutputFormat(brewfileReadConfiguration.GetFormat()),
        brewfile.Debug(brewfileReadConfiguration.GetDebug()),
    )

    if err != nil {
        rootLogger.Fatal(err)
        return
    }

    if err := p.Read(); err != nil {
        rootLogger.Fatal(err)
    }
}

func brewfileWriteRun(cmd *cobra.Command, args []string) {
    p, err := brewfile.NewBrewFileProcessor(
        brewfile.Logger(rootLogger),
        brewfile.Path(brewfileWriteConfiguration.GetPath()),
        brewfile.Debug(brewfileWriteConfiguration.GetDebug()),
    )

    if err != nil {
        rootLogger.Fatal(err)
        return
    }

    l := &brewfile.BrewFileLine{
        Name: brewfile.String2BrewfileLineType(args[0]),
        Value: args[1],
    }

    if len(args) > 2 {
        l.Comment = args[2]
    }

    if err := p.Write(*l); err != nil {
        rootLogger.Fatal(err)
    }
}

func brewfileWriteArgs(cmd *cobra.Command, args []string) error {
    if len(args) < 2 {
        return errors.Errorf("missing args")
    }

    t := brewfile.String2BrewfileLineType(args[0])

    if t == brewfile.BrewfileLineType_Unknown {
        return errors.Errorf("invalid line type: %s.  must be one of: [%s]", ts(","))
    }

    return nil
}
