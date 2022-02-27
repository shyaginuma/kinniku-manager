package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConnection() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@/%s",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_DATABASE"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to convert rows to Trainingexercise: %v", err)
	}
	return db, nil
}
