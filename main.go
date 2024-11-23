package main

import (
	"log"
	"net/http"

	"go-mongodb-api/config"
	"go-mongodb-api/db"
	"go-mongodb-api/handlers"
)

func main() {

	port := config.GetConfig("APP_PORT")
	client, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	handlers.InitUserHandlers(client)

	http.HandleFunc("/users", handlers.GetUsers())
	http.HandleFunc("/create", handlers.CreateUser())

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
