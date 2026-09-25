package main

import (
	"log"

	"example.com/assesment-app/config"
	"example.com/assesment-app/controllers"
	"example.com/assesment-app/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: File .env tidak ditemukan, menggunakan environment variable dari sistem/Docker")
	}

	config.ConnectDatabase()
	if err := config.DB.AutoMigrate(&models.Event{}); err != nil {
		log.Fatalf("Gagal membuat tabel database: %v", err)
	}

	server := gin.Default()

	// Routes
	api := server.Group("/api")
	{
		api.POST("/events", controllers.CreateEvent)
		api.GET("/events", controllers.GetEvents)
	}

	server.Run(":3002")
}
