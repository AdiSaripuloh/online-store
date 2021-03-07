package handlers

import (
	"github.com/AdiSaripuloh/online-store/common/resolvers"
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/modules/product/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	resolver *resolvers.OrderResolver
}

func NewOrderHandler(resolver *resolvers.OrderResolver) *OrderHandler {
	handler := &OrderHandler{
		resolver: resolver,
	}

	return handler
}

func (uh *OrderHandler) Get(ctx *gin.Context) {
	userID := ctx.GetString("UserID")
	order, err := uh.resolver.OrderService.GetOrderByUserID(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, mappers.ResponseFailed("Order not found!"))
		return
	}

	ctx.JSON(http.StatusOK, mappers.ResponseSuccess(order))
}

func (uh *OrderHandler) Show(ctx *gin.Context) {
	orderId := ctx.Param("orderID")
	order, err := uh.resolver.OrderService.GetOrderByID(orderId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, mappers.ResponseFailed("Order not found!"))
		return
	}

	ctx.JSON(http.StatusOK, mappers.ResponseSuccess(order))
}

func (uh *OrderHandler) Pay(ctx *gin.Context) {
	var request requests.PayOrder
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, mappers.ResponseFailed(err.Error()))
		return
	}

	orderId := ctx.Param("orderID")
	userID := ctx.GetString("UserID")
	order, err := uh.resolver.OrderService.Pay(orderId, userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, mappers.ResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, mappers.ResponseSuccess(order))
}
