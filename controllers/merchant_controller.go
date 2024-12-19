package controllers

import (
	database "mlm/db"
	"mlm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMerchant creates a new merchant
func CreateMerchant(c *gin.Context) {
	var input struct {
		UserID        uint    `json:"user_id" binding:"required"`
		Country       string  `json:"country" binding:"required"`
		DisplayName   string  `json:"display_name" binding:"required"`
		City          string  `json:"city" binding:"required"`
		Status        string  `json:"status"`
		MinimumAmount float64 `json:"minimum_amount" binding:"required"`
		MaximumAmount float64 `json:"maximum_amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merchant := models.Merchant{
		UserID:        input.UserID,
		Country:       input.Country,
		DisplayName:   input.DisplayName,
		City:          input.City,
		Status:        input.Status,
		MinimumAmount: input.MinimumAmount,
		MaximumAmount: input.MaximumAmount,
	}

	if err := database.DB.Create(&merchant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create merchant"})
		return
	}

	c.JSON(http.StatusCreated, merchant)
}

// GetMerchants retrieves all merchants
func GetMerchants(c *gin.Context) {
	var merchants []models.Merchant
	if err := database.DB.Preload("User").Find(&merchants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch merchants"})
		return
	}

	c.JSON(http.StatusOK, merchants)
}

// GetMerchant retrieves a single merchant by ID
func GetMerchant(c *gin.Context) {
	id := c.Param("id")
	var merchant models.Merchant

	if err := database.DB.Preload("User").First(&merchant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

// UpdateMerchant updates a merchant
func UpdateMerchant(c *gin.Context) {
	id := c.Param("id")
	var merchant models.Merchant

	if err := database.DB.First(&merchant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	var input struct {
		Country       string  `json:"country"`
		DisplayName   string  `json:"display_name"`
		City          string  `json:"city"`
		Status        string  `json:"status"`
		MinimumAmount float64 `json:"minimum_amount"`
		MaximumAmount float64 `json:"maximum_amount"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merchant.Country = input.Country
	merchant.DisplayName = input.DisplayName
	merchant.City = input.City
	merchant.Status = input.Status
	merchant.MinimumAmount = input.MinimumAmount
	merchant.MaximumAmount = input.MaximumAmount

	if err := database.DB.Save(&merchant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update merchant"})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

// DeleteMerchant deletes a merchant
func DeleteMerchant(c *gin.Context) {
	id := c.Param("id")
	var merchant models.Merchant

	if err := database.DB.First(&merchant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	if err := database.DB.Delete(&merchant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete merchant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Merchant deleted successfully"})
}
