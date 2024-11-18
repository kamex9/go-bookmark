package bookmark

import (
	"go-bookmark/core/constants"
)

type Repository interface {
	Save(*BookmarkEntity) error
	FindAll() ([]*BookmarkEntity, error)
	FindById(string) (*BookmarkEntity, error)
	DeleteAll()
}

func NewBookmarkRepository(mode constants.StoreMode) Repository {
	switch mode {
	case constants.MEMORY:
		return NewBookmarkMemoryStore()
	// case constants.FILE:
	// 	return NewFileStore()
	// case constants.DATABASE:
	// 	return NewDatabaseStore()
	default:
		return NewBookmarkMemoryStore()
	}
}
