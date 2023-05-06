package migrations

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func MigrateUpSuppliers(db *gorm.DB) {
	db.AutoMigrate(&models.Suppliers{})
}

func MigrateDownSuppliers(db *gorm.DB) {
	db.Migrator().DropTable("suppliers")
}
