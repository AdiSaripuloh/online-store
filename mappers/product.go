package mappers

import (
	"github.com/AdiSaripuloh/online-store/dto"
	"github.com/AdiSaripuloh/online-store/models"
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
