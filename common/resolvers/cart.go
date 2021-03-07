package resolvers

import (
	"github.com/AdiSaripuloh/online-store/modules/product/repositories/mysql"
	"github.com/AdiSaripuloh/online-store/modules/product/services"
	"github.com/jinzhu/gorm"
)

type CartResolver struct {
	CartService services.CartService
}

func NewCartResolver(db *gorm.DB) *CartResolver {
	cartRepository := mysql.NewCartRepository(db)
	cartItemRepository := mysql.NewCartItemRepository(db)
	productRepository := mysql.NewProductRepository(db)
	orderRepository := mysql.NewOrderRepository(db)
	cartService := services.NewCartService(cartRepository, cartItemRepository, productRepository, orderRepository)
	return &CartResolver{
		CartService: cartService,
	}
}
