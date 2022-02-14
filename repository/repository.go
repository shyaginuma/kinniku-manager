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
			os.Getenv("DBUSER"),
			os.Getenv("DBPASS"),
			os.Getenv("DBNAME"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to convert rows to Trainingexercise: %v", err)
	}
	return db, nil
}
