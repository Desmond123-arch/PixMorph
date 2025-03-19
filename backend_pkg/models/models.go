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

type resize struct{
	width int
	height int
}
type crop struct {
	width int
	height int
	x int
	y int
}
type filters struct{
	greyscale bool
	sepia bool
}
type Transform struct {
	resize resize
	crop crop
	filters filters
	rotate int
	format string
}