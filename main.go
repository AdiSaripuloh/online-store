package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func init() {
	// TODO create resolver
}

func main() {
	port := os.Getenv("APP_PORT")
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("/products", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "List Products",
				})
			})
		}
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
