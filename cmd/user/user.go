package user

import (
	"log"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User file management",
	Long: `Manage files in the user's configuration directory.
	`,
}

func GetCommand(logger *log.Logger) (*cobra.Command, error) {
	return userCmd, nil
}
