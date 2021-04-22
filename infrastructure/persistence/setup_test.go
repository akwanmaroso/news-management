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
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	return LocalDatabase()
	//}
}

func LocalDatabase() (*gorm.DB, error) {
	DBDriver := os.Getenv("TEST_DB_DRIVER")
	DBHost := os.Getenv("TEST_DB_HOST")
	DBPassword := os.Getenv("TEST_DB_PASSWORD")
	DBUser := os.Getenv("TEST_DB_USER")
	DBName := os.Getenv("TEST_DB_NAME")
	DBPort := os.Getenv("TEST_DB_PORT")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
	conn, err := gorm.Open(DBDriver, DBURL)
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
