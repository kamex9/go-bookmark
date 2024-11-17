package repository

import (
	"go-bookmark/core/models"
)

type StoreMode uint8

const (
	_ StoreMode = iota
	MEMORY
	FILE
	DATABASE
)

type Repository interface {
	Save(*models.Bookmark) error
	FindAll() ([]*models.Bookmark, error)
	FindById(string) (*models.Bookmark, error)
	DeleteAll()
}

func NewRepository(mode StoreMode) Repository {
	switch mode {
	case MEMORY:
		return NewMemoryStore()
	// case FILE:
	// 	return NewFileStore()
	// case DATABASE:
	// 	return NewDatabaseStore()
	default:
		return NewMemoryStore()
	}
}
