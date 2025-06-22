package repository

import (
	"errors"
	"gorm.io/gorm"
	"jubeliotesting/internal/domain"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (p *ProductRepository) FindByName(name string) (*domain.Product, error) {
	var product domain.Product
	if err := p.DB.Where("name = ?", name).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &ProductRepository{DB: db}
}

func (p *ProductRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	if err := p.DB.Find(&products).Error; err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

func (p *ProductRepository) Create(product domain.Product) error {

	var name string
	if err := p.DB.Model(&domain.Product{}).Select("name").Where("name = ?", product.Name).Scan(&name).Error; err != nil {
		return err
	}

	if name != "" {
		return errors.New("product already exists")
	}

	if err := p.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) FindByID(id uint) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
