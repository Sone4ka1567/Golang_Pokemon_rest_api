package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"pokemon-rest-api/repository"
	"pokemon-rest-api/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}

	pokemonRepository := repository.NewPokemonRepository(db)
	pokemonServer := server.NewPokemonServer(pokemonRepository)

	router := httprouter.New()
	router.GET("/pokemons", pokemonServer.GetAllPokemons)
	router.GET("/pokemons/:id", pokemonServer.GetPokemonByID)
	router.POST("/pokemons", pokemonServer.AddPokemon)

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func setupDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}