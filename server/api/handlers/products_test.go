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
	"gorm.io/gorm"
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

	r := utils.Setup()

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			scenario.run(t, r)
		})
	}
}

func TestGetProductByID(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetProductByID = func(id uint32) (models.Products, error) {
					return models.Products{
						ID:     1,
						Name:   "Shirt",
						Brand:  "Zalora",
						Sku:    "ZaloraShirt",
						Size:   "M",
						Color:  "Red",
						Status: boolPtr(false),
						Stock:  10,
					}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/1", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `{
					"id": 1,
					"name": "Shirt",
					"brand": "Zalora",
					"sku": "ZaloraShirt",
					"size": "M",
					"color": "Red",
					"status": false,
					"stock": 10
				}`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetProductByID = func(id uint32) (models.Products, error) {
					return models.Products{}, errors.New("error")
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/1", nil)
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

	r := utils.Setup()

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			scenario.run(t, r)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.UpdateProduct = func(product *models.Products, id uint32) error {
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

				req, _ := http.NewRequest("PUT", "/api/products/1", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				a.Equal(200, resp.Code)
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.UpdateProduct = func(product *models.Products, id uint32) error {
					return gorm.ErrRecordNotFound
				}

				a := assert.New(t)
				payload := `{
					"name": "Meoww",
					"brand": "Zalora",
					"size": "2",
					"color": "black",
					"status": true
				}`
				req, _ := http.NewRequest("PUT", "/api/products/1", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				a.Equal(404, resp.Code)
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

func TestDeleteProduct(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.DeleteProduct = func(id uint32) error {
					return nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("DELETE", "/api/products/1", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(200, resp.Code)
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.DeleteProduct = func(id uint32) error {
					return gorm.ErrRecordNotFound
				}

				a := assert.New(t)
				req, _ := http.NewRequest("DELETE", "/api/products/1", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(404, resp.Code)
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

func TestGetAttributeOfProducts(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetAttributeOfProducts = func(attribute string) ([]map[string]interface{}, error) {
					return []map[string]interface{}{
						{
							"brand": "Zalora",
						},
						{
							"brand": "Uniqlo",
						},
					}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/attribute/brand", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `[
					{
						"brand": "Zalora"
					},
					{
						"brand": "Uniqlo"
					}
				]`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetAttributeOfProducts = func(attribute string) ([]map[string]interface{}, error) {
					return nil, errors.New("error")
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/attribute/brand", nil)
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

func TestGetProductStocks(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetProductStocks = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{
						{
							"id":    1,
							"name":  "test 1",
							"stock": 10,
						},
						{
							"id":    2,
							"name":  "test 2",
							"stock": 20,
						},
					}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/1/stocks", nil)
				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				expected := `[
					{
						"id": 1,
						"name": "test 1",
						"stock": 10
					},
					{
						"id": 2,
						"name": "test 2",
						"stock": 20
					}
				]`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetProductStocks = func(id uint32) ([]map[string]interface{}, error) {
					return nil, errors.New("error")
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/1/stocks", nil)
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

func TestGetProductMissingSuppliers(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetProductMissingSuppliers = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{
						{
							"id":   1,
							"name": "test 1",
						},
						{
							"id":   2,
							"name": "test 2",
						},
					}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/1/suppliers/missing", nil)
				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				expected := `[
					{
						"id": 1,
						"name": "test 1"
					},
					{
						"id": 2,
						"name": "test 2"
					}
				]`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetProductMissingSuppliers = func(id uint32) ([]map[string]interface{}, error) {
					return nil, errors.New("error")
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/1/suppliers/missing", nil)
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

func TestGetProductSuppliers(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSuppliersByProductID = func(id string) ([]map[string]interface{}, error) {
					return []map[string]interface{}{
						{
							"id":   1,
							"name": "test 1",
						},
						{
							"id":   2,
							"name": "test 2",
						},
					}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/1/suppliers", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `[
					{
						"id": 1,
						"name": "test 1"
					},
					{
						"id": 2,
						"name": "test 2"
					}
				]`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSuppliersByProductID = func(id string) ([]map[string]interface{}, error) {
					return nil, errors.New("error")
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/products/1/suppliers", nil)
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
