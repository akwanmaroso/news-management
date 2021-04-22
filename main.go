package main

import (
	"github.com/akwanmaroso/news/infrastructure/persistence"
	"github.com/akwanmaroso/news/interfaces"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("not env gotten")
	}
}

func main() {
	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	services, err := persistence.NewRepositories(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	err = services.AutoMigrate()
	if err != nil {
		panic(err)
	}

	news := interfaces.NewNews(services.News, services.Tag)
	tag := interfaces.NewTag(services.Tag, services.News)

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/news", news.GetAllNews)
	api.GET("/news/:news_id", news.GetNews)
	api.POST("/news", news.SaveNews)
	api.PUT("/news/:news_id", news.UpdateNews)
	api.DELETE("/news/:news_id", news.DeleteNews)

	api.GET("/tags", tag.GetAllTag)
	api.GET("/tags/:tag_id", tag.GetTag)
	api.PUT("/tags/:tag_id", tag.UpdateTag)
	api.DELETE("/tags/:tag_id", tag.DeleteTag)
	api.POST("/tags", tag.SaveTag)

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
