package models

var ActionPath string = "Models/Actions.json"

type Actions struct {
	AreaId     string `json:"area_id"`
	ActionType int    `json:"action_type"`
	Symbole    string `json:"symbole"`
	Devise     string `json:"devise"`
	Value      int    `json:"value"`
}
