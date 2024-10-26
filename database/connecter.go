package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Connecter struct {
	DB *sql.DB
}

func OpenOrCreate(name string) (*Connecter, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting working directory: %v", err)
	}

	dbFile := filepath.Join(currentDir, name)

	_, err = os.Stat(dbFile)
	var create bool
	if err != nil {
		create = true
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if create {
		err = createDatabase(db)
		if err != nil {
			return nil, fmt.Errorf("error creating database: %v", err)
		}
	}

	return &Connecter{DB: db}, nil
}

func (c *Connecter) Close() {
	c.DB.Close()
}

func createDatabase(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS series
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		link TEXT NOT NULL UNIQUE,
		imdb TEXT,
		start_year INTEGER,
		end_year INTEGER,
		poster TEXT NOT NULL,
		country TEXT
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("query execution error")
	}

	return nil
}
