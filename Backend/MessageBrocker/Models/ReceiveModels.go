package models

type ReceivedActionToReactions struct {
	UserToken          string `json:"user_token"`
	ServiceSenderId    int    `json:"service_sender_id"`
	ServiceReceiverId  int    `json:"service_receiver_id"`
	ActionId           int    `json:"action_id"`
	ReactionIdentifyer string `json:"reaction_identifyer"`
	ReactionType       int    `json:"reaction_type"`
	Message            string `json:"message"`
	ChannelId          string `json:"channel_id"`
}
