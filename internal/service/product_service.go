package service

import (
	"jubeliotesting/internal/domain"
	"jubeliotesting/internal/dto"
)

type ProductService struct {
	ProductRepo domain.ProductRepository
}

func NewProductService(productRepo domain.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: productRepo}
}

func (p *ProductService) CreateProduct(productDto dto.ProductDTO) error {
	var product domain.Product
	product.Name = productDto.Name
	product.Manufacturer = productDto.Manufacturer
	product.SubCategoryID = productDto.SubCategoryID
	return p.ProductRepo.Create(product)
}

func (p *ProductService) GetProductsByName(productName string) (*domain.Product, error) {
	product, err := p.ProductRepo.FindByName(productName)
	if err != nil {
		return nil, err
	}
	return product, nil
}
