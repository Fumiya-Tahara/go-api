package main

import (
	"fmt"
)

type Pokemon struct {
	ID int		`json:"id"`
	Name string	`json:"name"`
	Type string	`json:"type"`
}

var pokemons = []Pokemon{{
	ID:		1,
	Name:	"Hitokage",
	Type:	"Fire",
},
{
	ID:		2,
	Name:	"Zenigame",
	Type:	"Water",
},
{
	ID:		2,
	Name:	"Fushigidane",
	Type:	"Grass",
},
}

func main()  {
	fmt.Println(pokemons[0])
}
