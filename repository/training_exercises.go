package repository

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/kinniku-manager/model"
)

type TrainingExerciseRepository struct {
	Database *sql.DB
}

type TrainingExerciseSearchOptions struct {
	searchLimit        int
	searchKeyword      []string
	targetMuscle       model.TargetMuscle
	trainingCategory   model.TrainingCategory
	trainingDifficulty model.TrainingDifficulty
}

type TrainingExerciseSearchOption func(*TrainingExerciseSearchOptions)

func WithSearchLimit(searchLimit int) TrainingExerciseSearchOption {
	return func(options *TrainingExerciseSearchOptions) {
		options.searchLimit = searchLimit
	}
}

func WithSearchKeyword(searchKeyword []string) TrainingExerciseSearchOption {
	return func(options *TrainingExerciseSearchOptions) {
		options.searchKeyword = searchKeyword
	}
}

func WithTargetMuscle(targetMuscle model.TargetMuscle) TrainingExerciseSearchOption {
	return func(options *TrainingExerciseSearchOptions) {
		options.targetMuscle = targetMuscle
	}
}

func WithTrainingCategory(trainingCategory model.TrainingCategory) TrainingExerciseSearchOption {
	return func(options *TrainingExerciseSearchOptions) {
		options.trainingCategory = trainingCategory
	}
}

func WithTrainingDifficulty(trainingDifficulty model.TrainingDifficulty) TrainingExerciseSearchOption {
	return func(options *TrainingExerciseSearchOptions) {
		options.trainingDifficulty = trainingDifficulty
	}
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

func (repository TrainingExerciseRepository) Search(options ...TrainingExerciseSearchOption) ([]model.TrainingExercise, error) {
	search_options := TrainingExerciseSearchOptions{
		searchLimit:        10,
		searchKeyword:      []string{""},
		targetMuscle:       "",
		trainingCategory:   "",
		trainingDifficulty: "",
	}
	for _, fn := range options {
		fn(&search_options)
	}

	base_query := "SELECT * FROM training_exercises WHERE 1=1"
	if !reflect.DeepEqual(search_options.searchKeyword, []string{""}) {
		for _, keyword := range search_options.searchKeyword {
			base_query += fmt.Sprintf(" AND (name LIKE '%%%s%%' or description LIKE '%%%s%%')", keyword, keyword)
		}
	}

	if len(search_options.targetMuscle) > 0 {
		base_query += fmt.Sprintf(" AND target = '%s'", search_options.targetMuscle)
	}

	if len(search_options.trainingCategory) > 0 {
		base_query += fmt.Sprintf(" AND category = '%s'", search_options.trainingCategory)
	}

	if len(search_options.trainingDifficulty) > 0 {
		base_query += fmt.Sprintf(" AND difficulty = '%s'", search_options.trainingDifficulty)
	}
	base_query += fmt.Sprintf(" LIMIT %d;", search_options.searchLimit)
	fmt.Println(base_query)

	rows, err := repository.Database.Query(base_query)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		exercises = append(exercises, exercise)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return exercises, nil
}
