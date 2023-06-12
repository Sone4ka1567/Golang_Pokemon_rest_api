package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pokemon-rest-api/listing"
	"pokemon-rest-api/repository"
)

type PokemonServer struct {
	pokemonRepository repository.PokemonRepository
}

func NewPokemonServer(pokemonRepository repository.PokemonRepository) *PokemonServer {
	return &PokemonServer{
		pokemonRepository: pokemonRepository,
	}
}

func (s *PokemonServer) GetAllPokemons(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pokemons, err := s.pokemonRepository.GetAllPokemons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(pokemons)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (s *PokemonServer) GetPokemonByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	pokemon, err := s.pokemonRepository.GetPokemonByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response, err := json.Marshal(pokemon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (s *PokemonServer) AddPokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var pokemon listing.Pokemon
	err := json.NewDecoder(r.Body).Decode(&pokemon)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newPokemon, err := s.pokemonRepository.AddPokemon(pokemon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(newPokemon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
