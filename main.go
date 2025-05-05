package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaiminbhaduri/codeprac/models"
	"github.com/jaiminbhaduri/codeprac/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env variable not set")
	}

	// === LOGGING TO FILE ===
	logFileName := fmt.Sprintf("logs/gin_%s.log", time.Now().Format("2006-01-02_15-04-05"))
	f, err := os.Create(logFileName)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	gin.DefaultWriter = f
	gin.DefaultErrorWriter = f
	log.SetOutput(f)

	// Connect to DB and run migrations
	models.ConnectDB()

	// Setup Gin router from routes package
	router := routes.SetupRouter()

	ip := os.Getenv("IPADDR")
	router.SetTrustedProxies([]string{ip})

	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
