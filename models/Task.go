package models

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model

	Title       string `gorm:"type:varchar(100);not null;unique_index" json:"title"`
	Description string `json:"description"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      uint   `json:"user_id"`
}
