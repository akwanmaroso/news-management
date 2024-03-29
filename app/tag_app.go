package app

import (
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/akwanmaroso/news/domain/repository"
)

type tagApps struct {
	tagApp repository.TagRepository
}

type TagAppInterface interface {
	SaveTag(*entity.Tag) (*entity.Tag, map[string]string)
	GetTag(uint64) (*entity.Tag, error)
	GetAllTag() ([]*entity.Tag, error)
	UpdateTag(*entity.Tag) (*entity.Tag, map[string]string)
	FindTagByName(name string) (*entity.Tag, error)
	GetNewsByTag(topic string) ([]*entity.Tag, error)
	DeleteTag(uint64) error
}

var _ TagAppInterface = &tagApps{}

func (t *tagApps) SaveTag(tag *entity.Tag) (*entity.Tag, map[string]string) {
	return t.tagApp.SaveTag(tag)
}

func (t *tagApps) GetTag(tagID uint64) (*entity.Tag, error) {
	return t.tagApp.GetTag(tagID)
}

func (t *tagApps) GetAllTag() ([]*entity.Tag, error) {
	return t.tagApp.GetAllTag()
}

func (t *tagApps) FindTagByName(name string) (*entity.Tag, error) {
	return t.tagApp.FindTagByName(name)
}

func (t *tagApps) UpdateTag(tag *entity.Tag) (*entity.Tag, map[string]string) {
	return t.tagApp.UpdateTag(tag)
}

func (t *tagApps) GetNewsByTag(topic string) ([]*entity.Tag, error) {
	return t.tagApp.GetNewsByTag(topic)
}

func (t *tagApps) DeleteTag(tagID uint64) error {
	return t.tagApp.DeleteTag(tagID)
}
