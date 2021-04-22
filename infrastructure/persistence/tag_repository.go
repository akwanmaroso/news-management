package persistence

import (
	"errors"
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/akwanmaroso/news/domain/repository"
	"github.com/jinzhu/gorm"
	"strings"
)

type TagRepo struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepo {
	return &TagRepo{db}
}

var _ repository.TagRepository = &TagRepo{}

func (t *TagRepo) SaveTag(tag *entity.Tag) (*entity.Tag, map[string]string) {
	dbErr := map[string]string{}
	err := t.db.Debug().Create(&tag).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_name"] = "tag already taken"
		}
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return tag, nil
}

func (t *TagRepo) GetTag(tagID uint64) (*entity.Tag, error) {
	var tag entity.Tag
	err := t.db.Debug().Where("id = ?", tagID).Take(&tag).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("record tag not found")
	}
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (t *TagRepo) GetAllTag() ([]*entity.Tag, error) {
	var tag []*entity.Tag
	err := t.db.Debug().Limit(100).Preload("News").Order("created_at desc").Find(&tag).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("record tag not found")
	}
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (t *TagRepo) UpdateTag(tag *entity.Tag) (*entity.Tag, map[string]string) {
	dbErr := map[string]string{}
	err := t.db.Debug().Save(&tag).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_name"] = "topic already taken"
			return nil, dbErr
		}
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return tag, nil
}

func (t *TagRepo) DeleteTag(tagID uint64) error {
	var tag entity.Tag

	err := t.db.Debug().Where("id = ?", tagID).Delete(&tag).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("record news not found")
	}
	if err != nil {
		return err
	}
	return nil
}

func (t *TagRepo) FindTagByName(name string) (*entity.Tag, error) {
	var tag entity.Tag

	err := t.db.Debug().Where("name = ?", name).Take(&tag).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("record news not found")
	}
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (t *TagRepo) GetNewsByTag(topic string) ([]*entity.Tag, error) {
	var tags []*entity.Tag

	err := t.db.Debug().Where("name LIKE ?", "%"+topic+"%").Preload("News").Take(&tags).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("record news not found")
	}
	if err != nil {
		return nil, err
	}
	return tags, nil
}
