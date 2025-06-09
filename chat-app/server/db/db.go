package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {

	db, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/chat-app?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

// Close db
func (d *Database) Close() {
	d.db.Close()
}

// get DB
func (d *Database) GetDB() *sql.DB {
	return d.db
}
