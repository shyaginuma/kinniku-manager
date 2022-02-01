package repository

import (
	"database/sql"
	"fmt"

	"github.com/kinniku-manager/model"
)

type TrainingExcerciseRepository struct {
	db *sql.DB
}

func (repository TrainingExcerciseRepository) ReadAll() ([]model.TrainingExcercise, error) {
	rows, err := repository.db.Query("SELECT * FROM training_excercises")
	if err != nil {
		return nil, fmt.Errorf("failed to read training excercises: %v", err)
	}
	defer rows.Close()

	var excercises []model.TrainingExcercise
	for rows.Next() {
		var excercise model.TrainingExcercise
		err := rows.Scan(
			&excercise.ID,
			&excercise.Name,
			&excercise.Description,
			&excercise.Target,
			&excercise.Category,
			&excercise.Difficulty,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to convert rows to TrainingExcercise: %v", err)
		}
		excercises = append(excercises, excercise)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	return excercises, nil
}
