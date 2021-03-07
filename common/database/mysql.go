package database

import (
	"github.com/AdiSaripuloh/online-store/config"
	loggerModels "github.com/AdiSaripuloh/online-store/modules/logger/models"
	productModels "github.com/AdiSaripuloh/online-store/modules/product/models"
	userModels "github.com/AdiSaripuloh/online-store/modules/user/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

func NewConnection() *gorm.DB {
	conf := config.BuildDbConfig()
	dbUrl := config.BuildDbUrl(conf)
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Panic(err.Error())
	}
	if config.HttpConfig.HttpDebug {
		db.LogMode(true)
	}
	log.Println("Connection Established")
	if *config.HttpConfig.DBMigration {
		Migration(db)
	}
	if *config.HttpConfig.DBSeed {
		Seed(db)
	}
	return db
}

func Migration(db *gorm.DB) {
	if !db.HasTable(userModels.User.TableName) {
		db.AutoMigrate(&userModels.User{})
	}
	if !db.HasTable(productModels.Product.TableName) {
		db.AutoMigrate(&productModels.Product{})
	}
	if !db.HasTable(productModels.Cart.TableName) {
		db.AutoMigrate(&productModels.Cart{})
		db.Model(&productModels.Cart{}).AddForeignKey("userID", "users (id)", "CASCADE", "CASCADE")
	}
	if !db.HasTable(productModels.CartItem.TableName) {
		db.AutoMigrate(&productModels.CartItem{})
		db.Model(&productModels.CartItem{}).AddForeignKey("cartID", "carts (id)", "CASCADE", "CASCADE")
		db.Model(&productModels.CartItem{}).AddForeignKey("productID", "products (id)", "CASCADE", "CASCADE")
	}
	if !db.HasTable(productModels.Order.TableName) {
		db.AutoMigrate(&productModels.Order{})
		db.Model(&productModels.Order{}).AddForeignKey("userID", "users (id)", "CASCADE", "CASCADE")
	}
	if !db.HasTable(productModels.OrderItem.TableName) {
		db.AutoMigrate(&productModels.OrderItem{})
		db.Model(&productModels.OrderItem{}).AddForeignKey("orderID", "orders (id)", "CASCADE", "CASCADE")
		db.Model(&productModels.OrderItem{}).AddForeignKey("productID", "products (id)", "CASCADE", "CASCADE")
	}
	if !db.HasTable(loggerModels.Logs.TableName) {
		db.AutoMigrate(&loggerModels.Logs{})
	}
}

func Seed(db *gorm.DB) {
	// User
	db.Create(&userModels.User{
		FullName: "Adi Saripuloh",
		Phone:    "1234567890",
		Email:    "adisaripuloh@gmail.com",
	})
	// Products
	for i := 1; i <= 20; i++ {
		go db.Create(&productModels.Product{
			Name:     "Product " + strconv.Itoa(i),
			Price:    10.00,
			Quantity: 10,
		})
	}
}
