package models

var ReactionModelPath string = "Models/Reactions.json"

type ReactionReceiveData struct {
	AreaId       string `json:"area_id"`
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	ChannelID    string `json:"channel_id"`
	Message      string `json:"message"`
	GuildID      string `json:"guild_id"`
}

type ReactionGet struct {
	ReactionId   int    `json:"reaction_id"`
	ReactionType int    `json:"reaction_type"`
	ChannelID    string `json:"channel_id"`
	Message      string `json:"message"`
	GuildID      string `json:"guild_id"`
}

type ActiveReactionData struct {
	ServiceId int `json:"service_id"`
	ActionId  int `json:"action_id"`
}

type Reactions struct {
	Message   string `json:"message"`
	ChannelId string `json:"channel_id"`
	GuildId   string `json:"guild_id"`
}
