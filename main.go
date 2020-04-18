package main

import (
	"log"
	"os"
)

func main() {
	// Set the router as the default one shipped with Gin
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := setupRouter()
	router.Run(":" + port)
}
