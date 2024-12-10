package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Pokemon struct {
	Name    string `json:"name"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type PokemonService struct{}

func NewPokemonService() *PokemonService {
	return &PokemonService{}
}

func (s *PokemonService) FetchPokemon(name string) (*Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch Pokémon %s: status code %d", name, resp.StatusCode)
	}

	var data Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *PokemonService) FetchMultiplePokemons(names []string) ([]*Pokemon, error) {
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	results := make([]*Pokemon, 0, len(names))
	var aggregatedErrors []error

	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			data, err := s.FetchPokemon(name)
			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				aggregatedErrors = append(aggregatedErrors, err)
			} else {
				results = append(results, data)
			}
		}(name)
	}

	wg.Wait()

	if len(aggregatedErrors) > 0 {
		return results, fmt.Errorf("errors occurred: %v", aggregatedErrors)
	}

	return results, nil
}
