package controllers

import (
	database "mlm/db"
	"mlm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers retrieves all users from the database
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUser retrieves a specific user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var user models.User

	// Bind JSON payload to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save to database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Check if the user exists
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind JSON payload to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save changes
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Check if the user exists
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user
	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
