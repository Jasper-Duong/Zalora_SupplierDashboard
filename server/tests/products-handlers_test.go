package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
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

	assert.Equal(t, 201, resp.Code)
}
