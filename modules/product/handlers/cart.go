package handlers

import (
	"github.com/AdiSaripuloh/online-store/common/resolvers/mysql"
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	resolver *resolvers.CartResolver
}

func NewCartHandler(resolver *resolvers.CartResolver) *CartHandler {
	handler := &CartHandler{
		resolver: resolver,
	}

	return handler
}

func (uh *CartHandler) Index(ctx *gin.Context) {
	userID := ctx.GetString("UserID")
	cart, err := uh.resolver.CartService.GetByUserID(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	response := responses.Success(cart)
	ctx.JSON(response.GetStatusCode(), gin.H{
		"status": response.GetType(),
		"data":   response.GetData(),
	})
}

func (uh *CartHandler) Store(ctx *gin.Context) {
	var request dto.CreateCartRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		err := responses.UnprocessableEntity(err.Error(), nil)
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	userID := ctx.GetString("UserID")
	cart, err := uh.resolver.CartService.Create(userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	response := responses.Created(cart)
	ctx.JSON(response.GetStatusCode(), gin.H{
		"status": response.GetType(),
		"data":   response.GetData(),
	})
}

func (uh *CartHandler) Checkout(ctx *gin.Context) {
	var request dto.CheckoutCartRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		err := responses.UnprocessableEntity(err.Error(), nil)
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	userID := ctx.GetString("UserID")
	cart, err := uh.resolver.CartService.Checkout(userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	response := responses.Success(cart)
	ctx.JSON(response.GetStatusCode(), gin.H{
		"status": response.GetType(),
		"data":   response.GetData(),
	})
}
