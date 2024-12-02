package main

import (
	"github.com/weiloon1234/gokit"
	"gokit-demo/routes"
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

	r := gokit.InitRouter(config)
	routes.Root(r)
	routes.ApiUserRoute(r)
	routes.ApiAdminRoute(r)

	err := r.Run(":" + config.AppConfig.AppPort)
	if err != nil {
		return
	}
}
