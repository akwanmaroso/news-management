package entity

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	NewsID uint64 `json:"news_id"`
	Name   string `json:"name"`
}
