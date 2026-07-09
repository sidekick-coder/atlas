package database

import (
	"os"

	"github.com/sidekick-coder/atlas/internal/config"
)

func Create(filename string) (*Database, error) {
	connection, err := Connect(filename)

	if err != nil {
		panic(err)
	}

	return New(connection), err
}

func CreateFromConfig(config *config.Config) (*Database, error) {
	dbPath, ok := config.Get("workspace.database_path")

	if !ok {
		return nil, os.ErrInvalid
	}

	if dbPath == "" {
		return nil, os.ErrInvalid
	}

	return Create(dbPath)
}
