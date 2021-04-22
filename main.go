package main

import (
	"github.com/akwanmaroso/news/infrastructure/persistance"
	"github.com/akwanmaroso/news/interfaces"
	"github.com/gin-gonic/gin"
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
	services, err := persistance.NewRepositories(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName)
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

	api.GET("/tag", tag.GetAllTag)
	api.GET("/tag/:tag_id", tag.GetTag)
	api.PUT("/tag/:tag_id", tag.UpdateTag)
	api.DELETE("/tag/:tag_id", tag.DeleteTag)
	api.POST("/tag", tag.SaveTag)

	log.Fatal(r.Run(":8000"))
}
