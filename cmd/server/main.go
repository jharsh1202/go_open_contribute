// cmd/server/main.go
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
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Organization{}, &models.Project{}, &models.Goal{}, &models.Task{})
	if err != nil {
		log.Fatalf("Could not auto migrate database: %v", err)
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	routes.SetupRoutes(router, db, cfg.JWTSecret)
	// routes.SetupUserRoutes(router, db, cfg.JWTSecret)
	// routes.SetupOrganizationRoutes(router, db, cfg.JWTSecret)

	if err := router.Run(cfg.Server.Address); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
