package models

type TimeActionReceive struct {
	AreaId    string `json:"area_id"`
	City      string `json:"city"`
	Continent string `json:"continent"`
	Hour      int    `json:"hour"`
	Minute    int    `json:"minute"`
}

type Database struct {
	Id        int
	AreaId    string `json:"area_id"`
	City      string `json:"city"`
	Continent string `json:"continent"`
	Hour      int    `json:"hour"`
	Minute    int    `json:"minute"`
}

type GetTimeAction struct {
	ActionId   int    `json:"action_id"`
	ActionType int    `json:"action_type"`
	City       string `json:"city"`
	Continent  string `json:"continent"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
}

// LAST
type DataReceive struct {
	Token             string `json:"token"`
	City              string `json:"city"`
	Continent         string `json:"continent"`
	Hour              int    `json:"hour"`
	Minute            int    `json:"minute"`
	ReactionType      int    `json:"reaction_type"`
	ReactionServiceId int    `json:"reaction_service_id"`
	Message           string `json:"message"`
	ServerId          string `json:"server_id"`
	ChannelId         string `json:"channel_id"`
}

type TimeModelSendReaction struct {
	ReactionId string `json:"area_id"`
}
