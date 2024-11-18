package bookmark

import (
	"go-bookmark/core/constants"
)

type Repository interface {
	Save(*BookmarkEntity) error
	FindAll() ([]*BookmarkEntity, error)
	FindById(string) (*BookmarkEntity, error)
	DeleteAll() error
}

func NewRepository(mode constants.StoreMode) Repository {
	switch mode {
	case constants.MEMORY:
		return NewMemoryStore()
	// case constants.FILE:
	// 	return NewFileStore()
	// case constants.DATABASE:
	// 	return NewDatabaseStore()
	default:
		return NewMemoryStore()
	}
}
