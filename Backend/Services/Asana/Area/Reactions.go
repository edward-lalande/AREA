package area

import (
	models "asana/Models"
	"asana/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createProject(info models.DatabaseReactions) (*http.Response, error) {
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"name":      info.ProjectName,
			"workspace": info.WorkSpaceId,
			"notes":     info.Note,
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://app.asana.com/api/1.0/projects", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+info.UserToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func createTask(info models.DatabaseReactions) (*http.Response, error) {
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"name":      info.Note,
			"workspace": info.WorkSpaceId,
			"projects":  []string{info.ProjectName},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://app.asana.com/api/1.0/tasks", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+info.UserToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func addMember(info models.DatabaseReactions) (*http.Response, error) {
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"members": []string{info.Note},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://app.asana.com/api/1.0/projects/%s/addMembers", info.ProjectName)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+info.UserToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func FindReactions(info models.DatabaseReactions) (*http.Response, error) {
	Reactions := map[int]func(models.DatabaseReactions) (*http.Response, error){
		0: createProject,
		1: createTask,
		2: addMember,
	}

	return Reactions[info.ReactionType](info)

}

// Asana Services
// @Summary Trigger an Area
// @Description Actions triggerd the reactions and call the trigger route
// @Tags Asana trigger
// @Accept json
// @Produce json
// @Param routes body models.TriggerModelGateway true "It contains the Area Id to the reactions"
// @Success 200 {object} map[string]string "Response of the reactions"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /trigger [post]
func Trigger(c *gin.Context) {
	receivedData := models.TriggerModelGateway{}
	database := models.DatabaseReactions{}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format", "details": err.Error()})
		return
	}
	db := utils.OpenDB(c)
	defer db.Close(c)

	row := db.QueryRow(c, "SELECT user_token, reaction_type, project_name, workspace_id, note, project_id, task_id FROM \"AsanaReactions\" WHERE area_id = $1", receivedData.AreaId)

	if err := row.Scan(&database.UserToken, &database.ReactionType, &database.ProjectName, &database.WorkSpaceId, &database.Note, &database.ProjectId, &database.TaskId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close(c)

	rep, err := FindReactions(database)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(rep.StatusCode, gin.H{
		"body": rep.Body,
	})
	defer rep.Body.Close()
}

// Asana Services
// @Summary Register an received Reactions
// @Description Register the reactions received by the message brocker with all informations nedded
// @Tags Asana Area
// @Accept json
// @Produce json
// @Param routes body models.Reactions true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reaction [post]
func StoreReactions(c *gin.Context) {
	receivedData := models.Reactions{}
	var user models.User

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format", "details": err.Error()})
		return
	}

	db := utils.OpenDB(c)
	defer db.Close(c)

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
		INSERT INTO "AsanaReactions" (
			user_token, 
			reaction_type, 
			area_id, 
			project_name, 
			workspace_id, 
			note, 
			project_id, 
			task_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err = db.Exec(c, query,
		user.AsanaToken,
		receivedData.ReactionType,
		receivedData.AreaId,
		receivedData.ProjectName,
		receivedData.WorkSpaceId,
		receivedData.Note,
		receivedData.ProjectId,
		receivedData.TaskId,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save reaction", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reaction stored successfully"})
}

// Asana Reactions
// @Summary send all the reactions
// @Description send all the reactions available on the Asana services as an object arrays with the names and the object needed
// @Tags Asana Area
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
