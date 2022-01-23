package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kinniku-manager/model"
)

type mysqlStore struct {
	db *sql.DB
}

func NewMysqlStore() mysqlStore {
	db, err := sql.Open("mysql", "/trainings")
	if err != nil {
		log.Fatal(err)
	}

	return mysqlStore{
		db: db,
	}
}

func (s *mysqlStore) GetAllTrainingExercises() ([]model.TrainingExcercise, error) {
	var exercises []model.TrainingExcercise

	rows, err := s.db.Query("SELECT * FROM training_exercises")
	if err != nil {
		return nil, fmt.Errorf("Couldn't fetch table")
	}
	defer rows.Close()

	for rows.Next() {
		var exercise model.TrainingExcercise
		if err := rows.Scan(
			&exercise.ID,
			&exercise.Name,
			&exercise.Description,
			&exercise.Target,
			&exercise.Category,
			&exercise.Difficulty,
		); err != nil {
			return nil, fmt.Errorf("Couldn't read data")
		}
		exercises = append(exercises, exercise)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Something wrong.")
	}
	return exercises, nil
}
