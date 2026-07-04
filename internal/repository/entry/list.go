package entry

import (
	"strings"
	"github.com/sidekick-coder/atlas/internal/models"
)

type ListOptions struct {
	Query []string
	Limit int 
	Offset int
}

func (r *Repository) List(options ...ListOptions) ([]models.Entry, error) {
	stmt := []string{
		"SELECT entries.id, entries.path",
		"FROM entries",
		"WHERE 1=1",
	}

	params := []interface{}{}

	if len(options) > 0 && len(options[0].Query) > 0 {
		node, err := ParseQuery(options[0].Query)
		if err != nil {
			return nil, err
		}

		if node != nil {
			condition, err := BuildSQL(node, &params)
			if err != nil {
				return nil, err
			}
			stmt = append(stmt, "AND", condition)
		}
	}

	if len(options) > 0 && options[0].Limit > 0 {
		stmt = append(stmt, "LIMIT ?")
		params = append(params, options[0].Limit)
	}

	if len(options) > 0 && options[0].Offset > 0 {
		stmt = append(stmt, "OFFSET ?")
		params = append(params, options[0].Offset)
	}

	stmtStr := strings.Join(stmt, " ")

	rows, err := r.Database.Query(stmtStr, params...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entries []models.Entry

	for rows.Next() {
		var entry models.Entry

		err := rows.Scan(&entry.ID, &entry.Path)

		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}
