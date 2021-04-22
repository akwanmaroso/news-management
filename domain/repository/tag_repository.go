package repository

import "github.com/akwanmaroso/news/domain/entity"

type TagRepository interface {
	SaveTag(*entity.Tag) (*entity.Tag, map[string]string)
	GetTag(uint64) (*entity.Tag, error)
	GetAllTag() ([]*entity.Tag, error)
	FindTagByName(name string) (*entity.Tag, error)
	UpdateTag(*entity.Tag) (*entity.Tag, map[string]string)
	DeleteTag(uint64) error
}
