package migrations

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func MigrateUpProducts(db *gorm.DB) {
	db.Table("products").AutoMigrate(&models.Products{})
}

func MigrateDownProducts(db *gorm.DB) {
	db.Migrator().DropTable("products")
}
