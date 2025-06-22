package dto

type SubCategoryDTO struct {
	Name         string `json:"name"`
	CategoryName string `json:"category_name"`
	CategoryID   uint   `json:"category_id"`
	Note         string `json:"note"`
}
