package models

type Pusher struct {
	Name string `json:"name"`
}

type Commit struct {
	Message  string   `json:"message"`
	Added    []string `json:"added"`
	Removed  []string `json:"removed"`
	Modified []string `json:"modified"`
}

type WebhookPush struct {
	Pusher  Pusher   `json:"pusher"`
	Commits []Commit `json:"commits"`
}

type User struct {
	Login string `json:"login"`
}

type Reactions struct {
	TotalCount int `json:"total_count"`
}

type Comment struct {
	User      User      `json:"user"`
	Body      string    `json:"body"`
	Reactions Reactions `json:"reactions"`
}

type WebhooksCommitComment struct {
	Comment Comment `json:"comment"`
}
