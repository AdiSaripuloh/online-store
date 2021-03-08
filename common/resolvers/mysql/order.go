package resolvers

import (
	"github.com/AdiSaripuloh/online-store/modules/product/repositories/mysql"
	"github.com/AdiSaripuloh/online-store/modules/product/services"
	"github.com/jinzhu/gorm"
)

type OrderResolver struct {
	OrderService services.OrderService
}

func NewOrderResolver(db *gorm.DB) *OrderResolver {
	orderRepository := mysql.NewOrderRepository(db)
	productRepository := mysql.NewProductRepository(db)
	orderService := services.NewOrderService(orderRepository, productRepository)
	return &OrderResolver{
		OrderService: orderService,
	}
}
