package models

import "time"

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Email      string `gorm:"unique;not null"`
	Phone      string
	Country    string
	IsAdmin    bool      `gorm:"default:false"`
	Username   string    `gorm:"unique;not null"`
	FullName   string    `gorm:"not null"`
	IsActive   bool      `gorm:"default:true"`
	IsStaff    bool      `gorm:"default:false"`
	Password   string    `gorm:"not null;default:testing321"` // Hashed password
	DateJoined time.Time `gorm:"autoCreateTime"`
}
