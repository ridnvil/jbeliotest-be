package repository

import (
	"gorm.io/gorm"
	"jubeliotesting/internal/domain"
	"log"
)

type SalesRepository struct {
	DB *gorm.DB
}

func (s SalesRepository) CountrySales() ([]domain.CountrySales, error) {
	var countrySales []domain.CountrySales
	if err := s.DB.Table("sales").
		Select("country as country, SUM(quantity) AS total_quantity").
		Group("country").
		Order("total_quantity DESC").
		Scan(&countrySales).Error; err != nil {
		return nil, err
	}
	return countrySales, nil
}

func NewSalesRepository(db *gorm.DB) domain.SalesRepository {
	return SalesRepository{DB: db}
}

func (s SalesRepository) Create(Sale domain.Sale) error {
	if err := s.DB.Create(&Sale).Error; err != nil {
		return err
	}
	return nil
}

func (s SalesRepository) FindAll() ([]domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (s SalesRepository) FindByID(id uint) (*domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (s SalesRepository) CorrelationSales(groupBy string) ([]domain.CorrelationPoint, error) {
	var correlation []domain.CorrelationPoint

	if groupBy == "category" {
		if err := s.DB.Table("sales").
			Select("sales.quantity, sales.discount, categories.name AS category").
			Joins("INNER JOIN categories ON sales.category_id = categories.id").
			Group("sales.product_id, sales.quantity, sales.discount, categories.name").
			Order("categories.name DESC").
			Scan(&correlation).Error; err != nil {
			return nil, err
		}
	} else if groupBy == "subcategory" {
		if err := s.DB.Table("sales").
			Select("sales.quantity, sales.discount, sub_categories.name AS category").
			Joins("INNER JOIN sub_categories ON sales.Sub_category_id = sub_categories.id").
			Group("sales.product_id, sales.quantity, sales.discount, sub_categories.name").
			Order("sub_categories.name DESC").
			Scan(&correlation).Error; err != nil {
			return nil, err
		}
	} else if groupBy == "region" {
		if err := s.DB.Table("sales").
			Select("sales.quantity, sales.discount, sales.region AS category").
			Group("sales.product_id, sales.quantity, sales.discount, sales.region").
			Order("sales.region DESC").
			Scan(&correlation).Error; err != nil {
			return nil, err
		}
	} else if groupBy == "segment" {
		if err := s.DB.Table("sales").
			Select("sales.quantity, sales.discount, sales.segment AS category").
			Group("sales.product_id, sales.quantity, sales.discount, sales.segment").
			Order("sales.segment DESC").
			Scan(&correlation).Error; err != nil {
			return nil, err
		}
	}
	return correlation, nil
}

func (s SalesRepository) Truncate() {
	if err := s.DB.Exec(`TRUNCATE TABLE sales`).Error; err != nil {
		log.Fatal(err)
	}
}

func (s SalesRepository) MostLeastSales() ([]domain.MostLeastSales, error) {
	var most []domain.MostLeastSales
	if err := s.DB.
		Table("sales").
		Select("sales.product_id, products.name AS product_name, SUM(sales.quantity) AS total_quantity_sold").
		Joins("INNER JOIN products ON sales.product_id = products.id").
		Group("sales.product_id, products.name").
		Order("total_quantity_sold DESC").
		Limit(10).
		Scan(&most).Error; err != nil {
		return nil, err
	}
	return most, nil
}
