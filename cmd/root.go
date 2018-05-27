package cmd

import (
	"fmt"
	"os"

	"github.com/boundedinfinity/devfile/cmd/project"
	"github.com/boundedinfinity/devfile/cmd/user"
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
	if cmd, err := project.GetCommand(rootLogger); err != nil {
		rootLogger.Fatal(err)
		return
	} else {
		RootCmd.AddCommand(cmd)
	}

	if cmd, err := user.GetCommand(rootLogger); err != nil {
		rootLogger.Fatal(err)
		return
	} else {
		RootCmd.AddCommand(cmd)
	}
}
