package usecase

import (
	"go-bookmark/app/web/dto"
	"go-bookmark/core/constants"
	"go-bookmark/core/domain/models"
	"go-bookmark/core/persistence/bookmark"
)

type CreateBookmarkUseCase interface {
	CreateBookmark(*dto.BookmarkCreateRequest) (*models.Bookmark, error)
}

type CreateBookmarkService struct {
	repository bookmark.Repository
}

func NewCreateBookmarkService(mode constants.StoreMode) *CreateBookmarkService {
	return &CreateBookmarkService{
		repository: bookmark.NewBookmarkRepository(mode),
	}
}

func (s *CreateBookmarkService) CreateBookmark(dto *dto.BookmarkCreateRequest) (*models.Bookmark, error) {
	entity := bookmark.NewBookmarkEntityForCreate(dto)
	s.repository.Save(entity)
	return entity.NewBookmarkFromEntity(), nil
}
