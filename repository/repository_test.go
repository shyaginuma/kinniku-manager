package repository

import (
	"os"
	"testing"

	"github.com/kinniku-manager/model"
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

func GetSampleSet() model.TrainingSet {
	return model.TrainingSet{
		ID:          1,
		Name:        "Barbell Curl",
		Description: "Barbell Curl",
		ExerciseID:  1,
		Reps:        12,
		Weight:      15,
		Interval:    3,
	}
}

func TestMain(m *testing.M) {
	db, err := NewDBConnection()
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()

	// insert sample training exercise
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

	// insert sample training set
	sample_set := GetSampleSet()
	stmt, err = db.Prepare("INSERT INTO training_sets VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		os.Exit(1)
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
		os.Exit(1)
	}

	status := m.Run()
	os.Exit(status)
}
