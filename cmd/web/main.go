package main

import (
	"context"
	"fmt"
	"time"

	"gokit-demo/ent/migrate"

	"github.com/weiloon1234/gokit"
	"github.com/weiloon1234/gokit/database"
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
			TranslationPaths:   []string{"lang"},
		},
	}

	gokit.Init(config)

	// Close everything that need to be closed
	defer gokit.Close()

	// Retrieve the ent client from gokit
	dbClient := database.GetDBClient()

	// Run the schema migration with project schema
	migrationCtx, cancelMigration := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelMigration()
	if err := dbClient.Schema.Create(
		migrationCtx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		panic(fmt.Errorf("failed to run schema migration: %v", err))
	}

	// Confirm that migration ran
	fmt.Println("ELLO")

	// Initialize router and start server
	r := gokit.InitRouter(config)
	err := r.Run(":" + config.AppConfig.AppPort)
	if err != nil {
		panic(err)
	}
}
