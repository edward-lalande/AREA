package models

type Attendee struct {
	Email string `json:"email"`
}

type GmailProfile struct {
	EmailAddress  string `json:"emailAddress"`
	MessagesTotal int    `json:"messagesTotal"`
	ThreadsTotal  int    `json:"threadsTotal"`
	HistoryId     string `json:"historyId"`
}

type SendMessageRequest struct {
	Raw string `json:"raw"`
}

type GoogleReaction struct {
	UserToken    string `json:"user_token"`
	AreaId       string `json:"area_id"`
	ReactionType int    `json:"reaction_type"`
	Summary      string `json:"summary"`
	Description  string `json:"description"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Attendees    string `json:"attendees"`
	Recipient    string `json:"recipient"`
	Subject      string `json:"subject"`
	Message      string `json:"message"`
}

type DateTime struct {
	DateTimeFields string `json:"dateTime"`
}

type GoogleReactionSend struct {
	Summary     string     `json:"summary"`
	Description string     `json:"description"`
	StartTime   DateTime   `json:"start"`
	EndTime     DateTime   `json:"end"`
	Attendees   []Attendee `json:"attendees"`
}
