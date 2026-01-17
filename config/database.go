package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "user:password@/go_products")
	if err != nil {
		panic(err)
	}

	DB = db
}
