package dto

type ProductDTO struct {
	Name            string `json:"name"`
	Manufacturer    string `json:"manufacturer"`
	SubCategoryID   uint   `json:"sub_category_id"`
	SubCategoryName string `json:"sub_category_name"`
}
