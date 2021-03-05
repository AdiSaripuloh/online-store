package mapper

import "github.com/gin-gonic/gin"

func success(data interface{}) gin.H {
	return gin.H{
		"status": "SUCCESS",
		"data":   data,
	}
}

func failed(data interface{}) gin.H {
	return gin.H{
		"status":  "FAILED",
		"message": data,
	}
}

func errors(data interface{}) gin.H {
	return gin.H{
		"status":  "ERROR",
		"message": data,
	}
}
