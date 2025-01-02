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

func createTextChannel(information models.Reactions) (*http.Response, error) {
	url := fmt.Sprintf("https://discord.com/api/v10/guilds/%s/channels", information.GuildId)
	tempBody := struct {
		ChannelId string `json:"name"`
	}{
		information.ChannelId,
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

func createVoiceChannel(information models.Reactions) (*http.Response, error) {
	url := fmt.Sprintf("https://discord.com/api/v10/guilds/%s/channels", information.GuildId)
	tempBody := struct {
		ChannelId string `json:"name"`
		Type      string `json:"type"`
	}{
		information.ChannelId,
		"2",
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

func deleteChannel(information models.Reactions) (*http.Response, error) {
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s", information.ChannelId)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+utils.GetEnvKey("BOT_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func pinMessage(information models.Reactions) (*http.Response, error) {
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/pins/%s", information.ChannelId, information.Message)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+utils.GetEnvKey("BOT_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func unpinMessage(information models.Reactions) (*http.Response, error) {
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/pins/%s", information.ChannelId, information.Message)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+utils.GetEnvKey("BOT_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func createRole(information models.Reactions) (*http.Response, error) {
	url := fmt.Sprintf("https://discord.com/api/v10/guilds/%s/roles", information.GuildId)

	tempBody := struct {
		Name string `json:"name"`
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

func deleteRole(information models.Reactions) (*http.Response, error) {
	url := fmt.Sprintf("https://discord.com/api/v10/guilds/%s/roles/%s", information.GuildId, information.Message)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bot "+utils.GetEnvKey("BOT_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func FindReactions(id int, information models.Reactions) (*http.Response, error) {
	Reactions := map[int]func(models.Reactions) (*http.Response, error){
		0: sendMessage,
		1: createTextChannel,
		2: createVoiceChannel,
		3: deleteChannel,
		4: pinMessage,
		5: unpinMessage,
		6: createRole,
		7: deleteRole,
	}

	return Reactions[id](information)
}
