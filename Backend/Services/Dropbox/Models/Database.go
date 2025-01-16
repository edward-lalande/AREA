package models

type Datbase struct {
	Id            int    `json:"id"`
	UserToken     string `json:"user_token"`
	ReactionType  int    `json:"reaction_type"`
	AreaId        string `json:"area_id"`
	FromPath      string `json:"from_path"`
	ToPath        string `json:"to_path"`
	FilepathShare string `json:"filepath_share"`
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
