package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/mrMaliosi/train-station/backend/internal/db"
	"github.com/mrMaliosi/train-station/backend/internal/routes"
)

func main() {
	dbConn := db.Init()
	defer dbConn.Close()

	r := gin.Default()
	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Укажите фронтенд адрес
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	routes.SetupRoutes(r, dbConn)

	serverPort := os.Getenv("PORT")
	log.Printf("🚀 Сервер запущен на порту %s", serverPort)
	if err := r.Run(":" + serverPort); err != nil {
		log.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}
