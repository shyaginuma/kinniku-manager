package repository

import (
	"os"
	"testing"

	"github.com/kinniku-manager/model"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func GetSampleExercise() model.TrainingExercise {
	return model.TrainingExercise{
		ID:          1,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Target:      model.Biceps,
		Category:    model.Barbell,
		Difficulty:  model.Beginner,
	}
}

func TestMain(m *testing.M) {
	db, err := NewDBConnection()
	if err != nil {
		os.Exit(1)
	}
	stmt, err := db.Prepare("INSERT INTO training_exercises VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		os.Exit(1)
	}
	sample_exercise := GetSampleExercise()
	if _, err := stmt.Exec(
		sample_exercise.ID,
		sample_exercise.Name,
		sample_exercise.Description,
		sample_exercise.Target,
		sample_exercise.Category,
		sample_exercise.Difficulty,
	); err != nil {
		os.Exit(1)
	}
	defer db.Close()
	status := m.Run()
	_, err = db.Exec("TRUNCATE training_exercises")
	if err != nil {
		os.Exit(1)
	}
	os.Exit(status)
}

func TestTrainingExerciseRepository_ReadAll(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_exercise := GetSampleExercise()

	// test
	repository := &TrainingExerciseRepository{Database: db}
	exercises, err := repository.ReadAll()
	if err != nil {
		t.Error(err.Error())
	}
	expected_response := []model.TrainingExercise{}
	expected_response = append(expected_response, sample_exercise)
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
