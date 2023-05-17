package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Happy_CreateProduct(t *testing.T) {
	a := assert.New(t)
	payload := `{
		"name": "Meoww",
		"brand": "Zalora",
		"size": "2",
		"color": "black",
		"status": true
	}`

	req, _ := http.NewRequest("POST", "/api/products/", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	a.Equal(201, resp.Code)
}

func Test_Happy_GetProducts(t *testing.T) {
	a := assert.New(t)

	req, _ := http.NewRequest("GET", "/api/products/?page=1&limit=10", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	expected := `{
		"total": 1,
		"products": [
			{
				"id": 1,
				"name": "Meoww",
				"brand": "Zalora",
				"sku": "Meo710-Zalora-black-2",
				"size": "2",
				"color": "black",
				"status": true,
				"stock": 0,
				"suppliers": []
			}
		]
	}`

	a.Equal(200, resp.Code)
	a.JSONEq(expected, resp.Body.String())
}

func Test_Happy_GetProductByID(t *testing.T) {
	a := assert.New(t)

	req, _ := http.NewRequest("GET", "/api/products/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	expected := `{
		"id": 1,
		"name": "Meoww",
		"brand": "Zalora",
		"sku": "Meo710-Zalora-black-2",
		"size": "2",
		"color": "black",
		"status": true,
		"stock": 0
	}`

	a.Equal(200, resp.Code)
	a.JSONEq(expected, resp.Body.String())
}

func Test_Happy_UpdateProduct(t *testing.T) {
	a := assert.New(t)
	payload := `{
		"name": "Meoww",
		"brand": "Zalora",
		"size": "2",
		"color": "black",
		"status": false
	}`
	req, _ := http.NewRequest("PUT", "/api/products/1", strings.NewReader(payload))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	a.Equal(200, resp.Code)
}

func Test_Happy_Delete(t *testing.T) {
	a := assert.New(t)
	req, _ := http.NewRequest("DELETE", "/api/products/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	a.Equal(200, resp.Code)
}
