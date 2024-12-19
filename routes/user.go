package routes

import (
	"mlm/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes defines the routes for the User model
func RegisterUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)         // GET /users - List all users
		userRoutes.GET("/:id", controllers.GetUser)       // GET /users/:id - Get a user by ID
		userRoutes.POST("/", controllers.CreateUser)      // POST /users - Create a new user
		userRoutes.PUT("/:id", controllers.UpdateUser)    // PUT /users/:id - Update a user
		userRoutes.DELETE("/:id", controllers.DeleteUser) // DELETE /users/:id - Delete a user
	}
}
