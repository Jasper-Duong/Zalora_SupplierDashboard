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
	"gorm.io/gorm"
)

func TestGetSuppliers(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSuppliers = func(query *models.SuppliersQueryParam) ([]models.SupplierWithAddresses, uint32, error) {
					return []models.SupplierWithAddresses{
						{
							Suppliers: models.Suppliers{
								ID:            1,
								Name:          "New supplier",
								Email:         "supplier@gmail.com",
								ContactNumber: "phone",
								Status:        new(bool),
								Stock:         10,
							},
							Addresses: []map[string]interface{}{
								{
									"id":           1,
									"type":         "office",
									"address_info": "address",
								},
							},
						},
					}, 1, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/?page=1&limit=10", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `{
					"suppliers": [
						{
							"id": 1,
							"name": "New supplier",
							"email": "supplier@gmail.com",
							"contact_number": "phone",
							"status": false,
							"stock": 10,
							"addresses": [
								{
									"id":   1,
									"type": "office",
									"address_info": "address"
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

func TestCreateSupplier(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.CreateSupplier = func(product *models.Suppliers) error {
					return nil
				}

				a := assert.New(t)
				payload := `{
					"name": "Supplier 2",
					"email": "supplier2@gmail.com",
					"contact_number": "+841234567890",
					"status": true
				}`

				req, _ := http.NewRequest("POST", "/api/suppliers/", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				a.Equal(201, resp.Code)
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.CreateSupplier = func(product *models.Suppliers) error {
					return errors.New("error")
				}

				a := assert.New(t)
				payload := `{
					"name": "Supplier 2",
					"email": "supplier2@gmail.com",
					"contact_number": "+841234567890",
					"status": true
				}`
				req, _ := http.NewRequest("POST", "/api/suppliers/", strings.NewReader(payload))
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

func TestUpdateSupplier(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.UpdateSupplier = func(product *models.Suppliers, id uint32) error {
					return nil
				}

				a := assert.New(t)
				payload := `{
					"name": "Supplier 2",
					"email": "supplier2@gmail.com",
					"contact_number": "+841234567890",
					"status": true
				}`

				req, _ := http.NewRequest("PUT", "/api/suppliers/1", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				a.Equal(200, resp.Code)
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.UpdateSupplier = func(product *models.Suppliers, id uint32) error {
					return errors.New("error")
				}

				a := assert.New(t)
				payload := `{
					"name": "Supplier 2",
					"email": "supplier2@gmail.com",
					"contact_number": "+841234567890",
					"status": true
				}`
				req, _ := http.NewRequest("PUT", "/api/suppliers/1", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				a.Equal(500, resp.Code)
			},
		},
		{
			name: "testcase 3",
			run: func(t *testing.T, r *gin.Engine) {
				services.UpdateSupplier = func(product *models.Suppliers, id uint32) error {
					return gorm.ErrRecordNotFound
				}

				a := assert.New(t)
				payload := `{
					"name": "Supplier 2",
					"email": "supplier2@gmail.com",
					"contact_number": "+841234567890",
					"status": true
				}`
				req, _ := http.NewRequest("PUT", "/api/suppliers/1", strings.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				a.Equal(404, resp.Code)
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

func TestDeleteSupplier(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.DeleteSupplier = func(id uint32) error {
					return nil
				}

				a := assert.New(t)

				req, _ := http.NewRequest("DELETE", "/api/suppliers/1", nil)
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				a.Equal(200, resp.Code)
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.DeleteSupplier = func(id uint32) error {
					return errors.New("error")
				}

				a := assert.New(t)

				req, _ := http.NewRequest("DELETE", "/api/suppliers/1", nil)
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				a.Equal(500, resp.Code)
			},
		},
		{
			name: "testcase 3",
			run: func(t *testing.T, r *gin.Engine) {
				services.DeleteSupplier = func(id uint32) error {
					return nil
				}

				a := assert.New(t)

				req, _ := http.NewRequest("DELETE", "/api/suppliers/", nil)
				req.Header.Set("Content-Type", "application/json")

				resp := httptest.NewRecorder()

				r.ServeHTTP(resp, req)

				a.Equal(404, resp.Code)
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

func TestGetSupplierByID(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSupplierByID = func(id uint32) (models.Suppliers, error) {
					return models.Suppliers{
						ID:            1,
						Name:          "New supplier",
						Email:         "supplier@gmail.com",
						ContactNumber: "phone",
						Status:        new(bool),
						Stock:         10,
					}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `{
					"id": 1,
					"name": "New supplier",
					"email": "supplier@gmail.com",
					"contact_number": "phone",
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
				services.GetSupplierByID = func(id uint32) (models.Suppliers, error) {
					return models.Suppliers{}, gorm.ErrRecordNotFound
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(404, resp.Code)
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

func TestGetSupplierAddresses(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSupplierAddresses = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{
						{
							"id":           1,
							"type":         "office",
							"address_info": "address",
						},
						{
							"id":           2,
							"type":         "office",
							"address_info": "address",
						},
					}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1/addresses", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `[
					{
						"id":           1,
						"type":         "office",
						"address_info": "address"
					},
					{
						"id":           2,
						"type":         "office",
						"address_info": "address"
					}
				]`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSupplierAddresses = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1/addresses", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `[
				]`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 3",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSupplierAddresses = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{}, gorm.ErrRecordNotFound
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1/addresses", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(404, resp.Code)
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

func TestGetSupplierStocks(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSupplierStocks = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{
						{
							"id":    1,
							"name":  "j",
							"stock": 10,
						},
					}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1/stocks", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `[
					{
						"id": 1,
						"name": "j",
						"stock": 10
					}
				]`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSupplierStocks = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1/stocks", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `[
				]`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 3",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSupplierStocks = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{}, gorm.ErrRecordNotFound
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1/stocks", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(404, resp.Code)
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

func TestGetSupplierMissingProducts(t *testing.T) {
	scenarios := []struct {
		name string
		run  func(t *testing.T, r *gin.Engine)
	}{
		{
			name: "testcase 1",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSupplierMissingProducts = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{
						{
							"id":   1,
							"name": "j",
						},
					}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1/products/missing", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `[
					{
						"id": 1,
						"name": "j"
					}
				]`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 2",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSupplierMissingProducts = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{}, nil
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1/products/missing", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				expected := `[
				]`

				a.Equal(200, resp.Code)
				a.JSONEq(expected, resp.Body.String())
			},
		},
		{
			name: "testcase 3",
			run: func(t *testing.T, r *gin.Engine) {
				services.GetSupplierMissingProducts = func(id uint32) ([]map[string]interface{}, error) {
					return []map[string]interface{}{}, gorm.ErrRecordNotFound
				}

				a := assert.New(t)
				req, _ := http.NewRequest("GET", "/api/suppliers/1/products/missing", nil)
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)

				a.Equal(404, resp.Code)
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
