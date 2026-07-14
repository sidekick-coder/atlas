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

func (r *Repository) List(payload ...ListOptions) ([]models.Entry, error) {
	options := ListOptions{
		Query:     []string{},
		Limit:     10,
		Offset:    0,
		LoadMetas: true,
	}

	if len(payload) > 0 {
		o := payload[0]

		if o.Limit > 0 {
			options.Limit = o.Limit
		}

		if o.Offset > 0 {
			options.Offset = o.Offset
		}

		if len(o.Query) > 0 {
			for _, q := range o.Query {
				if q != "" {
					options.Query = append(options.Query, q)
				}
			}
		}

		if o.LoadMetas == false {
			options.LoadMetas = false
		}
	}

	stmt := []string{
		"SELECT entries.id, entries.path",
		"FROM entries",
		"WHERE 1=1",
	}

	params := []any{}

	if len(options.Query) > 0{
		node, err := ParseQuery(options.Query)

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

	if options.Limit > 0 {
		stmt = append(stmt, "LIMIT ?")
		params = append(params, options.Limit)
	}

	if options.Offset > 0 {
		stmt = append(stmt, "OFFSET ?")
		params = append(params, options.Offset)
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

	if options.LoadMetas {
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
