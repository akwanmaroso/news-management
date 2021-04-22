package interfaces

import (
	"fmt"
	"github.com/akwanmaroso/news/app"
	"github.com/akwanmaroso/news/domain/entity"
	. "github.com/akwanmaroso/news/infrastructure/responses"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type News struct {
	newsApp app.NewsAppInterface
	tagApp  app.TagAppInterface
}

func NewNews(newsApp app.NewsAppInterface, tagApp app.TagAppInterface) *News {
	return &News{newsApp, tagApp}
}

func (n *News) SaveNews(c *gin.Context) {
	type NewsTagPayload struct {
		Topic   string   `json:"topic"`
		Content string   `json:"content"`
		Status  string   `json:"status"`
		Tags    []string `json:"tags"`
	}

	var newsTagPayload NewsTagPayload
	err := c.BindJSON(&newsTagPayload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorResponse(http.StatusBadRequest, "Error bind json"))
		return
	}

	log.Println(newsTagPayload)

	var tags []*entity.Tag
	var foundTag *entity.Tag
	for _, v := range newsTagPayload.Tags {
		foundTag, err = n.tagApp.FindTagByName(v)
		if err != nil {
			log.Println("Create tag")
			foundTag = &entity.Tag{}
			foundTag.Name = v
			n.tagApp.SaveTag(foundTag)
		}
		tags = append(tags, foundTag)
	}

	statusType := map[string]string{
		"Draft":   "Draft",
		"Publish": "Publish",
		"Deleted": "Deleted",
	}

	emptyNews := entity.News{}
	emptyNews.Topic = newsTagPayload.Topic
	emptyNews.Content = newsTagPayload.Content
	emptyNews.Status = statusType[newsTagPayload.Status]
	emptyNews.Tags = tags
	validateNews := emptyNews.Validate()
	if len(validateNews) > 0 {
		c.JSON(http.StatusUnprocessableEntity, ErrorResponse(http.StatusBadRequest, validateNews))
		return
	}

	newNews := entity.News{}
	newNews.Topic = newsTagPayload.Topic
	newNews.Content = newsTagPayload.Content
	newNews.Status = statusType[newsTagPayload.Status]
	newNews.Tags = tags
	log.Println(emptyNews.Status)
	savedNews, saveErr := n.newsApp.SaveNews(&newNews)
	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, saveErr))
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse(http.StatusOK, savedNews))
}

func (n *News) GetAllNews(c *gin.Context) {
	if c.Query("topic") != "" {
		news, err := n.newsApp.GetByTopic(c.Query("topic"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, err.Error()))
			return
		}
		c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, news))
		return
	} else if c.Query("status") != "" {
		news, err := n.newsApp.GetByStatus(c.Query("status"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, err.Error()))
			return
		}
		c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, news))
		return
	}

	news, err := n.newsApp.GetAllNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, news))
}

func (n *News) GetNews(c *gin.Context) {
	newsID, err := strconv.ParseUint(c.Param("news_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	news, err := n.newsApp.GetNews(newsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, news))
}

func (n *News) UpdateNews(c *gin.Context) {
	newsID, err := strconv.ParseUint(c.Param("news_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	type NewsTagPayload struct {
		Topic   string   `json:"topic"`
		Content string   `json:"content"`
		Status  string   `json:"status"`
		Tag     []string `json:"tags"`
	}

	// var news entity.News
	var newsTagPayload NewsTagPayload
	err = c.BindJSON(&newsTagPayload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorResponse(http.StatusBadRequest, "Error bind json"))
		return
	}
	// emptyNews := entity.News{}
	// emptyNews.Topic = news.Topic
	// emptyNews.Content = news.Content
	// emptyNews.Status = news.Status
	// emptyNews.Tag = news.Tag
	// validateNews := emptyNews.Validate()
	// if len(validateNews) > 0 {
	//	c.JSON(http.StatusUnprocessableEntity, ErrorResponse(http.StatusBadRequest, validateNews))
	//	return
	// }

	log.Println(newsTagPayload)

	var tags []*entity.Tag
	var foundTag *entity.Tag
	for _, v := range newsTagPayload.Tag {
		foundTag, err = n.tagApp.FindTagByName(v)
		if err != nil {
			log.Println("Create tag")
			foundTag = &entity.Tag{}
			foundTag.Name = v
			n.tagApp.SaveTag(foundTag)
		}
		tags = append(tags, foundTag)
	}

	// emptyNews := entity.News{}
	// emptyNews.Topic = news.Topic
	// emptyNews.Content = news.Content
	// emptyNews.Status = news.Status
	// emptyNews.TagID = news.TagID
	// emptyNews.UpdatedAt = time.Now()
	// validateNews := emptyNews.Validate()
	// if len(validateNews) > 0 {
	//	c.JSON(http.StatusBadRequest, validateNews)
	//	return
	// }

	newsToUpdate, err := n.newsApp.GetNews(newsID)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse(http.StatusNotFound, fmt.Sprintf("news with id: %d is not exists", newsID)))
		return
	}

	newsToUpdate.Topic = newsTagPayload.Topic
	newsToUpdate.Content = newsTagPayload.Content
	newsToUpdate.Status = newsTagPayload.Status
	newsToUpdate.Tags = tags
	updatedNews, updateError := n.newsApp.UpdateNews(newsToUpdate)
	if updateError != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, updateError))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, updatedNews))
}

func (n *News) DeleteNews(c *gin.Context) {
	newsID, err := strconv.ParseUint(c.Param("news_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err = n.newsApp.GetNews(newsID)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse(http.StatusNotFound, fmt.Sprintf("news with id: %d is not exists", newsID)))
		return
	}

	err = n.newsApp.DeleteNews(newsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, fmt.Sprintf("News With ID: %d is successfully deleted", newsID)))
}
