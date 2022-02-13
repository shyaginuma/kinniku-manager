package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kinniku-manager/repository"
)

func main() {
	r := gin.Default()
	r.GET("/training_exercise/read_all", readAllTrainingExcercise)
	r.Run("localhost:8080")
}

func readAllTrainingExcercise(c *gin.Context) {
	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	repository := &repository.TrainingExcerciseRepository{Database: db}
	excercises, err := repository.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, excercises)
}
