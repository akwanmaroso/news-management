package entity

import (
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Status int

type News struct {
	gorm.Model
	Topic   string `gorm:"not null;unique" json:"topic"`
	Content string `gorm:"not null" json:"content"`
	Status  string `json:"status"`
	Tags    []*Tag `gorm:"many2many:news_tags" json:"tags"`
}

func (news *News) BeforeSave() {
	news.Topic = html.EscapeString(strings.TrimSpace(news.Topic))
	news.Content = html.EscapeString(strings.TrimSpace(news.Content))
	news.Status = html.EscapeString(strings.TrimSpace(news.Status))
}

func (news *News) Prepare() {
	news.Topic = html.EscapeString(strings.TrimSpace(news.Topic))
	news.Content = html.EscapeString(strings.TrimSpace(news.Content))
	news.Status = html.EscapeString(strings.TrimSpace(news.Status))
	news.CreatedAt = time.Now()
	news.UpdatedAt = time.Now()
}

func (news *News) Validate() map[string]string {
	var errorMessage = make(map[string]string)
	if news.Topic == "" {
		errorMessage["topic_required"] = "Topic can't be null"
	}
	if news.Content == "" {
		errorMessage["content_required"] = "Content can't be null"
	}
	return errorMessage
}
