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
	sample_set_a := GetSampleSet()
	sample_set_b := GetSampleSet()
	sample_set_b.ID = 2

	// test
	repository := &TrainingSetRepository{Database: db}
	sets, err := repository.ReadAll()
	if err != nil {
		t.Error(err.Error())
	}
	expected_response := []model.TrainingSet{}
	expected_response = append(expected_response, sample_set_a)
	expected_response = append(expected_response, sample_set_b)
	assert.Equal(t, expected_response, sets)
}

func TestTrainingSetRepository_Read(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_set := GetSampleSet()

	// test
	repository := &TrainingSetRepository{Database: db}
	set, err := repository.Read(sample_set.ID)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, sample_set, set)
}

func TestTrainingSetRepository_Create(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_set := GetSampleSet()
	sample_set.ID = 10

	// test
	repository := &TrainingSetRepository{Database: db}
	if err := repository.Create(sample_set); err != nil {
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
	sample_set := GetSampleSet()

	// test
	repository := &TrainingSetRepository{Database: db}
	sample_set.Weight = 20
	if err := repository.Update(sample_set); err != nil {
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
	sample_set := GetSampleSet()

	// test
	repository := &TrainingSetRepository{Database: db}
	if err := repository.Delete(sample_set.ID); err != nil {
		t.Error(err.Error())
	}
}
