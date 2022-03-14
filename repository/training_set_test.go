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
	stmt, err := db.Prepare("INSERT INTO training_sets VALUES(?, ?, ?, ?, ?, ?, ?)")
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
	sets, err := repository.ReadAll()
	if err != nil {
		t.Error(err.Error())
	}
	expected_response := []model.TrainingSet{}
	expected_response = append(expected_response, sample_data)
	assert.Equal(t, expected_response, sets)
}

func TestTrainingSetRepository_Read(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_data := model.TrainingSet{
		ID:          10,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		ExerciseID:  1,
		Reps:        12,
		Weight:      15,
		Interval:    3,
	}
	stmt, err := db.Prepare("INSERT INTO training_sets VALUES(?, ?, ?, ?, ?, ?, ?)")
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
	set, err := repository.Read(sample_data.ID)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, sample_data, set)
}

func TestTrainingSetRepository_Create(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_data := model.TrainingSet{
		ID:          0,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		ExerciseID:  1,
		Reps:        12,
		Weight:      15,
		Interval:    3,
	}

	// test
	repository := &TrainingSetRepository{Database: db}
	if err := repository.Create(sample_data); err != nil {
		t.Error(err.Error())
	}
}

func TestTrainingSetRepository_Update(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_data := model.TrainingSet{
		ID:          100,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		ExerciseID:  1,
		Reps:        12,
		Weight:      15,
		Interval:    3,
	}
	stmt, err := db.Prepare("INSERT INTO training_sets VALUES(?, ?, ?, ?, ?, ?, ?)")
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
	sample_data.Weight = 20
	if err := repository.Update(sample_data); err != nil {
		t.Error(err.Error())
	}
}

func TestTrainingSetRepository_Delete(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_data := model.TrainingSet{
		ID:          1000,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		ExerciseID:  1,
		Reps:        12,
		Weight:      15,
		Interval:    3,
	}
	stmt, err := db.Prepare("INSERT INTO training_sets VALUES(?, ?, ?, ?, ?, ?, ?)")
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
	if err := repository.Delete(sample_data.ID); err != nil {
		t.Error(err.Error())
	}
}
