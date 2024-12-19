package models

import "encoding/json"

type BaseAction struct {
	// Action ID
	// required: true
	ActionID int `json:"action_id"`
}

type TypeTimeAction struct {
	BaseAction
	// ActionType of dateTime
	// required: true
	ActionType int `json:"action_type"`

	// Continent of the hour for the actions
	// required: true
	Continent string `json:"continent"`

	// City of the hour for the actions
	// required: true
	City string `json:"city"`

	// Hour of the the actions (0-24)
	// required: true
	Hour int `json:"hour"`

	// Minute of the the actions (0-59)
	// required: true
	Minute int `json:"minute"`
}

type BaseReaction struct {
	// Reactions ID
	// required: true
	ReactionID int `json:"reaction_id"`
}

type TypeDiscordReaction struct {
	BaseReaction
	// Type of reactions (here is for discord send_message reactions)
	// required: true
	ReactionType int `json:"reaction_type"`

	// Channel Id of the message to send
	// required: true
	ChannelID string `json:"channel_id"`

	// Message to send
	// required: true
	Message string `json:"message"`

	// Guild to create a channel
	// required: true
	GuildID string `json:"guild_id"`
}

type PayloadItem struct {
	// User token
	// required: true
	UserToken string `json:"user_token"`

	// Actions
	// required: true
	Action *json.RawMessage `json:"action,omitempty"`

	// Reactions
	// required: true
	Reactions []*json.RawMessage `json:"reactions,omitempty"`
}

type AreaDatabase struct {
	Id                int    `json:"id"`
	UserToken         string `json:"user_token"`
	AreaId            string `json:"area_id"`
	ServiceActionId   int    `json:"service_action_id"`
	ServiceReactionId int    `json:"service_reaction_id"`
}

type OauthCode struct {
	OauthCode string `json:"code"`
}
