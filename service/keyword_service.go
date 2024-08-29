package service

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/repository"
    "errors"
    "go.mongodb.org/mongo-driver/mongo"
)

type KeywordService struct {
    repo repository.KeywordRepository
}

func NewKeywordService(repo repository.KeywordRepository) *KeywordService {
    return &KeywordService{repo: repo}
}

func (s *KeywordService) validateKeyword(kw *models.Keyword) error {
    if kw.CourseId == "" {
        return errors.New("courseID é obrigatório")
    }
    if kw.Keyword == "" {
        return errors.New("keyword é obrigatório")
    }
    if kw.ClassMaterialId == "" {
        return errors.New("classMaterialId é obrigatório")
    }
    if kw.TranscriptTimeId == "" {
        return errors.New("transcriptTimeId é obrigatório")
    }
    return nil
}

func (s *KeywordService) Create(ctx context.Context, keyword *models.Keyword) error {
    if err := s.validateKeyword(keyword); err != nil {
        return err
    }
    
    // Verifica se já existe um ClassMaterial com o mesmo UUID
    existingKW, _ := s.repo.FindByID(ctx, keyword.ID)
    if existingKW != nil {
        return errors.New("keyword já existe")
    }
    
    return s.repo.Create(ctx, keyword)
}

func (s *KeywordService) FindAll(ctx context.Context) ([]models.Keyword, error) {
    return s.repo.FindAll(ctx)
}

func (s *KeywordService) FindByID(ctx context.Context, id string) (*models.Keyword, error) {
    if id == "" {
        return nil, errors.New("id é obrigatório")
    }
    
    keyword, err := s.repo.FindByID(ctx, id)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, errors.New("keyword não encontrado")
        }
        return nil, err
    }
    
    return keyword, nil
}

func (s *KeywordService) Update(ctx context.Context, id string, keyword *models.Keyword) error {
    if id == "" {
        return errors.New("id é obrigatório")
    }
    
    if err := s.validateKeyword(keyword); err != nil {
        return err
    }
    
    // Verifica se o Keyword existe antes de atualizar
    existingKW, _ := s.repo.FindByID(ctx, id)
    if existingKW == nil {
        return errors.New("keyword não encontrado")
    }
    
    return s.repo.Update(ctx, id, keyword)
}

func (s *KeywordService) Delete(ctx context.Context, id string) error {
    if id == "" {
        return errors.New("id é obrigatório")
    }
    
    // Verifica se o Keyword existe antes de deletar
    existingKW, _ := s.repo.FindByID(ctx, id)
    if existingKW == nil {
        return errors.New("keyword não encontrado")
    }
    
    return s.repo.Delete(ctx, id)
}

func (s *KeywordService) UpdateKeywordsByTranscriptTimeID(ctx context.Context, transcriptTimeId string, keywords []string) error {
    if transcriptTimeId == "" {
        return errors.New("transcriptTimeId é obrigatório")
    }
    
    // Encontre todos os documentos com o transcriptTimeId dado
    keywordDocs, err := s.repo.FindByTranscriptTimeID(ctx, transcriptTimeId)
    if err != nil {
        return err
    }
    
    if len(keywords) != len(keywordDocs) {
        return errors.New("o número de keywords fornecidas não corresponde ao número de documentos encontrados")
    }
    
    for i, keywordDoc := range keywordDocs {
        keywordDoc.Keyword = keywords[i]
        if err := s.repo.Update(ctx, keywordDoc.ID, &keywordDoc); err != nil {
            return err
        }
    }
    
    return nil
}