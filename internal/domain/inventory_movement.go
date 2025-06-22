package domain

import "time"

type InventoryMovement struct {
	ID         uint `gorm:"primaryKey"`
	ProductID  uint
	OrderID    string `gorm:"size:100"`
	Product    Product
	OrderDate  time.Time
	MovementAt time.Time
	Quantity   int
	Note       string
}

type InventoryMovementRepository interface {
	Create(InventoryMovement) error
	FindAll() ([]InventoryMovement, error)
	FindByID(id uint) (*InventoryMovement, error)
	Truncate()
}
