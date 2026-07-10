package entry

import (
	"strings"

	"github.com/sidekick-coder/atlas/internal/models"
)

type ListOptions struct {
	Query     []string
	Limit     int
	Offset    int
	LoadMetas bool
}

func (r *Repository) List(options ...ListOptions) ([]models.Entry, error) {
	limit := 100
	offset := 0
	query := []string{}
	loadMetas := true

	if len(options) > 0 {
		if options[0].Limit > 0 {
			limit = options[0].Limit
		}

		if options[0].Offset > 0 {
			offset = options[0].Offset
		}

		if len(options[0].Query) > 0 {
			query = options[0].Query
		}


		loadMetas = options[0].LoadMetas
	}

	stmt := []string{
		"SELECT entries.id, entries.path",
		"FROM entries",
		"WHERE 1=1",
	}

	params := []any{}

	if len(query) > 0 && query[0] != "" {
		node, err := ParseQuery(query)

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

	stmt = append(stmt, "ORDER BY entries.path ASC")

	if limit > 0 {
		stmt = append(stmt, "LIMIT ?")
		params = append(params, limit)
	}

	if offset > 0 {
		stmt = append(stmt, "OFFSET ?")
		params = append(params, offset)
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

	if loadMetas {
		entryMetas, err := r.ListMetas(entries...)

		if err != nil {
			return nil, err
		}

		for i, entry := range entries {
			entries[i].Metas = r.GetEntryMetasMap(entryMetas, entry.ID)
		}

		return entries, nil
	}

	return entries, nil
}
