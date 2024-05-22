package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `gorm:"not null" json:"title"`
	Body   string `gorm:"type:text" json:"body"`
	UserID uint   `gorm:"foreignkey:UserID" json:"userID"`
	User   User   `gorm:"foreignkey:UserID"`
}
