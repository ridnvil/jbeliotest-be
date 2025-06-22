package domain

type MostLeastSales struct {
	ProductID         uint   `json:"product_id"`
	ProductName       string `json:"product_name"`
	TotalQuantitySold int    `json:"total_quantity_sold"`
}
