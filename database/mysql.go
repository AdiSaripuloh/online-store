package database

import (
	"database/sql"
	"github.com/AdiSaripuloh/online-store/config"
	"log"
)

var db *sql.DB

func Connect(DBConfig *config.DBConfig) {
	DBUrl := config.BuildDbUrl(DBConfig)
	con, err := sql.Open("mysql", DBUrl)
	if err != nil {
		panic(err.Error())
	}
	db = con
	log.Println("Connection Established")
}

func Close() {
	db.Close()
}
