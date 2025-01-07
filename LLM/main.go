package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gage-technologies/mistral-go"
)

func main() {
	client := mistral.NewMistralClientDefault("HCZtt06PoHGhiRb1Wjed8b5v4zeQ874H")

	actions, err := http.Get("http://127.0.0.1:8080/actions")

	if err != nil {
		return
	}
	defer actions.Body.Close()

	body, err := io.ReadAll(actions.Body)

	if err != nil {
		return
	}

	request := "je veux que quand il est 17 heures 45 minutes à Paris en Europe tu m'écrives un message discord dans le channel id : 1308006396213727232 en disant le message coucou les amis !"

	prompt := "You are an agent for an action and reaction creation application." +
		"Your goal is to generate a JSON object for me that is exactly this format 'action: { action_id: $, action_type: $, ... }, reactions: [{ reaction_id: $, reaction_type: $, ... }]' ." +
		"In order to replace the $, you will need to use the information provided in the user's request which is: " + request +
		"So you need to identify in the request what the action and reaction is in order to correctly replace it in the JSON object." +
		"You must also replace the ... in the JSON format that I am asking you for with the action and reaction arguments." +
		"For these arguments, you will have to look for them in the user request with the display of the list and inform me in the JSON object the name of the argument (name in the list) followed by its value (indicated by l 'user)." +
		"In order to know the types, ids and arguments of each available action and reaction you must use this list: " + string(body) +
		"Example: '{ action: { action_id: 0, action_type: 0, hour: 17, minute: 45 }, reactions: [{ reaction_id: 0, reaction_type: 0, channel_id: 123456, message_id: This is an example }]'." +
		"Just send me back the only JSON object, I don't want ANY other text."

	chatRes, err := client.Chat("mistral-tiny", []mistral.ChatMessage{{Content: prompt, Role: mistral.RoleUser}}, nil)
	if err != nil {
		log.Fatalf("Error getting chat completion: %v", err)
	}
	fmt.Printf(chatRes.Choices[0].Message.Content)
}
