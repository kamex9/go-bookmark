package repository

import (
	"fmt"
	"go-bookmark/core/models"
)

type MemoryStore struct {
	data []*models.Bookmark
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([]*models.Bookmark, 0, 256),
	}
}

func (ms *MemoryStore) Save(bm *models.Bookmark) error {
	ms.data = append(ms.data, bm)
	return nil
}

func (ms *MemoryStore) FindAll() ([]*models.Bookmark, error) {
	return ms.data, nil
}

func (ms *MemoryStore) FindById(id string) (*models.Bookmark, error) {
	for _, bm := range ms.data {
		if bm.ID == id {
			return bm, nil
		}
	}
	return nil, fmt.Errorf("specified id '%s' not found", id)
}

func (ms *MemoryStore) DeleteAll() {
	ms.data = make([]*models.Bookmark, 0, 256)
}
