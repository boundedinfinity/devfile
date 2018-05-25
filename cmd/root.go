package cmd

import (
	"fmt"
	"os"

	"github.com/boundedinfinity/devfile/cmd/brewfile"
	"github.com/boundedinfinity/devfile/config"
	"github.com/spf13/cobra"
)

var rootLogger = config.CreateLogger()

var RootCmd = &cobra.Command{
	Use:   "devfile",
	Short: "Manage your development files",
	Long:  `Manage your development files`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	brewCmd, err := brewfile.GetCommand(rootLogger)

	if err != nil {
		rootLogger.Fatal(err)
		return
	}

	RootCmd.AddCommand(brewCmd)
}
