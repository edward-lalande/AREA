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
