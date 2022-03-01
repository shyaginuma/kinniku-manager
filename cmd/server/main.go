package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kinniku-manager/model"
	"github.com/kinniku-manager/repository"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.GET("/training_exercise/read_all", readAllTrainingExercises)
	r.POST("/training_exercise/save", createTrainingExercise)
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
