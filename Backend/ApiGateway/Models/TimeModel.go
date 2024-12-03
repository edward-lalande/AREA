package models

type DateTimeResponse struct {
	Routes   string `json:"routes"`
	Datetime string `json:"datetime"`
}

type TimeDataReceive struct {
	Routes            string `json:"routes"`
	Token             string `json:"token"`
	City              string `json:"city"`
	Continent         string `json:"continent"`
	Hour              int    `json:"hour"`
	Minute            int    `json:"minute"`
	ReactionServiceId int    `json:"reaction_service_id"`
	Message           string `json:"message"`
}

type TimeDataToSend struct {
	Token             string `json:"token"`
	City              string `json:"city"`
	Continent         string `json:"continent"`
	Hour              int    `json:"hour"`
	Minute            int    `json:"minute"`
	ReactionServiceId int    `json:"reaction_service_id"`
	Message           string `json:"message"`
}
