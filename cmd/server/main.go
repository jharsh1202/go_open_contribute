package main

import (
	"log"
	"open-contribute/config"
	"open-contribute/models"
	"open-contribute/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Initialize database connection
	log.Printf("Connecting to the database: %v", cfg.Database.DSN)
	db, err := gorm.Open(mysql.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	log.Printf("Database connection established: %v", db)

	// Create a new Gin router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "message": "Database connection error"})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(500, gin.H{"status": "error", "message": "Database ping error"})
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Setup routes
	routes.SetupRoutes(router, db)

	// Auto migrate database schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Could not auto migrate database: %v", err)
	}

	// Start the server
	if err := router.Run(cfg.Server.Address); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
