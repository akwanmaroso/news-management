package mock

import "github.com/akwanmaroso/news/domain/entity"

type NewsAppInterface struct {
	SaveNewsFn    func(*entity.News) (*entity.News, map[string]string)
	GetAllNewsFn  func() ([]*entity.News, error)
	GetNewsFn     func(uint64) (*entity.News, error)
	GetByTopicFn  func(string) ([]*entity.News, error)
	GetByStatusFn func(string) ([]*entity.News, error)
	UpdateNewsFn  func(*entity.News) (*entity.News, map[string]string)
	DeleteNewsFn  func(uint64) error
}

func (n *NewsAppInterface) SaveNews(news *entity.News) (*entity.News, map[string]string) {
	return n.SaveNewsFn(news)
}

func (n *NewsAppInterface) GetAllNews() ([]*entity.News, error) {
	return n.GetAllNewsFn()
}

func (n *NewsAppInterface) GetNews(newsID uint64) (*entity.News, error) {
	return n.GetNewsFn(newsID)
}

func (n *NewsAppInterface) GetByTopic(topic string) ([]*entity.News, error) {
	return n.GetByTopicFn(topic)
}

func (n *NewsAppInterface) GetByStatus(status string) ([]*entity.News, error) {
	return n.GetByStatusFn(status)
}

func (n *NewsAppInterface) UpdateNews(news *entity.News) (*entity.News, map[string]string) {
	return n.UpdateNewsFn(news)
}

func (n *NewsAppInterface) DeleteNews(newsID uint64) error {
	return n.DeleteNewsFn(newsID)
}

type TagAppInterface struct {
	SaveTagFn       func(*entity.Tag) (*entity.Tag, map[string]string)
	GetTagFn        func(uint64) (*entity.Tag, error)
	GetAllTagFn     func() ([]*entity.Tag, error)
	FindTagByNameFn func(name string) (*entity.Tag, error)
	UpdateTagFn     func(*entity.Tag) (*entity.Tag, map[string]string)
	DeleteTagFn     func(uint64) error
}

func (t *TagAppInterface) SaveTag(tag *entity.Tag) (*entity.Tag, map[string]string) {
	return t.SaveTagFn(tag)
}

func (t *TagAppInterface) GetTag(tagID uint64) (*entity.Tag, error) {
	return t.GetTag(tagID)
}

func (t *TagAppInterface) GetAllTag() ([]*entity.Tag, error) {
	return t.GetAllTagFn()
}

func (t *TagAppInterface) FindTagByName(name string) (*entity.Tag, error) {
	return t.FindTagByNameFn(name)
}

func (t *TagAppInterface) UpdateTag(tag *entity.Tag) (*entity.Tag, map[string]string) {
	return t.UpdateTagFn(tag)
}

func (t *TagAppInterface) DeleteTag(tagID uint64) error {
	return t.DeleteTagFn(tagID)
}
