package handlers

import (
	"encoding/json"
	"go-mongodb-api/services"
	"net/http"
)

type PokemonHandlers struct {
	service *services.PokemonService
}

func NewPokemonHandlers(service *services.PokemonService) *PokemonHandlers {
	return &PokemonHandlers{service: service}
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
