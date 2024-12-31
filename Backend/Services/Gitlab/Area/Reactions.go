package area

import (
	"bytes"
	models "gitlab/Models"
	"gitlab/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
    Commentaire automatique:
        curl --request POST \
  --header "Authorization: Bearer d8ebd73da344e0ed67766099114f97df9bf9a22e80e0cb08aabfa91449456048" \
  --header "Content-Type: application/json" \
  --data '{"body": "Merci pour cette MR ! Nous allons la vérifier rapidement."}' \
  "https://gitlab.com/api/v4/projects/65448666/merge_requests/2/notes"
                                                ID-PROJ       ID MR

    Labeliser une MR:
curl --request PUT \
  --header "Authorization: Bearer d8ebd73da344e0ed67766099114f97df9bf9a22e80e0cb08aabfa91449456048" \
  --header "Content-Type: application/json" \
  --data '{"labels": "To Review"}' \
  "https://gitlab.com/api/v4/projects/65448666/merge_requests/2"
                                      ID-PROJ                 ID MR
  nb de MR:
    curl --head --header "Authorization: Bearer d8ebd73da344e0ed67766099114f97df9bf9a22e80e0cb08aabfa91449456048" \
"https://gitlab.com/api/v4/projects/65448666/merge_requests?state=all&per_page=1"
                                    ID-PROJ
*/

func findNumberOfMr(data models.Database) int {
	client := &http.Client{}
	url := "https://gitlab.com/api/v4/projects/" + data.ProjectId + "/merge_requests?state=all&per_page=1"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return -1
	}

	req.Header.Set("Authorization", "Bearer "+data.UserToken)
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

func labeliseMr(data models.Database) (*http.Response, error) {
	lastMrId := findNumberOfMr(data)
	url := "https://gitlab.com/api/v4/projects/" + data.ProjectId + "/merge_requests/" + strconv.Itoa(lastMrId)

	requestBody := `{"labels": "` + data.Body + `"}`

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(requestBody)))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+data.UserToken)
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func automatiqueCommentary(data models.Database) (*http.Response, error) {
	lastMrId := findNumberOfMr(data)
	url := "https://gitlab.com/api/v4/projects/" + data.ProjectId + "/merge_requests/" + strconv.Itoa(lastMrId) + "/notes"

	requestBody := `{"body": "` + data.Body + `"}`

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+data.UserToken)
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func FindReaction(data models.Database) (*http.Response, error) {
	reactions := map[int]func(models.Database) (*http.Response, error){
		0: automatiqueCommentary,
		1: labeliseMr,
	}

	return reactions[data.ReactionType](data)
}
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
