CREATE TABLE entry_metas (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    entry_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    value TEXT
);

-- unique constraint to ensure that each entrie_id and name combination is unique
CREATE UNIQUE INDEX idx_entrie_metas_entry_id_name ON entry_metas (entry_id, name);
