package entrymeta

import (
	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
)

type Repository struct {
	Database *database.Database
	EntryRepository *entry.Repository
}

func New(db *database.Database) *Repository {
	EntryRepo := entry.New(db)

	return &Repository{Database: db, EntryRepository: EntryRepo}
}
