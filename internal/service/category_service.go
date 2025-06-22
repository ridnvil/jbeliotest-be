package service

import (
	"jubeliotesting/internal/domain"
	"jubeliotesting/internal/dto"
)

type CategoryService struct {
	CategoryRepo domain.CategoryRepository
}

func NewCategoryService(categoryRepo domain.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: categoryRepo}
}

func (r *CategoryService) CreateCategory(categoryDto dto.CategoryDTO) error {
	var category domain.Category
	category.Name = categoryDto.Name
	return r.CategoryRepo.Create(category)
}

func (r *CategoryService) GetCategoryByName(name string) (*domain.Category, error) {
	category, err := r.CategoryRepo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return category, nil
}
