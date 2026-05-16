package main

import (
	"goGin/config"
	"goGin/db"
	"goGin/middleware"
	"goGin/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	cfg := config.Load()
	db.Connect(cfg)

	r := gin.New()
	
	r.Use(gin.Recovery())
	r.Use(middleware.Logger)

	routes.SetupRoutes(r)

	r.Run(":" + cfg.Port)
}
