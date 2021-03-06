package resolvers

import (
	repo "github.com/AdiSaripuloh/online-store/repositories"
	svc "github.com/AdiSaripuloh/online-store/services"
	"github.com/jinzhu/gorm"
)

type ProductResolver struct {
	ProductService svc.IProductService
}

func NewProductResolver(db *gorm.DB) *ProductResolver {
	productRepository := repo.NewProductRepository(db)
	productService := svc.NewProductService(productRepository)
	return &ProductResolver{
		ProductService: productService,
	}
}
