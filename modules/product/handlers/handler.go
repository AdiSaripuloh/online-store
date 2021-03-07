package handlers

import (
	"github.com/AdiSaripuloh/online-store/common/database"
	"github.com/AdiSaripuloh/online-store/common/resolvers"
	"github.com/AdiSaripuloh/online-store/config"
	"github.com/AdiSaripuloh/online-store/modules/user/handlers"
	"sync"
)

type HttpHandler struct {
	UserHandler    *handlers.UserHandler
	ProductHandler *ProductHandler
	CartHandler    *CartHandler
	OrderHandler   *OrderHandler
}

var (
	handlerLock sync.Once
	h           *HttpHandler
)

func NewHandler() *HttpHandler {
	if config.HttpConfig.DBDriver == config.MYSQL {
		db := database.NewConnection()
		// User
		userResolver := resolvers.NewUserResolver(db)
		userHandler := handlers.NewUserHandler(userResolver)
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
			h = &HttpHandler{
				UserHandler:    userHandler,
				ProductHandler: productHandler,
				CartHandler:    cartHandler,
				OrderHandler:   orderHandler,
			}
		})
		return h
	}
	return nil
}
