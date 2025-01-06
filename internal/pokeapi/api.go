package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	NextUrl     string
	PreviousUrl string
}

var config Config

const baseUrl = "https://pokeapi.co/api/v2/location-area/"

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
	url := baseUrl

	res, err := http.Get(url + locationName)
	if err != nil {
		return LocationDetails{}, fmt.Errorf("error getting location details: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, fmt.Errorf("error reading body from location details response: %v", err)
	}

	locationsDetails := LocationDetails{}
	err = json.Unmarshal(body, &locationsDetails)
	if err != nil {
		return LocationDetails{}, fmt.Errorf("error using unmarshal on json code: %v", err)
	}

	return locationsDetails, nil
}

func GetLocations(usePrevious bool) (Locations, error) {
	url := baseUrl

	if usePrevious {
		if config.PreviousUrl == "" {
			return Locations{}, fmt.Errorf("you're on the first page")
		}
		url = config.PreviousUrl
	} else if config.NextUrl != "" {
		url = config.NextUrl
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

	locations := Locations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return Locations{}, fmt.Errorf("error using unmarshal on json code: %v", err)
	}

	config.NextUrl = locations.Next
	config.PreviousUrl = locations.Previous

	return locations, nil
}
