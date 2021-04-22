package persistance

import (
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewsRepo_SaveNews(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	news := entity.News{}
	news.Topic = "News Topic"
	news.Content = "News Content"
	news.Status = "News Status"
	news.Tags = []*entity.Tag{{
		Name: "News Tag",
	}}

	repo := NewNewsRepository(conn)

	n, saveErr := repo.SaveNews(&news)
	assert.Nil(t, saveErr)
	assert.EqualValues(t, n.Topic, "News Topic")
	assert.EqualValues(t, n.Content, "News Content")
	assert.EqualValues(t, n.Status, "News Status")
}

func TestNewsRepo_GetNews(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	fakeNews := &entity.News{
		Topic:   "News Topic",
		Content: "News Content",
		Status:  "News Status",
	}

	err = conn.Create(&fakeNews).Error
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	repo := NewNewsRepository(conn)

	n, getErr := repo.GetNews(uint64(fakeNews.ID))
	assert.Nil(t, getErr)
	assert.EqualValues(t, n.Topic, "News Topic")
	assert.EqualValues(t, n.Content, "News Content")
	assert.EqualValues(t, n.Status, "News Status")
}

func TestNewsRepo_GetAllNews(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	fakeNews := []*entity.News{
		{
			Topic:   "First News Topic",
			Content: "First News Content",
			Status:  "First News Status",
		}, {
			Topic:   "Second News Topic",
			Content: "Second News Content",
			Status:  "Second News Status",
		},
	}

	for _, v := range fakeNews {
		err = conn.Create(&v).Error
		if err != nil {
			t.Fatalf("want non error, got %v", err.Error())
		}
	}

	repo := NewNewsRepository(conn)

	n, getAllErr := repo.GetAllNews()
	assert.Nil(t, getAllErr)
	assert.EqualValues(t, len(n), 2)
}

func TestNewsRepo_UpdateNews(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	fakeNews := &entity.News{
		Topic:   "News Topic",
		Content: "News Content",
		Status:  "News Status",
	}

	err = conn.Create(&fakeNews).Error
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	fakeNews.Topic = "News Topic Updated"
	fakeNews.Content = "News Content Updated"
	fakeNews.Status = "News Status Updated"

	repo := NewNewsRepository(conn)
	n, updateErr := repo.UpdateNews(fakeNews)
	assert.Nil(t, updateErr)
	assert.EqualValues(t, n.Topic, "News Topic Updated")
	assert.EqualValues(t, n.Content, "News Content Updated")
	assert.EqualValues(t, n.Status, "News Status Updated")
}

func TestNewsRepo_DeleteNews(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	fakeNews := &entity.News{
		Topic:   "News Topic",
		Content: "News Content",
		Status:  "News Status",
	}

	err = conn.Create(&fakeNews).Error
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	repo := NewNewsRepository(conn)

	err = repo.DeleteNews(uint64(fakeNews.ID))
	assert.Nil(t, err)
}
