package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Pokemon struct {
	ID int		`json:"id"`
	Name string	`json:"name"`
	Type string	`json:"type"`
}

var pokemons = []Pokemon{{
	ID:		1,
	Name:	"ヒトカゲ",
	Type:	"ほのお",
},
{
	ID:		2,
	Name:	"ゼニガメ",
	Type:	"みず",
},
{
	ID:		3,
	Name:	"フシギダネ",
	Type:	"くさ",
},
}

type PokemonJSON struct {
	Status int `json:"status"`
	Pokemons *[]Pokemon 
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello, World!!")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
