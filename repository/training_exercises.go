package repository

import (
	"database/sql"
	"fmt"

	"github.com/kinniku-manager/model"
)

type TrainingExerciseRepository struct {
	Database *sql.DB
}

func (repository TrainingExerciseRepository) ReadAll() ([]model.TrainingExercise, error) {
	rows, err := repository.Database.Query("SELECT * FROM training_exercises")
	if err != nil {
		return nil, fmt.Errorf("failed to read training exercises: %v", err)
	}
	defer rows.Close()

	var exercises []model.TrainingExercise
	for rows.Next() {
		var exercise model.TrainingExercise
		err := rows.Scan(
			&exercise.ID,
			&exercise.Name,
			&exercise.Description,
			&exercise.Target,
			&exercise.Category,
			&exercise.Difficulty,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to convert rows to Trainingexercise: %v", err)
		}
		exercises = append(exercises, exercise)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	return exercises, nil
}
