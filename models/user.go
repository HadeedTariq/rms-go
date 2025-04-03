package models

import (
	"gorm.io/gorm"
)

// UserRole defines the possible roles for a user
type UserRole string

// Constants for the UserRole enum
const (
	Customer UserRole = "customer"
	Manager  UserRole = "manager"
	Admin    UserRole = "admin"
)

type User struct {
	gorm.Model
	Username string   `json:"username" gorm:"unique;not null"`
	Email    string   `json:"email" gorm:"unique;not null"`
	Password string   `json:"password" gorm:"not null"`
	FullName string   `json:"full_name" gorm:"not null"`
	Phone    string   `json:"phone"`
	Role     UserRole `json:"role" gorm:"default:'customer'"`
}
