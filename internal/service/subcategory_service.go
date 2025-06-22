package service

import (
	"jubeliotesting/internal/domain"
	"jubeliotesting/internal/dto"
)

type SubCategoryService struct {
	SubCategoryRepo domain.SubCategoryRepository
}

func NewSubCategoryService(subCategoryRepo domain.SubCategoryRepository) *SubCategoryService {
	return &SubCategoryService{SubCategoryRepo: subCategoryRepo}
}

func (r *SubCategoryService) CreateSubCategory(subcategoryDto dto.SubCategoryDTO) error {
	var subcategory domain.SubCategory
	subcategory.Name = subcategoryDto.Name
	subcategory.CategoryID = subcategoryDto.CategoryID
	return r.SubCategoryRepo.Create(subcategory)
}

func (r *SubCategoryService) GetSubCategoryByName(subcategoryname string) (domain.SubCategory, error) {
	subcategory, err := r.SubCategoryRepo.FindByName(subcategoryname)
	if err != nil {
		return domain.SubCategory{}, err
	}
	return subcategory, nil
}
