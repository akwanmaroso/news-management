package persistence

import (
	"fmt"
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/akwanmaroso/news/domain/repository"
	"github.com/jinzhu/gorm"
	"log"
)

type Repositories struct {
	News repository.NewsRepository
	Tag  repository.TagRepository
	db   *gorm.DB
}

func NewRepositories(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	db, err := gorm.Open(dbDriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &Repositories{
		News: NewNewsRepository(db),
		Tag:  NewTagRepository(db),
		db:   db,
	}, nil
}
func (repository *Repositories) Close() error {
	return repository.db.Close()
}

func (repository *Repositories) AutoMigrate() error {
	err := repository.db.Debug().AutoMigrate(&entity.News{}, &entity.Tag{}).Error
	if err != nil {
		log.Fatal(err)
	}

	return err
}
