package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	loggedInUserKey = "logged_in_user_id"
)

func (app *application) readIntWithDefault(r *http.Request, key string, dValue int) int {
	v, err := strconv.Atoi(r.URL.Query().Get(key))
	if err != nil {
		return dValue
	}
	return v
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	filter := Filter{
		Query:    r.URL.Query().Get("q"),
		OrderBy:  r.URL.Query().Get("order_by"),
		Page:     app.readIntWithDefault(r, "page", 1),
		PageSize: app.readIntWithDefault(r, "page_size", 50),
	}

	posts, metadata, err := app.postRepo.GetAll(filter)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("\nMetadata: %+v\n", metadata)
	// app.infoLog.Printf("\nPosts: %s and len %d\n", posts[0].GetVoteCountHuman(), len(posts))

	app.render(w, r, "index.html", &templateData{
		Posts:    posts,
		Metadata: metadata,
		NextLink: fmt.Sprintf("/?q=%s&order_by=%s&page=%d&page_size=%d",
			filter.Query, filter.OrderBy, metadata.NextPage, filter.PageSize),
		PreviousLink: fmt.Sprintf("/?q=%s&order_by=%s&page=%d&page_size=%d",
			filter.Query, filter.OrderBy, metadata.PrevPage, filter.PageSize),
	})
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {

	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("email", "password").
			MaxLength("email", 255).
			MaxLength("password", 255).
			MinLength("password", 8).
			MinLength("email", 7).
			IsEmail("email")

		if !form.Valid() {
			form.Errors.Add("generic", "The data you submitted was not valid")
			app.render(w, r, "login.html", &templateData{
				Form: form,
			})
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		_, err := app.userRepo.Authenticate(email, password)
		if err != nil {
			form.Errors.Add("generic", err.Error())
			app.render(w, r, "login.html", &templateData{
				Form: form,
			})
			return
		}
		app.sessionManager.Put(r.Context(), loggedInUserKey, email)
		app.sessionManager.Put(r.Context(), "flash", "You are logged in")

		app.infoLog.Println("Logged in")

		http.Redirect(w, r, "/submit", http.StatusSeeOther)
		return
	}

	app.render(w, r, "login.html", &templateData{
		Form: NewForm(r.PostForm),
	})
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	app.sessionManager.Remove(r.Context(), loggedInUserKey)
	app.sessionManager.Put(r.Context(), "flash", "You are logged out")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app *application) register(w http.ResponseWriter, r *http.Request) {

	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("email", "password", "name").
			MaxLength("email", 255).
			MaxLength("password", 255).
			MinLength("password", 8).
			MinLength("email", 7).
			MinLength("name", 3).
			IsEmail("email")

		if !form.Valid() {
			form.Errors.Add("generic", "The data you submitted was not valid")
			app.render(w, r, "register.html", &templateData{
				Form: form,
			})
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")
		name := r.FormValue("name")
		avatar := r.FormValue("avatar")

		_, err := app.userRepo.CreateUser(name, email, password, avatar)
		if err != nil {
			form.Errors.Add("generic", err.Error())
			app.render(w, r, "register.html", &templateData{
				Form: form,
			})
			return
		}
		app.sessionManager.Put(r.Context(), "flash", "You are registered")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	app.render(w, r, "register.html", &templateData{
		Form: NewForm(r.PostForm),
	})

}

func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("title", "url").MaxLength("title", 255).MaxLength("url", 255).MinLength("url", 3)

		if !form.Valid() {
			form.Errors.Add("generic", "The data you submitted was not valid")
			app.render(w, r, "submit.html", &templateData{
				Form: form,
			})
			return
		}

		title := r.FormValue("title")
		url := r.FormValue("url")

		user := app.getUserFromContext(r.Context())
		id, err := app.postRepo.CreatePost(title, url, user.ID)
		if err != nil {
			app.errorLog.Printf("error creating post: %s\n", err.Error())
			form.Errors.Add("generic", "creation of post failed")
			app.render(w, r, "submit.html", &templateData{
				Form: form,
			})
			return
		}
		app.sessionManager.Put(r.Context(), "flash", "post created")
		app.infoLog.Printf("post created with %d", id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.render(w, r, "submit.html", &templateData{
		Form: NewForm(r.PostForm),
	})

}

func (app *application) vote(w http.ResponseWriter, r *http.Request) {
	postId := app.readIntWithDefault(r, "post_id", 0)
	u := app.getUserFromContext(r.Context())

	err := app.postRepo.AddVote(u.ID, postId)
	if err != nil {
		app.errorLog.Printf("error adding vote: %s\n", err.Error())
		app.sessionManager.Put(r.Context(), "flash", "voting failed")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", fmt.Sprintf("You voted for post with id #%d", postId))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app *application) comments(w http.ResponseWriter, r *http.Request) {
	postId := app.readIntWithDefault(r, "post_id", 0)
	u := app.getUserFromContext(r.Context())

	post, err := app.postRepo.GetById(postId)
	if err != nil {
		app.errorLog.Printf("error getting comments: %s\n", err.Error())
		app.sessionManager.Put(r.Context(), "flash", "post not found")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	comments, err := app.postRepo.GetComment(postId)
	if err != nil {
		app.errorLog.Printf("error getting comments: %s\n", err.Error())
		app.sessionManager.Put(r.Context(), "flash", "error getting comments")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("comment").MinLength("comment", 5).MaxLength("comment", 255)

		if !form.Valid() {
			form.Errors.Add("generic", "The data you submitted was not valid")
			app.render(w, r, "comments.html", &templateData{
				Form: form,
			})
			return
		}
		_, err := app.postRepo.AddComment(u.ID, postId, r.FormValue("comment"))
		if err != nil {
			app.errorLog.Printf("error adding comments: %s\n", err.Error())
			app.sessionManager.Put(r.Context(), "flash", "error adding comments")
			http.Redirect(w, r, fmt.Sprintf("/comments?post_id=%d", post.ID), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/comments?post_id=%d", post.ID), http.StatusSeeOther)
		return
	}
	app.render(w, r, "comments.html", &templateData{
		Form:     NewForm(r.PostForm),
		Comments: comments,
		Post:     post,
	})
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "about.html", nil)

}
func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.html", nil)

}
