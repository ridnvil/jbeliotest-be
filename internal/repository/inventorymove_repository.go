package repository

import (
	"gorm.io/gorm"
	"jubeliotesting/internal/domain"
	"log"
)

type InventoryMoveRepository struct {
	DB *gorm.DB
}

func (r *InventoryMoveRepository) Truncate() {
	if err := r.DB.Exec("TRUNCATE TABLE inventory_movements RESTART IDENTITY CASCADE").Error; err != nil {
		log.Fatal(err)
	}
}

func NewInventoryMoveRepository(db *gorm.DB) domain.InventoryMovementRepository {
	return &InventoryMoveRepository{DB: db}
}

func (r *InventoryMoveRepository) Create(inventory domain.InventoryMovement) error {
	if err := r.DB.Create(&inventory).Error; err != nil {
		return err
	}
	return nil
}

func (r *InventoryMoveRepository) FindAll() ([]domain.InventoryMovement, error) {
	var inventoryMovements []domain.InventoryMovement
	if err := r.DB.Find(&inventoryMovements).Error; err != nil {
		return inventoryMovements, err
	}
	return inventoryMovements, nil
}

func (r *InventoryMoveRepository) FindByID(id uint) (*domain.InventoryMovement, error) {
	//TODO implement me
	panic("implement me")
}
