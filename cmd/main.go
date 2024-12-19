package main

import (
	"log"
	"mlm/config"
	database "mlm/db"
	"mlm/routes"

	"github.com/gin-gonic/gin"
)

// func createUser() {

// 	// Create a new user
// 	newUser := models.User{
// 		Email:      "example@gmail.com",
// 		Phone:      "+123456789",
// 		Country:    "USA",
// 		IsAdmin:    false,
// 		Username:   "johndoe",
// 		FullName:   "John Doe",
// 		IsActive:   true,
// 		IsStaff:    false,
// 		DateJoined: time.Now(),
// 	}

// Save the user to the database
// 	result := database.DB.Create(&newUser)
// 	if result.Error != nil {
// 		log.Fatalf("Failed to create user: %v", result.Error)
// 	}

// 	log.Println("User created successfully!")
// }

func main() {

	// Connect to the database
	database.ConnectDatabase()

	// Initialize the Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(router)
	routes.RegisterAuthRoutes(router)
	routes.RegisterMerchantRoutes(router)

	// Load app configuration
	config.InitConfig()

	// Start the server
	port := config.AppConfig.Port
	log.Printf("Server is running on port %s", port)
	router.Run(":" + port)
}
