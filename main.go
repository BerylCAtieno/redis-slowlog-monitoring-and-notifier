package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BerylCAtieno/redis-slowlog-monitor/api"
	"github.com/joho/godotenv"
)

func init() {
    // Load .env file from current directory
    if err := godotenv.Load(); err != nil {
        log.Printf("Warning: No .env file found: %v", err)
    }
}


func main() {
	http.HandleFunc("/integration.json", api.HandleIntegrationConfig)
	http.HandleFunc("/format-alert", api.HandleIncomingMessage)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}