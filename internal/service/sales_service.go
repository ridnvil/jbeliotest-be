package service

import (
	"jubeliotesting/internal/converter"
	"jubeliotesting/internal/domain"
	"jubeliotesting/internal/dto"
)

type SalesService struct {
	SalesDomainRepository domain.SalesRepository
}

func NewSalesService(salesRepo domain.SalesRepository) *SalesService {
	return &SalesService{SalesDomainRepository: salesRepo}
}

func (s *SalesService) FindAll() ([]domain.Sale, error) {
	return s.SalesDomainRepository.FindAll()
}

func (s *SalesService) Create(sale dto.SaleDTO) error {
	sales := converter.ConvertSaleDTOToSale(sale)
	return s.SalesDomainRepository.Create(sales)
}

func (s *SalesService) TruncateSales() {
	s.SalesDomainRepository.Truncate()
}

func (s *SalesService) GetMostLeastSales() ([]domain.MostLeastSales, error) {
	return s.SalesDomainRepository.MostLeastSales()
}

func (s *SalesService) GetCorrelationSales(groupBy string) ([]domain.CorrelationPoint, error) {
	return s.SalesDomainRepository.CorrelationSales(groupBy)
}

func (s *SalesService) GetCountrySales() ([]domain.CountrySales, error) {
	return s.SalesDomainRepository.CountrySales()
}
