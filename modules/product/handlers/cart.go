package handlers

import (
	"github.com/AdiSaripuloh/online-store/common/resolvers"
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/modules/product/requests"
	"github.com/gin-gonic/gin"
	"net/http"
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
		ctx.AbortWithStatusJSON(http.StatusNotFound, mappers.ResponseFailed("Cart not found!"))
		return
	}

	ctx.JSON(http.StatusOK, mappers.ResponseSuccess(cart))
}

func (uh *CartHandler) Store(ctx *gin.Context) {
	var request requests.CreateCart
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, mappers.ResponseFailed(err.Error()))
		return
	}

	userID := ctx.GetString("UserID")
	cart, err := uh.resolver.CartService.Create(userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, mappers.ResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, mappers.ResponseSuccess(cart))
}

func (uh *CartHandler) Checkout(ctx *gin.Context) {
	var request requests.Checkout
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, mappers.ResponseFailed(err.Error()))
		return
	}

	userID := ctx.GetString("UserID")
	cart, err := uh.resolver.CartService.Checkout(userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, mappers.ResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, mappers.ResponseSuccess(cart))
}
