package persistence

import (
	"github.com/akwanmaroso/news/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagRepo_SaveTag(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	tag := entity.Tag{}
	tag.Name = "Tag Name"

	repo := NewTagRepository(conn)

	n, saveErr := repo.SaveTag(&tag)
	assert.Nil(t, saveErr)
	assert.EqualValues(t, n.Name, "Tag Name")
}

func TestTagRepo_GetTag(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	fakeTag := &entity.Tag{
		Name: "Name Tag",
	}

	err = conn.Create(&fakeTag).Error
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	repo := NewTagRepository(conn)

	n, getErr := repo.GetTag(uint64(fakeTag.ID))
	assert.Nil(t, getErr)
	assert.EqualValues(t, n.Name, "Name Tag")
}

func TestTagRepo_GetAllTag(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	fakeTag := []*entity.Tag{
		{
			Name: "First Tag Name",
		}, {
			Name: "Second Tag Name",
		},
	}

	for _, v := range fakeTag {
		err = conn.Create(&v).Error
		if err != nil {
			t.Fatalf("want non error, got %v", err.Error())
		}
	}

	repo := NewTagRepository(conn)

	n, getAllErr := repo.GetAllTag()
	assert.Nil(t, getAllErr)
	assert.EqualValues(t, len(n), 2)
}

func TestTagRepo_UpdateTag(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	fakeTag := &entity.Tag{
		Name: "Name Tag",
	}

	err = conn.Create(&fakeTag).Error
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	fakeTag.Name = "Name Tag Updated"

	repo := NewTagRepository(conn)
	n, updateErr := repo.UpdateTag(fakeTag)
	assert.Nil(t, updateErr)
	assert.EqualValues(t, n.Name, "Name Tag Updated")
}

func TestTagRepo_DeleteTag(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	fakeTag := &entity.Tag{
		Name: "Name Tag",
	}

	err = conn.Create(&fakeTag).Error
	if err != nil {
		t.Fatalf("want non error, got %v", err.Error())
	}

	repo := NewTagRepository(conn)

	err = repo.DeleteTag(uint64(fakeTag.ID))
	assert.Nil(t, err)
}
