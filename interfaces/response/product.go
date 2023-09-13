package response

import "backend/domain/entity"

type Product struct {
	ID                string `json:"id"`
	ProductCategoryID string `json:"product_category_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Price             int64  `json:"price"`
	Stock             int64  `json:"stock"`
}

func ListProducts(e []entity.Product) []Product {
	var products []Product

	for _, productEntity := range e {
		product := Product{
			ID:                productEntity.ID,
			ProductCategoryID: productEntity.ProductCategoryID,
			Name:              productEntity.Name,
			Description:       productEntity.Description,
			Price:             productEntity.Price,
			Stock:             productEntity.Stock,
		}
		products = append(products, product)
	}

	return products
}
