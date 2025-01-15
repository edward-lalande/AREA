package models

var ReactionsModelsPath string = "Models/Reactions.json"

type Reactions struct {
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	AreaId       string `json:"area_id"`
	Name         string `json:"name"`
}
