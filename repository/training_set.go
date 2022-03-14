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
