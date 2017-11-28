package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/boundedinfinity/devfile/config"
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
}