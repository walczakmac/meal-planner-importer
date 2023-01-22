package main

import (
	"flag"
	"log"
)

type Flags struct {
	database Database
	Filepath string
}

func NewFlags() *Flags {
	host := flag.String("host", "", "Host")
	port := flag.Int("port", 5432, "Port")
	user := flag.String("user", "", "User")
	password := flag.String("password", "", "Password")
	dbname := flag.String("dbname", "", "Database name")
	filePath := flag.String("filepath", "", "Path to CSV file with recipes")

	if flag.Parsed() == false {
		flag.Parse()
	}

	if *host == "" {
		log.Fatalln("Database host must be provided and not empty")
	}
	if *user == "" {
		log.Fatalln("Database user must be provided and not empty")
	}
	if *password == "" {
		log.Fatalln("Database password must be provided and not empty")
	}
	if *dbname == "" {
		log.Fatalln("Database name must be provided and not empty")
	}
	if *filePath == "" {
		log.Fatalln("CSV file path must be provided and not empty")
	}

	return &Flags{
		database: Database{
			Host:     *host,
			Name:     *dbname,
			Port:     *port,
			User:     *user,
			Password: *password,
		},
		Filepath: *filePath,
	}
}

func (flags Flags) Database() Database {
	return flags.database
}
