package models

type DataReceive struct {
	Token     string `json:"token"`
	City      string `json:"city"`
	Continent string `json:"continent"`
	Hour      string `json:"hour"`
	Minute    string `json:"minute"`
}

type Database struct {
	Id        int
	Mail      string
	City      string
	Continent string
	Hour      string
	Minute    string
}
