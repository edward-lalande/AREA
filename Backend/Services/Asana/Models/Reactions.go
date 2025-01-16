package models

type Reactions struct {
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	AreaId       string `json:"area_id"`
	ProjectName  string `json:"project_name"`
	WorkSpaceId  string `json:"workspace_id"`
	Note         string `json:"note"`
	ProjectId    string `json:"project_id"`
	TaskId       string `json:"task_id"`
}

type DatabaseReactions struct {
	Id           int    `json:"id"`
	UserToken    string `json:"user_token"`
	ReactionType int    `json:"reaction_type"`
	AreaId       string `json:"area_id"`
	ProjectName  string `json:"project_name"`
	WorkSpaceId  string `json:"workspace_id"`
	Note         string `json:"note"`
	ProjectId    string `json:"project_id"`
	TaskId       string `json:"task_id"`
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
