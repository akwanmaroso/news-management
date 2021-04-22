package persistence

import (
	"errors"
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/akwanmaroso/news/domain/repository"
	"github.com/jinzhu/gorm"
	"strings"
)

type NewsRepo struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepo {
	return &NewsRepo{db}
}

var _ repository.NewsRepository = &NewsRepo{}

func (n *NewsRepo) SaveNews(news *entity.News) (*entity.News, map[string]string) {
	dbErr := map[string]string{}
	err := n.db.Debug().Create(&news).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_name"] = "news already taken"
		}
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return news, nil
}

func (n *NewsRepo) GetNews(newsID uint64) (*entity.News, error) {
	var news entity.News
	err := n.db.Debug().Where("id = ?", newsID).Preload("Tags").Find(&news).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("news not found")
	}
	if err != nil {
		return nil, err
	}

	if len(news.Tags) > 0 {
		for i := range news.Tags {
			err := n.db.Debug().Where("id = ?", news.Tags[i].ID).Find(&news.Tags[i]).Error
			if err != nil {
				return nil, err
			}
		}
	}

	return &news, nil
}

func (n *NewsRepo) GetAllNews() ([]*entity.News, error) {
	var news []*entity.News

	err := n.db.Debug().Order("created_at desc").Preload("Tags").Limit(100).Find(&news).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("record news not found")
	}
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n *NewsRepo) GetByTopic(topic string) ([]*entity.News, error) {
	var news []*entity.News
	err := n.db.Debug().Order("created_at desc").Preload("Tags").Limit(100).Find(&news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n *NewsRepo) GetByStatus(status string) ([]*entity.News, error) {
	var news []*entity.News
	err := n.db.Debug().Order("created_at desc").Preload("Tags").Where("status = ?", status).Limit(100).Find(&news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n *NewsRepo) UpdateNews(news *entity.News) (*entity.News, map[string]string) {
	dbErr := map[string]string{}
	err := n.db.Debug().Save(&news).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_name"] = "topic already taken"
			return nil, dbErr
		}
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return news, nil
}

func (n *NewsRepo) DeleteNews(newsID uint64) error {
	var news entity.News

	err := n.db.Debug().Where("id = ?", newsID).Delete(&news).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("record news not found")
	}
	if err != nil {
		return err
	}

	return nil
}
