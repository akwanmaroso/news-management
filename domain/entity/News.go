package entity

import (
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type Status int

const (
	Draft Status = iota + 1
	Publish
	Deleted
)

type News struct {
	gorm.Model
	Topic   string `json:"topic"`
	Content string `json:"content"`
	Status  Status `json:""`
}


