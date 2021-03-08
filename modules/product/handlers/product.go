package handlers

import (
	"github.com/AdiSaripuloh/online-store/common/resolvers/mysql"
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	resolver *resolvers.ProductResolver
}

func NewProductHandler(resolver *resolvers.ProductResolver) *ProductHandler {
	handler := &ProductHandler{
		resolver: resolver,
	}

	return handler
}

func (uh *ProductHandler) Index(ctx *gin.Context) {
	products, err := uh.resolver.ProductService.All()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, mappers.ResponseErr(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, mappers.ResponseSuccess(products))
}
