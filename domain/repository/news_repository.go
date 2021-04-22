package repository

import "github.com/akwanmaroso/news/domain/entity"

type NewsRepository interface {
	SaveNews(*entity.News) (*entity.News, map[string]string)
	GetNews(uint64) (*entity.News, error)
	GetAllNews() ([]*entity.News, error)
	GetByStatus(string) ([]*entity.News, error)
	UpdateNews(*entity.News) (*entity.News, map[string]string)
	DeleteNews(uint64) error
}
