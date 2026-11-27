package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"german-grammar-academy/backend/internal/academy"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://academy:academy_dev_password@localhost:5432/german_academy?sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}
	if err := academy.MigrateAndSeed(db); err != nil {
		log.Fatalf("migrate/seed: %v", err)
	}

	r := gin.Default()
	origin := os.Getenv("FRONTEND_ORIGIN")
	if origin == "" { origin = "http://localhost:4200" }
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if c.Request.Method == "OPTIONS" { c.AbortWithStatus(204); return }
		c.Next()
	})
	academy.RegisterRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	log.Printf("German Grammar Academy API listening on :%s", port)
	if err := r.Run(":" + port); err != nil { log.Fatal(err) }
}
