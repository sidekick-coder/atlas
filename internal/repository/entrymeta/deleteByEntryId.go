package entrymeta

func (r * Repository) DeleteByEntryID(entryId int64) error {
	// SELECT ...
	smtmt := `
	DELETE FROM entry_metas
	WHERE entry_id = $1;
	`

	_, err := r.Database.Exec(smtmt, entryId)

	if err != nil {
		return err
	}

	return nil
}
