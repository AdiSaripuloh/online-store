package main

import (
	"flag"
	"github.com/AdiSaripuloh/online-store/common/handlers"
	"github.com/AdiSaripuloh/online-store/common/middlewares"
	"github.com/AdiSaripuloh/online-store/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	migration := flag.Bool("migration", false, "Migrate Database")
	seed := flag.Bool("seed", false, "Seed Database")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	config.NewHttpConfig(port, migration, seed)
}

func main() {
	router := gin.Default()

	router.Use(middlewares.JSONLogMiddleware())
	router.Use(gin.Recovery())
	router.Use(middlewares.RequestID(middlewares.RequestIDOptions{AllowSetting: true}))
	router.Use(middlewares.CORS(middlewares.CORSOptions{}))

	gin.SetMode(config.HttpConfig.HttpDebugMode)

	handler := handlers.NewHandler()
	if handler == nil {
		log.Fatal("Setup .env properly")
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    "Not Found",
			"message": "Page not found"},
		)
	})

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
			v1.GET("/user", handler.UserHandler.Index)
			v1.GET("/product", handler.ProductHandler.Index)
			v1.Use(middlewares.Auth()).GET("/cart", handler.CartHandler.Index)
			v1.Use(middlewares.Auth()).POST("/cart", handler.CartHandler.Store)
			v1.Use(middlewares.Auth()).POST("/cart/checkout", handler.CartHandler.Checkout)
			v1.Use(middlewares.Auth()).GET("/order", handler.OrderHandler.Index)
			v1.Use(middlewares.Auth()).GET("/order/:orderID", handler.OrderHandler.Show)
			v1.Use(middlewares.Auth()).POST("/order/:orderID/pay", handler.OrderHandler.Pay)
		}
	}

	if err := router.Run(":" + config.HttpConfig.HttpPort); err != nil {
		log.Fatal(err)
	}
}
