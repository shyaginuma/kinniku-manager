package repository

import (
	"testing"

	"github.com/kinniku-manager/model"
	"github.com/stretchr/testify/assert"
)

func TestTrainingMenuRepository_ReadAll(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_data := GetSampleMenu()

	// test
	repository := &TrainingMenuRepository{Database: db}
	menus, err := repository.ReadAll()
	if err != nil {
		t.Error(err.Error())
	}
	expected_response := []model.TrainingMenu{}
	expected_response = append(expected_response, sample_data)
	assert.Equal(t, expected_response, menus)
}

func TestTrainingMenuRepository_Read(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_data := GetSampleMenu()

	// test
	repository := &TrainingMenuRepository{Database: db}
	menu, err := repository.Read(sample_data.ID)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, sample_data, menu)
}

func TestTrainingMenuRepository_Create(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
}

func TestTrainingMenuRepository_Update(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
}

func TestTrainingMenuRepository_Delete(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
}
