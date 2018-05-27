package project

import (
	"log"

	"github.com/boundedinfinity/devfile/cmd/brewfile"
	"github.com/boundedinfinity/devfile/cmd/bumpversion"
	"github.com/boundedinfinity/devfile/cmd/dockercompose"
	"github.com/boundedinfinity/devfile/cmd/dockerfile"
	"github.com/boundedinfinity/devfile/cmd/editorconfig"
	"github.com/boundedinfinity/devfile/cmd/gitignore"
	"github.com/boundedinfinity/devfile/cmd/makefile"
	"github.com/boundedinfinity/devfile/cmd/readme"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "project",
	Short: "Project file management",
	Long: `Manage files in a project.
	`,
}

func GetCommand(logger *log.Logger) (*cobra.Command, error) {
	if cmd, err := brewfile.GetCommand(logger); err != nil {
		return nil, err
	} else {
		rootCmd.AddCommand(cmd)
	}

	if cmd, err := bumpversion.GetCommand(logger); err != nil {
		return nil, err
	} else {
		rootCmd.AddCommand(cmd)
	}

	if cmd, err := dockercompose.GetCommand(logger); err != nil {
		return nil, err
	} else {
		rootCmd.AddCommand(cmd)
	}

	if cmd, err := dockerfile.GetCommand(logger); err != nil {
		return nil, err
	} else {
		rootCmd.AddCommand(cmd)
	}

	if cmd, err := editorconfig.GetCommand(logger); err != nil {
		return nil, err
	} else {
		rootCmd.AddCommand(cmd)
	}

	if cmd, err := gitignore.GetCommand(logger); err != nil {
		return nil, err
	} else {
		rootCmd.AddCommand(cmd)
	}

	if cmd, err := makefile.GetCommand(logger); err != nil {
		return nil, err
	} else {
		rootCmd.AddCommand(cmd)
	}

	if cmd, err := readme.GetCommand(logger); err != nil {
		return nil, err
	} else {
		rootCmd.AddCommand(cmd)
	}

	return rootCmd, nil
}
