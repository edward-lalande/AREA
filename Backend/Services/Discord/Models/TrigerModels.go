package models

type TriggerdModels struct {
	AreaId string `json:"area_id"`
}

type TriggerdUserModel struct {
	UserEmail    string `json:"user_email"`
	Message      string `json:"message"`
	ReactionType int    `json:"reaction_type"`
	Channel      string `json:"channel"`
	Guild        string `json:"guild_id"`
}
