package models

import "time"

type GlobalData struct {
	NewConfirmed   int `json:"NewConfirmed"`
	TotalConfirmed int `json:"TotalConfirmed"`
	NewDeaths      int `json:"NewDeaths"`
	TotalDeaths    int `json:"TotalDeaths"`
	NewRecovered   int `json:"NewRecovered"`
	TotalRecovered int `json:"TotalRecovered"`
}

type CountriesData []struct {
	Country        string    `json:"Country"`
	CountryCode    string    `json:"CountryCode"`
	Slug           string    `json:"Slug"`
	NewConfirmed   int       `json:"NewConfirmed"`
	TotalConfirmed int       `json:"TotalConfirmed"`
	NewDeaths      int       `json:"NewDeaths"`
	TotalDeaths    int       `json:"TotalDeaths"`
	NewRecovered   int       `json:"NewRecovered"`
	TotalRecovered int       `json:"TotalRecovered"`
	Date           time.Time `json:"Date"`
}

type GlobalCovid struct {
	Global    GlobalData    `json:"Global"`
	Countries CountriesData `json:"Countries"`
}
