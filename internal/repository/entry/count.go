package entry

type CountOptions struct {
	Query []string
}

func (r *Repository) Count(options ...CountOptions) (int, error) {
	stmt := "SELECT COUNT(*) FROM entries WHERE 1=1"
	params := []interface{}{}

	if len(options) > 0 && len(options[0].Query) > 0 {
		node, err := ParseQuery(options[0].Query)
		if err != nil {
			return 0, err
		}
		if node != nil {
			condition, err := BuildSQL(node, &params)
			if err != nil {
				return 0, err
			}
			stmt += " AND " + condition
		}
	}

	rows, err := r.Database.Query(stmt, params...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}
