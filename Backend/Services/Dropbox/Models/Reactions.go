package models

var ReactionsModelPath string = "Models/Reactions.json"

type DropBoxReactions struct {
	UserToken     string `json:"user_token"`
	ReactionType  int    `json:"reaction_type"`
	AreaId        string `json:"area_id"`
	FromPath      string `json:"from_path"`
	ToPath        string `json:"to_path"`
	FilepathShare string `json:"filepath_share"`
}
