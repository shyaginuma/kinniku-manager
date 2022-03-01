package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConnection() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "@tcp(db:3306)"
	database := os.Getenv("MYSQL_DATABASE")

	dbconf := user + ":" + password + PROTOCOL + "/" + database

	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		return nil, fmt.Errorf("failed to convert rows to Trainingexercise: %v", err)
	}
	return db, nil
}
