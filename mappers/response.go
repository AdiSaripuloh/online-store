package mappers

import "github.com/gin-gonic/gin"

func Success(data interface{}) gin.H {
	return gin.H{
		"status": "SUCCESS",
		"data":   data,
	}
}

func Failed(data interface{}) gin.H {
	return gin.H{
		"status":  "FAILED",
		"message": data,
	}
}

func Err(data interface{}) gin.H {
	return gin.H{
		"status":  "ERROR",
		"message": data,
	}
}
