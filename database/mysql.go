package database

import (
	"github.com/AdiSaripuloh/online-store/config"
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	model2 "github.com/AdiSaripuloh/online-store/modules/user/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strconv"
)

var Mysql *gorm.DB

func Connect(conf *config.DBConfig) {
	dbUrl := config.BuildDbUrl(conf)
	con, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Panic(err.Error())
	}
	Mysql = con
	if os.Getenv("APP_MODE") == gin.DebugMode {
		Mysql.LogMode(true)
	}
	log.Println("Connection Established")
}

func Migration() {
	if !Mysql.HasTable(model2.User.TableName) {
		Mysql.AutoMigrate(&model2.User{})
	}
	if !Mysql.HasTable(models.Product.TableName) {
		Mysql.AutoMigrate(&models.Product{})
	}
	if !Mysql.HasTable(models.Cart.TableName) {
		Mysql.AutoMigrate(&models.Cart{})
		Mysql.Model(&models.Cart{}).AddForeignKey("userID", "users (id)", "CASCADE", "CASCADE")
	}
	if !Mysql.HasTable(models.CartItem.TableName) {
		Mysql.AutoMigrate(&models.CartItem{})
		Mysql.Model(&models.CartItem{}).AddForeignKey("cartID", "carts (id)", "CASCADE", "CASCADE")
		Mysql.Model(&models.CartItem{}).AddForeignKey("productID", "products (id)", "CASCADE", "CASCADE")
	}
	if !Mysql.HasTable(models.Order.TableName) {
		Mysql.AutoMigrate(&models.Order{})
		Mysql.Model(&models.Order{}).AddForeignKey("userID", "users (id)", "CASCADE", "CASCADE")
	}
	if !Mysql.HasTable(models.OrderItem.TableName) {
		Mysql.AutoMigrate(&models.OrderItem{})
		Mysql.Model(&models.OrderItem{}).AddForeignKey("orderID", "orders (id)", "CASCADE", "CASCADE")
		Mysql.Model(&models.OrderItem{}).AddForeignKey("productID", "products (id)", "CASCADE", "CASCADE")
	}
}

func Seed() {
	// User
	Mysql.Create(&model2.User{
		FullName: "Adi Saripuloh",
		Phone:    "1234567890",
		Email:    "adisaripuloh@gmail.com",
	})
	// Products
	for i := 1; i <= 20; i++ {
		go Mysql.Create(&models.Product{
			Name:     "Product " + strconv.Itoa(i),
			Price:    10.00,
			Quantity: 10,
		})
	}
}
