package entry 
//
// import (
// 	"github.com/sidekick-coder/atlas/internal/models"
// )
//
// func (r *Repository) Upsert(path string, is_dir bool) (*models.Entry, error) {
// 	smtmt := `
// 	INSERT INTO entries (path, is_dir)
// 	VALUES ($1, $2)
// 	ON CONFLICT (path) DO UPDATE SET is_dir = EXCLUDED.is_dir;
// 	`
//
// 	_, err := r.Database.Query(smtmt, path, is_dir)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return GetByPath(path)
// }
