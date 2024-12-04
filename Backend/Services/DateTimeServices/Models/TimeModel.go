package models

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

type Database struct {
	Id                int
	Mail              string
	City              string
	Continent         string
	Hour              int
	Minute            int
	ReactionServiceId int
	ReactionId        string
}
type TimeModelSendReaction struct {
	ReactionServiceId int    `json:"reaction_service_id"`
	ReactionId        string `json:"reaction_id"`
}
