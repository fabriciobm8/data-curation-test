package service

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/repository"
)

type ClassMaterialService struct {
    repo repository.ClassMaterialRepository
}

func NewClassMaterialService(repo repository.ClassMaterialRepository) *ClassMaterialService {
    return &ClassMaterialService{repo: repo}
}

func (s *ClassMaterialService) Create(ctx context.Context, classMaterial *models.ClassMaterial) error {
    return s.repo.Create(ctx, classMaterial)
}

func (s *ClassMaterialService) FindAll(ctx context.Context) ([]models.ClassMaterial, error) {
    return s.repo.FindAll(ctx)
}

func (s *ClassMaterialService) FindByID(ctx context.Context, id string) (*models.ClassMaterial, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *ClassMaterialService) Update(ctx context.Context, id string, classMaterial *models.ClassMaterial) error {
    return s.repo.Update(ctx, id, classMaterial)
}    

func (s *ClassMaterialService) Delete(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}