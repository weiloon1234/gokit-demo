package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/localization"
)

func Root(r *gin.Engine) {
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello2": localization.Translate(c, "Hello world with :attribute", map[string]string{
				"attribute": "You",
			}),
			"Hello3": localization.Translate(c, "ZAZA"),
		})
	})
}
