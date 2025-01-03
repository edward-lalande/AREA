package oauth

import (
	"net/http"
	"spotify/utils"

	"github.com/gin-gonic/gin"
)

// Spotify OAUTH2
// @Summary Send the url to redirect to for the OAUTH2 Spotify
// @Description Send the url to redirect to for the OAUTH2 Spotify
// @Tags Spotify OAUTH2
// @Accept json
// @Produce json
// @Success 200 {string} string "the URL to redirect to for the OAUTH2 Spotify"
// @Router /oauth [get]
func OAuthFront(c *gin.Context) {
	authUrl := "https://accounts.spotify.com/authorize?&client_id=" + utils.GetEnvKey("CLIENT_ID") +
		"&redirect_uri=" + utils.GetEnvKey("REDIRECT_URI") +
		"&response_type=code" +
		"&scope=user-library-read%20user-library-modify%20user-read-recently-played%20user-top-read%20user-read-playback-position%20user-follow-modify%20playlist-modify-public%20playlist-modify-private%20playlist-read-collaborative%20playlist-read-private%20streaming%20user-read-currently-playing%20user-modify-playback-state%20user-read-playback-state%20user-read-email%20user-follow-read"
	c.String(http.StatusOK, authUrl)
}
