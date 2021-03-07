package handlers

import (
	"github.com/AdiSaripuloh/online-store/resolvers"
	"github.com/jinzhu/gorm"
	"sync"
)

type Handler struct {
	UserHandler    *UserHandler
	ProductHandler *ProductHandler
	CartHandler    *CartHandler
	OrderHandler   *OrderHandler
}

var (
	handlerLock sync.Once
	h           *Handler
)

func NewHandler(db *gorm.DB) *Handler {
	// User
	userResolver := resolvers.NewUserResolver(db)
	userHandler := NewUserHandler(userResolver)
	// Product
	productResolver := resolvers.NewProductResolver(db)
	productHandler := NewProductHandler(productResolver)
	// Cart
	cartResolver := resolvers.NewCartResolver(db)
	cartHandler := NewCartHandler(cartResolver)
	// Order
	orderResolver := resolvers.NewOrderResolver(db)
	orderHandler := NewOrderHandler(orderResolver)

	handlerLock.Do(func() {
		h = &Handler{
			UserHandler:    userHandler,
			ProductHandler: productHandler,
			CartHandler:    cartHandler,
			OrderHandler:   orderHandler,
		}
	})

	return h
}
