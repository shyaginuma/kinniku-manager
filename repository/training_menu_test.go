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
	sample_data := model.TrainingMenu{
		ID:          1,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Menu:        []int64{1},
	}

	// insert to table training_menus
	stmt_menus, err := db.Prepare("INSERT INTO training_menus VALUES(?, ?, ?)")
	if err != nil {
		t.Error(err.Error())
	}
	if _, err := stmt_menus.Exec(
		sample_data.ID,
		sample_data.Name,
		sample_data.Description,
	); err != nil {
		t.Error(err.Error())
	}

	// insert to table training_menus
	stmt_relations, err := db.Prepare("INSERT INTO training_menu_set_relations VALUES(?, ?, ?)")
	if err != nil {
		t.Error(err.Error())
	}
	if _, err := stmt_relations.Exec(
		0,
		sample_data.ID,
		sample_data.Menu[0],
	); err != nil {
		t.Error(err.Error())
	}

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
