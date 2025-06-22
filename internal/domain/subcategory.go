package domain

type SubCategory struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100;not null;unique"`
	CategoryID uint
	Category   Category
}

type SubCategoryRepository interface {
	Create(SubCategory) error
	FindAll() ([]SubCategory, error)
	FindByID(id uint) (SubCategory, error)
	FindByName(name string) (SubCategory, error)
}
