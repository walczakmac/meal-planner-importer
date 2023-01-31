package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfigErrors(t *testing.T) {
	var tests = []struct {
		host          string
		port          int
		user          string
		password      string
		dbname        string
		filepath      string
		expectedError string
	}{
		{
			"", 0, "", "", "", "",
			"database host must be provided and not empty",
		},
		{
			"localhost", 0, "", "", "", "",
			"database port must be provided and not equal 0",
		},
		{
			"localhost", 5432, "", "", "", "",
			"database user must be provided and not empty",
		},
		{
			"localhost", 5432, "user1", "", "", "",
			"database password must be provided and not empty",
		},
		{
			"localhost", 5432, "user1", "qwerty123", "", "",
			"database name must be provided and not empty",
		},
		{
			"localhost", 5432, "user1", "qwerty123", "my_db", "",
			"CSV file path must be provided and not empty",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Expected error: %s", test.expectedError), func(t *testing.T) {
			config, err := NewConfig(test.host, test.port, test.user, test.password, test.dbname, test.filepath)
			assert.Nil(t, config)
			assert.Equal(t, test.expectedError, err.Error())
		})
	}
}

func TestNewConfigSuccess(t *testing.T) {
	config, err := NewConfig(
		"localhost",
		5432,
		"user1",
		"qwerty123",
		"my_db",
		"/path/to/file",
	)
	assert.NotNil(t, config)
	assert.Nil(t, err)
	assert.Equal(t, "localhost", config.Database().Host)
	assert.Equal(t, 5432, config.Database().Port)
	assert.Equal(t, "user1", config.Database().User)
	assert.Equal(t, "qwerty123", config.Database().Password)
	assert.Equal(t, "my_db", config.Database().Name)
	assert.Equal(t, "/path/to/file", config.Filepath)
}
