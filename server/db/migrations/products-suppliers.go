package migrations

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func MigrateUpProductsSuppliers(db *gorm.DB) {
	db.Table("products_suppliers").AutoMigrate(&models.ProductsSuppliers{})
}

func MigrateDownProductsSuppliers(db *gorm.DB) {
	db.Migrator().DropTable("products_suppliers")
}
