package models

import (
	"time"
)

type Merchant struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`                     // Foreign key referencing User
	User          User      `gorm:"constraint:OnDelete:CASCADE;"` // Reference to the User model
	Country       string    `gorm:"not null"`                     // Country where the merchant is located
	DisplayName   string    `gorm:"not null"`                     // Merchant's display name
	City          string    `gorm:"not null"`                     // City where the merchant is located
	CreatedAt     time.Time `gorm:"autoCreateTime"`               // Timestamp when the merchant was created
	Status        string    `gorm:"default:'inactive'"`           // Status (e.g., active, inactive)
	MinimumAmount float64   `gorm:"not null;default:0"`           // Minimum transaction amount
	MaximumAmount float64   `gorm:"not null;default:1000000"`     // Maximum transaction amount
}
