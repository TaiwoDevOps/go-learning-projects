package main

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	buf := new(bytes.Buffer)
	testApp.infoLog = log.New(buf, "", 0)
	handler := testApp.logger(testHandler)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
	assert.Contains(t, buf.String(), "HTTP/1.1 GET /test")

}

func TestRecover(t *testing.T) {
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})

	handler := testApp.recover(panicHandler)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "Internal Server Error\n", w.Body.String())
	assert.Equal(t, "close", w.Header().Get("Connection"))

}

func TestRequireAuth_Authenticated(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("protected"))
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req = req.WithContext(contextWithAuth(req.Context(), true))
	w := httptest.NewRecorder()

	handler := testApp.requireAuth(testHandler)

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "protected", w.Body.String())
	assert.Equal(t, "no-cache", w.Header().Get("Cache-Control"))
}

func TestRequireAuth_NotAuthenticated(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("protected"))
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	handler := testApp.requireAuth(testHandler)

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusSeeOther, w.Code)
	assert.Equal(t, "/login?redirectTo=/test", w.Header().Get("Location"))
}

func TestAuthenticate_ValidSession(t *testing.T) {
	defer cleanupTestData(t)
	userID, err := testApp.userRepo.CreateUser(
		"session user",
		"valid@session.com",
		"password",
		"avatar",
	)
	assert.NoError(t, err)
	assert.Greater(t, userID, int64(0))

	setupHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testApp.sessionManager.Put(r.Context(), loggedInUserKey, "valid@session.com")
		w.WriteHeader(http.StatusOK)
	})

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, testApp.isAuthenticated(r))
		user := testApp.getUserFromContext(r.Context())
		assert.Equal(t, userID, int64(user.ID))
		assert.Equal(t, "valid@session.com", user.Email)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	setupChain := testApp.sessionManager.LoadAndSave(setupHandler)
	req1 := httptest.NewRequest(http.MethodGet, "/setup", nil)
	w1 := httptest.NewRecorder()
	setupChain.ServeHTTP(w1, req1)

	testChain := testApp.sessionManager.LoadAndSave(testApp.authenticate(testHandler))
	req2 := httptest.NewRequest(http.MethodGet, "/test", nil)

	if cookies := w1.Result().Cookies(); len(cookies) > 0 {
		for _, cookie := range cookies {
			req2.AddCookie(cookie)
		}
	}

	w2 := httptest.NewRecorder()
	testChain.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, "OK", w2.Body.String())

}

func contextWithAuth(ctx context.Context, isAuth interface{}) context.Context {
	return context.WithValue(ctx, contextAuthKey, isAuth)
}
