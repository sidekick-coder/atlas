package entrymeta

import (
	"strings"
)

func (r * Repository) InsertMany(entryId int64, metas map[string]string) error {
	smtmt := []string{}
	params := []any{}
	values := []string{}

	smtmt = append(smtmt, "INSERT INTO entry_metas (entry_id, name, value) VALUES")

	for name, value := range metas {
		values = append(values, "(?, ?, ?)")
		params = append(params, entryId, name, value)
	}

	smtmt = append(smtmt, strings.Join(values, ","))

	_, err := r.Database.Exec(strings.Join(smtmt, " "), params...)

	if err != nil {
		return err
	}


	return nil
}
