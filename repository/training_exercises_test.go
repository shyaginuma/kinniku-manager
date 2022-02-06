package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestTrainingExcerciseRepository_ReadAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	returnRows := sqlmock.NewRows([]string{"id", "name", "description", "target", "category", "difficulty"})
	returnRows.AddRow("1", "Barbell Curl", "Barbell Curl", "biceps", "barbell", "beginner")
	mock.ExpectQuery("select * from training_exercises").WillReturnRows(returnRows)

	repository := &TrainingExcerciseRepository{Database: db}
}
