package services

import (
	"server/db"
	"server/internal/models"
	"strconv"
)

func GetProductStocks(id string) ([]map[string]interface{}, error) {
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	stocks, err := models.GetSuppliersByProductID(db.DB, uint32(idInt))
	return stocks, err
}

func GetProductMissingSuppliers(id string) ([]map[string]interface{}, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	suppliers, err := models.GetSuppliersByProductID(db.DB, uint32(idInt))
	if err != nil {
		return nil, err
	}

	var ids []uint32 = make([]uint32, 0)
	ids = append(ids, 0)
	for _, supplier := range suppliers {
		ids = append(ids, supplier["id"].(uint32))
	}

	missing, err := models.GetProductMissingSuppliers(db.DB, ids)
	if err != nil {
		return nil, err
	}

	return missing, nil
}

func GetSupplierStocks(id string) ([]map[string]interface{}, error) {
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	stocks, err := models.GetSuppliersByProductID(db.DB, uint32(idInt))
	return stocks, err
}

func CreateStock(stock models.StockRequest) error {
	product_supplier := &models.ProductsSuppliers{
		ProductID:  stock.ProductID,
		SupplierID: stock.SupplierID,
		Stock:      stock.Stock,
	}
	return models.CreateStock(db.DB, product_supplier)
}

func UpdateStock(stock models.StockRequest) error {
	return models.UpdateStock(db.DB, stock.ProductID, stock.SupplierID, stock.Stock)
}

func DeleteStock(product_id string, supplier_id string) error {
	product_id_int, err := strconv.Atoi(product_id)
	if err != nil {
		return err
	}

	supplier_id_int, err := strconv.Atoi(supplier_id)
	if err != nil {
		return err
	}

	return models.DeleteStock(db.DB, uint32(product_id_int), uint32(supplier_id_int))
}
