package main

import (
	"github.com/AdiSaripuloh/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBConfig := config.db.BuildConfig()
	_ := config.db.BuildDbUrl(DBConfig)

	// TODO
	// - create migrations
	// - create resolver
}

func main() {
	port := os.Getenv("APP_PORT")
	router := gin.Default()

	if port == "" {
		port = "8000"
	}

	// Home
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Online Store Challenge",
		})
	})

	// API
	api := router.Group("api")
	{
		// V1
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
