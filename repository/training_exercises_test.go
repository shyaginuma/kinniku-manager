package repository

import (
	"regexp"
	"testing"

	"github.com/kinniku-manager/model"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestTrainingExerciseRepository_ReadAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	returnRows := sqlmock.NewRows([]string{"id", "name", "description", "target", "category", "difficulty"})
	returnRows.AddRow("1", "Barbell Curl", "Barbell Curl", "biceps", "barbell", "beginner")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM training_exercises`)).WillReturnRows(returnRows)

	repository := &TrainingExerciseRepository{Database: db}
	exercises, err := repository.ReadAll()
	if err != nil {
		t.Error(err.Error())
	}
	expected_response := []model.TrainingExercise{}
	data := model.TrainingExercise{
		ID:          "1",
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Target:      model.Biceps,
		Category:    model.Barbell,
		Difficulty:  model.Beginner,
	}
	expected_response = append(expected_response, data)
	assert.Equal(t, expected_response, exercises)
}

func TestTrainingExerciseRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	data := model.TrainingExercise{
		ID:          "1",
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		Target:      model.Biceps,
		Category:    model.Barbell,
		Difficulty:  model.Beginner,
	}

	prep := mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO training_exercises VALUES(?, ?, ?, ?, ?, ?)`))
	prep.ExpectExec().
		WithArgs(data.ID, data.Name, data.Description, data.Target, data.Category, data.Difficulty).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := &TrainingExerciseRepository{Database: db}
	if err := repository.Create(data); err != nil {
		t.Error(err.Error())
	}
}
