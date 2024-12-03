package models

type ReactionReceiveData struct {
	UserToken          string `json:"user_token"`
	ServiceId          int    `json:"service_id"`
	ActionId           int    `json:"action_id"`
	ReactionIdentifyer int    `json:"reaction_identifyer"`
	UserEmail          string `json:"user_email"`
	Message            string `json:"message"`
}

type ActiveReactionData struct {
	ServiceId int `json:"service_id"`
	ActionId  int `json:"action_id"`
}
