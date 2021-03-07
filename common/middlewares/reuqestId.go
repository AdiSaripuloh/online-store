package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

type RequestIDOptions struct {
	AllowSetting bool
}

func RequestID(options RequestIDOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string

		if options.AllowSetting {
			// If Set-Request-Id header is set on request, use that for
			// Request-Id response header. Otherwise, generate a new one.
			requestID = c.Request.Header.Get("Set-Request-Id")
		}

		if requestID == "" {
			id, err := uuid.NewV4()
			if err != nil {
				requestID = time.Now().String()
			}
			requestID = id.String()
		}

		c.Writer.Header().Set("RequestId", requestID)
		c.Next()
	}
}
