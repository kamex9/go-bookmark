package dto

import (
	"go-bookmark/core/domain/models"
)

type BookmarkCreateRequest struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type BookmarkResponse struct {
	ID        string     `json:"id"`
	URL       string     `json:"url"`
	Title     string     `json:"title"`
	CreatedAt CustomTime `json:"created_at"`
}

func NewBookmarkResponse(bm *models.Bookmark) *BookmarkResponse {
	return &BookmarkResponse{
		ID:        bm.ID,
		URL:       bm.URL,
		Title:     bm.Title,
		CreatedAt: CustomTime(bm.CreatedAt),
	}
}
