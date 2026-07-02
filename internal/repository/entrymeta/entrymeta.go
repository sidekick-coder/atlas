package entrymeta

import (
    "github.com/sidekick-coder/atlas/internal/database"
)

type Repository struct {
	Database *database.Database
}

func New(db *database.Database) *Repository {
	return &Repository{Database: db}
}
