package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MousaZa/e-vet/models"
	"github.com/MousaZa/e-vet/server"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	s := server.New()
	w := httptest.NewRecorder()
	b := `
	{
		"username":"test",
		"password":"mousa1234",
		"email":"test@gmail.com"
	}
	`
	r, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(b))

	s.R.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLogin(t *testing.T) {
	s := server.New()
	w := httptest.NewRecorder()
	b := `
	{
		"password":"mousa1234",
		"email":"test@gmail.com"
	}
	`
	r, err := http.NewRequest(http.MethodPost, "/user/login", strings.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	s.R.ServeHTTP(w, r)
	rd, _ := io.ReadAll(w.Body)
	var lr models.LoginResponse
	err = json.Unmarshal(rd, &lr)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "test", lr.Username)
	assert.Equal(t, http.StatusOK, w.Code)
}
