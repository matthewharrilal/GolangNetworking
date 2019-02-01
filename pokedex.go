package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Define structures from base level up

type Region struct {
	Descriptions []Description `json:"descriptions"` // Slice of description objects

	Identification int `json:"id"`

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

	PokemonSpecies Language `json:"pokemon_species"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	url := "http://pokeapi.co/api/v2/pokedex/kanto/"

	// Allows you to add added configurations to the http request and execute you can construct the request and execute with the client which add specified configurations
	client := http.Client{
		// If for some reason the request to the pokedex lags it will take awhile timeout allowed for request ... no more than a minute
		Timeout: time.Second * 60,
	}

	// To construct a new request HTTP Method, desired url, and nil?
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		// Checking if error present
		log.Fatal(err)
	}

	// Execute request and assign callback status
	res, getErr := client.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	// All available data is contained to res.body use the input output utilities to read the contents until the end of file EOF occurs
	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		// See if error exists after reading the  contents of the file
		log.Fatal(readErr)
	}

	region := &Region{}

	fmt.Println("This is the response body >>>> ", &region)
	// Unmarshal json data and writes to the team object at that spot in memory
	parseErr := json.Unmarshal([]byte(body), *region)

	if parseErr != nil {
		fmt.Println(parseErr)
		return
	}

	for _, region := range region.Descriptions {
		fmt.Println("REGIONS ", region)
	}
}
