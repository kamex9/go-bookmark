package usecase

import (
	"go-bookmark/app/web/dto"
	"go-bookmark/core/constants"
	"go-bookmark/core/domain/models"
	"go-bookmark/core/persistence/bookmark"
)

type CrudBookmarkUseCase interface {
	Create(*dto.BookmarkCreateRequest) (*models.Bookmark, error)
	FindAll() ([]*models.Bookmark, error)
	FindById(string) (*models.Bookmark, error)
	DeleteAll() error
}

type CreateBookmarkService struct {
	repo bookmark.Repository
}

func NewCreateBookmarkService(mode constants.StoreMode) *CreateBookmarkService {
	return &CreateBookmarkService{
		repo: bookmark.NewBookmarkRepository(mode),
	}
}

func (s *CreateBookmarkService) Create(dto *dto.BookmarkCreateRequest) (*models.Bookmark, error) {
	e := bookmark.NewBookmarkEntityForCreate(dto)
	s.repo.Save(e)
	return e.NewBookmarkFromEntity(), nil
}

func (s *CreateBookmarkService) FindAll() ([]*models.Bookmark, error) {
	es, err := s.repo.FindAll()
	return convertBookmarkEntitiesToModels(es), err
}

func (s *CreateBookmarkService) FindById(id string) (*models.Bookmark, error) {
	e, err := s.repo.FindById(id)
	if e == nil {
		return nil, err
	}
	return e.NewBookmarkFromEntity(), err
}

func (s *CreateBookmarkService) DeleteAll() error {
	s.repo.DeleteAll()
	return nil
}

func convertBookmarkEntitiesToModels(es []*bookmark.BookmarkEntity) []*models.Bookmark {
	results := make([]*models.Bookmark, 0, len(es))
	for _, e := range es {
		results = append(results, e.NewBookmarkFromEntity())
	}
	return results
}
