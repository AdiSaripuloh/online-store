package resolvers

import (
	repo "github.com/AdiSaripuloh/online-store/repositories"
	svc "github.com/AdiSaripuloh/online-store/services"
	"github.com/jinzhu/gorm"
)

type OrderResolver struct {
	OrderService svc.IOrderService
}

func NewOrderResolver(db *gorm.DB) *OrderResolver {
	orderRepository := repo.NewOrderRepository(db)
	productRepository := repo.NewProductRepository(db)
	orderService := svc.NewOrderService(orderRepository, productRepository)
	return &OrderResolver{
		OrderService: orderService,
	}
}
