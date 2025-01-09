package pokeapi

import (
	"bootdev/emiel/pokedex/internal/pokecache"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var CaughtPokemonNames = make(map[string]struct{})

type Config struct {
	NextUrl     string
	PreviousUrl string
}

var config Config
var cache = pokecache.NewCache(10 * time.Minute)

const locationBaseUrl = "https://pokeapi.co/api/v2/location-area/"
const pokemonBaseUrl = "https://pokeapi.co/api/v2/pokemon/"

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationDetails struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func Init(initValues Config) {
	config.NextUrl = initValues.NextUrl
	config.PreviousUrl = initValues.PreviousUrl
}

func GetLocationDetails(locationName string) (LocationDetails, error) {
	cacheKey := "LocationDetails:" + locationName
	cachedResponse, found := cache.Get(cacheKey)
	if found {
		var locationDetails LocationDetails
		err := json.Unmarshal(cachedResponse, &locationDetails)
		if err == nil {
			return locationDetails, nil
		}
	}

	res, err := http.Get(locationBaseUrl + locationName)
	if err != nil {
		return LocationDetails{}, fmt.Errorf("error getting location details: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, fmt.Errorf("error reading body from location details response: %v", err)
	}
	cache.Add(cacheKey, body)

	locationsDetails := LocationDetails{}
	err = json.Unmarshal(body, &locationsDetails)
	if err != nil {
		return LocationDetails{}, fmt.Errorf("this location does not exist")
	}

	return locationsDetails, nil
}

func GetLocations(usePrevious bool) (Locations, error) {
	url := locationBaseUrl

	if usePrevious {
		if config.PreviousUrl == "" {
			return Locations{}, fmt.Errorf("you're on the first page")
		}
		url = config.PreviousUrl
	} else if config.NextUrl != "" {
		url = config.NextUrl
	}

	cacheKey := url + strconv.FormatBool(usePrevious)
	cachedResponse, found := cache.Get(cacheKey)
	if found {
		var locations Locations
		err := json.Unmarshal(cachedResponse, &locations)
		if err == nil {
			return locations, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return Locations{}, fmt.Errorf("error using GET on pokeapi location: %v", err)
	}
	if res.StatusCode > 299 {
		return Locations{}, fmt.Errorf("using GET on API endpoint resulted in status code: %v", res.StatusCode)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, fmt.Errorf("using GET on API endpoint resulted in error: %v", err)
	}
	cache.Add(cacheKey, body)

	locations := Locations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return Locations{}, fmt.Errorf("error using unmarshal on json code: %v", err)
	}

	config.NextUrl = locations.Next
	config.PreviousUrl = locations.Previous

	return locations, nil
}

type PokemonDetails struct {
	Height         int    `json:"height"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Species        struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func GetPokemonDetails(pokemonName string) (PokemonDetails, error) {

	cacheKey := pokemonBaseUrl + ":" + pokemonName
	cachedResponse, found := cache.Get(cacheKey)
	if found {
		var pokemonDetails PokemonDetails
		err := json.Unmarshal(cachedResponse, &pokemonDetails)
		if err == nil {
			return pokemonDetails, nil
		}
	}

	res, err := http.Get(pokemonBaseUrl + pokemonName)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("error using GET on pokeapi pokemon: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("error reading body from pokemon details response: %v", err)
	}

	pokemonDetails := PokemonDetails{}
	err = json.Unmarshal(body, &pokemonDetails)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("this pokemon does not exist")
	}

	return pokemonDetails, nil
}
