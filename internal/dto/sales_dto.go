package dto

import "time"

type SaleDTO struct {
	CategoryID     uint      `json:"category_id"`
	City           string    `json:"city"`
	Country        string    `json:"country"`
	CustomerName   string    `json:"customer_name"`
	Manufacturer   string    `json:"manufacturer"`
	OrderDate      time.Time `json:"order_date"`
	OrderID        string    `json:"order_id"`
	PostalCode     string    `json:"postal_code"`
	ProductID      uint      `json:"product_id"`
	Region         string    `json:"region"`
	Segment        string    `json:"segment"`
	ShipDate       time.Time `json:"ship_date"`
	ShipMode       string    `json:"ship_mode"`
	State          string    `json:"state"`
	SubCategoryID  uint      `json:"sub_category_id"`
	Discount       float64   `json:"discount"`
	NumberOfRecord int       `json:"number_of_record"`
	Profit         float64   `json:"profit"`
	ProfitRatio    float64   `json:"profit_ratio"`
	Quantity       int       `json:"quantity"`
	Sales          float64   `json:"sales"`
}

type SupermarketDTO struct {
	Category       string    `json:"category"`
	City           string    `json:"city"`
	Country        string    `json:"country"`
	CustomerName   string    `json:"customer_name"`
	Manufacturer   string    `json:"manufacturer"`
	OrderDate      time.Time `json:"order_date"`
	OrderID        string    `json:"order_id"`
	PostalCode     string    `json:"postal_code"`
	ProductName    string    `json:"product_name"`
	Region         string    `json:"region"`
	Segment        string    `json:"segment"`
	ShipDate       time.Time `json:"ship_date"`
	ShipMode       string    `json:"ship_mode"`
	State          string    `json:"state"`
	SubCategory    string    `json:"sub_category"`
	Discount       float64   `json:"discount"`
	NumberOfRecord int       `json:"number_of_record"`
	Profit         float64   `json:"profit"`
	ProfitRatio    float64   `json:"profit_ratio"`
	Quantity       int       `json:"quantity"`
	Sales          float64   `json:"sales"`
}
