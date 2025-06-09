package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

// Todo: avoid duplicate user by verifying if email and username already exist
func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertId)

	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil

}

// Todo: Test if the select statement works
// Todo: test if the returned user struct in the error msg is nil or not
func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {

	u := User{}
	query := "SELECT * FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return &u, nil
	}

	return &u, nil
}
