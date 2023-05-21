package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"server/api/routes"
	"server/internal/models"
	"server/internal/services"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func boolPtr(b bool) *bool {
	return &b
}
func TestGetProducts(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetProducts = func(query *models.ProductsQueryParams) ([]models.ProductsWithSuppliers, uint32, error) {
					return []models.ProductsWithSuppliers{
						{
							Products: models.Products{
								ID:     1,
								Name:   "Shirt",
								Brand:  "Zalora",
								Sku:    "ZaloraShirt",
								Size:   "M",
								Color:  "Red",
								Status: boolPtr(false),
								Stock:  10,
							},
							Suppliers: []map[string]interface{}{
								{
									"id":   1,
									"name": "test",
								},
							},
						},
					}, 1, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/?page=1&limit=10", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `{
					"products": [
						{
							"id": 1,
							"name": "Shirt",
							"brand": "Zalora",
							"sku": "ZaloraShirt",
							"size": "M",
							"color": "Red",
							"status": false,
							"stock": 10,
							"suppliers": [
								{
									"id": 1,
									"name": "test"
								}
							]
						}
					],
					"total": 1
				}`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
	}

	r := gin.Default()
	routes.LoadRoutes(r)

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			scenario.run(t, r)
		})
	}
}

func TestCreateProduct(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.CreateProduct = func(product *models.Products) error {
					return nil
				}

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
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.CreateProduct = func(product *models.Products) error {
					return errors.New("error")
				}

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

				a.Equal(500, resp.Code)
			},
		},
	}

	r := gin.Default()
	routes.LoadRoutes(r)

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			scenario.run(t, r)
		})
	}
}
