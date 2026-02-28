package main

import (
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, filename string, data *templateData) {
	if app.tp == nil {
		http.Error(w, "Template renderer not initialized correctly", http.StatusInternalServerError)
		return
	}

	app.tp.Render(w, filename, app.appDefaultTemplate(data, r))

}

func (app *application) appDefaultTemplate(data *templateData, r *http.Request) *templateData {
	if data == nil {
		data = &templateData{}
	}

	data.Flash = app.sessionManager.PopString(r.Context(), "flash")
	data.IsAuthenticated = app.isAuthenticated(r)
	app.infoLog.Printf("Flash cookie: %s; called by %s\n", data.Flash, r.URL.Path)

	return data

}
