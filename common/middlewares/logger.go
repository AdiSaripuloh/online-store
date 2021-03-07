package middlewares

import (
	"github.com/AdiSaripuloh/online-store/common/database"
	"github.com/AdiSaripuloh/online-store/modules/logger/models"
	"github.com/AdiSaripuloh/online-store/modules/logger/repositories/mysql"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func GetDurationInMilliseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}

func GetClientIP(c *gin.Context) string {
	requester := c.Request.Header.Get("X-Forwarded-For")

	if len(requester) == 0 {
		requester = c.Request.Header.Get("X-Real-IP")
	}

	if len(requester) == 0 {
		requester = c.Request.RemoteAddr
	}

	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}

func GetUserID(c *gin.Context) string {
	userID, exists := c.Get("UserID")
	if exists {
		return userID.(string)
	}
	return ""
}

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := GetDurationInMilliseconds(start)

		fields := log.Fields{
			"clientIP":  GetClientIP(c),
			"duration":  duration,
			"method":    c.Request.Method,
			"path":      c.Request.RequestURI,
			"status":    c.Writer.Status(),
			"userId":    GetUserID(c),
			"referrer":  c.Request.Referer(),
			"requestID": c.Writer.Header().Get("RequestID"),
		}
		entry := log.WithFields(fields)

		db := database.NewConnection()
		logger := mysql.NewLoggerRepository(db)
		_, err := logger.Create(&models.Logs{
			ClientIP:  GetClientIP(c),
			Duration:  duration,
			Method:    c.Request.Method,
			Path:      c.Request.RequestURI,
			Status:    c.Writer.Status(),
			UserID:    GetUserID(c),
			Referer:   c.Request.Referer(),
			RequestID: c.Writer.Header().Get("RequestID"),
		})
		if err != nil {
			log.Error(err.Error())
		}

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else if c.Writer.Status() >= 400 {
			entry.Warn("")
		} else {
			entry.Info("")
		}
	}
}
