package main

import (
	"fmt"
	"log"

	"github.com/kinniku-manager/repository"
)

func main() {
	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	repository := &repository.TrainingExcerciseRepository{Database: db}
	excercises, err := repository.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(excercises)
}
