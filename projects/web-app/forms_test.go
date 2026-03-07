package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewForm(t *testing.T) {
	email := "josh@tee.com"
	values := url.Values{}
	values.Add("email", email)

	form := NewForm(values)
	assert.NotNil(t, form)
	assert.Equal(t, email, form.Get("email"))
	assert.NotNil(t, form.Errors)
	assert.Len(t, form.Errors, 0)
}

func TestForm_Required(t *testing.T) {
	email := "josh@tee.com"
	values := url.Values{}
	values.Add("email", email)
	values.Add("empty", " ")

	form := NewForm(values)
	form.Required("email", "password", "empty")
	assert.NotNil(t, form)
	assert.Equal(t, "", form.Errors.Get("email"), "email is required")
	assert.Contains(t, form.Errors.Get("password"), "password is required")
	assert.Contains(t, form.Errors.Get("empty"), "empty is required")

}
