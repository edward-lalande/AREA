package models

var ReactionsModelsPath string = "Models/Reactions.json"

type Reactions struct {
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	AreaId       string `json:"area_id"`
	Name         string `json:"name"`
}

type User struct {
	Id           string  `json:"id"`
	Mail         *string `json:"mail"`
	Password     *string `json:"password"`
	Login        *string `json:"name"`
	Lastname     *string `json:"lastname"`
	AsanaToken   *string `json:"asana_token"`
	DiscordToken *string `json:"discord_token"`
	DropboxToken *string `json:"dropbox_token"`
	GithubToken  *string `json:"github_token"`
	GitlabToken  *string `json:"gitlab_token"`
	GoogleToken  *string `json:"google_token"`
	MiroToken    *string `json:"miro_token"`
	SpotifyToken *string `json:"spotify_token"`
}
