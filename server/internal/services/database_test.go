package services

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBtest *gorm.DB

func renewDatabase() {
	dsn := "root:pwd@tcp(127.0.0.1:3306)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("cannot connect to MySQL server:", err)
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS dashboard_test")
	if err != nil {
		log.Fatal("failed to drop database:", err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS dashboard_test")
	if err != nil {
		log.Fatal("failed to create database:", err)
	}
}

func TestMain(m *testing.M) {
	var err error
	renewDatabase()
	DBtest, err = gorm.Open(mysql.Open("root:pwd@tcp(127.0.0.1:3306)/dashboard_test"), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect", err)
	}
	os.Exit(m.Run())
}
