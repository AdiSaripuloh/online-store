package dto

import (
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Quantity int64     `json:"quantity"`
}

func ProductsResponse(products []*models.Product) []*Product {
	var response []*Product
	for _, product := range products {
		response = append(response, &Product{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return response
}
