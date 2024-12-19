package routes

import (
	"mlm/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMerchantRoutes(router *gin.Engine) {
	merchantRoutes := router.Group("/merchants")
	{
		merchantRoutes.POST("/", controllers.CreateMerchant)      // Create a new merchant
		merchantRoutes.GET("/", controllers.GetMerchants)         // List all merchants
		merchantRoutes.GET("/:id", controllers.GetMerchant)       // Get a merchant by ID
		merchantRoutes.PUT("/:id", controllers.UpdateMerchant)    // Update a merchant
		merchantRoutes.DELETE("/:id", controllers.DeleteMerchant) // Delete a merchant
	}
}
