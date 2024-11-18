package bookmark

import (
	"fmt"
)

type BookmarkMemoryStore struct {
	data []*BookmarkEntity
}

func NewMemoryStore() *BookmarkMemoryStore {
	return &BookmarkMemoryStore{
		data: make([]*BookmarkEntity, 0, 256),
	}
}

func (ms *BookmarkMemoryStore) Save(bm *BookmarkEntity) error {
	ms.data = append(ms.data, bm)
	return nil
}

func (ms *BookmarkMemoryStore) FindAll() ([]*BookmarkEntity, error) {
	return ms.data, nil
}

func (ms *BookmarkMemoryStore) FindById(id string) (*BookmarkEntity, error) {
	for _, bm := range ms.data {
		if bm.ID == id {
			return bm, nil
		}
	}
	return nil, fmt.Errorf("specified id '%s' not found", id)
}

func (ms *BookmarkMemoryStore) DeleteAll() error {
	ms.data = make([]*BookmarkEntity, 0, 256)
	return nil
}
