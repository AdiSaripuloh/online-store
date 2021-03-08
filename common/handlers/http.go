package handlers

import (
	"github.com/AdiSaripuloh/online-store/common/database"
	resolvers "github.com/AdiSaripuloh/online-store/common/resolvers/mysql"
	"github.com/AdiSaripuloh/online-store/config"
	handlersProduct "github.com/AdiSaripuloh/online-store/modules/product/handlers"
	handlersUser "github.com/AdiSaripuloh/online-store/modules/user/handlers"
	"sync"
)

type HttpHandler struct {
	UserHandler    *handlersUser.UserHandler
	ProductHandler *handlersProduct.ProductHandler
	CartHandler    *handlersProduct.CartHandler
	OrderHandler   *handlersProduct.OrderHandler
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
		userHandler := handlersUser.NewUserHandler(userResolver)
		// Product
		productResolver := resolvers.NewProductResolver(db)
		productHandler := handlersProduct.NewProductHandler(productResolver)
		// Cart
		cartResolver := resolvers.NewCartResolver(db)
		cartHandler := handlersProduct.NewCartHandler(cartResolver)
		// Order
		orderResolver := resolvers.NewOrderResolver(db)
		orderHandler := handlersProduct.NewOrderHandler(orderResolver)

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
