package resolvers

import (
	"github.com/AdiSaripuloh/online-store/modules/product/repositories/mysql"
	"github.com/AdiSaripuloh/online-store/modules/product/services"
	"github.com/jinzhu/gorm"
)

type ProductResolver struct {
	ProductService services.ProductService
}

func NewProductResolver(db *gorm.DB) *ProductResolver {
	productRepository := mysql.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	return &ProductResolver{
		ProductService: productService,
	}
}
