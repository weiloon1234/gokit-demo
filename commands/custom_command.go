package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CustomCommand = &cobra.Command{
	Use:   "custom_command",
	Short: "Run a custom command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Running custom command!")
		return nil
	},
}
