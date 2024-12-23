package models

type Database struct {
	Id           int    `json:"id"`
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	AreaId       string `json:"area_id"`
	ProjectId    string `json:"project_id"`
	Body         string `json:"body"`
}

type ReceivedReactions struct {
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	AreaId       string `json:"area_id"`
	ProjectId    string `json:"project_id"`
	Body         string `json:"body"`
}
