package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/localization"
	"gokit-demo/models"
)

func Root(r *gin.Engine) {
	r.GET("/hello", func(c *gin.Context) {
		// Declare userResponse as an interface{} to handle both nil and *User
		var userResponse interface{}
		user, err := model.GetUserByID(1)
		if err == nil {
			userResponse = user
		} else {
			userResponse = nil // Optional: nil represents no user found
		}

		c.JSON(200, gin.H{
			"Hello2": localization.Translate(c, "Hello world with :attribute", map[string]string{
				"attribute": "You",
			}),
			"Hello3": localization.Translate(c, "ZAZA"),
			"User":   userResponse,
		})
	})
}
