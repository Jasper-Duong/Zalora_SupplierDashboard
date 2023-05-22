package migrations

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func MigrateUpAddresses(db *gorm.DB) {
	db.AutoMigrate(&models.Addresses{})
}

func MigrateDownAddresses(db *gorm.DB) {
	db.Migrator().DropTable("addresses")
}
