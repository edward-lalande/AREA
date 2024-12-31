package models

type GithubAction struct {
	AreaId     string `json:"area_id"`
	ActionType int    `json:"action_type"`
	UserToken  string `json:"user_token"`
	Pusher     string `json:"pusher"`
	Value      string `json:"value"`
	Number     int    `json:"number"`
}

type GithubSendReaction struct {
	ReactionId string `json:"area_id"`
}
