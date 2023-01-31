package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	Host     string
	Name     string
	Port     int
	User     string
	Password string
}

func NewDatabase(flags Config) *Database {
	return &Database{
		Host:     flags.Database().Host,
		Name:     flags.Database().Name,
		Port:     flags.Database().Port,
		User:     flags.Database().User,
		Password: flags.Database().Password,
	}
}

func (database Database) OpenConnection() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		database.Host,
		database.Port,
		database.User,
		database.Password,
		database.Name,
	))
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}
