package bookmark

import (
	"go-bookmark/app/web/dto"
	"go-bookmark/core/domain/models"
	"time"

	"github.com/rs/xid"
)

type BookmarkEntity struct {
	ID        string
	URL       string
	Title     string
	CreatedAt time.Time
}

func NewBookmarkEntityForCreate(r *dto.BookmarkCreateRequest) *BookmarkEntity {
	return &BookmarkEntity{
		ID:        xid.New().String(),
		URL:       r.URL,
		Title:     r.Title,
		CreatedAt: time.Now(),
	}
}

func (be *BookmarkEntity) NewBookmarkFromEntity() *models.Bookmark {
	return &models.Bookmark{
		ID:        be.ID,
		URL:       be.URL,
		Title:     be.Title,
		CreatedAt: be.CreatedAt,
	}
}
