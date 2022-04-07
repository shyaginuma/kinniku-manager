package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kinniku-manager/model"
	"github.com/kinniku-manager/repository"
)

func main() {
	r := gin.Default()
	// Training Exercises
	r.GET("/training_exercise/", readAllTrainingExercises)
	r.GET("/training_exercise/:id", readTrainingExercise)
	r.GET("/training_exercise/search", searchTrainingExercises)
	r.POST("/training_exercise/save", createTrainingExercise)
	r.PUT("/training_exercise/edit", updateTrainingExercise)
	r.DELETE("/training_exercise/delete/:id", deleteTrainingExercise)

	// Training Sets
	r.GET("/training_set/", readAllTrainingSets)
	r.GET("/training_set/:id", readTrainingSet)
	r.POST("/training_set/save", createTrainingSet)
	r.PUT("/training_set/edit", updateTrainingSet)
	r.DELETE("/training_set/delete/:id", deleteTrainingSet)

	// Training Menus
	r.GET("/training_menu/", readAllTrainingMenus)
	r.GET("/training_menu/:id", readTrainingMenu)
	r.POST("/training_menu/save", createTrainingMenu)
	r.PUT("/training_menu/edit", updateTrainingMenu)
	r.DELETE("/training_menu/delete/:id", deleteTrainingMenu)

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

func readTrainingExercise(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("failed to read id parameter from path: %v", err)
	}
	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	repository := &repository.TrainingExerciseRepository{Database: db}
	exercises, err := repository.Read(id)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, exercises)
}

func searchTrainingExercises(c *gin.Context) {
	searchLimit, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Fatal(err)
	}
	searchKeyword := strings.Split(c.Query("keyword"), " ")
	targetMuscle := model.TargetMuscle(c.Query("target"))
	trainingCategory := model.TrainingCategory(c.Query("category"))
	trainingDifficulty := model.TrainingDifficulty(c.Query("difficulty"))

	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	repo := &repository.TrainingExerciseRepository{Database: db}
	exercises, err := repo.Search(
		repository.WithSearchLimit(searchLimit),
		repository.WithSearchKeyword(searchKeyword),
		repository.WithTargetMuscle(targetMuscle),
		repository.WithTrainingCategory(trainingCategory),
		repository.WithTrainingDifficulty(trainingDifficulty),
	)
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

func readAllTrainingSets(c *gin.Context) {
	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	repository := &repository.TrainingSetRepository{Database: db}
	sets, err := repository.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, sets)
}

func readTrainingSet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("failed to read id parameter from path: %v", err)
	}
	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	repository := &repository.TrainingSetRepository{Database: db}
	set, err := repository.Read(id)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, set)
}

func createTrainingSet(c *gin.Context) {
	var newTrainingSet model.TrainingSet
	if err := c.BindJSON(&newTrainingSet); err != nil {
		log.Fatalf("failed to bind json: %v", err)
	}

	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatalf("failed to establish connection with db: %v", err)
	}
	repository := &repository.TrainingSetRepository{Database: db}
	if err := repository.Create(newTrainingSet); err != nil {
		log.Fatalf("failed to insert record into db: %v", err)
	}
	c.IndentedJSON(http.StatusCreated, newTrainingSet)
}

func updateTrainingSet(c *gin.Context) {
	var updatedTrainingSet model.TrainingSet
	if err := c.BindJSON(&updatedTrainingSet); err != nil {
		log.Fatalf("failed to bind json: %v", err)
	}

	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatalf("failed to establish connection with db: %v", err)
	}
	repository := &repository.TrainingSetRepository{Database: db}
	if err := repository.Update(updatedTrainingSet); err != nil {
		log.Fatalf("failed to update data: %v", err)
	}
	c.IndentedJSON(http.StatusCreated, updatedTrainingSet)
}

func deleteTrainingSet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("failed to read id parameter from path: %v", err)
	}
	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatalf("failed to establish connection with db: %v", err)
	}
	repository := &repository.TrainingSetRepository{Database: db}
	if err := repository.Delete(id); err != nil {
		log.Fatalf("failed to delete data: %v", err)
	}
	c.String(http.StatusOK, "successfully delete training set.")
}

func readAllTrainingMenus(c *gin.Context) {

}

func readTrainingMenu(c *gin.Context) {

}

func createTrainingMenu(c *gin.Context) {

}

func updateTrainingMenu(c *gin.Context) {

}

func deleteTrainingMenu(c *gin.Context) {

}
