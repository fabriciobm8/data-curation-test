package service

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/repository"
)

type KeywordService interface {
    Create(ctx context.Context, keyword *models.Keyword) error
    FindAll(ctx context.Context) ([]models.Keyword, error)
    FindByID(ctx context.Context, id string) (*models.Keyword, error)
    Update(ctx context.Context, id string, keyword *models.Keyword) error
    Delete(ctx context.Context, id string) error
}

type keywordService struct {
    repo repository.KeywordRepository
}

func NewKeywordService(repo repository.KeywordRepository) KeywordService {
    return &keywordService{repo: repo}
}

func (s *keywordService) Create(ctx context.Context, keyword *models.Keyword) error {
    return s.repo.Create(ctx, keyword)
}

func (s *keywordService) FindAll(ctx context.Context) ([]models.Keyword, error) {
    return s.repo.FindAll(ctx)
}

func (s *keywordService) FindByID(ctx context.Context, id string) (*models.Keyword, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *keywordService) Update(ctx context.Context, id string, keyword *models.Keyword) error {
    return s.repo.Update(ctx, id, keyword)
}

func (s *keywordService) Delete(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}
