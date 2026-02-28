package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const dbName = "../users_table.db"

var schema = `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	hashed_password TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
`

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"user_name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

func main() {
	// Note a practice in GO
	_ = os.Remove(dbName)
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	// once you open any file - including DB, always close
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection to database established")

	createTable(db)

	createUser(db, "User One", "user@one.com", "P@ssw0rd")
	createUser(db, "User Two", "user@Two.com", "P@ssw0rd")
	createUser(db, "User Three", "user@Three.com", "P@ssw0rd")
	createUser(db, "User Four", "user@Four.com", "P@ssw0rd")
	createUser(db, "User Five", "user@Five.com", "P@ssw0rd")
	createUser(db, "User Six", "user@six.com", "P@ssw0rd")
}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal("error creating table ", err)
	}
}

func createUser(db *sql.DB, name, email, plainPassword string) (int64, error) {
	stmt, err := db.Prepare(`INSERT INTO users (name, email, hashed_password) VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}

	hp, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, email, string(hp))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
