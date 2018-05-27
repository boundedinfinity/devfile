package cmd

import (
	"fmt"
	"os"

	"github.com/boundedinfinity/devfile/cmd/brewfile"
	"github.com/boundedinfinity/devfile/cmd/bumpversion"
	"github.com/boundedinfinity/devfile/cmd/dockercompose"
	"github.com/boundedinfinity/devfile/cmd/readme"
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
	if cmd, err := brewfile.GetCommand(rootLogger); err != nil {
		rootLogger.Fatal(err)
		return
	} else {
		RootCmd.AddCommand(cmd)
	}

	if cmd, err := bumpversion.GetCommand(rootLogger); err != nil {
		rootLogger.Fatal(err)
		return
	} else {
		RootCmd.AddCommand(cmd)
	}

	if cmd, err := dockercompose.GetCommand(rootLogger); err != nil {
		rootLogger.Fatal(err)
		return
	} else {
		RootCmd.AddCommand(cmd)
	}

	if cmd, err := readme.GetCommand(rootLogger); err != nil {
		rootLogger.Fatal(err)
		return
	} else {
		RootCmd.AddCommand(cmd)
	}
}
