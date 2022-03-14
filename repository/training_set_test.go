package repository

import (
	"testing"

	"github.com/kinniku-manager/model"
	"github.com/stretchr/testify/assert"
)

func TestTrainingSetRepository_ReadAll(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_data := model.TrainingSet{
		ID:          1,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		ExerciseID:  1,
		Reps:        12,
		Weight:      15,
		Interval:    3,
	}
	stmt, err := db.Prepare("INSERT INTO training_exercises VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		t.Error(err.Error())
	}
	if _, err := stmt.Exec(
		sample_data.ID,
		sample_data.Name,
		sample_data.Description,
		sample_data.ExerciseID,
		sample_data.Reps,
		sample_data.Weight,
		sample_data.Interval,
	); err != nil {
		t.Error(err.Error())
	}

	// test
	repository := &TrainingSetRepository{Database: db}
	exercises, err := repository.ReadAll()
	if err != nil {
		t.Error(err.Error())
	}
	expected_response := []model.TrainingSet{}
	expected_response = append(expected_response, sample_data)
	assert.Equal(t, expected_response, exercises)
}

func TestTrainingSetRepository_Read(t *testing.T) {

}

func TestTrainingSetRepository_Create(t *testing.T) {

}

func TestTrainingSetRepository_Update(t *testing.T) {

}

func TestTrainingSetRepository_Delete(t *testing.T) {

}