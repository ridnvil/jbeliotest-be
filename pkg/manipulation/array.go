package manipulation

import "jubeliotesting/internal/dto"

func AddUniqueCategory(data []dto.CategoryDTO, value string) []dto.CategoryDTO {
	for _, v := range data {
		if v.Name == value {
			return data
		}
	}
	return append(data, dto.CategoryDTO{Name: value})
}

func AddUniqueSubCategory(data []dto.SubCategoryDTO, value string, category string) []dto.SubCategoryDTO {
	for _, v := range data {
		if v.Name == value {
			v.CategoryName = category
			return data
		}
	}
	return append(data, dto.SubCategoryDTO{Name: value, CategoryName: category})
}

func AddUniqueProduct(data []dto.ProductDTO, productName string, subcategory string, manufacture string) []dto.ProductDTO {
	for _, v := range data {
		if v.Name == productName {
			v.SubCategoryName = subcategory
			v.Manufacturer = manufacture
			return data
		}
	}
	return append(data, dto.ProductDTO{Name: productName, Manufacturer: manufacture, SubCategoryName: subcategory})
}
