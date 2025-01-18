package routes

import (
	"api-gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serviceList struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Color       string `json:"color"`
}

type serviceSendList struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

// Get Up services
// @Summary Get all services up
// @Description Get services up with name to display, routes to call it to the api-gateway, url and color to display
// @Tags Area api-gateway
// @Accept json
// @Produce json
// @Success 200 {object} serviceList "services up with name to display, routes to call it to the api-gateway, url and color to display"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /services [get]
func Services(c *gin.Context) {
	servicesArray := []serviceList{
		{"Date Time", "Add a reactions when Every hour at a minute, every day at a certains minute in a certain hour.", utils.GetEnvKey("TIME_API"), "#b3b3b3"},
		{"Discord", `Whether you're part of a school club, gaming group, worldwide art community, or just a handful of friends that want to spend time together, Discord makes it easy to talk every day and hang out more often.
To add the Bot, the user needs to have the Manage Server permission. Note: Revoking Discord Permissions will not stop the Bot from running. The Bot needs to be removed from the server or have its permissions disabled.
`, utils.GetEnvKey("DISCORD_API"), "#7289da"},
		{"Dropbox", `Dropbox lets people bring their documents, photos and videos everywhere and share them easily. Use this services to sync your Dropbox uploads with other services, quickly add new files, and keep track of all your important photos, documents, and data — automatically.`, utils.GetEnvKey("DROPBOX_API"), "#0061FE"},
		{"Github", `GitHub is the best place to share code with friends, co-workers, classmates, and complete strangers. Turn the services to automatically track issues, pull requests, repositories, and to quickly create issues with Webhooks.`, utils.GetEnvKey("GITHUB_API"), "black"},
		{"Gitlab", `Gitlab is the best place to share code with friends, co-workers, classmates, and complete strangers. Turn the services to automatically track issues, pull requests, repositories, and to quickly create issues with Webhooks. To set the Webhook you need to go to the settings of your repositories in Webhooks page, add a new webhooks and add the link of our API. `, utils.GetEnvKey("GITLAB_API"), "#fc6d26"},
		{"Google", "Google Calendar is a free time-management web application offered by Google. Turn on the services to add the most important information, right into your calendar, automatically — and get custom notifications about the events that matter the most to you, also for mails and tasks", utils.GetEnvKey("GOOGLE_API"), "#0F9D58"},
		{"Meteo", "Add a reactions when the weather is hot, cold or if it's rain", utils.GetEnvKey("METEO_API"), "#4285F4"},
		{"Spotify", `Spotify is a digital music service that gives you access to millions of songs. Applets can help you save your Discover Weekly and Release Radar playlists, share your favorite tunes, and much more.`, utils.GetEnvKey("SPOTIFY_API"), "#1db954"},
		{"Asana", `Asana is the easiest way for teams to track their work. From tasks and projects to conversations and dashboards, Asana enables teams to move work from start to finish — and get great results.`, utils.GetEnvKey("ASANA_API"), "#ff80e1"},
		{"Ticket Master", "Create a reactions when there is your favorite kind of music concert in your city", utils.GetEnvKey("TICKET_MASTER_API"), "#1c24ff"},
		{"CryptoMoney", "Create a reactions when your favorite cryptomoney is up to or down to or equal to a certains value", utils.GetEnvKey("CRYPTOMONEY_API"), "#f7931a"},
		{"Miro", "Create a reactions when you need to create a share board to work", utils.GetEnvKey("MIRO_API"), "#faca00"},
	}
	var uppedServices []serviceSendList

	for _, service := range servicesArray {
		_, err := http.Get(service.Url + "ping")
		if err != nil {
			continue
		}
		uppedServices = append(uppedServices, serviceSendList{service.Name, service.Description, service.Color})
	}

	c.JSON(http.StatusOK, gin.H{"services": uppedServices})
}
