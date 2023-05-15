package services

import (
	"log"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBtest *gorm.DB

func TestMain(m *testing.M) {
	var err error
	DBtest, err = gorm.Open(mysql.Open("root:pwd@tcp(127.0.0.1:3306)/dashboard"), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect", err)
	}
	os.Exit(m.Run())
}
