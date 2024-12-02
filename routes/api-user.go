package routes

import "github.com/gin-gonic/gin"

func ApiUserRoute(r *gin.Engine) {
	g := r.Group("/api/user")
	{
		g.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "user login",
			})
		})
	}
}
