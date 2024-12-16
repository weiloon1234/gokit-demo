package main

import (
	"gokit-demo/commands"

	"github.com/spf13/cobra"
	"github.com/weiloon1234/gokit"
	"github.com/weiloon1234/gokit/cli"
)

func main() {
	config := gokit.Config{
		AppConfig: gokit.AppConfig{
			AppName: "GoKit-Boilerplate",
			AppEnv:  "local",
			AppPort: "8081",
		},
		FeatureConfig: gokit.FeatureConfig{
			EnableDB:     true,
			EnableLocale: true,
			EnableRedis:  true,
		},
		DBConfig: gokit.DBConfig{
			Driver:       "mysql",
			Host:         "localhost",
			Port:         "3306",
			Username:     "root",
			DatabaseName: "go_test1",
		},
		RedisConfig: gokit.RedisConfig{
			Host: "127.0.0.1",
			Port: "6379",
		},
		LocalizationConfig: gokit.LocaleConfig{
			DefaultLanguage:    "en",
			SupportedLanguages: []string{"en", "zh"},
			TranslationPaths:   []string{"locales"},
		},
	}

	// Initialize gokit for CLI use
	gokit.Init(config)

	// Register a custom seeder
	// gokitCommand.RegisterSeeder("country_seeder", seeds.CountrySeeder)

	// Add custom command here
	cli.Init([]*cobra.Command{
		commands.CustomCommand,
	})

	// Execute CLI with custom commands
	cli.ExecuteCLI()
}
