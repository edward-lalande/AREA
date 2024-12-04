package models

type TriggerdModels struct {
	ReactionIdentifyer string `json:"reaction_identifyer"`
}

type TriggerdUserModel struct {
	UserEmail    string
	Message      string
	ReactionType int
	Channel      string
}
