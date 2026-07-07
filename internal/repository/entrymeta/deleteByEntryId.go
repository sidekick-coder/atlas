package entrymeta

func (r *Repository) DeleteByEntryIDSql(entryId int64) (string, []any) {
	smtmt := `DELETE FROM entry_metas	WHERE entry_id = $1`

	params := []any{entryId}

	return smtmt, params
}

func (r *Repository) DeleteByEntryID(entryId int64) error {
	smtmt, _ := r.DeleteByEntryIDSql(entryId)

	_, err := r.Database.Exec(smtmt, entryId)

	if err != nil {
		return err
	}

	return nil
}
