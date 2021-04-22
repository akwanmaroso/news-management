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

func TestTag_SaveTag_Success(t *testing.T) {
	tagApp.SaveTagFn = func(tag *entity.Tag) (*entity.Tag, map[string]string) {
		return &entity.Tag{
			Name: "First Tag",
		}, nil
	}

	r := gin.Default()
	r.POST("/api/tags", newTagApp.SaveTag)
	fakeInputJSON := `{
		"name": "First Tag"
	}`

	req, err := http.NewRequest(http.MethodPost, "/api/tags", bytes.NewBufferString(fakeInputJSON))
	if err != nil {
		t.Errorf("Cannot make request: %v\n", err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	type DataEmbed struct {
		Name string `json:"name"`
	}

	type TagPayload struct {
		Data    DataEmbed `json:"data"`
		Success bool      `json:"success"`
		Status  uint64    `json:"status"`
	}

	var tagPayload TagPayload
	err = json.Unmarshal(rr.Body.Bytes(), &tagPayload)
	if err != nil {
		t.Errorf("Error unmarshal body: %v\n", err)
	}

	assert.Equal(t, rr.Code, 200)
	assert.EqualValues(t, tagPayload.Data.Name, "First Tag")
}

func TestTag_GetAllTag_Success(t *testing.T) {
	tagApp.GetAllTagFn = func() ([]*entity.Tag, error) {
		return []*entity.Tag{
			{
				Name: "First Tag",
			},
			{
				Name: "Second Tag",
			},
		}, nil
	}

	r := gin.Default()
	r.GET("/api/tags", newTagApp.GetAllTag)

	req, err := http.NewRequest(http.MethodGet, "/api/tags", nil)
	if err != nil {
		t.Errorf("Cannot make request: %v\n", err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	type DataEmbed struct {
		Name string `json:"name"`
	}

	type TagPayload struct {
		Data    []DataEmbed `json:"data"`
		Success bool        `json:"success"`
		Status  uint64      `json:"status"`
	}

	var tagPayload TagPayload
	err = json.Unmarshal(rr.Body.Bytes(), &tagPayload)
	if err != nil {
		t.Errorf("Error unmarshal body: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 200)
	assert.EqualValues(t, len(tagPayload.Data), 2)
}

func TestTag_GetTag_Success(t *testing.T) {
	tagApp.GetTagFn = func(tagID uint64) (*entity.Tag, error) {
		return &entity.Tag{
			Name: "First Tag",
		}, nil
	}

	r := gin.Default()
	tagID := strconv.Itoa(1)
	r.GET("/api/tags/:tag_id", newTagApp.GetTag)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/tags/%s", tagID), nil)
	if err != nil {
		t.Errorf("Cannot make request: %v\n", err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	type DataEmbed struct {
		Name string `json:"name"`
	}

	type TagPayload struct {
		Data    DataEmbed `json:"data"`
		Success bool      `json:"success"`
		Status  uint64    `json:"status"`
	}

	var tagPayload TagPayload
	err = json.Unmarshal(rr.Body.Bytes(), &tagPayload)
	if err != nil {
		t.Errorf("Error unmarshal body: %v\n", err)
	}

	assert.Equal(t, rr.Code, 200)
	assert.EqualValues(t, tagPayload.Data.Name, "First Tag")
}
