package services

import (
	"database/sql"
	"fmt"
	"os"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() (*sql.DB, error) {
	if dbUser == "" {
		log.Println("DB_USER is not set")
	}
	if dbPassword == "" {
		log.Println("DB_PASSWORD is not set")
	}
	if dbDatabase == "" {
		log.Println("DB_NAME is not set")
	}

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
