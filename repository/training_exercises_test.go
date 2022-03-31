package repository

import (
	"testing"

	"github.com/kinniku-manager/model"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestTrainingExerciseRepository_ReadAll(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_exercise_a := GetSampleExercise()
	sample_exercise_b := GetSampleExercise()
	sample_exercise_b.ID = 2

	// test
	repository := &TrainingExerciseRepository{Database: db}
	exercises, err := repository.ReadAll()
	if err != nil {
		t.Error(err.Error())
	}
	expected_response := []model.TrainingExercise{}
	expected_response = append(expected_response, sample_exercise_a)
	expected_response = append(expected_response, sample_exercise_b)
	assert.Equal(t, expected_response, exercises)
}

func TestTrainingExerciseRepository_Read(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_exercise := GetSampleExercise()

	// test
	repository := &TrainingExerciseRepository{Database: db}
	exercise, err := repository.Read(sample_exercise.ID)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, sample_exercise, exercise)
}
func TestTrainingExerciseRepository_Create(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_exercise := GetSampleExercise()
	sample_exercise.ID = 10 // Avoid id duplication

	// test
	repository := &TrainingExerciseRepository{Database: db}
	if err := repository.Create(sample_exercise); err != nil {
		t.Error(err.Error())
	}
}

func TestTrainingExerciseRepository_Update(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_exercise := GetSampleExercise()

	// test
	repository := &TrainingExerciseRepository{Database: db}
	sample_exercise.Difficulty = model.Intermediate
	if err := repository.Update(sample_exercise); err != nil {
		t.Error(err.Error())
	}
}

func TestTrainingExerciseRepository_Delete(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_exercise := GetSampleExercise()

	// test
	repository := &TrainingExerciseRepository{Database: db}
	if err := repository.Delete(sample_exercise.ID); err != nil {
		t.Error(err.Error())
	}
}

func TestTrainingExerciseRepository_Search(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_exercise := GetSampleExercise()

	// test
	repository := &TrainingExerciseRepository{Database: db}
	results, err := repository.Search(
		WithSearchLimit(30),
		WithSearchKeyword([]string{"Curl"}),
		WithTargetMuscle(model.Biceps),
		WithTrainingCategory(model.Barbell),
		WithTrainingDifficulty(model.Beginner),
	)
	if err != nil {
		t.Error(err.Error())
	}

	expected_response := []model.TrainingExercise{}
	expected_response = append(expected_response, sample_exercise)
	assert.Equal(t, expected_response, results)
}
