package domain

type Category struct {
	ID            uint          `gorm:"primaryKey"`
	Name          string        `gorm:"size:100;not null;unique"`
	SubCategories []SubCategory `gorm:"foreignKey:CategoryID"`
}

type CategoryRepository interface {
	Create(Category) error
	FindAll() ([]Category, error)
	FindByID(id uint) (*Category, error)
	FindByName(name string) (*Category, error)
}
