package area

import (
	"bytes"
	models "discord-service/Models"
	utils "discord-service/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendMessage(information models.Reactions) (*http.Response, error) {
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages", information.ChannelId)
	tempBody := struct {
		Message string `json:"content"`
	}{
		information.Message,
	}
	body, _ := json.Marshal(tempBody)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+utils.GetEnvKey("BOT_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func FindReactions(id int, information models.Reactions) (*http.Response, error) {
	var Reactions map[int]func(models.Reactions) (*http.Response, error) = make(map[int]func(models.Reactions) (*http.Response, error))
	Reactions[0] = sendMessage
	return Reactions[id](information)
}
