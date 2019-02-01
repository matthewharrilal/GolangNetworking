package main
import (
	"net/http"
	"fmt"
)
// Define structures from base level up


type Region struct {
	Descriptions []Description `json:"descriptions"` // Slice of description objects

	Identification int 	`json:"id"`

	IsMainSeries bool `json:"bool"`

	Name string `json:"name"`

	Names []Language `json:"names"`

	PokemonEntries []PokemonSpecies `json:"pokemon_entries"`
}

type Description struct { 
	Description string `json:"description"`

	Language Language `json:"language"`

}


type PokemonSpecies struct {
	EntryNumber int `json:"entry_number"`
	
	PokemonSpecies []Language `json:"pokemon_species"`
}


type Language struct {
	Name string `json:"name"`
	URL string 	`json:"url"`
}
