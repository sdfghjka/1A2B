package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/golang-migrate/migrate/v4"
	// "github.com/golang-migrate/migrate/v4/database/mysql"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func MysqlDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MYSQL_USER_NAME := os.Getenv("MYSQL_USER_NAME")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_URL := os.Getenv("MYSQL_URL")
	MYSQL_DATABASE_NAME := os.Getenv("MYSQL_DATABASE_NAME")
	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", MYSQL_USER_NAME, MYSQL_PASSWORD, MYSQL_URL, MYSQL_DATABASE_NAME)
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal("MYSQLDB struct error:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("MYSQLDB ping error:", err)
	}
	log.Println("Connected to MYSQL!")
	return db
}
