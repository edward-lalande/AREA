package models

type Attendee struct {
	Email string `json:"email"`
}

type GoogleCalendarReaction struct {
	UserToken    string `json:"user_token"`
	AreaId       string `json:"area_id"`
	ReactionType int    `json:"reaction_type"`
	Summary      string `json:"summary"`
	Description  string `json:"description"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Attendees    string `json:"attendees"`
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
