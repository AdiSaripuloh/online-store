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
	Mysql.AutoMigrate(&models.User{})
	Mysql.AutoMigrate(&models.Cart{})
}

func Seed() {
	Mysql.Create(models.DummyUser())
}
