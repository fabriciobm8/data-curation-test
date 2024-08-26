package service

import (
    "context"
    "data-curation-test/models"
    "data-curation-test/repository"
)

type ClassMaterialService interface {
    Create(ctx context.Context, classMaterial *models.ClassMaterial) error
    FindAll(ctx context.Context) ([]models.ClassMaterial, error)
    FindByID(ctx context.Context, id string) (*models.ClassMaterial, error)
    Update(ctx context.Context, id string, classMaterial *models.ClassMaterial) error
    Delete(ctx context.Context, id string) error
}

type classMaterialService struct {
    repo repository.ClassMaterialRepository
}

func NewClassMaterialService(repo repository.ClassMaterialRepository) ClassMaterialService {
    return &classMaterialService{repo: repo}
}

func (s *classMaterialService) Create(ctx context.Context, classMaterial *models.ClassMaterial) error {
    return s.repo.Create(ctx, classMaterial)
}

func (s *classMaterialService) FindAll(ctx context.Context) ([]models.ClassMaterial, error) {
    return s.repo.FindAll(ctx)
}

func (s *classMaterialService) FindByID(ctx context.Context, id string) (*models.ClassMaterial, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *classMaterialService) Update(ctx context.Context, id string, classMaterial *models.ClassMaterial) error {
    return s.repo.Update(ctx, id, classMaterial)
}

func (s *classMaterialService) Delete(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}
