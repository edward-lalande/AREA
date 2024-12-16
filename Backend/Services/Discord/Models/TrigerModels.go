package models

type TriggerdModels struct {
	ReactionIdentifyer string `json:"area_id"`
}

type TriggerdUserModel struct {
	UserEmail    string
	Message      string
	ReactionType int
	Channel      string
	Guild		 string
}
