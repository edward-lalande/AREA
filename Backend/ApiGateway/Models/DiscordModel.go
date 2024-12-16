package models

type DiscordGet struct {
	Routes string `json:"routes"`
}

type DiscordPost struct {
	Routes string `json:"routes"`
	Code   string `json:"code"`
	Token  string `json:"token"`
}

type DiscordPostOatuh struct {
	Code  string `json:"code"`
	Token string `json:"token"`
}
type DiscordReactionDatabase struct {
	Id           int    `json:"id"`
	AreaId       string `json:"area_id"`
	ReactionType int    `json:"reaction_type"`
	UserToken    string `json:"user_token"`
	ChannelId    string `json:"channel_id"`
	Message      string `json:"message"`
}
