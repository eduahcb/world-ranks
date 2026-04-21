package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://restcountries.com/v3.1"
)

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Country struct {
	Area       float64 `json:"area"`
	Population int     `json:"population"`
	Region     string  `json:"region"`
	Name       struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	}
	Flags struct {
		Svg string `json:"svg"`
	}
	Independent bool                `json:"independent"`
	UnMember    bool                `json:"unMember"`
	Code        string              `json:"cca3"`
	Capital     []string            `json:"capital"`
	Subregion   string              `json:"subregion"`
	Languages   map[string]string   `json:"languages"`
	Currencies  map[string]Currency `json:"currencies"`
	Continents  []string            `json:"continents"`
}

func GetCountries() ([]Country, error) {
	url := fmt.Sprintf("%s/all?fields=name,population,area,region,flags,independent,unMember,cca3", baseURL)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch countries")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var countries []Country

	err = json.Unmarshal(body, &countries)

	if err != nil {
		return nil, err
	}

	return countries, nil
}

func GetCountryByCode(code string) (Country, error) {
	url := fmt.Sprintf("%s/alpha/%s/?fields=name,flags,population,area,capital,subregion,languages,currencies,continents", baseURL, code)

	res, err := http.Get(url)

	if err != nil {
		return Country{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Country{}, errors.New("failed to fetch country")
	}

	var country Country

	err = json.NewDecoder(res.Body).Decode(&country)

	if err != nil {
		return Country{}, err
	}

	return country, nil
}
