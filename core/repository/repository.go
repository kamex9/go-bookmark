package repository

import (
	"go-bookmark/core/models"
)

// var bookmarkData = map[uint16]BM{
// 	1: {
// 		ID:    1,
// 		URL:   "https://zenn.dev/hsaki/articles/go-time-cheatsheet",
// 		Title: "Goで時刻を扱うチートシート",
// 	},
// 	2: {
// 		ID:    2,
// 		URL:   "https://qiita.com/twrcd1227/items/1a05ffa459f45b2968e4",
// 		Title: "【2024】Go言語おすすめライブラリ15選",
// 	},
// }

type StoreMode uint8

const (
	_ StoreMode = iota
	MEMORY
	FILE
	RDBMS
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
	default:
		return nil
	}
}
