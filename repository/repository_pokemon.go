package repository

import (
	"gorm.io/gorm"

	"pokemon-rest-api/listing"
)

type PokemonRepository interface {
	GetAllPokemons() ([]listing.Pokemon, error)
	GetPokemonByID(id int) (listing.Pokemon, error)
	AddPokemon(pokemon listing.Pokemon) (listing.Pokemon, error)
}

type pokemonRepository struct {
	db *gorm.DB
}

func NewPokemonRepository(db *gorm.DB) PokemonRepository {
	err := db.AutoMigrate(&listing.Pokemon{})
	if err != nil {
		return nil
	}

	return &pokemonRepository{
		db: db,
	}
}

func (r *pokemonRepository) GetAllPokemons() ([]listing.Pokemon, error) {
	var pokemons []listing.Pokemon
	result := r.db.Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}
	return pokemons, nil
}

func (r *pokemonRepository) GetPokemonByID(id int) (listing.Pokemon, error) {
	var pokemon listing.Pokemon
	result := r.db.First(&pokemon, id)
	if result.Error != nil {
		return listing.Pokemon{}, result.Error
	}
	return pokemon, nil
}

func (r *pokemonRepository) AddPokemon(pokemon listing.Pokemon) (listing.Pokemon, error) {
	result := r.db.Create(&pokemon)
	if result.Error != nil {
		return listing.Pokemon{}, result.Error
	}
	return pokemon, nil
}
