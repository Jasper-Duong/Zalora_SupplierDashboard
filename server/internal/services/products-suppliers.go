package services

import (
	"server/db"
	"server/internal/models"
	"strconv"
)

func GetProductStocks(id string) ([]map[string]interface{}, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	stocks, err := models.SelectProductStocks(db.DB, idInt)
	var stockMap []map[string]interface{}
	for _, stock := range stocks {
		stockMap = append(stockMap, map[string]interface{}{
			"id":    stock.Suppliers,
			"stock": stock.Stock,
		})
	}

	return stockMap, err
}

func GetSupplierStocks(id string) ([]map[string]interface{}, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	stocks, err := models.SelectSupplierStocks(db.DB, idInt)
	var stockMap []map[string]interface{}
	for _, stock := range stocks {
		stockMap = append(stockMap, map[string]interface{}{
			"id":    stock.Products,
			"stock": stock.Stock,
		})
	}

	return stockMap, err
}
