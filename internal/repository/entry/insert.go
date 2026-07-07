package entry

import (
	"fmt"
	"strings"
)


func (r *Repository) InsertManyPaths(paths []string) error {
	smtmt := []string{}
	params := []any{}

	smtmt = append(smtmt, "INSERT INTO entries (path) VALUES")

	values := []string{}

	for _, p := range paths {
		values = append(values, "(?)")
		params = append(params, p)
	}

	smtmt = append(smtmt, strings.Join(values, ",\n"))

	_, err := r.Database.Exec(strings.Join(smtmt, " "), params...)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) InsertByPathWithMetas(path string, metas map[string]string) error {
	tx, err := r.Database.Connection.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	res, err := tx.Exec("INSERT INTO entries (path) VALUES (?);", path)

	if err != nil {
		return err
	}

	if len(metas) == 0 {
		return tx.Commit()
	}

	entryId, err := res.LastInsertId()

	values := []string{}
	params := []any{}

	for name, value := range metas {
		values = append(values, "(?, ?, ?)")
		params = append(params, entryId, name, value)
	}

	smtmt := fmt.Sprintf("INSERT INTO entry_metas (entry_id, name, value) VALUES %s", strings.Join(values, ","))

	_, err = tx.Exec(smtmt, params...)

	if err != nil {
		return err
	}

	return tx.Commit()

}
