package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MousaZa/e-vet/server"
	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	s := server.NewServer()
	w := httptest.NewRecorder()
	b := `
	{
	"user_id":1,
	"name":"test",
	"quantity":20,
	"last_price":4.99
	}
	`
	r, _ := http.NewRequest(http.MethodPost, "/stock/product", strings.NewReader(b))
	s.R.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
