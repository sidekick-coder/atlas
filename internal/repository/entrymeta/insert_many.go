package entrymeta

import (
	"strings"
)

func (r * Repository) InsertManySql(entryId int64, metas map[string]string) (string, []any) {
	smtmt := []string{}
	params := []any{}
	values := []string{}

	smtmt = append(smtmt, "INSERT INTO entry_metas (entry_id, name, value) VALUES")

	for name, value := range metas {
		values = append(values, "(?, ?, ?)")
		params = append(params, entryId, name, value)
	}

	smtmt = append(smtmt, strings.Join(values, ","))

	return strings.Join(smtmt, " "), params
}

func (r * Repository) InsertMany(entryId int64, metas map[string]string) error {
	smtmt, params := r.InsertManySql(entryId, metas)

	_, err := r.Database.Exec(smtmt, params...)

	if err != nil {
		return err
	}


	return nil
}

func (r * Repository) InsertManyByPath(path string, metas map[string]string) error {
	entry, err := r.EntryRepository.GetByPath(path)

	if err != nil {
		return err
	}

	return r.InsertMany(entry.ID, metas)
}

