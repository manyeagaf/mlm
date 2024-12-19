package routes

import (
	"mlm/controllers"
	"mlm/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.Signup) // Signup route
	router.POST("/login", controllers.Login)   // Login route

	protected := router.Group("/protected")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID := c.GetString("user_id")
			c.JSON(200, gin.H{"message": "Welcome!", "user_id": userID})
		})
	}
}
