package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"server/internal/models"
	"server/internal/services"
	"server/utils"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateStock(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.CreateStock = func(stock models.StockRequest) error {
					return nil
				}

				a := assert.New(t)
				payload := `{
					"product_id": 1,
					"supplier_id": 1,
					"stock": 1
				}`

				req, _ := http.NewRequest("POST", "/api/products-suppliers/", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(201, resp.Code)
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.CreateStock = func(stock models.StockRequest) error {
					return errors.New("error")
				}

				a := assert.New(t)
				payload := `{
					"product_id": 1,
					"supplier_id": 1,
					"stock": 1
				}`

				req, _ := http.NewRequest("POST", "/api/products-suppliers/", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(500, resp.Code)
			},
		},
	}

	r := utils.Setup()

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			scenario.run(t, r)
		})
	}
}

func TestUpdateStock(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.UpdateStock = func(stock models.StockRequest) error {
					return nil
				}

				a := assert.New(t)
				payload := `{
					"product_id": 1,
					"supplier_id": 1,
					"stock": 1
				}`

				req, _ := http.NewRequest("PUT", "/api/products-suppliers/", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(200, resp.Code)
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.UpdateStock = func(stock models.StockRequest) error {
					return errors.New("error")
				}

				a := assert.New(t)
				payload := `{
					"product_id": 1,
					"supplier_id": 1,
					"stock": 1
				}`

				req, _ := http.NewRequest("PUT", "/api/products-suppliers/", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(500, resp.Code)
			},
		},
	}

	r := utils.Setup()

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			scenario.run(t, r)
		})
	}
}

func TestDeleteStock(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.DeleteStock = func(product_id uint32, supplier_id uint32) error {
					return nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("DELETE", "/api/products-suppliers/1/1", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(200, resp.Code)
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.DeleteStock = func(product_id uint32, supplier_id uint32) error {
					return errors.New("error")
				}

				a := assert.New(t)
				req, _ := http.NewRequest("DELETE", "/api/products-suppliers/1/1", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(500, resp.Code)
			},
		},
	}

	r := utils.Setup()

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			scenario.run(t, r)
		})
	}
}
