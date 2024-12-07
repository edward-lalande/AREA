package models

type ReactionReceiveData struct {
	AreaId       string `json:"area_id"`
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	ChannelID    string `json:"channel_id"`
	Message      string `json:"message"`
}

type ActiveReactionData struct {
	ServiceId int `json:"service_id"`
	ActionId  int `json:"action_id"`
}

type Reactions struct {
	Message   string `json:"message"`
	ChannelId string `json:"channel_id"`
}
