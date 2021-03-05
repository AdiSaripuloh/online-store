package database

import (
	"database/sql"
	"github.com/AdiSaripuloh/online-store/config"
	"log"
)

var db *sql.DB

func Connect(conf *config.DBConfig) {
	DBUrl := config.BuildDBUrl(conf)
	con, err := sql.Open("mysql", DBUrl)
	if err != nil {
		log.Panic(err.Error())
	}
	db = con
	log.Println("Connection Established")
}

func Close() {
	db.Close()
}
