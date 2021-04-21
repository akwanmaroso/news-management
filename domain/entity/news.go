package entity

import (
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
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

func (news *News) BeforeSave() {
	news.Topic = html.EscapeString(strings.TrimSpace(news.Topic))
	news.Content = html.EscapeString(strings.TrimSpace(news.Content))
}

func (news *News) Prepare() {
	news.Topic = html.EscapeString(strings.TrimSpace(news.Topic))
	news.Content = html.EscapeString(strings.TrimSpace(news.Content))
	news.CreatedAt = time.Now()
	news.UpdatedAt = time.Now()
}

func (news *News) Validate() map[string]string {
	var errorMessage = make(map[string]string)

	if news.Topic == "" || news.Topic == "null" {
		errorMessage["topic_required"] = "Topic can't be null"
	}
	if news.Content == "" || news.Content == "null" {
		errorMessage["content_required"] = "Content can't be null"
	}
	return errorMessage
}
