package models

type Database struct {
	Id          int    `json:"id"`
	AreaId      string `json:"area_id"`
	ActionType  int    `json:"action_type"`
	IsPlaying   int    `json:"is_playing"`
	AccessToken string `json:"user_token"`
	MusicName   string `json:"music_name"`
}

type ActionsData struct {
	AreaId      string `json:"area_id"`
	ActionType  int    `json:"action_type"`
	IsPlaying   int    `json:"is_playing"`
	AccessToken string `json:"user_token"`
	MusicName   string `json:"music_name"`
}
