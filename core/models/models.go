package models

import (
	"time"

	"github.com/rs/xid"
)

// TODO: 重複を許容してしまっている
// TODO: ID識別がないのでエンティティとしては機能しない
type Bookmark struct {
	ID        string
	URL       string
	Title     string
	CreatedAt CustomTime
}

func NewBookmark(r BookmarkPostRequest) *Bookmark {
	return &Bookmark{
		ID:        xid.New().String(),
		URL:       r.URL,
		Title:     r.Title,
		CreatedAt: CustomTime(time.Now()),
	}
}
