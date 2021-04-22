package entity

import (
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type Tag struct {
	gorm.Model
	Name string  `gorm:"not null;unique" json:"name"`
	News []*News `gorm:"many2many:news_tags" json:"news"`
}

func (tag *Tag) BeforeSave() {
	tag.Name = html.EscapeString(strings.TrimSpace(tag.Name))
}

func (tag *Tag) Validate() map[string]string {
	var errorMessage = make(map[string]string)

	if tag.Name == "" || tag.Name == "null" {
		errorMessage["name_required"] = "Name of tag can't be null"
	}

	return errorMessage
}
