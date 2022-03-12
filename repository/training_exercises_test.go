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
	sample_data := model.TrainingExercise{
		ID:          1,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Target:      model.Biceps,
		Category:    model.Barbell,
		Difficulty:  model.Beginner,
	}
	stmt, err := db.Prepare("INSERT INTO training_exercises VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		t.Error(err.Error())
	}
	if _, err := stmt.Exec(
		sample_data.ID,
		sample_data.Name,
		sample_data.Description,
		sample_data.Target,
		sample_data.Category,
		sample_data.Difficulty,
	); err != nil {
		t.Error(err.Error())
	}

	// test
	repository := &TrainingExerciseRepository{Database: db}
	exercises, err := repository.ReadAll()
	if err != nil {
		t.Error(err.Error())
	}
	expected_response := []model.TrainingExercise{}
	expected_response = append(expected_response, sample_data)
	assert.Equal(t, expected_response, exercises)
}

func TestTrainingExerciseRepository_Read(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_data := model.TrainingExercise{
		ID:          10,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Target:      model.Biceps,
		Category:    model.Barbell,
		Difficulty:  model.Beginner,
	}
	stmt, err := db.Prepare("INSERT INTO training_exercises VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		t.Error(err.Error())
	}
	if _, err := stmt.Exec(
		sample_data.ID,
		sample_data.Name,
		sample_data.Description,
		sample_data.Target,
		sample_data.Category,
		sample_data.Difficulty,
	); err != nil {
		t.Error(err.Error())
	}

	// test
	repository := &TrainingExerciseRepository{Database: db}
	exercise, err := repository.Read(sample_data.ID)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, sample_data, exercise)
}
func TestTrainingExerciseRepository_Create(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_data := model.TrainingExercise{
		ID:          0,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Target:      model.Biceps,
		Category:    model.Barbell,
		Difficulty:  model.Beginner,
	}

	// test
	repository := &TrainingExerciseRepository{Database: db}
	if err := repository.Create(sample_data); err != nil {
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
	sample_data := model.TrainingExercise{
		ID:          100,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Target:      model.Biceps,
		Category:    model.Barbell,
		Difficulty:  model.Beginner,
	}
	stmt, err := db.Prepare("INSERT INTO training_exercises VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		t.Error(err.Error())
	}
	if _, err := stmt.Exec(
		sample_data.ID,
		sample_data.Name,
		sample_data.Description,
		sample_data.Target,
		sample_data.Category,
		sample_data.Difficulty,
	); err != nil {
		t.Error(err.Error())
	}

	// test
	repository := &TrainingExerciseRepository{Database: db}
	sample_data.Difficulty = model.Intermediate
	if err := repository.Update(sample_data); err != nil {
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
	sample_data := model.TrainingExercise{
		ID:          1000,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Target:      model.Biceps,
		Category:    model.Barbell,
		Difficulty:  model.Beginner,
	}
	stmt, err := db.Prepare("INSERT INTO training_exercises VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		t.Error(err.Error())
	}
	if _, err := stmt.Exec(
		sample_data.ID,
		sample_data.Name,
		sample_data.Description,
		sample_data.Target,
		sample_data.Category,
		sample_data.Difficulty,
	); err != nil {
		t.Error(err.Error())
	}

	// test
	repository := &TrainingExerciseRepository{Database: db}
	if err := repository.Delete(sample_data.ID); err != nil {
		t.Error(err.Error())
	}
}
