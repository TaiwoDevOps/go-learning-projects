package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type webAppContextKey string

const (
	contextAuthKey webAppContextKey = webAppContextKey("isAuthKey")
	contextUserKey webAppContextKey = webAppContextKey("auth_user")
)

func (app *application) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *application) recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !app.isAuthenticated(r) {
			http.Redirect(w, r, fmt.Sprintf("/login?redirectTo=%s", r.URL.Path), http.StatusSeeOther)
			return
		}
		w.Header().Set("Cache-Control", "no-cache")
		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exists := app.sessionManager.Exists(r.Context(), loggedInUserKey)
		if !exists {
			next.ServeHTTP(w, r)
			return
		}
		u, err := app.userRepo.GetUserByEmail(app.sessionManager.GetString(r.Context(), loggedInUserKey))
		if errors.Is(err, sql.ErrNoRows) {
			app.sessionManager.Remove(r.Context(), loggedInUserKey)
			next.ServeHTTP(w, r)
			return
		} else if err != nil {
			app.serverError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), contextAuthKey, true)
		ctx = context.WithValue(ctx, contextUserKey, u)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) isAuthenticated(r *http.Request) bool {
	isAuth, ok := r.Context().Value(contextAuthKey).(bool)
	if !ok {
		return false
	}

	return isAuth

}
