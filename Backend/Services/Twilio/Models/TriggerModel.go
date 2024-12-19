package models

type TriggerdModels struct {
	ReactionIdentifyer string `json:"area_id"`
}

type TriggerdUserModel struct {
	AreaId       string
	ReactionType int
	UserToken    string
	PhoneNumber  string
	Message      string
}
