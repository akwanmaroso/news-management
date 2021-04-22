package app

import (
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/akwanmaroso/news/domain/repository"
)

type newsApps struct {
	newsApp repository.NewsRepository
}

type NewsAppInterface interface {
	SaveNews(*entity.News) (*entity.News, map[string]string)
	GetAllNews() ([]*entity.News, error)
	GetNews(uint64) (*entity.News, error)
	GetByTopic(string) ([]*entity.News, error)
	GetByStatus(string) ([]*entity.News, error)
	UpdateNews(*entity.News) (*entity.News, map[string]string)
	DeleteNews(uint64) error
}

var _ NewsAppInterface = &newsApps{}

func (n *newsApps) SaveNews(news *entity.News) (*entity.News, map[string]string) {
	return n.newsApp.SaveNews(news)
}

func (n *newsApps) GetAllNews() ([]*entity.News, error) {
	return n.newsApp.GetAllNews()
}

func (n *newsApps) GetNews(newsId uint64) (*entity.News, error) {
	return n.newsApp.GetNews(newsId)
}

func (n *newsApps) GetByTopic(topic string) ([]*entity.News, error) {
	return n.newsApp.GetByTopic(topic)
}

func (n *newsApps) GetByStatus(status string) ([]*entity.News, error) {
	return n.newsApp.GetByStatus(status)
}

func (n *newsApps) UpdateNews(news *entity.News) (*entity.News, map[string]string) {
	return n.newsApp.UpdateNews(news)
}

func (n *newsApps) DeleteNews(newsId uint64) error {
	return n.newsApp.DeleteNews(newsId)
}
