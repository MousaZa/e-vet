package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
