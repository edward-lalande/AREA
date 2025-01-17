package models

var ActionsModelsPath string = "Models/Actions.json"

type ReceivedActions struct {
	UserToken  string `json:"user_token"`
	AreaId     string `json:"area_id"`
	ActionType int    `json:"action_type"`
	NbEvents   int    `json:"nb_events"`
}
