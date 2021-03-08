package handlers

import (
	"github.com/AdiSaripuloh/online-store/common/resolvers/mysql"
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/gin-gonic/gin"
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

func (uh *OrderHandler) Index(ctx *gin.Context) {
	userID := ctx.GetString("UserID")
	order, err := uh.resolver.OrderService.GetOrderByUserID(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	response := responses.Success(order)
	ctx.JSON(response.GetStatusCode(), gin.H{
		"status": response.GetType(),
		"data":   response.GetData(),
	})
}

func (uh *OrderHandler) Show(ctx *gin.Context) {
	orderId := ctx.Param("orderID")
	order, err := uh.resolver.OrderService.GetOrderByID(orderId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	response := responses.Success(order)
	ctx.JSON(response.GetStatusCode(), gin.H{
		"status": response.GetType(),
		"data":   response.GetData(),
	})
}

func (uh *OrderHandler) Pay(ctx *gin.Context) {
	var request dto.PayOrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		err := responses.UnprocessableEntity(err.Error(), nil)
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	orderId := ctx.Param("orderID")
	userID := ctx.GetString("UserID")
	order, err := uh.resolver.OrderService.Pay(orderId, userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	response := responses.Success(order)
	ctx.JSON(response.GetStatusCode(), gin.H{
		"status": response.GetType(),
		"data":   response.GetData(),
	})
}
