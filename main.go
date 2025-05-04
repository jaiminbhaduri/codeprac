package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jaiminbhaduri/codeprac/db"
	"github.com/jaiminbhaduri/codeprac/routes"
	"github.com/joho/godotenv"
)

func init() {
	//tmpl, _ := template.ParseGlob("templates/*.html")
}

func main() { // Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT env not set")
		return
	}

	dbcon, dbconerr := db.DB()
	if dbconerr != nil {
		fmt.Println("DB_CONN_ERR", dbconerr)
		return
	}
	defer dbcon.Close()

	if err := db.InitDB(); err != nil {
		fmt.Println("INITDB_ERROR", err)
		return
	}

	// Setup routes from external file
	router := routes.SetupRoutes()

	// Start server
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
