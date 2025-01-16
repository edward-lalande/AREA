package models

var ActionsModelsPath string = "Models/Actions.json"

type Action struct {
	AreaID     string `json:"area_id"`
	ActionType int    `json:"action_type"`
	Name       string `json:"name"`
	Venue      string `json:"venue"`
	City       string `json:"city"`
	NbEvents   int    `json:"nb_events"`
}
