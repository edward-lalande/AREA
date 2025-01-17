package models

var ModelPath string = "Models/Actions.json"

type Message struct {
	Pinned		bool	`json:"pinned"`
}

type Database struct {
	Id        	int
	Type  	  	int	   `json:"action_type"`
	AreaId    	string `json:"area_id"`
	ChannelId   string `json:"channel_id"`
	MessageId 	string `json:"message_id"`
	UserToken 	string `json:"user_token"`
}

type DiscordActionReceive struct {
	Id		  int
	Type  	  int	 `json:"action_type"`
	AreaId    string `json:"area_id"`
	ChannelId   string `json:"channel_id"`
	MessageId string `json:"message_id"`
	UserToken string `json:"user_token"`
}

type DiscordModelSendReaction struct {
	ReactionId string `json:"area_id"`
}