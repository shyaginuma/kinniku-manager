package repository

import (
	"database/sql"
	"fmt"
	"os"
)

func NewDBConnection() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@localhost/%s",
			os.Getenv("DBUSER"),
			os.Getenv("DBPASS"),
			os.Getenv("DBNAME"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to convert rows to TrainingExcercise: %v", err)
	}
	return db, nil
}
