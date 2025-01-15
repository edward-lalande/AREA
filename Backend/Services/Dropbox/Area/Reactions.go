package area

import (
	"bytes"
	models "dropbox/Models"
	"dropbox/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context, token string) string {

	id := utils.ParseToken(token)

	db := utils.OpenDB(c)
	if db == nil {
		return ""
	}

	var discord_token string

	query := `SELECT discord_token FROM "User" WHERE id = $1`
	err := db.QueryRow(c, query, id).Scan(&discord_token)
	if err != nil {
		return ""
	}

	defer db.Close(c)
	fmt.Println("discord token => " + discord_token)
	return discord_token

}

func renameFile(c *gin.Context, info models.DropBoxReactions) (*http.Response, error) {
	url := "https://api.dropboxapi.com/2/files/move_v2"

	body := map[string]interface{}{
		"from_path": info.FromPath,
		"to_path":   info.ToPath,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+GetToken(c, info.UserToken))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func shareFile(c *gin.Context, info models.DropBoxReactions) (*http.Response, error) {
	url := "https://api.dropboxapi.com/2/sharing/create_shared_link_with_settings"

	body := map[string]interface{}{
		"path": info.FilepathShare,
		"settings": map[string]interface{}{
			"requested_visibility": "public",
		},
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+GetToken(c, info.UserToken))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func findReactions(c *gin.Context, info models.DropBoxReactions) (*http.Response, error) {
	Reactions := map[int]func(*gin.Context, models.DropBoxReactions) (*http.Response, error){
		0: renameFile,
		1: shareFile,
	}

	return Reactions[info.ReactionType](c, info)
}

// Dropbox Services
// @Summary Trigger an Area
// @Description Actions triggerd the reactions and call the trigger route
// @Tags Dropbox trigger
// @Accept json
// @Produce json
// @Param routes body models.TriggerdModels true "It contains the Area Id to the reactions"
// @Success 200 {object} map[string]string "Response of the reactions"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /trigger [post]
func Trigger(c *gin.Context) {
	var (
		receivedData models.TriggerdModels
		database     models.DropBoxReactions
	)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := utils.OpenDB(c)
	row := db.QueryRow(c, "SELECT user_token, reaction_type, from_path, to_path, filepath_share FROM \"DropboxReactions\" WHERE area_id = $1", receivedData.AreaId)

	if err := row.Scan(&database.UserToken, &database.ReactionType, &database.FromPath, &database.ToPath, &database.FilepathShare); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close(c)

	resp, err := findReactions(c, database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	defer resp.Body.Close()

	c.JSON(resp.StatusCode, gin.H{
		"body": resp.Body,
	})
}

// Dropbox Services
// @Summary Register an received Reactions
// @Description Register the reactions received by the message brocker with all informations nedded
// @Tags Dropbox Area
// @Accept json
// @Produce json
// @Param routes body models.DropBoxReactions true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reaction [post]
func StoreReactions(c *gin.Context) {
	receivedData := models.DropBoxReactions{}
	db := utils.OpenDB(c)
	defer db.Close(c)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	query := `INSERT INTO "DropboxReactions" (user_token, reaction_type, area_id, from_path, to_path, filepath_share) 
	          VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(c, query, receivedData.UserToken, receivedData.ReactionType, receivedData.AreaId, receivedData.FromPath, receivedData.ToPath, receivedData.FilepathShare)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reaction stored successfully"})
}

// Dropbox Reactions
// @Summary send all the reactions
// @Description send all the reactions available on the Dropbox services as an object arrays with the names and the object needed
// @Tags Dropbox Area
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reactions [get]
func GetReactions(c *gin.Context) {
	b, err := utils.OpenFile(models.ReactionsModelPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}
