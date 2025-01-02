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
