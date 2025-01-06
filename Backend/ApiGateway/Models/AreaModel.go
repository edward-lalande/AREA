package models

import "encoding/json"

type BaseAction struct {
	// Action ID
	// required: true
	ActionID  int    `json:"action_id"`
	UserToken string `json:"user_token"`
	Name      string `json:"name"`
}

type GoogleAction struct {
	BaseAction
	UserToken  string `json:"user_token"`
	AreaId     string `json:"area_id"`
	ActionType int    `json:"action_type"`
	NbEvents   int    `json:"nb_events"`
}

type MeteoActions struct {
	BaseAction
	ActionType int    `json:"action_type"`
	AreaId     string `json:"area_id"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Value      int    `json:"value"`
}

type GitlabAction struct {
	BaseAction
	ActionType int    `json:"action_type"`
	AreaId     string `json:"area_id"`
}

type SpotifyActions struct {
	AreaId      string `json:"area_id"`
	ActionType  int    `json:"action_type"`
	AccessToken string `json:"user_token"`
}

type DropBoxReactions struct {
	AreaId        string `json:"area_id"`
	UserToken     string `json:"user_token"`
	ReactionType  int    `json:"reaction_type"`
	FromPath      string `json:"from_path"`
	ToPath        string `json:"to_path"`
	FilepathShare string `json:"filepath_share"`
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

type TypeDiscordAction struct {
	BaseAction
	ActionType int    `json:"action_type"`
	ChannelId  string `json:"channel_id"`
	MessageId  string `json:"message_id"`
}

type CryptoMoneyActions struct {
	BaseAction
	AreaId     string `json:"area_id"`
	ActionType int    `json:"action_type"`
	Symbole    string `json:"symbole"`
	Devise     string `json:"devise"`
	Value      int    `json:"value"`
}

type TicketMasterAction struct {
	BaseAction
	AreaID     string `json:"area_id"`
	ActionType int    `json:"action_type"`
	Name       string `json:"name"`
	Venue      string `json:"venue"`
	City       string `json:"city"`
	NbEvents   int    `json:"nb_events"`
}

type AsanaReactions struct {
	BaseAction
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	AreaId       string `json:"area_id"`
	ProjectName  string `json:"project_name"`
	WorkSpaceId  string `json:"workspace_id"`
	Note         string `json:"note"`
	ProjectId    string `json:"project_id"`
	TaskId       string `json:"task_id"`
}

type BaseReaction struct {
	// Reactions ID
	// required: true
	ReactionID int    `json:"reaction_id"`
	Name       string `json:"name"`
}

type SpotifyReactions struct {
	BaseReaction
	AreaId       string `json:"area_id"`
	ReactionType int    `json:"reaction_type"`
	AccessToken  string `json:"user_token"`
}

type GitlabReactions struct {
	BaseReaction
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	AreaId       string `json:"area_id"`
	ProjectId    string `json:"project_id"`
	Body         string `json:"body"`
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

type Pusher struct {
	Name string `json:"name"`
}

type Commit struct {
	Message  string   `json:"message"`
	Added    []string `json:"added"`
	Removed  []string `json:"removed"`
	Modified []string `json:"modified"`
}

type User struct {
	Login string `json:"login"`
}

type Reactions struct {
	TotalCount int `json:"total_count"`
}

type Comment struct {
	User      User      `json:"user"`
	Body      string    `json:"body"`
	Reactions Reactions `json:"reactions"`
}

type TypeGithubAction struct {
	BaseAction

	ActionType int    `json:"action_type"`
	UserToken  int    `json:"user_token"`
	Pusher     string `json:"pusher"`
	Value      string `json:"value"`
	Number     int    `json:"number"`
}

type GoogleReaction struct {
	BaseReaction
	UserToken    string `json:"user_token"`
	AreaId       string `json:"area_id"`
	ReactionType int    `json:"reaction_type"`
	Summary      string `json:"summary"`
	Description  string `json:"description"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Attendees    string `json:"attendees"`
	Recipient    string `json:"recipient"`
	Subject      string `json:"subject"`
	Message      string `json:"message"`
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
	ActionName        string `json:"action_name"`
	ReactionName      string `json:"reaction_name"`
	ServiceActionId   int    `json:"service_action_id"`
	ServiceReactionId int    `json:"service_reaction_id"`
}

type OauthCode struct {
	OauthCode string `json:"code"`
}
