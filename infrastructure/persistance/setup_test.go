package persistance

import (
	"fmt"
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DBConn() (*gorm.DB, error) {
	//if _, err := os.Stat("../../.env"); !os.IsNotExist(err) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	return LocalDatabase()
	//}
}

func LocalDatabase() (*gorm.DB, error) {
	DbDriver := os.Getenv("TEST_DB_DRIVER")
	DbHost := os.Getenv("TEST_DB_HOST")
	DbPassword := os.Getenv("TEST_DB_PASSWORD")
	DbUser := os.Getenv("TEST_DB_USER")
	DbName := os.Getenv("TEST_DB_NAME")
	DbPort := os.Getenv("TEST_DB_PORT")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	conn, err := gorm.Open(DbDriver, DBURL)
	if err != nil {
		return nil, err
	}

	err = conn.DropTableIfExists(&entity.News{}, &entity.Tag{}).Error
	if err != nil {
		return nil, err
	}
	err = conn.Debug().AutoMigrate(entity.News{}, entity.Tag{}).Error
	if err != nil {
		return nil, err
	}

	return conn, nil
}
