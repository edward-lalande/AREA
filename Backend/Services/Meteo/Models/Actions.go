package models

var GetActionsPath string = "Models/Actions.json"

type MeteoActions struct {
	ActionType int    `json:"action_type"`
	AreaId     string `json:"area_id"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Value      int    `json:"value"`
}

type MeteoDatabase struct {
	Id         int    `json:"id"`
	ActionType int    `json:"action_type"`
	AreaId     string `json:"area_id"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Value      int    `json:"value"`
}
