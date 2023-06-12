# Golang_Pokemon_rest_api

The task is to write a local REST API server with a local PostgreSQL database for storing Pokémon. The server should have only one endpoint: /pokemons. This endpoint should support the following operations:
- Add a Pokémon
- Get a list of all Pokémon
- Get a Pokémon by its ID

## Requirements

- https://github.com/julienschmidt/httprouter
- https://github.com/go-gorm/gorm
- The Pokémon structure is in listing.pokemon
- Return the correct HTTP status codes (see tests)
- The server listens on port 8080
