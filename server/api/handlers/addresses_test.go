package handlers_test

import (
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

type Scenario struct {
	name string
	run  func(t *testing.T, r *gin.Engine)
}

type CreateAddressService func(data *models.AddressCreate) error
type UpdateAddressService func(data *models.AddressUpdate, id uint32) error
type DeleteAddressService func(id uint32) error

var runTests = func(t *testing.T, scenarios []Scenario) {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.LoadRoutes(r)

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			scenario.run(t, r)
		})
	}
}

func TestCreateAddress(t *testing.T) {
	var runFunc = func(t *testing.T, r *gin.Engine, mockService CreateAddressService, payload string, expStatusCode int) {
		services.CreateAddress = mockService
		req, _ := http.NewRequest("POST", "/api/addresses/", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		a := assert.New(t)
		a.Equal(expStatusCode, res.Code)
	}
	scenarios := []Scenario{
		{
			name: "Success",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(address *models.AddressCreate) error {
					return nil
				}
				payload := `{
					"supplier_id": 1,
					"type": "warehouse",
					"address_info": "Street 1, District 1"
					}`

				runFunc(t, r, mockService, payload, http.StatusCreated)
			},
		},
		{
			name: "Supplier ID not found",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(address *models.AddressCreate) error {
					return gorm.ErrRecordNotFound
				}
				payload := `{
					"supplier_id": 10000,
					"type": "warehouse",
					"address_info": "Street 10000, District 10000"
					}`
				runFunc(t, r, mockService, payload, http.StatusNotFound)
			},
		},
		{
			name: "Invalid Payload",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(address *models.AddressCreate) error {
					return gorm.ErrRecordNotFound
				}
				// Missing type
				payload := `{
					"supplier_id": 10000,
					"address_info": "Street 10000, District 10000"
					}`
				runFunc(t, r, mockService, payload, http.StatusBadRequest)
			},
		},
		{
			name: "Bad Request",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(address *models.AddressCreate) error {
					return nil
				}
				// Wrong type sup_id, missing address_info
				payload := `{
					"supplier_id": "THIS IS WRONG",
					"type": "office",
				}`
				runFunc(t, r, mockService, payload, http.StatusBadRequest)
			},
		},
	}
	runTests(t, scenarios)
}

func TestUpdateAddress(t *testing.T) {
	var runFunc = func(t *testing.T, r *gin.Engine, id string, mockService UpdateAddressService, payload string, expStatusCode int) {
		services.UpdateAddress = mockService
		req, _ := http.NewRequest("PUT", "/api/addresses/"+id, strings.NewReader(payload))
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		a := assert.New(t)
		a.Equal(expStatusCode, res.Code)
	}
	scenarios := []Scenario{
		{
			name: "Success",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(address *models.AddressUpdate, id uint32) error {
					return nil
				}
				payload := `{
					"type": "warehouse",
					"address_info": "Street X, District Y"
				}`
				addressID := "1"
				runFunc(t, r, addressID, mockService, payload, http.StatusOK)
			},
		},
		{
			name: "Address Not Found",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(address *models.AddressUpdate, id uint32) error {
					return gorm.ErrRecordNotFound
				}
				payload := `{
					"type": "warehouse",
					"address_info": "abc"
				}`
				addressID := "10000"
				runFunc(t, r, addressID, mockService, payload, http.StatusNotFound)
			},
		},
		{
			name: "Invalid Payload",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(address *models.AddressUpdate, id uint32) error {
					return nil
				}
				payload := `{
					"type": "warehouse",
				}`
				addressID := "1"
				runFunc(t, r, addressID, mockService, payload, http.StatusBadRequest)
			},
		},
		{
			name: "Bad Request",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(address *models.AddressUpdate, id uint32) error {
					return gorm.ErrRecordNotFound
				}
				payload := `{
					"type": "warehouse",
					"address_info": null
				}`
				addressID := "10000"
				runFunc(t, r, addressID, mockService, payload, http.StatusBadRequest)
			},
		},
	}
	runTests(t, scenarios)
}

func TestDeleteAddress(t *testing.T) {
	var runFunc = func(t *testing.T, r *gin.Engine, id string, mockService func(id uint32) error, expStatusCode int) {
		services.DeleteAddress = mockService
		req, _ := http.NewRequest("DELETE", "/api/addresses/"+id, nil)
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		a := assert.New(t)
		a.Equal(expStatusCode, res.Code)

	}
	scenarios := []Scenario{
		{
			name: "Success",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(id uint32) error {
					return nil
				}
				addressID := "2"
				runFunc(t, r, addressID, mockService, http.StatusOK)
			},
		},
		{
			name: "AddressID Not Found",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(id uint32) error {
					return gorm.ErrRecordNotFound
				}
				addressID := "10000"
				runFunc(t, r, addressID, mockService, http.StatusNotFound)
			},
		},
		{
			name: "Bad Request",
			run: func(t *testing.T, r *gin.Engine) {
				mockService := func(id uint32) error {
					return nil
				}
				addressID := "hi"
				runFunc(t, r, addressID, mockService, http.StatusBadRequest)
			},
		},
	}
	runTests(t, scenarios)
}
