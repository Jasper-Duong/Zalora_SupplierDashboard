package migrations

import "gorm.io/gorm"

func MigrateUp(DB *gorm.DB) {
	MigrateUpSuppliers(DB)
	MigrateUpProducts(DB)
	MigrateUpProductsSuppliers(DB)
}
