package middlewares

import (
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header["Authorization"]
		if len(auth) <= 0 {
			ctx.JSON(http.StatusUnauthorized, mappers.ResponseErr("Unauthorized"))
			ctx.Abort()
		}
		authSlice := strings.Split(auth[0], " ")
		if len(authSlice) < 2 {
			ctx.JSON(http.StatusUnauthorized, mappers.ResponseErr("Unauthorized"))
			ctx.Abort()
		}
		userID := strings.Split(auth[0], " ")[1]
		ctx.Set("UserID", userID)
		ctx.Next()
	}
}
