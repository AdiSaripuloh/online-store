package http

import (
	"github.com/AdiSaripuloh/online-store/database"
	"github.com/AdiSaripuloh/online-store/middlewares"
	"github.com/AdiSaripuloh/online-store/modules/product/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var (
	port string
)

func Http() {
	port = os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	migration := false
	seed := false
	database.InitDatabase(&migration, &seed)
	defer database.Mysql.Close()

	handler := handlers.NewHandler(database.Mysql)

	router := gin.Default()
	gin.SetMode(os.Getenv("APP_MODE"))

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
			v1.GET("/product", handler.ProductHandler.GetAll)
			v1.Use(middlewares.Auth()).GET("/cart", handler.CartHandler.Get)
			v1.Use(middlewares.Auth()).POST("/cart", handler.CartHandler.Create)
			v1.Use(middlewares.Auth()).POST("/cart/checkout", handler.CartHandler.Checkout)
			v1.Use(middlewares.Auth()).GET("/order", handler.OrderHandler.Get)
			v1.Use(middlewares.Auth()).GET("/order/:orderID", handler.OrderHandler.Show)
			v1.Use(middlewares.Auth()).POST("/order/:orderID/pay", handler.OrderHandler.Pay)
		}
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
