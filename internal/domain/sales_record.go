package domain

import "time"

type Sale struct {
	ID             uint `gorm:"primaryKey"`
	CategoryID     uint
	City           string `gorm:"size:100"`
	Country        string `gorm:"size:100"`
	CustomerName   string `gorm:"size:255"`
	Manufacturer   string `gorm:"size:255"`
	OrderDate      time.Time
	OrderID        string `gorm:"size:50;not null"`
	PostalCode     string `gorm:"size:20"`
	ProductID      uint
	Region         string `gorm:"size:100"`
	Segment        string `gorm:"size:100"`
	ShipDate       time.Time
	ShipMode       string `gorm:"size:100"`
	State          string `gorm:"size:100"`
	SubCategoryID  uint
	Discount       float64
	NumberOfRecord int
	Profit         float64
	ProfitRatio    float64
	Quantity       int
	Sales          float64

	Category    Category    `gorm:"foreignKey:CategoryID"`
	Product     Product     `gorm:"foreignKey:ProductID"`
	SubCategory SubCategory `gorm:"foreignKey:SubCategoryID"`
}

type SalesRepository interface {
	Create(Sale Sale) error
	FindAll() ([]Sale, error)
	FindByID(id uint) (*Sale, error)

	CountrySales() ([]CountrySales, error)
	CorrelationSales(groupBy string) ([]CorrelationPoint, error)
	MostLeastSales() ([]MostLeastSales, error)
	Truncate()
}
