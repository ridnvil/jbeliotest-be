package dto

import "time"

type InventoryMoveDTO struct {
	ProductID  uint      `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Note       string    `json:"note"`
	OrderID    string    `json:"order_id"`
	MovementAt time.Time `json:"movement_at"`
	OrderDate  time.Time `json:"order_date"`
}
