package area

import (
	"context"
	models "cryptomoney/Models"
	"cryptomoney/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CryptoMoney Services Services
// @Summary Register an received Actions
// @Description Register the Actions received by the message brocker with all informations nedded
// @Tags CryptoMoney Services Area
// @Accept json
// @Produce json
// @Param routes body models.Actions true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /action [post]
func StoreActions(c *gin.Context) {
	receivedData := models.Actions{}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}
	db := utils.OpenDB(c)
	defer db.Close(context.Background())

	query := `INSERT INTO "CryptoMoneyActions" (area_id, action_type, symbole, devise, value) VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(context.Background(), query, receivedData.AreaId, receivedData.ActionType, receivedData.Symbole, receivedData.Devise, receivedData.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert action", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Action stored successfully"})
}

// Get Actions of CryptoMoney Services
// @Summary Get Actions from CryptoMoney Services
// @Description Get Actions from CryptoMoney Services
// @Tags Actions CryptoMoney Services
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Reactions name with parameters of it as object"
// @Router /actions [get]
func GetActions(c *gin.Context) {
	b, err := utils.OpenFile(models.ActionPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}
