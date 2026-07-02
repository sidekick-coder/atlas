package database 

import (
	"errors"
	"github.com/sidekick-coder/atlas/internal/utils"
)

func Create(filename string) (*Database, error) {
	exists, err := utils.PathExists(filename)

	if err != nil || !exists {
		return nil, errors.New("database file don't found")
	}

	if err != nil {
		panic(err)
	}

	connection, err := Connect(filename)

	if err != nil {
		panic(err)
	}

	return New(connection), err
}
