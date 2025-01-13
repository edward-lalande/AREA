package models

type LoginId struct {
	Id int `json:"id"`
}

type Login struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
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

type UserInformation struct {
	Id       int    `json:"id"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}
