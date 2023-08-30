package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// InitDatabase initializes the SQLite database and creates necessary tables.
func InitDatabase(dbPath string) (*sql.DB, error) {
	// Open or create the database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create tables
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS file (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				path TEXT NOT NULL,
				hash TEXT,
				inFlight BOOLEAN NOT NULL DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS blobHash (
			fileHash INTEGER PRIMARY KEY AUTOINCREMENT,
			hash TEXT NOT NULL,
			container TEXT NOT NULL,
			FOREIGN KEY(fileHash) REFERENCES hash(hash),
			FOREIGN KEY(container) REFERENCES container(container)
		);

		CREATE TABLE IF NOT EXISTS container (
			container INTEGER PRIMARY KEY AUTOINCREMENT
		);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
