package database

func (db *Database) Migrate() error {
	smpt := []string{}

	smpt = append(smpt, `
		CREATE TABLE entries (
			id INTEGER PRIMARY KEY,
			path TEXT NOT NULL UNIQUE
		);
	`)

	smpt = append(smpt, `
		CREATE TABLE entry_metas (
			id INTEGER PRIMARY KEY,
			entry_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			value TEXT NOT NULL
		);
	`)

	smpt = append(smpt, `CREATE UNIQUE INDEX idx_entrie_metas_entry_id_name ON entry_metas (entry_id, name);`)

	for _, s := range smpt {
		if _, err := db.Connection.Exec(s); err != nil {
			return err
		}
	}
	return nil
}
