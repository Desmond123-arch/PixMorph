// MIGHT CHANGE THIS LATER
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint	`gorm:"primaryKey" form:"id"`
	Password string  `gorm:"not null" form:"password" binding:"required"`
	Username string  `gorm:"unique;not null" form:"username" binding:"required"`
	Images   []Image `gorm:"foreignKey:UserID;references:ID"`
}

type Image struct {
	gorm.Model
	ID uint `gorm:"primaryKey" form:"id"`
	UserID uint
	Url    string `gorm:"unique;not null"`
}