package service

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/repository"
    "errors"
    "go.mongodb.org/mongo-driver/mongo"
)

type TranscriptTimeService struct {
    repo repository.TranscriptTimeRepository
}

func NewTranscriptTimeService(repo repository.TranscriptTimeRepository) *TranscriptTimeService {
    return &TranscriptTimeService{repo: repo}
}

// Valida se o TranscriptTime possui todos os campos obrigatórios
func (s *TranscriptTimeService) validateTranscriptTime(tt *models.TranscriptTime) error {
    if tt.Transcript == "" {
        return errors.New("transcript é obrigatório")
    }
    if tt.ClassMaterialId == "" {
        return errors.New("classMaterialId é obrigatório")
    }    
    return nil
}

// Cria um novo TranscriptTime
func (s *TranscriptTimeService) Create(ctx context.Context, transcriptTime *models.TranscriptTime) error {
    if err := s.validateTranscriptTime(transcriptTime); err != nil {
        return err
    }
    
    // Verifica se já existe um TranscriptTime com o mesmo UUID
    existingTT, _ := s.repo.FindByID(ctx, transcriptTime.ID)
    if existingTT != nil {
        return errors.New("transcriptTime já existe")
    }
    
    return s.repo.Create(ctx, transcriptTime)
}

func (s *TranscriptTimeService) FindAll(ctx context.Context) ([]models.TranscriptTime, error) {
    return s.repo.FindAll(ctx)
}

func (s *TranscriptTimeService) FindByID(ctx context.Context, id string) (*models.TranscriptTime, error) {
    if id == "" {
        return nil, errors.New("id é obrigatório")
    }
    
    transcriptTime, err := s.repo.FindByID(ctx, id)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, errors.New("transcriptTime não encontrado")
        }
        return nil, err
    }
    
    return transcriptTime, nil
}

func (s *TranscriptTimeService) Update(ctx context.Context, id string, transcriptTime *models.TranscriptTime) error {
    if id == "" {
        return errors.New("iD é obrigatório")
    }
    
    if err := s.validateTranscriptTime(transcriptTime); err != nil {
        return err
    }
    
    // Verifica se o TranscriptTime existe antes de atualizar
    existingTT, _ := s.repo.FindByID(ctx, id)
    if existingTT == nil {
        return errors.New("transcriptTime não encontrado")
    }
    
    return s.repo.Update(ctx, id, transcriptTime)
}

func (s *TranscriptTimeService) Delete(ctx context.Context, id string) error {
    if id == "" {
        return errors.New("iD é obrigatório")
    }
    
    // Verifica se o TranscriptTime existe antes de deletar
    existingTT, _ := s.repo.FindByID(ctx, id)
    if existingTT == nil {
        return errors.New("transcriptTime não encontrado")
    }
    
    return s.repo.Delete(ctx, id)
}