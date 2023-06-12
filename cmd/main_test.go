package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os/exec"
	"testing"

	"pokemon-rest-api/listing"

	client "github.com/anatolio-deb/pokemon-async-client"
)

func TestAddPokemon(t *testing.T) {
	c := exec.Command("go", "run", "main.go")

	if err := c.Start(); err != nil {
		t.Error(err)
	}

	defer func() {
		if err := c.Process.Kill(); err != nil {
			t.Error(err)
		}
	}()

	pokemons := client.GetPokemons()

	for _, p := range pokemons {
		j, err := json.Marshal(p)

		if err != nil {
			t.Error(err)
		}

		r, err := http.Post("http://localhost:8080/pokemons", "application/json", bytes.NewBuffer(j))

		if err != nil {
			t.Error(err)
		}

		r.Body.Close()

		if r.StatusCode != http.StatusCreated {
			t.Error(r.StatusCode)
		}
	}

	var pokemons2 []listing.Pokemon
	res, err := http.Get("http://localhost:8080/pokemons")

	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Error(err)
	}

	if err := json.Unmarshal(body, &pokemons2); err != nil {
		t.Error(err)
	}

	if len(pokemons) != len(pokemons2) {
		t.Error("not enough pokemons mined")
	}

	for _, i := range pokemons2 {
		var x interface{} = i
		p, ok := x.(listing.Pokemon)
		if !ok {
			t.Error("Неверный тип покемона", p)
		}
	}
}

func TestGetPokemon(t *testing.T) {
	r, err := http.Get("http://localhost:8080/pokemons/1")

	if err != nil {
		t.Error(err)
	}

	b, err := io.ReadAll(r.Body)

	if err != nil {
		t.Error(err)
	}

	var p listing.Pokemon
	json.Unmarshal(b, &p)

	if p.ID != 1 {
		t.Errorf("wrong pokemon.ID=%v; want=1", p.ID)
	}
}
