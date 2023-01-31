package main

import (
	"context"
	"flag"
	_ "github.com/lib/pq"
	"log"
)

var ctx = context.Background()

func main() {
	host := flag.String("host", "", "Host")
	port := flag.Int("port", 5432, "Port")
	user := flag.String("user", "", "User")
	password := flag.String("password", "", "Password")
	dbname := flag.String("dbname", "", "Database name")
	filePath := flag.String("filepath", "", "Path to CSV file with recipes")
	flag.Parse()

	config, err := NewConfig(
		*host,
		*port,
		*user,
		*password,
		*dbname,
		*filePath,
	)

	if err != nil {
		log.Fatalln(err)
	}

	recipes := NewCsv(config.Filepath).readFile()

	db := NewDatabase(*config)
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
