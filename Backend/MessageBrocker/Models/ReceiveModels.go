package models

type ReceivedAction struct {
	UserToken          string `json:"user_token"`
	Service            string `json:"service"`
	ActionId           int    `json:"action_id"`
	ReactionIdentifyer int    `json:"reaction_identifyer"`
}
