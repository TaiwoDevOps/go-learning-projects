package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLUserRepository_CreateUser(t *testing.T) {
	defer cleanupTestData(t)

	repo := NewSQLUserRepository(testDB)
	userId, err := repo.CreateUser(
		"Josh Tee A",
		"josh@tee_a.com",
		"P@ssw0rd",
		"avatar",
	)

	assert.Nil(t, err)
	assert.Greater(t, userId, int64(0))

	user, err := repo.GetUserByEmail("josh@tee_a.com")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, int64(userId), int64(user.ID))

	userId, err = repo.CreateUser(
		"Josh Tee A",
		"josh@tee_a.com",
		generateString(73),
		"avatar",
	)
	assert.Error(t, err)
	assert.Equal(t, int64(userId), int64(0))

}

func TestSQLUserRepository_CreateUser_DuplicateEmail(t *testing.T) {
	defer cleanupTestData(t)

	repo := NewSQLUserRepository(testDB)
	userId, err := repo.CreateUser(
		"Josh Tee A",
		"josh@tee_a.com",
		"P@ssw0rd",
		"avatar",
	)
	assert.Nil(t, err)
	assert.Greater(t, userId, int64(0))

	_, err = repo.CreateUser(
		"Josh Tee A",
		"josh@tee_a.com",
		"P@ssw0rd",
		"avatar",
	)
	assert.Error(t, err)
}

func TestSQLUserRepository_Authenticate_Success(t *testing.T) {
	defer cleanupTestData(t)

	repo := NewSQLUserRepository(testDB)
	userId, err := repo.CreateUser(
		"Josh Tee A",
		"josh@tee_a.com",
		"P@ssw0rd",
		"avatar",
	)
	assert.Nil(t, err)
	assert.Greater(t, userId, int64(0))

	authId, err := repo.Authenticate("josh@tee_a.com", "P@ssw0rd")
	assert.NoError(t, err)
	assert.Equal(t, userId, int64(authId))

}

func TestSQLUserRepository_Authenticate_WrongPassword(t *testing.T) {
	defer cleanupTestData(t)

	repo := NewSQLUserRepository(testDB)
	userId, err := repo.CreateUser(
		"Josh Tee A",
		"josh@tee_a.com",
		"P@ssw0rd",
		"avatar",
	)
	assert.Nil(t, err)
	assert.Greater(t, userId, int64(0))

	_, err = repo.Authenticate("josh@tee_a.com", "invalidPassword")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidCredential, err)

}

func generateString(n int) string {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = 'e'
	}

	return string(buf)
}
