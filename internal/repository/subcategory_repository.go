package repository

import (
	"gorm.io/gorm"
	"jubeliotesting/internal/domain"
)

type SubCategoryRepository struct {
	DB *gorm.DB
}

func (r *SubCategoryRepository) FindByName(name string) (domain.SubCategory, error) {
	var subCategory domain.SubCategory
	if err := r.DB.Where("name = ?", name).First(&subCategory).Error; err != nil {
		return subCategory, err
	}
	return subCategory, nil
}

func (r *SubCategoryRepository) FindAll() ([]domain.SubCategory, error) {
	var categories []domain.SubCategory
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func NewSubCategoryRepository(db *gorm.DB) domain.SubCategoryRepository {
	return &SubCategoryRepository{DB: db}
}

func (r *SubCategoryRepository) Create(category domain.SubCategory) error {
	if err := r.DB.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *SubCategoryRepository) FindByID(id uint) (domain.SubCategory, error) {
	//TODO implement me
	panic("implement me")
}
