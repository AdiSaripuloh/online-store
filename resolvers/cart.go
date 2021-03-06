package resolvers

import (
	repo "github.com/AdiSaripuloh/online-store/repositories"
	svc "github.com/AdiSaripuloh/online-store/services"
	"github.com/jinzhu/gorm"
)

type CartResolver struct {
	CartService svc.ICartService
}

func NewCartResolver(db *gorm.DB) *CartResolver {
	cartRepository := repo.NewCartRepository(db)
	productRepository := repo.NewProductRepository(db)
	orderRepository := repo.NewOrderRepository(db)
	cartService := svc.NewCartService(cartRepository, productRepository, orderRepository)
	return &CartResolver{
		CartService: cartService,
	}
}
