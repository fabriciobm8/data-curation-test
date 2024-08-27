package service

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/repository"
)

type TranscriptTimeService struct {
    repo repository.TranscriptTimeRepository
}

func NewTranscriptTimeService(repo repository.TranscriptTimeRepository) *TranscriptTimeService {
    return &TranscriptTimeService{repo: repo}
}

func (s *TranscriptTimeService) Create(ctx context.Context, transcriptTime *models.TranscriptTime) error {
    return s.repo.Create(ctx, transcriptTime)
}

func (s *TranscriptTimeService) FindAll(ctx context.Context) ([]models.TranscriptTime, error) {
    return s.repo.FindAll(ctx)
}

func (s *TranscriptTimeService) FindByID(ctx context.Context, id string) (*models.TranscriptTime, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *TranscriptTimeService) Update(ctx context.Context, id string, transcriptTime *models.TranscriptTime) error {
    return s.repo.Update(ctx, id, transcriptTime)
}    

func (s *TranscriptTimeService) Delete(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}