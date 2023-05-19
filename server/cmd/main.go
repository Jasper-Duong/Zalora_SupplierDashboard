package main

import (
	"log"

	"server/api/routes"
	"server/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	env := godotenv.Load()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	routes.LoadRoutes(r)
	r.Run()
}
