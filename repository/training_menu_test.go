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
	sample_menu := model.TrainingMenu{
		ID:          1,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Menu:        []int64{1100},
	}

	// insert to table training_menus
	stmt_menus, err := db.Prepare("INSERT INTO training_menus VALUES(?, ?, ?)")
	if err != nil {
		t.Error(err.Error())
	}
	if _, err := stmt_menus.Exec(
		sample_menu.ID,
		sample_menu.Name,
		sample_menu.Description,
	); err != nil {
		t.Error(err.Error())
	}

	sample_set := model.TrainingSet{
		ID:          1100,
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
		sample_set.ID,
		sample_set.Name,
		sample_set.Description,
		sample_set.ExerciseID,
		sample_set.Reps,
		sample_set.Weight,
		sample_set.Interval,
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
		sample_menu.ID,
		sample_menu.Menu[0],
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
	expected_response = append(expected_response, sample_menu)
	assert.Equal(t, expected_response, menus)
}

func TestTrainingMenuRepository_Read(t *testing.T) {
	// set up db & sample data
	db, err := NewDBConnection()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()
	sample_menu := model.TrainingMenu{
		ID:          10,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Menu:        []int64{1200},
	}

	// insert to table training_menus
	stmt_menus, err := db.Prepare("INSERT INTO training_menus VALUES(?, ?, ?)")
	if err != nil {
		t.Error(err.Error())
	}
	if _, err := stmt_menus.Exec(
		sample_menu.ID,
		sample_menu.Name,
		sample_menu.Description,
	); err != nil {
		t.Error(err.Error())
	}

	sample_set := model.TrainingSet{
		ID:          1200,
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
		sample_set.ID,
		sample_set.Name,
		sample_set.Description,
		sample_set.ExerciseID,
		sample_set.Reps,
		sample_set.Weight,
		sample_set.Interval,
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
		sample_menu.ID,
		sample_menu.Menu[0],
	); err != nil {
		t.Error(err.Error())
	}

	// test
	repository := &TrainingMenuRepository{Database: db}
	menu, err := repository.Read(sample_menu.ID)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, sample_menu, menu)
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
