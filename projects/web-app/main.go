package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	userRepo       UserRepository
	postRepo       PostRepository
	templateDir    string
	publicDir      string
	tp             *TemplateRenderer
	sessionManager *scs.SessionManager
}

func main() {

	db, err := connectToDatabase("hk_news.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Secure = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	app := &application{
		errorLog:       log.New(os.Stderr, "ERROR\t", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog:        log.New(os.Stdout, "INFO\t", log.Ltime|log.LstdFlags),
		userRepo:       NewSQLUserRepository(db),
		postRepo:       NewSQLPostRepository(db),
		templateDir:    "../web-app/templates",
		publicDir:      "../web-app/public/",
		sessionManager: sessionManager,
	}
	app.tp = NewTemplateRender(app.templateDir, true)

	log.Println("Listening on :8080")
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
