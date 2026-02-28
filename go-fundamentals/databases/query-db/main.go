package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"user_name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

func main() {

	db, err := sql.Open("sqlite3", "../users_table.db")
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

	ctx := context.Background()
	user, err := GetUsers(db, ctx)
	if err != nil {
		log.Fatal(err)
	}

	byteUser, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteUser))

}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt, err := db.Prepare(`SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`)
	if err != nil {
		return nil, err
	}

	record := stmt.QueryRow(email)
	var user User
	err = record.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsers(db *sql.DB, ctx context.Context) ([]User, error) {
	stmt, err := db.Prepare(`SELECT id, name, email, hashed_password, created_at FROM users`)
	if err != nil {
		return nil, err
	}
	records, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer records.Close()
	var users []User
	for records.Next() {
		var user User
		if err := records.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.HashedPassword,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := records.Err(); err != nil {
		return nil, err
	}

	return users, nil

}
