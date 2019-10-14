package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
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
	fmt.Println(p.Say)

	printMessage(p.Pokemon)
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

func printMessage(name string) {
	l := len(name) + 31
	fmt.Println("+" + strings.Repeat("-", l) + "+")
	fmt.Println("|"+ strings.Repeat(" ", l) + "|")
	fmt.Println("|       A wild " + name + " appeared!       |")
	fmt.Println("|"+ strings.Repeat(" ", l) + "|")
	fmt.Println("+" + strings.Repeat("-", l) + "+")
}