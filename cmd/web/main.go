package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit"
	"github.com/weiloon1234/gokit/validator"
	"regexp"
	"time"
)

func main() {
	// gokit configuration
	config := gokit.Config{
		DBConfig: &gokit.DBConfig{
			Host:         "localhost",
			Port:         "3306",
			Username:     "root",
			DatabaseName: "go_test1",
		},
		RedisConfig: &gokit.RedisConfig{
			Host: "127.0.0.1",
			Port: "6379",
		},
		LocalizationConfig: gokit.LocaleConfig{
			DefaultLanguage:    "en",
			SupportedLanguages: []string{"en", "zh"},
			TranslationPaths:   []string{"locales"},
		},
		Timezone: "Asia/Kuala_Lumpur",
		Features: gokit.NewDefaultFeatures(),
	}

	// Initialize gokit for application runtime
	gokit.Init(config)

	// Add custom validation rule: alphanumeric
	validator.RegisterRule("alphanumeric", func(value interface{}, param string, _ *gin.Context) bool {
		alphaNumRegex := `^[a-zA-Z0-9]+$`
		re := regexp.MustCompile(alphaNumRegex)
		return re.MatchString(fmt.Sprintf("%v", value))
	})

	ginConfig := gokit.GinConfig{
		LogFile:                       "tmp/server.log",
		FatalErrorInformTelegram:      true,
		TelegramBotToken:              "123",
		TelegramChatId:                "123",
		TelegramMessageThrottleMinute: 10 * time.Minute,
	}

	r := gokit.InitRouter(ginConfig)
	r.Run(":8080")
}
