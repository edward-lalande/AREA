package models

var ActionsModelsPath string = "Models/Actions.json"

type Actions struct {
	ActionType int `json:"action_type"`
	AreaId     string `json:"area_id"`
}
