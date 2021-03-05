package handlers

import (
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/resolvers"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (uh *UserHandler) GetAll(ctx *gin.Context) {
	users, err := uh.resolver.UserService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, mappers.Err(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, mappers.Success(users))
}
