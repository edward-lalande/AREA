package area

import (
	"bytes"
	models "dropbox/Models"
	"dropbox/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func renameFile(info models.DropBoxReactions) (*http.Response, error) {
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

	req.Header.Set("Authorization", "Bearer "+info.UserToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func shareFile(info models.DropBoxReactions) (*http.Response, error) {
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

	req.Header.Set("Authorization", "Bearer "+info.UserToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func findReactions(info models.DropBoxReactions) (*http.Response, error) {
	Reactions := map[int]func(models.DropBoxReactions) (*http.Response, error){
		0: renameFile,
		1: shareFile,
	}

	return Reactions[info.ReactionType](info)
}

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

	resp, err := findReactions(database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	defer resp.Body.Close()

	c.JSON(resp.StatusCode, gin.H{
		"body": resp.Body,
	})
}

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
