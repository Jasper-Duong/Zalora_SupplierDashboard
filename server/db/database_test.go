package db

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	dsn := os.Getenv("CONNECTION_URL")
	fmt.Print(dsn)
	/*_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db")
	}*/
	os.Exit(m.Run())
}
