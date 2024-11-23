package routes

import (
	"go-mongodb-api/handlers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(client *mongo.Client) *mux.Router {
	router := mux.NewRouter()

	//start handlers for modules
	handlers.InitUserHandlers(client)

	// Grouping routes by modules
	users := router.PathPrefix("/users").Subrouter()
	users.HandleFunc("", handlers.GetUsers()).Methods("GET")
	users.HandleFunc("", handlers.CreateUser()).Methods("POST")
	users.HandleFunc("", handlers.UpdateUser()).Methods("PATCH")

	return router
}
