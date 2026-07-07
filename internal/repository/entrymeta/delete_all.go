package entrymeta

func (r * Repository) DeleteAll() error {
	// SELECT ...
	smtmt := "DELETE FROM entry_metas;"

    _, err := r.Database.Exec(smtmt)

	if err != nil {
		return err
	}

    return nil
}
