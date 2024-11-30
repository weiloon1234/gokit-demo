package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit"
	"github.com/weiloon1234/gokit/validator"
	"regexp"
)

func main() {
	// gokit configuration
	config := gokit.Config{
		DBConfig: gokit.DBConfig{
			DSN: "root@tcp(127.0.0.1:3306)/go_test1",
		},
		RedisConfig: gokit.RedisConfig{
			Addr: "127.0.0.1:6379",
		},
		LocalizationConfig: gokit.LocaleConfig{
			DefaultLanguage:    "en",
			SupportedLanguages: []string{"en", "zh"},
			TranslationPaths:   []string{"locales"},
		},
		Timezone: "Asia/Singapore",
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

	// Set up Gin router
	r := gin.Default()

	// Example route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Start Gin server
	r.Run(":8080")
}
