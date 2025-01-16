package area

import (
	"bytes"
	"fmt"
	models "gitlab/Models"
	"gitlab/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context, token string) string {

	id := utils.ParseToken(token)

	db := utils.OpenDB(c)
	if db == nil {
		return ""
	}

	var discord_token string

	query := `SELECT gitlab_token FROM "User" WHERE id = $1`
	err := db.QueryRow(c, query, id).Scan(&discord_token)
	if err != nil {
		return ""
	}

	defer db.Close(c)
	fmt.Println("gitlab token => " + discord_token)
	return discord_token

}

func findNumberOfMr(c *gin.Context, data models.Database) int {
	client := &http.Client{}
	url := "https://gitlab.com/api/v4/projects/" + data.ProjectId + "/merge_requests?state=all&per_page=1"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return -1
	}

	req.Header.Set("Authorization", "Bearer "+GetToken(c, data.UserToken))
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return -1
	}
	defer response.Body.Close()

	xTotal := response.Header.Get("x-total")
	if xTotal == "" {
		return -1
	}

	nbOfMr, err := strconv.Atoi(xTotal)
	if err != nil {
		return -1
	}

	return nbOfMr
}

func labeliseMr(c *gin.Context, data models.Database) (*http.Response, error) {
	lastMrId := findNumberOfMr(c, data)
	url := "https://gitlab.com/api/v4/projects/" + data.ProjectId + "/merge_requests/" + strconv.Itoa(lastMrId)

	requestBody := `{"labels": "` + data.Body + `"}`

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(requestBody)))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+GetToken(c, data.UserToken))
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func automatiqueCommentary(c *gin.Context, data models.Database) (*http.Response, error) {
	lastMrId := findNumberOfMr(c, data)
	url := "https://gitlab.com/api/v4/projects/" + data.ProjectId + "/merge_requests/" + strconv.Itoa(lastMrId) + "/notes"

	requestBody := `{"body": "` + data.Body + `"}`

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+GetToken(c, data.UserToken))
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func FindReaction(c *gin.Context, data models.Database) (*http.Response, error) {
	reactions := map[int]func(*gin.Context, models.Database) (*http.Response, error){
		0: automatiqueCommentary,
		1: labeliseMr,
	}

	return reactions[data.ReactionType](c, data)
}

// Gitlab Services
// @Summary Register an received Reactions
// @Description Register the reactions received by the message brocker with all informations nedded
// @Tags Gitlab Area
// @Accept json
// @Produce json
// @Param routes body models.ReceivedReactions true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reaction [post]
func StoreReactions(c *gin.Context) {
	receivedData := models.ReceivedReactions{}
	db := utils.OpenDB(c)
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer db.Close(c)

	_, err := db.Exec(c, "INSERT INTO \"GitlabReactions\" (user_token, reaction_type, area_id, project_id, body)"+
		" VALUES ($1, $2, $3, $4, $5)", receivedData.UserToken, receivedData.ReactionType, receivedData.AreaId, receivedData.ProjectId, receivedData.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reaction stored successfully"})
}

// Gitlab Reactions
// @Summary send all the Reactions
// @Description send all the Reactions available on the Gitlab services as an object arrays with the names and the object needed
// @Tags Gitlab Area
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reactions [get]
func GetReactions(c *gin.Context) {
	b, err := utils.OpenFile(models.ReactionsModelsPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}
