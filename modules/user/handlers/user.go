package handlers

import (
	"github.com/AdiSaripuloh/online-store/common/resolvers/mysql"
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	resolver *resolvers.UserResolver
}

func NewUserHandler(resolver *resolvers.UserResolver) *UserHandler {
	handler := &UserHandler{
		resolver: resolver,
	}

	return handler
}

func (uh *UserHandler) Index(ctx *gin.Context) {
	users, err := uh.resolver.UserService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(err.GetStatusCode(), gin.H{
			"status":  err.GetType(),
			"message": err.GetMessage(),
			"data":    err.GetData(),
		})
		return
	}

	response := responses.Success(users)
	ctx.JSON(response.GetStatusCode(), gin.H{
		"status": response.GetType(),
		"data":   response.GetData(),
	})
}
