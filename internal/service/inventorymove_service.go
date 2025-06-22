package service

import (
	"jubeliotesting/internal/domain"
	"jubeliotesting/internal/dto"
)

type InventoryMoveService struct {
	InventoryMoveRepo domain.InventoryMovementRepository
}

func NewInventoryMoveService(inventoryMoveRepo domain.InventoryMovementRepository) *InventoryMoveService {
	return &InventoryMoveService{InventoryMoveRepo: inventoryMoveRepo}
}

func (r *InventoryMoveService) CreateInventoryMove(inventoryDto dto.InventoryMoveDTO) error {
	var inventoryMove domain.InventoryMovement
	inventoryMove.ProductID = inventoryDto.ProductID
	inventoryMove.OrderID = inventoryDto.OrderID
	inventoryMove.OrderDate = inventoryDto.OrderDate
	inventoryMove.MovementAt = inventoryDto.MovementAt
	inventoryMove.Quantity = inventoryDto.Quantity
	inventoryMove.Note = inventoryDto.Note

	return r.InventoryMoveRepo.Create(inventoryMove)
}

func (r *InventoryMoveService) TruncateInventory() {
	r.InventoryMoveRepo.Truncate()
}
