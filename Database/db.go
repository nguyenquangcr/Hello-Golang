package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Kết nối thành công đến cơ sở dữ liệu")
}
