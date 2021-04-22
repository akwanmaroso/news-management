package app

import (
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	saveTagRepo       func(*entity.Tag) (*entity.Tag, map[string]string)
	getTagRepo        func(uint64) (*entity.Tag, error)
	getAllTagRepo     func() ([]*entity.Tag, error)
	updateTagRepo     func(*entity.Tag) (*entity.Tag, map[string]string)
	findTagByNameRepo func(name string) (*entity.Tag, error)
	deleteTagRepo     func(uint64) error
)

type fakeTagApp struct{}

func (f *fakeTagApp) SaveTag(tag *entity.Tag) (*entity.Tag, map[string]string) {
	return saveTagRepo(tag)
}

func (f *fakeTagApp) GetTag(tagID uint64) (*entity.Tag, error) {
	return getTagRepo(tagID)
}

func (f *fakeTagApp) GetAllTag() ([]*entity.Tag, error) {
	return getAllTagRepo()
}

func (f *fakeTagApp) FindTagByName(name string) (*entity.Tag, error) {
	return findTagByNameRepo(name)
}

func (f *fakeTagApp) UpdateTag(tag *entity.Tag) (*entity.Tag, map[string]string) {
	return updateTagRepo(tag)
}

func (f *fakeTagApp) DeleteTag(tagID uint64) error {
	return deleteTagRepo(tagID)
}

var tagAppFake TagAppInterface = &fakeTagApp{}

func TestTagApps_SaveTag(t *testing.T) {
	saveTagRepo = func(tag *entity.Tag) (*entity.Tag, map[string]string) {
		return &entity.Tag{
			Name: "First Tag",
		}, nil
	}

	tag := &entity.Tag{
		Name: "First Tag",
	}

	u, err := tagAppFake.SaveTag(tag)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Name, "First Tag")
}

func TestTagApps_GetTag(t *testing.T) {
	getTagRepo = func(tagID uint64) (*entity.Tag, error) {
		return &entity.Tag{
			Name: "First Tag",
		}, nil
	}

	tagID := uint64(1)
	u, err := tagAppFake.GetTag(tagID)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Name, "First Tag")
}

func TestTagApps_GetAllTag(t *testing.T) {
	getAllTagRepo = func() ([]*entity.Tag, error) {
		return []*entity.Tag{
			{
				Name: "First Tag",
			}, {
				Name: "Second Tag",
			},
		}, nil
	}
	n, err := tagAppFake.GetAllTag()
	assert.Nil(t, err)
	assert.EqualValues(t, len(n), 2)
}

func TestTagApps_UpdateTag(t *testing.T) {
	updateTagRepo = func(tag *entity.Tag) (*entity.Tag, map[string]string) {
		return &entity.Tag{
			Name: "Tag Updated",
		}, nil
	}

	tag := &entity.Tag{
		Name: "Tag Updated",
	}

	n, err := tagAppFake.UpdateTag(tag)
	assert.Nil(t, err)
	assert.EqualValues(t, n.Name, "Tag Updated")
}

func TestTagApps_DeleteTag(t *testing.T) {
	deleteTagRepo = func(tagID uint64) error {
		return nil
	}
	tagID := uint64(1)
	err := tagAppFake.DeleteTag(tagID)
	assert.Nil(t, err)
}
