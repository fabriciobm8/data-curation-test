package service

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/repository"
)

type KeywordService struct {
    repo repository.KeywordRepository
}

func NewKeywordService(repo repository.KeywordRepository) *KeywordService {
    return &KeywordService{repo: repo}
}

func (s *KeywordService) Create(ctx context.Context, keyword *models.Keyword) error {
    return s.repo.Create(ctx, keyword)
}

func (s *KeywordService) FindAll(ctx context.Context) ([]models.Keyword, error) {
    return s.repo.FindAll(ctx)
}

func (s *KeywordService) FindByID(ctx context.Context, id string) (*models.Keyword, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *KeywordService) Update(ctx context.Context, id string, keyword *models.Keyword) error {
    return s.repo.Update(ctx, id, keyword)
}

func (s *KeywordService) Delete(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}