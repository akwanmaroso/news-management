package interfaces

import (
	"fmt"
	"github.com/akwanmaroso/news/app"
	"github.com/akwanmaroso/news/domain/entity"
	. "github.com/akwanmaroso/news/infrastructure/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Tag struct {
	tagApp  app.TagAppInterface
	newsApp app.NewsAppInterface
}

func NewTag(tagApp app.TagAppInterface, newsApp app.NewsAppInterface) *Tag {
	return &Tag{tagApp: tagApp, newsApp: newsApp}
}

func (t *Tag) SaveTag(c *gin.Context) {
	var tag entity.Tag
	err := c.BindJSON(&tag)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorResponse(http.StatusUnprocessableEntity, "JSON not valid"))
		return
	}

	emptyTag := entity.Tag{}
	emptyTag.Name = tag.Name
	validateTag := emptyTag.Validate()
	if len(validateTag) > 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse(http.StatusBadRequest, validateTag))
		return
	}

	newTag := entity.Tag{}
	newTag.Name = tag.Name
	savedTag, saveErr := t.tagApp.SaveTag(&newTag)
	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, saveErr))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, savedTag))
}

func (t *Tag) GetAllTag(c *gin.Context) {
	tags, err := t.tagApp.GetAllTag()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, tags))
}

func (t *Tag) GetTag(c *gin.Context) {
	tagID, err := strconv.ParseUint(c.Param("tag_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	tag, err := t.tagApp.GetTag(tagID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, tag))
}

func (t *Tag) UpdateTag(c *gin.Context) {
	tagID, err := strconv.ParseUint(c.Param("tag_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	var tag entity.Tag
	err = c.BindJSON(&tag)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorResponse(http.StatusUnprocessableEntity, "JSON not valid"))
		return
	}

	emptyTag := entity.Tag{}
	emptyTag.Name = tag.Name
	validateTag := emptyTag.Validate()
	if len(validateTag) > 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse(http.StatusBadRequest, validateTag))
		return
	}

	tagToUpdate, err := t.tagApp.GetTag(tagID)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse(http.StatusNotFound, fmt.Sprintf("tag with id: %d is not exists", tagID)))
		return
	}

	tagToUpdate.Name = tag.Name
	updatedTag, updateError := t.tagApp.UpdateTag(tagToUpdate)
	if updateError != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, updateError))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, updatedTag))
}

func (t *Tag) DeleteTag(c *gin.Context) {
	tagID, err := strconv.ParseUint(c.Param("tag_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err = t.tagApp.DeleteTag(tagID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(http.StatusOK, fmt.Sprintf("News With Tag: %d is successfully deleted", tagID)))
}
