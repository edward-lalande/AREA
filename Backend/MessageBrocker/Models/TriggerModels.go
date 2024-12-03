package models

type TriggerModelGateway struct {
	ReactionServiceId int `json:"reaction_service_id"`
	ReactionId        int `json:"reaction_id"`
}

type TriggerdModelsSending struct {
	ReactionIdentifyer int `json:"reaction_identifyer"`
}
