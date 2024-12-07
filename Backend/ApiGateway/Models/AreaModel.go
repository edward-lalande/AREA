package models

import (
	"encoding/json"
)

type BaseAction struct {
	ActionID int `json:"action_id"`
}

type TypeTimeAction struct {
	BaseAction
	ActionType int    `json:"action_type"`
	Continent  string `json:"continent"`
	City       string `json:"city"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
}

type BaseReaction struct {
	ReactionID int `json:"reaction_id"`
}

type TypeDiscordReaction struct {
	BaseReaction
	ReactionType int    `json:"reaction_type"`
	ChannelID    string `json:"channel_id"`
	Message      string `json:"message"`
}

type PayloadItem struct {
	UserToken string           `json:"user_token"`
	Action    *json.RawMessage `json:"action,omitempty"`
	Reaction  *json.RawMessage `json:"reaction,omitempty"`
}
