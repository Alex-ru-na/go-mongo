package routes

import (
	"go-mongodb-api/handlers"
	"go-mongodb-api/middlewares"
	"go-mongodb-api/pkg/redis"
	"go-mongodb-api/pkg/websocket"
	"go-mongodb-api/repositories"
	"go-mongodb-api/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(client *mongo.Client) *mux.Router {
	router := mux.NewRouter()

	manager := websocket.NewWebSocketManager()
	socketHandler := websocket.NewWebSocketHandler(manager)

	redisService := redis.NewRedisService("localhost:6379", "", 0) // Redis configuration

	userRepo := repositories.NewUserRepository(client)
	userService := services.NewUserService(userRepo, manager)
	userHandlers := handlers.NewUserHandlers(userService)

	authService := services.AuthService{Repo: userRepo}
	authHandlers := handlers.NewAuthHandlers(&authService)

	pokemonService := services.NewPokemonService()
	pokemonHandlers := handlers.NewPokemonHandlers(pokemonService, redisService)

	// Create the AuthMiddleware
	authMiddleware := middlewares.NewAuthMiddleware(&authService)

	// Grouping routes users
	users := router.PathPrefix("/users").Subrouter()
	users.HandleFunc("", userHandlers.GetUsers()).Methods("GET")
	users.HandleFunc("/{id}", userHandlers.GetUser()).Methods("GET")
	users.HandleFunc("", userHandlers.CreateUser()).Methods("POST")
	users.HandleFunc("", userHandlers.UpdateUser()).Methods("PATCH")

	// Protect user routes with AuthMiddleware
	users.Use(authMiddleware.Protect)

	// Grouping routes auth
	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", authHandlers.Login()).Methods("POST")

	// WebSocket route
	router.HandleFunc("/ws", socketHandler.HandleConnection)

	// pokemon
	router.HandleFunc("/pokemons", pokemonHandlers.Pokemons()).Methods("GET")
	router.HandleFunc("/pokemons/{name}", pokemonHandlers.PokemonCache()).Methods("GET")

	return router
}
