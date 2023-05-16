package tests

import (
	"log"
	"server/db"
	"server/db/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	dsn := "root:pwd@tcp(127.0.0.1:3306)/"
	db_name := "dashboard_test"
	db.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db")
	}

	err = db.DB.Exec("DROP DATABASE IF EXISTS " + db_name).Error
	if err != nil {
		log.Fatal("Failed to drop database")
	}

	err = db.DB.Exec("CREATE DATABASE IF NOT EXISTS " + db_name).Error
	if err != nil {
		log.Fatal("Failed to create database")
	}

	dsn = dsn + db_name
	db.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	migrations.MigrateUp(db.DB)
}
