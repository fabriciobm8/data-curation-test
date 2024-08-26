package service

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/repository"
)

type TranscriptTimeService interface {
    Create(ctx context.Context, transcriptTime *models.TranscriptTime) error
    FindAll(ctx context.Context) ([]models.TranscriptTime, error)
    FindByID(ctx context.Context, id string) (*models.TranscriptTime, error)
    Update(ctx context.Context, id string, transcriptTime *models.TranscriptTime) error
    Delete(ctx context.Context, id string) error
}

type transcriptTimeService struct {
    repo repository.TranscriptTimeRepository
}

func NewTranscriptTimeService(repo repository.TranscriptTimeRepository) TranscriptTimeService {
    return &transcriptTimeService{repo: repo}
}

func (s *transcriptTimeService) Create(ctx context.Context, transcriptTime *models.TranscriptTime) error {
    return s.repo.Create(ctx, transcriptTime)
}

func (s *transcriptTimeService) FindAll(ctx context.Context) ([]models.TranscriptTime, error) {
    return s.repo.FindAll(ctx)
}

func (s *transcriptTimeService) FindByID(ctx context.Context, id string) (*models.TranscriptTime, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *transcriptTimeService) Update(ctx context.Context, id string, transcriptTime *models.TranscriptTime) error {
    return s.repo.Update(ctx, id, transcriptTime)
}

func (s *transcriptTimeService) Delete(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}
