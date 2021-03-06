package database

import (
	"github.com/AdiSaripuloh/online-store/config"
	"github.com/AdiSaripuloh/online-store/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Mysql *gorm.DB

func Connect(conf *config.DBConfig) {
	dbUrl := config.BuildDbUrl(conf)
	con, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Panic(err.Error())
	}
	Mysql = con
	log.Println("Connection Established")
}

func Migration() {
	if !Mysql.HasTable(models.User.TableName) {
		Mysql.AutoMigrate(&models.User{})
	}
	if !Mysql.HasTable(models.Product.TableName) {
		Mysql.AutoMigrate(&models.Product{})
	}
	if !Mysql.HasTable(models.Cart.TableName) {
		Mysql.AutoMigrate(&models.Cart{})
	}
	if !Mysql.HasTable(models.CartItem.TableName) {
		Mysql.AutoMigrate(&models.CartItem{})
	}
	if !Mysql.HasTable(models.Order.TableName) {
		Mysql.AutoMigrate(&models.Order{})
	}
	if !Mysql.HasTable(models.OrderItem.TableName) {
		Mysql.AutoMigrate(&models.OrderItem{})
	}
}

func Seed() {
	Mysql.Create(models.DummyUser())
}
