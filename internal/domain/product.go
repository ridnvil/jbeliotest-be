package domain

type Product struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"type:text;not null"`
	Manufacturer  string `gorm:"size:100"`
	SubCategoryID uint
	SubCategory   SubCategory
}

type ProductRepository interface {
	Create(Product) error
	FindAll() ([]Product, error)
	FindByID(id uint) (*Product, error)
	FindByName(name string) (*Product, error)
}
