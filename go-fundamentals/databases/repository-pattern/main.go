package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/taiwodevops/go-fundamentals/databases/repository-pattern/repository"
)

func main() {
	db, err := connectToDatabase("../users_table.db")

	checkErr(err)
	defer db.Close()
	pingDBConnection(db)

	userRepo := repository.NewSQLUserRepository(db)
	_, err = userRepo.CreateUser("New User", "new@user.com", "P@ssw0rd", "https://avatar.jpeg")
	checkErr(err)
	printAllUsers(userRepo)

}

func printAllUsers(repo repository.UserRepository) {
	users, err := repo.GetUsers()
	checkErr(err)
	for _, user := range users {
		fmt.Println(user.ID, user.Email)
	}
}

func pingDBConnection(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection to database established")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func connectToDatabase(dbName string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}
