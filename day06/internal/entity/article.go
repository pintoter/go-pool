package entity

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"type:text;notnull"`
	Content string `gorm:"type:text;notnull"`
}
