package models

type TriggerdModels struct {
	ReactionIdentifyer int `json:"reaction_identifyer"`
}

type TriggerdUserModel struct {
	UserEmail string
	Message   string
}
