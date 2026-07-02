package entry 

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func (r *Repository) Upsert(path string) (*models.Entry, error) {
	smtmt := `
	INSERT INTO entries (path)
	VALUES ($1)
	ON CONFLICT (path) DO NOTHING;
	`

	_, err := r.Database.Query(smtmt, path)

	if err != nil {
		return nil, err
	}

	return r.GetByPath(path)
}
