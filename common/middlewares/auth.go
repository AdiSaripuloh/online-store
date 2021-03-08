package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
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

		if len(match) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "FAILED",
				"message": "Unauthorized",
			})
			ctx.Abort()
		}
		tokenString := match[1]

		if len(tokenString) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "FAILED",
				"message": "Unauthorized",
			})
			ctx.Abort()
		}

		auth := strings.Split(match[0], " ")
		if len(auth) < 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "FAILED",
				"message": "Unauthorized",
			})
			ctx.Abort()
		}
		userID := auth[1]

		_, err := uuid.FromString(userID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "FAILED",
				"message": "Unauthorized",
			})
			ctx.Abort()
		}

		ctx.Set("UserID", userID)
		ctx.Next()
	}
}
