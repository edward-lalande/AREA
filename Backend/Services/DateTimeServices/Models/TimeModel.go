package models

type DataReceive struct {
	Token             string `json:"token"`
	City              string `json:"city"`
	Continent         string `json:"continent"`
	Hour              int    `json:"hour"`
	Minute            int    `json:"minute"`
	ReactionServiceId int    `json:"reaction_service_id"`
}

type Database struct {
	Id        int
	Mail      string
	City      string
	Continent string
	Hour      int
	Minute    int
}
