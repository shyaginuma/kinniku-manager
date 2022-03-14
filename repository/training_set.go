package repository

import (
	"database/sql"

	"github.com/kinniku-manager/model"
)

type TrainingSetRepository struct {
	Database *sql.DB
}

func (repository TrainingSetRepository) ReadAll() ([]model.TrainingSet, error) {
	rows, err := repository.Database.Query("SELECT * FROM training_sets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sets []model.TrainingSet
	for rows.Next() {
		var set model.TrainingSet
		err := rows.Scan(
			&set.ID,
			&set.Name,
			&set.Description,
			&set.ExerciseID,
			&set.Reps,
			&set.Weight,
			&set.Interval,
		)
		if err != nil {
			return nil, err
		}
		sets = append(sets, set)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return sets, nil
}

func (repository TrainingSetRepository) Read(id int64) (model.TrainingSet, error) {
	row := repository.Database.QueryRow("SELECT * FROM training_sets where id = ?", id)
	var set model.TrainingSet
	err := row.Scan(
		&set.ID,
		&set.Name,
		&set.Description,
		&set.ExerciseID,
		&set.Reps,
		&set.Weight,
		&set.Interval,
	)
	if err != nil {
		return model.TrainingSet{}, err
	}
	return set, nil
}

func (repository TrainingSetRepository) Create(newTrainingSet model.TrainingSet) error {
	stmt, err := repository.Database.Prepare("INSERT INTO training_sets VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(
		newTrainingSet.ID,
		newTrainingSet.Name,
		newTrainingSet.Description,
		newTrainingSet.ExerciseID,
		newTrainingSet.Reps,
		newTrainingSet.Weight,
		newTrainingSet.Interval,
	); err != nil {
		return err
	}
	return nil
}

func (repository TrainingSetRepository) Update(modifiedTrainingSet model.TrainingSet) error {
	stmt, err := repository.Database.Prepare(`
		UPDATE training_sets
		SET name = ?,
			description = ?,
			exercise_id = ?,
			reps = ?,
			weight_kg = ?,
			interval_min = ?
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(
		modifiedTrainingSet.Name,
		modifiedTrainingSet.Description,
		modifiedTrainingSet.ExerciseID,
		modifiedTrainingSet.Reps,
		modifiedTrainingSet.Weight,
		modifiedTrainingSet.Interval,
		modifiedTrainingSet.ID,
	); err != nil {
		return err
	}
	return nil
}

func (repository TrainingSetRepository) Delete(trainingSetID int64) error {
	stmt, err := repository.Database.Prepare("DELETE FROM training_sets WHERE id = ?")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(trainingSetID); err != nil {
		return err
	}
	return nil
}
