package handlers

import (
	"encoding/json"
	"go-mongodb-api/pkg/redis"
	"go-mongodb-api/services"
	"net/http"

	"github.com/gorilla/mux"
)

type PokemonHandlers struct {
	service      *services.PokemonService
	redisService *redis.RedisService
}

func NewPokemonHandlers(service *services.PokemonService, redisService *redis.RedisService) *PokemonHandlers {
	return &PokemonHandlers{service: service, redisService: redisService}
}

func (h *PokemonHandlers) Pokemons() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		names := []string{"pikachu", "bulbasaur", "charmander"}

		pokemons, err := h.service.FetchMultiplePokemons(names)
		if err != nil {
			http.Error(w, "Failed to fetch Pokémon data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pokemons)
	}
}

func (h *PokemonHandlers) PokemonCache() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		if name == "" {
			http.Error(w, "Missing pokemon name", http.StatusBadRequest)
			return
		}
		pokemons, err := h.service.FetchPokemonWithCache(name, h.redisService)
		if err != nil {
			http.Error(w, "Failed to fetch Pokémon data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pokemons)
	}
}
