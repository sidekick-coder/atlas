package entry

func (r * Repository) DeleteAll() error {
	// SELECT ...
	smtmt := "DELETE FROM entries;"

    _, err := r.Database.Query(smtmt)

	if err != nil {
		return err
	}

    return nil
}
