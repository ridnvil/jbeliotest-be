package domain

type CorrelationPoint struct {
	Quantity int     `json:"quantity"`
	Discount float64 `json:"discount"`
	Category string  `json:"group"`
}
