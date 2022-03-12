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

func (repository TrainingExerciseRepository) Read(id int64) (model.TrainingExercise, error) {
	row := repository.Database.QueryRow("SELECT * FROM training_exercises where id = ?", id)
	var exercise model.TrainingExercise
	err := row.Scan(
		&exercise.ID,
		&exercise.Name,
		&exercise.Description,
		&exercise.Target,
		&exercise.Category,
		&exercise.Difficulty,
	)
	if err != nil {
		return model.TrainingExercise{}, err
	}
	return exercise, nil
}

func (repository TrainingExerciseRepository) Create(newTrainingExercise model.TrainingExercise) error {
	stmt, err := repository.Database.Prepare("INSERT INTO training_exercises VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(
		newTrainingExercise.ID,
		newTrainingExercise.Name,
		newTrainingExercise.Description,
		newTrainingExercise.Target,
		newTrainingExercise.Category,
		newTrainingExercise.Difficulty,
	); err != nil {
		return err
	}
	return nil
}

func (repository TrainingExerciseRepository) Update(modifiedTrainingExercise model.TrainingExercise) error {
	stmt, err := repository.Database.Prepare(`
		UPDATE training_exercises
		SET name = ?,
			description = ?,
			target = ?,
			category = ?,
			difficulty = ?
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(
		modifiedTrainingExercise.Name,
		modifiedTrainingExercise.Description,
		modifiedTrainingExercise.Target,
		modifiedTrainingExercise.Category,
		modifiedTrainingExercise.Difficulty,
		modifiedTrainingExercise.ID,
	); err != nil {
		return err
	}
	return nil
}

func (repository TrainingExerciseRepository) Delete(trainingExerciseID int64) error {
	stmt, err := repository.Database.Prepare("DELETE FROM training_exercises WHERE id = ?")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(trainingExerciseID); err != nil {
		return err
	}
	return nil
}
