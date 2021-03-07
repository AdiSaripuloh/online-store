package mappers

import (
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/AdiSaripuloh/online-store/modules/product/models"
)

func ProductsResponse(products []models.Product) []dto.Product {
	var response []dto.Product
	for _, product := range products {
		response = append(response, dto.Product{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return response
}
