package routes

import "github.com/gin-gonic/gin"

func ApiAdminRoute(r *gin.Engine) {
	g := r.Group("/api/admin")
	{
		g.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "admin login",
			})
		})
	}
}
