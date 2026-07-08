package database

func (d *Database) Cleanup() error {
	tx, err := d.Connection.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM entries;")

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM entry_metas;")

	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}
