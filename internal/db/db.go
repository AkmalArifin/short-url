package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	dataSource := DB_USER + ":" + DB_PASSWORD + "@/" + DB_NAME

	var err error
	DB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Panic("could not connect DB")
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
