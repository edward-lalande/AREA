package area

import (
	"bytes"
	"context"
	"encoding/json"
	models "miro/Models"
	"miro/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Miro Services
// @Summary Register an received Reactions
// @Description Register the reactions received by the message brocker with all informations nedded
// @Tags Miro Area
// @Accept json
// @Produce json
// @Param routes body models.Reactions true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reaction [post]
func StoreReactions(c *gin.Context) {
	var user models.User
	receivedData := models.Reactions{}
	db := utils.OpenDB(c)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request", "details": err.Error()})
		return
	}
	defer db.Close(context.Background())

	id := utils.ParseToken(receivedData.UserToken)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	row := db.QueryRow(c, "SELECT * FROM \"User\" WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.Mail, &user.Password, &user.Login, &user.Lastname, &user.AsanaToken, &user.DiscordToken,
		&user.DropboxToken, &user.GithubToken, &user.GitlabToken, &user.GoogleToken, &user.MiroToken, &user.SpotifyToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	query := `
		INSERT INTO "MiroReactions" (
			user_token, 
			reaction_type, 
			area_id, 
			name
		) VALUES ($1, $2, $3, $4)
	`

	token := *user.MiroToken

	_, err = db.Exec(context.Background(),
		query,
		token,
		receivedData.ReactionType,
		receivedData.AreaId,
		receivedData.Name,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save reaction", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reaction stored successfully"})
}

func createBoard(reactions models.Reactions) (*http.Response, error) {

	url := "https://api.miro.com/v2/boards"
	tempBody := struct {
		Name string `json:"name"`
	}{
		reactions.Name,
	}
	body, _ := json.Marshal(tempBody)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", "Bearer "+reactions.UserToken)
	req.Header.Set("content-type", "application/json")

	return client.Do(req)
}

func FindReactions(reactions models.Reactions) (*http.Response, error) {
	Reactions := map[int]func(models.Reactions) (*http.Response, error){
		0: createBoard,
	}
	return Reactions[reactions.ReactionType](reactions)
}

// Miro Reactions
// @Summary send all the reactions
// @Description send all the reactions availablef6ff00 on the Miro services as an object arrays with the names and the object needed
// @Tags Miro Area
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reactions [get]
func GetReactions(c *gin.Context) {
	b, err := utils.OpenFile("Models/Reactions.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}

func Trigger(c *gin.Context) {
	receivedData := models.TriggerModelGateway{}
	reactions := models.Reactions{}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := utils.OpenDB(c)
	row := db.QueryRow(c, "SELECT user_token, reaction_type, name FROM \"MiroReactions\" WHERE area_id = $1", receivedData.AreaId)

	if err := row.Scan(&reactions.UserToken, &reactions.ReactionType, &reactions.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close(c)

	rep, err := FindReactions(reactions)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(rep.StatusCode, gin.H{
		"body": rep.Body,
	})
}
