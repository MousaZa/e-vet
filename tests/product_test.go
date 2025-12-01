package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MousaZa/e-vet/models"
	"github.com/MousaZa/e-vet/server"
	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	s := server.New()
	w := httptest.NewRecorder()
	b := `
	{
		"user_id":1,
		"name":"test",
		"quantity":20,
		"last_price":4.99
	}
	`
	r, _ := http.NewRequest(http.MethodPost, "/stock/products", strings.NewReader(b))
	s.R.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProducts(t *testing.T) {
	s := server.New()
	w := httptest.NewRecorder()

	r, _ := http.NewRequest(http.MethodGet, "/stock/products", nil)
	s.R.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteProduct(t *testing.T) {
	s := server.New()
	gw := httptest.NewRecorder()

	gr, _ := http.NewRequest(http.MethodGet, "/stock/products", nil)
	s.R.ServeHTTP(gw, gr)
	var p []models.Product
	reader, err := io.ReadAll(gw.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(reader, &p)
	if err != nil {
		t.Fatal(err)
	}
	id := p[0].ID

	dw := httptest.NewRecorder()
	dr, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/stock/products/%s", id), nil)
	s.R.ServeHTTP(dw, dr)
	assert.Equal(t, http.StatusOK, dw.Code)
}
