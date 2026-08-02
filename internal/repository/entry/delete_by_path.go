package entry

func (r * Repository) DeleteByPath(path string) error {
	entry, err := r.GetByPath(path)

	if err == nil {
		_, err = r.Database.Exec("DELETE FROM entry_metas WHERE entry_id = ?;", entry.ID)
		
		if err != nil {
			return err
		}
	}

    _, err = r.Database.Exec("DELETE FROM entries WHERE path = ?;", path)

	if err != nil {
		return err
	}

    return nil
}
