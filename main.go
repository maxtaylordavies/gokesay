package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/maxtaylordavies/gokesay/config"
	"github.com/gobuffalo/packr/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	pokemon, err := parseJson()
	if err != nil {
		log.Fatalln(err)
	}

	rng := createRng()
	i := rng.Intn(len(pokemon)-1)

	p := pokemon[i]
	display(p.Say)
	printMessage(p.Pokemon)
}

func parseJson() ([]config.Pokemon, error) {
	var pokemon []config.Pokemon

	box := packr.New("StaticBox", "./static")
	b, err := box.Find("pokemon.json")
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

func display(img string) {
	fmt.Print("\033]1337;")
	fmt.Printf("File=inline=1")
	fmt.Print(";width=50")
	fmt.Print(";height=50")
	fmt.Print(":")
	fmt.Print(img)
	fmt.Print("\a\n")
}

func printPng(fn string) {
	r, err := os.Open(fn)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("\033]1337;")
	fmt.Printf("File=inline=1")
	fmt.Print(";width=50")
	fmt.Print(";height=50")
	fmt.Print(":")
	fmt.Printf("%s", base64.StdEncoding.EncodeToString(data))
	fmt.Print("\a\n")
}




