package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

type Pokemon struct {
	Number int `json:"number"`
	Pokemon string `json:"pokemon"`
	Form string `json:"form,omitempty"`
	Say string `json:"say"`
}

func main() {
	pokemon, err := parseJson()
	if err != nil {
		log.Fatalln(err)
	}

	rng := createRng()
	i := rng.Intn(len(pokemon)-1)

	p := pokemon[i]
	fmt.Println("A wild " + p.Pokemon + " appeared!")
	fmt.Println(p.Say)
}

func parseJson() ([]Pokemon, error) {
	var pokemon []Pokemon

	f, err := os.Open("pokemon.json")
	if err != nil {
		return pokemon, err
	}
	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(b, &pokemon)

	return pokemon, err
}

func createRng() *rand.Rand {
	src := rand.NewSource(time.Now().UnixNano())
	return rand.New(src)
}