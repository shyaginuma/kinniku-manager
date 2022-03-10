package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kinniku-manager/model"
	"github.com/kinniku-manager/repository"
)

func main() {
	r := gin.Default()
	r.GET("/training_exercise/read_all", readAllTrainingExercises)
	r.POST("/training_exercise/save", createTrainingExercise)
	r.PUT("/training_exercise/edit", updateTrainingExercise)
	r.DELETE("/training_exercise/delete/:id", deleteTrainingExercise)
	r.Run()
}

func readAllTrainingExercises(c *gin.Context) {
	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	repository := &repository.TrainingExerciseRepository{Database: db}
	exercises, err := repository.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, exercises)
}

func createTrainingExercise(c *gin.Context) {
	var newTrainingExercise model.TrainingExercise
	if err := c.BindJSON(&newTrainingExercise); err != nil {
		log.Fatalf("failed to bind json: %v", err)
	}

	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatalf("failed to establish connection with db: %v", err)
	}
	repository := &repository.TrainingExerciseRepository{Database: db}
	if err := repository.Create(newTrainingExercise); err != nil {
		log.Fatalf("failed to insert record into db: %v", err)
	}
	c.IndentedJSON(http.StatusCreated, newTrainingExercise)
}

func updateTrainingExercise(c *gin.Context) {
	var updatedTrainingExercise model.TrainingExercise
	if err := c.BindJSON(&updatedTrainingExercise); err != nil {
		log.Fatalf("failed to bind json: %v", err)
	}

	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatalf("failed to establish connection with db: %v", err)
	}
	repository := &repository.TrainingExerciseRepository{Database: db}
	if err := repository.Update(updatedTrainingExercise); err != nil {
		log.Fatalf("failed to update data: %v", err)
	}
	c.IndentedJSON(http.StatusCreated, updatedTrainingExercise)
}

func deleteTrainingExercise(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("failed to read id parameter from path: %v", err)
	}
	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatalf("failed to establish connection with db: %v", err)
	}
	repository := &repository.TrainingExerciseRepository{Database: db}
	if err := repository.Delete(id); err != nil {
		log.Fatalf("failed to delete data: %v", err)
	}
	c.String(http.StatusOK, "successfully delete training exercise.")
}
