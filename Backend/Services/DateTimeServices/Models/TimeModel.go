package models

type DataReceive struct {
	Token             string `json:"token"`
	City              string `json:"city"`
	Continent         string `json:"continent"`
	Hour              int    `json:"hour"`
	Minute            int    `json:"minute"`
	ReactionServiceId int    `json:"reaction_service_id"`
	Message           string `json:"message"`
}

type Database struct {
	Id                int
	Mail              string
	City              string
	Continent         string
	Hour              int
	Minute            int
	ReactionServiceId int
}

/*
{
    "service_id": 0,
    "action_id": 1,
    "reaction_identifyer": 2,
    "user_email": "ccaca",
    "message": "test"
}
*/
