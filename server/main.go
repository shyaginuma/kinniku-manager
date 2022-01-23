package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db := store.NewMysqlStore()
	// exercises, err := db.GetAllTrainingExercises()
	// if err != nil {
	// 	log.Fatal()
	// }
	// fmt.Println(exercises)
	db, err := sql.Open("mysql", "shika:deer@/trainings")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name FROM training_exercises")
	if err != nil {
		log.Fatal(err)
	}
	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", name)
	}
}
