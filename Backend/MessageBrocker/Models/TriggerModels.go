package models

type TriggerModelGateway struct {
	ReactionServiceId int    `json:"reaction_service_id"`
	ReactionId        string `json:"reaction_id"`
}

type TriggerdModelsSending struct {
	ReactionIdentifyer string `json:"reaction_identifyer"`
}
