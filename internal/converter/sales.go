package converter

import (
	"jubeliotesting/internal/domain"
	"jubeliotesting/internal/dto"
)

func ConvertSaleDTOToSale(dto dto.SaleDTO) domain.Sale {
	return domain.Sale{
		CategoryID:     dto.CategoryID,
		City:           dto.City,
		Country:        dto.Country,
		CustomerName:   dto.CustomerName,
		Manufacturer:   dto.Manufacturer,
		OrderDate:      dto.OrderDate,
		OrderID:        dto.OrderID,
		PostalCode:     dto.PostalCode,
		ProductID:      dto.ProductID,
		Region:         dto.Region,
		Segment:        dto.Segment,
		ShipDate:       dto.ShipDate,
		ShipMode:       dto.ShipMode,
		State:          dto.State,
		SubCategoryID:  dto.SubCategoryID,
		Discount:       dto.Discount,
		NumberOfRecord: dto.NumberOfRecord,
		Profit:         dto.Profit,
		ProfitRatio:    dto.ProfitRatio,
		Quantity:       dto.Quantity,
		Sales:          dto.Sales,
	}
}
