package models

type ReactionReceiveData struct {
	UserToken          string `json:"user_token"`
	ReactionIdentifyer string `json:"reaction_identifyer"`
	ReactionType       int    `json:"reaction_type"`
	Message            string `json:"message"`
	ChannelId          string `json:"channel_id"`
}

type ActiveReactionData struct {
	ServiceId int `json:"service_id"`
	ActionId  int `json:"action_id"`
}

type Reactions struct {
	Message   string `json:"message"`
	ChannelId string `json:"channel_id"`
}
