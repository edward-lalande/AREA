package models

var ActionsModelsPath string = "Models/Actions.json"

type Database struct {
	Id          int    `json:"id"`
	AreaId      string `json:"area_id"`
	ActionType  int    `json:"action_type"`
	IsPlaying   int    `json:"is_playing"`
	AccessToken string `json:"user_token"`
	UserId      string `json:"user_id"`
	NbPlaylists int    `json:"nb_playlists"`
}

type ActionsData struct {
	AreaId      string `json:"area_id"`
	ActionType  int    `json:"action_type"`
	IsPlaying   int    `json:"is_playing"`
	AccessToken string `json:"user_token"`
}
