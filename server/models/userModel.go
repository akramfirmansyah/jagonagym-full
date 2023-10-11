package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:256" json:"username"`
	Password string `json:"password"`
}
