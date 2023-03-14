package dialectquery

import "fmt"

type sqlite3 struct {
	table string
}

func (s *sqlite3) CreateTable() string {
	q := `CREATE TABLE %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		version_id INTEGER NOT NULL,
		is_applied INTEGER NOT NULL,
		tstamp TIMESTAMP DEFAULT (datetime('now'))
	)`
	return fmt.Sprintf(q, s.table)
}

func (s *sqlite3) InsertVersion() string {
	q := `INSERT INTO %s (version_id, is_applied) VALUES (?, ?)`
	return fmt.Sprintf(q, s.table)
}

func (s *sqlite3) DeleteVersion() string {
	q := `DELETE FROM %s WHERE version_id=?`
	return fmt.Sprintf(q, s.table)
}

func (s *sqlite3) GetMigrationByVersion() string {
	q := `SELECT tstamp, is_applied FROM %s WHERE version_id=? ORDER BY tstamp DESC LIMIT 1`
	return fmt.Sprintf(q, s.table)
}

func (s *sqlite3) ListMigrations() string {
	q := `SELECT version_id, is_applied from %s ORDER BY id DESC`
	return fmt.Sprintf(q, s.table)
}
