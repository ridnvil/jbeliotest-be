package repository

import (
	"errors"
	"gorm.io/gorm"
	"jubeliotesting/internal/domain"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (r *CategoryRepository) FindByName(name string) (*domain.Category, error) {
	var category domain.Category
	if err := r.DB.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) FindAll() ([]domain.Category, error) {
	var categories []domain.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) Create(category domain.Category) error {
	var name string
	if err := r.DB.Model(&domain.Category{}).Select("name").Where("name = ?", category.Name).Scan(&name).Error; err != nil {
		return err
	}

	if name != "" {
		return errors.New("category already exists")
	}

	if err := r.DB.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) FindByID(id uint) (*domain.Category, error) {
	//TODO implement me
	panic("implement me")
}
