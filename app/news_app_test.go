package app

import (
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var (
	saveNewsRepo    func(*entity.News) (*entity.News, map[string]string)
	getAllNewsRepo  func() ([]*entity.News, error)
	getNewsRepo     func(uint64) (*entity.News, error)
	getByTopicRepo  func(string) ([]*entity.News, error)
	getByStatusRepo func(string) ([]*entity.News, error)
	updateNewsRepo  func(*entity.News) (*entity.News, map[string]string)
	deleteNewsRepo  func(uint64) error
)

type fakeNewsRepo struct{}

func (f *fakeNewsRepo) SaveNews(news *entity.News) (*entity.News, map[string]string) {
	return saveNewsRepo(news)
}

func (f *fakeNewsRepo) GetNews(newsID uint64) (*entity.News, error) {
	return getNewsRepo(newsID)
}

func (f *fakeNewsRepo) GetAllNews() ([]*entity.News, error) {
	return getAllNewsRepo()
}

func (f *fakeNewsRepo) GetByTopic(topic string) ([]*entity.News, error) {
	return getByTopicRepo(topic)
}

func (f *fakeNewsRepo) GetByStatus(status string) ([]*entity.News, error) {
	return getByStatusRepo(status)
}

func (f *fakeNewsRepo) UpdateNews(news *entity.News) (*entity.News, map[string]string) {
	return updateNewsRepo(news)
}

func (f *fakeNewsRepo) DeleteNews(newsID uint64) error {
	return deleteNewsRepo(newsID)
}

var newsAppFake NewsAppInterface = &fakeNewsRepo{}

func TestNewsApps_SaveNews(t *testing.T) {
	fakeNewTag := &entity.Tag{
		Name: "First Tag",
	}
	saveNewsRepo = func(*entity.News) (*entity.News, map[string]string) {
		return &entity.News{
			Topic:   "First Topic",
			Content: "First Content",
			Status:  "Draft",
			Tags: []*entity.Tag{
				fakeNewTag,
			},
		}, nil
	}

	news := &entity.News{
		Topic:   "First Topic",
		Content: "First Content",
		Status:  "Draft",
		Tags: []*entity.Tag{
			fakeNewTag,
		},
	}

	n, err := newsAppFake.SaveNews(news)
	assert.Nil(t, err)
	assert.EqualValues(t, n.Topic, "First Topic")
	assert.EqualValues(t, n.Content, "First Content")
	assert.EqualValues(t, n.Status, "Draft")
	assert.EqualValues(t, n.Status, "Draft")
	assert.EqualValues(t, n.Tags, []*entity.Tag{{Name: "First Tag"}})
}

func TestNewsApps_GetNews(t *testing.T) {
	fakeNewTag := &entity.Tag{
		Name: "First Tag",
	}
	getNewsRepo = func(newsID uint64) (*entity.News, error) {
		return &entity.News{
			Topic:   "First Topic",
			Content: "First Content",
			Status:  "Draft",
			Tags: []*entity.Tag{
				fakeNewTag,
			},
		}, nil
	}
	newsID := uint64(1)
	n, err := newsAppFake.GetNews(newsID)
	assert.Nil(t, err)
	assert.EqualValues(t, n.Topic, "First Topic")
	assert.EqualValues(t, n.Content, "First Content")
	assert.EqualValues(t, n.Status, "Draft")
	assert.EqualValues(t, n.Tags, []*entity.Tag{{Name: "First Tag"}})
}

func TestNewsApps_GetAllNews(t *testing.T) {
	fakeFirstNewTag := &entity.Tag{Name: "First Tag"}
	fakeSecondNewTag := &entity.Tag{Name: "Second Tag"}
	getAllNewsRepo = func() ([]*entity.News, error) {
		return []*entity.News{
			{
				Topic:   "First Topic",
				Content: "First Content",
				Status:  "Draft",
				Tags: []*entity.Tag{
					fakeFirstNewTag,
				},
			}, {
				Topic:   "Second Topic",
				Content: "Second Content",
				Status:  "Draft",
				Tags: []*entity.Tag{
					fakeSecondNewTag,
				},
			},
		}, nil
	}

	n, err := newsAppFake.GetAllNews()
	assert.Nil(t, err)
	assert.EqualValues(t, len(n), 2)
}

func TestNewsApps_UpdateNews(t *testing.T) {
	updateNewsRepo = func(news *entity.News) (*entity.News, map[string]string) {
		return &entity.News{
			Topic:   "First Topic Updated",
			Content: "First Content Updated",
			Status:  "Draft Updated",
		}, nil
	}

	news := &entity.News{
		Topic:   "First Topic Updated",
		Content: "First Content Updated",
		Status:  "Draft Updated",
	}

	n, err := newsAppFake.UpdateNews(news)
	log.Println(n.Topic)
	assert.Nil(t, err)
	assert.EqualValues(t, n.Topic, "First Topic Updated")
	assert.EqualValues(t, n.Content, "First Content Updated")
	assert.EqualValues(t, n.Status, "Draft Updated")
}

func TestNewsApps_DeleteNews(t *testing.T) {
	deleteNewsRepo = func(newsID uint64) error {
		return nil
	}

	newsID := uint64(1)
	err := deleteNewsRepo(newsID)
	assert.Nil(t, err)
}
