package handlers

import (
	"github.com/AdiSaripuloh/online-store/common/resolvers/mysql"
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/gin-gonic/gin"
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
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	response := responses.Success(products)
	ctx.JSON(response.GetStatusCode(), gin.H{
		"status": response.GetType(),
		"data":   response.GetData(),
	})
}
