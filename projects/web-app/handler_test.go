package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin_GET_NotAuthenticated(t *testing.T) {
	defer cleanupTestData(t)

	handler := testApp.sessionManager.LoadAndSave(testApp.authenticate(http.HandlerFunc(testApp.login)))

	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLogin_GET_AlreadyAuthenticated(t *testing.T) {
	defer cleanupTestData(t)

	_, err := testApp.userRepo.CreateUser(
		"login",
		"login@test.com",
		"goodpassword",
		"avatar",
	)

	assert.NoError(t, err)

	setupHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testApp.sessionManager.Put(r.Context(), loggedInUserKey, "login@test.com")
		w.WriteHeader(http.StatusOK)
	})

	setupChain := testApp.sessionManager.LoadAndSave(setupHandler)
	req1 := httptest.NewRequest(http.MethodGet, "/setup", nil)
	w1 := httptest.NewRecorder()
	setupChain.ServeHTTP(w1, req1)

	handler := testApp.sessionManager.LoadAndSave(testApp.authenticate(http.HandlerFunc(testApp.login)))
	req2 := httptest.NewRequest(http.MethodGet, "/test", nil)
	if cookies := w1.Result().Cookies(); len(cookies) > 0 {
		for _, cookie := range cookies {
			req2.AddCookie(cookie)
		}
	}

	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusSeeOther, w2.Code)
	assert.Equal(t, "/", w2.Header().Get("Location"))
}

func TestLogin_POST_ValidCredentials(t *testing.T) {

	defer cleanupTestData(t)

	_, err := testApp.userRepo.CreateUser(
		"login",
		"login@test.com",
		"goodpassword",
		"avatar",
	)
	assert.NoError(t, err)

	handler := testApp.sessionManager.LoadAndSave(testApp.authenticate(http.HandlerFunc(testApp.login)))
	formData := "email=login@test.com&password=goodpassword"
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusSeeOther, w.Code)
	assert.Equal(t, "/submit", w.Header().Get("Location"))

}

func TestLogin_POST_InvalidFormData(t *testing.T) {
	defer cleanupTestData(t)

	handler := testApp.sessionManager.LoadAndSave(testApp.authenticate(http.HandlerFunc(testApp.login)))
	formData := "email=&password="
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	body := w.Body.String()

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, body, "The data you submitted was not valid")
	assert.Contains(t, body, "This field email is required")
	assert.Contains(t, body, "This field password is required")
}

func TestLogin_POST_InvalidAuthenticationData(t *testing.T) {
	defer cleanupTestData(t)

	handler := testApp.sessionManager.LoadAndSave(testApp.authenticate(http.HandlerFunc(testApp.login)))
	formData := "email=test@test.com&password=goodpassword"
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	body := w.Body.String()

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, body, "sql: no rows in result set")

}
