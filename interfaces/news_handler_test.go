package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestNews_SaveNews_Success(t *testing.T) {
	newsApp.SaveNewsFn = func(news *entity.News) (*entity.News, map[string]string) {
		return &entity.News{
			Topic:   "First Topic",
			Content: "First Content",
			Status:  "First Status",
		}, nil
	}

	r := gin.Default()
	r.POST("/api/news", newNewsApp.SaveNews)
	fakeInputJSON := `{
		"topic": "First Topic",
		"content": "First Content",
		"status": "First Status"
	}`

	req, err := http.NewRequest(http.MethodPost, "/api/news", bytes.NewBufferString(fakeInputJSON))
	if err != nil {
		t.Errorf("Cannot make request: %v\n", err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	type DataEmbed struct {
		Topic   string `json:"topic"`
		Content string `json:"content"`
		Status  string `json:"status"`
	}

	type NewsTagPayload struct {
		Data    DataEmbed `json:"data"`
		Success bool      `json:"success"`
		Status  uint64    `json:"status"`
	}

	var newsTagPayload NewsTagPayload
	err = json.Unmarshal(rr.Body.Bytes(), &newsTagPayload)
	if err != nil {
		t.Errorf("Error unmarshal body: %v\n", err)
	}

	assert.Equal(t, rr.Code, 201)
	assert.EqualValues(t, newsTagPayload.Data.Topic, "First Topic")
	assert.EqualValues(t, newsTagPayload.Data.Content, "First Content")
	assert.EqualValues(t, newsTagPayload.Data.Status, "First Status")
}

func TestNews_UpdateNews_Success(t *testing.T) {
	newsApp.GetNewsFn = func(newsID uint64) (*entity.News, error) {
		return &entity.News{
			Topic:   "First Topic",
			Content: "First Content",
			Status:  "First Status",
		}, nil
	}

	newsApp.UpdateNewsFn = func(news *entity.News) (*entity.News, map[string]string) {
		return &entity.News{
			Topic:   "First Topic Updated",
			Content: "First Content Updated",
			Status:  "First Status Updated",
		}, nil
	}

	newsID := strconv.Itoa(1)
	r := gin.Default()
	r.PUT("/api/news/:news_id", newNewsApp.UpdateNews)
	fakeInputJSON := `{
		"topic": "First Topic Updated",
		"content": "First Content Updated",
		"status": "First Status Updated"
	}`

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/news/%s", newsID), bytes.NewBufferString(fakeInputJSON))
	if err != nil {
		t.Errorf("Cannot make request: %v\n", err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	type DataEmbed struct {
		Topic   string `json:"topic"`
		Content string `json:"content"`
		Status  string `json:"status"`
	}

	type NewsTagPayload struct {
		Data    DataEmbed `json:"data"`
		Success bool      `json:"success"`
		Status  uint64    `json:"status"`
	}

	var newsTagPayload NewsTagPayload
	err = json.Unmarshal(rr.Body.Bytes(), &newsTagPayload)
	if err != nil {
		t.Errorf("Error unmarshal body: %v\n", err)
	}

	assert.Equal(t, rr.Code, 200)
	assert.EqualValues(t, newsTagPayload.Data.Topic, "First Topic Updated")
	assert.EqualValues(t, newsTagPayload.Data.Content, "First Content Updated")
	assert.EqualValues(t, newsTagPayload.Data.Status, "First Status Updated")
}

func TestNews_GetAllNews_Success(t *testing.T) {
	newsApp.GetAllNewsFn = func() ([]*entity.News, error) {
		return []*entity.News{
			{
				Topic:   "First Topic",
				Content: "First Content",
				Status:  "First Status",
			}, {
				Topic:   "Second Topic",
				Content: "Second Content",
				Status:  "Second Status",
			},
		}, nil
	}

	r := gin.Default()
	r.GET("/api/news", newNewsApp.GetAllNews)

	req, err := http.NewRequest(http.MethodGet, "/api/news", nil)
	if err != nil {
		t.Errorf("Cannot make request: %v\n", err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	type DataEmbed struct {
		Topic   string `json:"topic"`
		Content string `json:"content"`
		Status  string `json:"status"`
	}

	type NewsTagPayload struct {
		Data    []DataEmbed `json:"data"`
		Success bool        `json:"success"`
		Status  uint64      `json:"status"`
	}

	var newsTagPayload NewsTagPayload
	err = json.Unmarshal(rr.Body.Bytes(), &newsTagPayload)
	if err != nil {
		t.Errorf("Error unmarshal body: %v\n", err)
	}

	assert.Equal(t, rr.Code, 200)
	assert.EqualValues(t, len(newsTagPayload.Data), 2)
}

func TestNews_GetNews_Success(t *testing.T) {
	newsApp.GetNewsFn = func(newsID uint64) (*entity.News, error) {
		return &entity.News{
			Topic:   "First Topic",
			Content: "First Content",
			Status:  "First Status",
		}, nil
	}

	r := gin.Default()
	newsID := strconv.Itoa(1)
	r.GET("/api/news/:news_id", newNewsApp.GetNews)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/news/%s", newsID), nil)
	if err != nil {
		t.Errorf("Cannot make request: %v\n", err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	type DataEmbed struct {
		Topic   string `json:"topic"`
		Content string `json:"content"`
		Status  string `json:"status"`
	}

	type NewsTagPayload struct {
		Data    DataEmbed `json:"data"`
		Success bool      `json:"success"`
		Status  uint64    `json:"status"`
	}

	var newsTagPayload NewsTagPayload
	err = json.Unmarshal(rr.Body.Bytes(), &newsTagPayload)
	if err != nil {
		t.Errorf("Error unmarshal body: %v\n", err)
	}

	assert.Equal(t, rr.Code, 200)
	assert.EqualValues(t, newsTagPayload.Data.Topic, "First Topic")
	assert.EqualValues(t, newsTagPayload.Data.Content, "First Content")
	assert.EqualValues(t, newsTagPayload.Data.Status, "First Status")
}
