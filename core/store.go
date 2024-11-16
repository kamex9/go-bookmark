package core

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomTime time.Time

// MarshalJSON implements the json.Marshaler interface.
// It formats the time according to the specified layout.
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ct).Format(time.DateTime))
}

// TODO: 重複を許容してしまっている
// TODO: ID識別がないのでエンティティとしては機能しない
type Bookmark struct {
	ID        uint64     `json:"id"`
	URL       string     `json:"url"`
	Title     string     `json:"title"`
	CreatedAt CustomTime `json:"created_at"`
}

var bookmarkData = map[uint64]Bookmark{
	1: {
		ID:    1,
		URL:   "https://zenn.dev/hsaki/articles/go-time-cheatsheet",
		Title: "Goで時刻を扱うチートシート",
	},
	2: {
		ID:    2,
		URL:   "https://qiita.com/twrcd1227/items/1a05ffa459f45b2968e4",
		Title: "【2024】Go言語おすすめライブラリ15選",
	},
}

func NewBookmark(id uint64) (*Bookmark, error) {
	bookmark, exists := bookmarkData[id]
	if !exists {
		return nil, fmt.Errorf("bookmark with ID %d does not exist", id)
	}
	bookmark.CreatedAt = CustomTime(time.Now())
	return &bookmark, nil
}

type Store struct {
	data []Bookmark
}

func NewStore() *Store {
	return &Store{
		data: make([]Bookmark, 0, 100),
	}
}

func (s *Store) Add(b Bookmark) {
	s.data = append(s.data, b)
}

func (s *Store) GetAll() []Bookmark {
	result := make([]Bookmark, len(s.data))
	copy(result, s.data)
	return result
}

func (s *Store) Reset() {
	s.data = make([]Bookmark, 0, 100)
}
