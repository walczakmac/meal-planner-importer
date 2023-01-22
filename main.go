package main

import (
	"context"
	_ "github.com/lib/pq"
	"log"
)

var ctx = context.Background()

func main() {
	flags := NewFlags()
	recipes := NewCsv(flags.Filepath).readFile()

	db := NewDatabase(*flags)
	mealService := NewMealService(db.OpenConnection())

	for i, details := range recipes {
		if i == 0 {
			continue
		}

		err := mealService.NewMeal(details)
		if err != nil {
			log.Println(err)
		}
	}
}
