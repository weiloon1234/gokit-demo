package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/weiloon1234/gokit"
	"github.com/weiloon1234/gokit/cli"
	"github.com/weiloon1234/gokit/cli/commands"
)

func main() {
	// Minimal gokit configuration for CLI
	config := gokit.Config{
		DBConfig: &gokit.DBConfig{
			Host:         "localhost",
			Port:         "3306",
			Username:     "root",
			DatabaseName: "go_test1",
		},
		LocalizationConfig: gokit.LocaleConfig{
			DefaultLanguage: "en",
		},
		Timezone: "UTC",
		Features: gokit.Features{EnableDB: true}, // Enable only DB feature
	}

	// Initialize gokit for CLI use
	gokit.Init(config)

	// Register a custom seeder
	commands.RegisterSeeder("example_seeder", func() {
		fmt.Println("Running example seeder...")
	})

	// Add a custom CLI command
	customCommand := &cobra.Command{
		Use:   "custom",
		Short: "Run a custom command",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Running custom command!")
			return nil
		},
	}

	cli.Init([]*cobra.Command{
		customCommand,
	})

	// Execute CLI with custom commands
	cli.ExecuteCLI()
}
