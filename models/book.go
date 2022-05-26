package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ISBN   string `gorm:"unique"`
	Title  string `gorm:"unique"`
	Author string
	Genre  string
}
