package models

import (
	"gorm.io/gorm"
)

type MenuCategory struct {
	gorm.Model
	Name        string     `gorm:"unique;not null"`
	Description string     `gorm:"type:text"`
	MenuItems   []MenuItem `gorm:"foreignKey:CategoryID"`
}

type MenuItem struct {
	gorm.Model
	Name        string  `gorm:"not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	ImageURL    string
	IsAvailable bool         `gorm:"default:true"`
	CategoryID  uint         `gorm:"not null"`
	Category    MenuCategory `gorm:"foreignKey:CategoryID"`
}
