// MIGHT CHANGE THIS LATER
package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	ID uint
	Password string `gorm:"unique;not null"`
	Username string `gorm:"unique;not null"`
	Images   []Image `gorm:"foreignKey:UserID;references:ID"`
}

type Image struct {
	gorm.Model
	UserID uint 
	Url string `gorm:"unique;not null"`
}