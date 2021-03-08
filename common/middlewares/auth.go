package middlewares

import (
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"regexp"
	"strings"
)

// Auth
// Because we are not using JWT, we are use UserID in header Authorization
// Check request has user id or not
// Authorization: Bearer {userID}
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		r, _ := regexp.Compile("^Bearer (.+)$")

		match := r.FindStringSubmatch(authHeader)
		unauthorized := responses.Unauthorized("Unauthorized", nil)
		responseError := gin.H{
			"status":  unauthorized.GetType(),
			"message": unauthorized.GetMessage(),
			"data":    unauthorized.GetData(),
		}

		if len(match) == 0 {
			ctx.AbortWithStatusJSON(unauthorized.GetStatusCode(), responseError)
		}
		tokenString := match[1]

		if len(tokenString) == 0 {
			ctx.AbortWithStatusJSON(unauthorized.GetStatusCode(), responseError)
		}

		auth := strings.Split(match[0], " ")
		if len(auth) < 2 {
			ctx.AbortWithStatusJSON(unauthorized.GetStatusCode(), responseError)
		}
		userID := auth[1]

		_, err := uuid.FromString(userID)
		if err != nil {
			ctx.AbortWithStatusJSON(unauthorized.GetStatusCode(), responseError)
		}

		ctx.Set("UserID", userID)
		ctx.Next()
	}
}
