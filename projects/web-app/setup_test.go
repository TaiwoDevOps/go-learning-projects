package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var testDB *sql.DB
var testApp *application

func TestMain(m *testing.M) {
	fmt.Println("This is before the testing but within the testing file")
	var err error
	testDB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	if err := testDB.Ping(); err != nil {
		panic(err)
	}

	testApp = setupApp(testDB)
	if err = setupTestSchema(testDB); err != nil {
		panic(err)
	}

	code := m.Run()
	testDB.Close()
	fmt.Println("This is after the testing but within the testing file")
	os.Exit(code)
}

func setupApp(db *sql.DB) *application {
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Secure = true

	app := &application{
		errorLog:       log.New(io.Discard, "", 0),
		infoLog:        log.New(io.Discard, "", 0),
		userRepo:       NewSQLUserRepository(db),
		postRepo:       NewSQLPostRepository(db),
		templateDir:    "../web-app/templates",
		publicDir:      "../web-app/public/",
		sessionManager: sessionManager,
	}

	app.tp = NewTemplateRender(app.templateDir, false)
	return app

}

func setupTestSchema(db *sql.DB) error {
	schema := `
		CREATE TABLE users (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name TEXT NOT NULL,
   email TEXT NOT NULL UNIQUE,
   hashed_password TEXT NOT NULL,
   created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE profile (
     user_id INTEGER PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
     avatar TEXT NOT NULL,
     created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE posts (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   url TEXT NOT NULL,
   title TEXT NOT NULL UNIQUE,
   user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
   created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  body TEXT NOT NULL,
  user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
  post_id INTEGER REFERENCES posts(post_id) ON DELETE CASCADE,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE votes (
   user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
   post_id INTEGER REFERENCES posts(post_id) ON DELETE CASCADE,
   created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (user_id, post_id)
);

	`

	_, err := db.Exec(schema)
	return err
}

func cleanupTestData(t *testing.T) {
	tables := []string{
		"profile",
		"votes",
		"comments",
		"posts",
		"users",
	}

	for _, table := range tables {
		_, err := testDB.Exec("DELETE FROM " + table)
		assert.NoError(t, err)
	}
}
