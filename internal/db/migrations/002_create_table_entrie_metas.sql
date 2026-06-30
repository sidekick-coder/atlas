CREATE TABLE entrie_metas (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    entrie_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    value TEXT
);

-- unique constraint to ensure that each entrie_id and name combination is unique
CREATE UNIQUE INDEX idx_entrie_metas_entrie_id_name ON entrie_metas (entrie_id, name);
