package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	hashed_password BLOB NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
`

func main() {
	dbName := "data.db"

	// Note a practice in GO
	_ = os.Remove(dbName)

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	// once you open any file - including DB, always close
	defer func() {
		fmt.Println("closing the database connection")
		if err := db.Close(); err != nil {
			log.Printf("error closing database connection: %v", err)
		}

	}()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection to database established")

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal("error creating table ", err)
	}

	fmt.Println("table was created")

}
