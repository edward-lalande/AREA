package models

type ReceivedActionToReactions struct {
	UserToken          string `json:"user_token"`
	ServiceSenderId    int    `json:"service_sender_id"`
	ServiceReceiverId  int    `json:"service_receiver_id"`
	ActionId           int    `json:"action_id"`
	ReactionIdentifyer int    `json:"reaction_identifyer"`
	Message            string `json:"message"`
}
