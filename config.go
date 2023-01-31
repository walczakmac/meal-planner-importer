package main

import (
	"errors"
)

type Config struct {
	database Database
	Filepath string
}

func NewConfig(
	host string,
	port int,
	user string,
	password string,
	dbname string,
	filePath string,
) (*Config, error) {
	if host == "" {
		return nil, errors.New("database host must be provided and not empty")
	}
	if port == 0 {
		return nil, errors.New("database port must be provided and not equal 0")
	}
	if user == "" {
		return nil, errors.New("database user must be provided and not empty")
	}
	if password == "" {
		return nil, errors.New("database password must be provided and not empty")
	}
	if dbname == "" {
		return nil, errors.New("database name must be provided and not empty")
	}
	if filePath == "" {
		return nil, errors.New("CSV file path must be provided and not empty")
	}

	return &Config{
		database: Database{
			Host:     host,
			Name:     dbname,
			Port:     port,
			User:     user,
			Password: password,
		},
		Filepath: filePath,
	}, nil
}

func (flags Config) Database() Database {
	return flags.database
}
