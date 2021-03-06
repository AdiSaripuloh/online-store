package mappers

import "github.com/gin-gonic/gin"

func ResponseSuccess(data interface{}) gin.H {
	return gin.H{
		"status": "SUCCESS",
		"data":   data,
	}
}

func ResponseFailed(data interface{}) gin.H {
	return gin.H{
		"status":  "FAILED",
		"message": data,
	}
}

func ResponseErr(data interface{}) gin.H {
	return gin.H{
		"status":  "ERROR",
		"message": data,
	}
}
