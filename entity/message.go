package entity

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID  int
	Content string `gorm:"size:10000"`
}
