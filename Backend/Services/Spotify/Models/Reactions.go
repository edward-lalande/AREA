package models

type Reactions struct {
	UserToken    string
	ReactionType int
}

type TriggerdModels struct {
	AreaId string `json:"area_id"`
}

type ReactionsReceived struct {
	AreaId       string `json:"area_id"`
	ReactionType int    `json:"reaction_type"`
	UserToken    string `json:"user_token"`
}
