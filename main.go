package main

import (
	"flag"
	"github.com/AdiSaripuloh/online-store/config"
	"github.com/AdiSaripuloh/online-store/database"
	"github.com/AdiSaripuloh/online-store/handlers"
	"github.com/AdiSaripuloh/online-store/resolvers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Handler struct {
	userHandler    *handlers.UserHandler
	productHandler *handlers.ProductHandler
}

var (
	port    string
	handler Handler
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	migration := flag.Bool("migration", false, "Migrate Database")
	seed := flag.Bool("seed", false, "Seed Database")
	flag.Parse()

	dbConfig := config.BuildDbConfig()
	database.Connect(dbConfig)
	if *migration {
		database.Migration()
	}
	if *seed {
		database.Seed()
	}

	gin.SetMode(os.Getenv("APP_MODE"))

	port = os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	// User
	userResolver := resolvers.NewUserResolver(database.Mysql)
	handler.userHandler = handlers.NewUserHandler(userResolver)
	// Product
	productResolver := resolvers.NewProductResolver(database.Mysql)
	handler.productHandler = handlers.NewProductHandler(productResolver)
}

func main() {
	defer database.Mysql.Close()

	router := gin.Default()

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
			v1.GET("/users", handler.userHandler.GetAll)
			v1.GET("/products", handler.productHandler.GetAll)
		}
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
