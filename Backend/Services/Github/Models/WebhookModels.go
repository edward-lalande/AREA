package models

type Repository struct {
	Name string `json:"name"`
}

type Sender struct {
	Login string `json:"login"`
}

type HeadCommit struct {
	Timestamp string `json:"timestamp"`
}

type Webhook struct {
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
	HeadCommit HeadCommit `json:"head_commit"`
}
