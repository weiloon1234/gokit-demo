package main

import (
	"context"
	"fmt"
	"gokit-demo/ent"
	"gokit-demo/ent/hooks"
	"gokit-demo/ent/migrate"
	"time"

	"entgo.io/ent/dialect"
	entSQL "entgo.io/ent/dialect/sql"
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

	if config.FeatureConfig.EnableDB {
		// Wrap the database connection with Ent's SQL driver
		entDriver := entSQL.OpenDB(dialect.MySQL, database.GetSQLDB())
		entClient := ent.NewClient(ent.Driver(entDriver))

		// Run the schema migration with context timeout
		migrationCtx, cancelMigration := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancelMigration()
		if err := entClient.Schema.Create(
			migrationCtx,
			migrate.WithDropColumn(true),
			migrate.WithDropIndex(true),
		); err != nil {
			panic(fmt.Errorf("failed to create schema resources: %w", err))
		}

		database.SetEntClient(entClient)

		/** FOR GOKIT AUTO REGISTER ENTITY HOOKS HERE, DON'T EDIT THIS LINE **/
		entClient.Use(hooks.SoftDeleteHook())
	}

	// Initialize router and start server
	r := gokit.InitRouter(config)
	err := r.Run(":" + config.AppConfig.AppPort)
	if err != nil {
		panic(err)
	}
}
