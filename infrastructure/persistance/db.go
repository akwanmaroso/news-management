package persistance

import (
	"fmt"
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/akwanmaroso/news/domain/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Repositories struct {
	News repository.NewsRepository
	Tag  repository.TagRepository
	db   *gorm.DB
}

func NewRepositories(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(DbDriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &Repositories{
		News: NewDrugRepository(db),
		Tag:  NewTagRepository(db),
		db:   db,
	}, nil
}
func (repository *Repositories) Close() error {
	return repository.Close()
}

func (repository *Repositories) AutoMigrate() error {
	err := repository.db.Debug().AutoMigrate(&entity.News{}, &entity.Tag{}).Error
	if err != nil {
		log.Fatal(err)
	}

	return err
}
