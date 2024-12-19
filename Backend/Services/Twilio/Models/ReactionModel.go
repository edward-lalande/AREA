package models

type ReactionReceiveData struct {
	AreaId       string `json:"area_id"`
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	PhoneNumber  string `json:"phone_number"`
	Message      string `json:"message"`
}

type Reactions struct {
	PhoneNumber string `json:"phone_number"`
	Message     string `json:"message"`
}
