package main

import (
	"bytes"
	"context"
	models "cryptomoney/Models"
	"cryptomoney/routes"
	"cryptomoney/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func fetchCryptoPrice(symbole, devise string) (float64, error) {
	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=%s", symbole, devise)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result map[string]float64
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}

	price := result[devise]

	return price, nil
}

func callMsgBrocker(areaId string) {
	send := struct {
		AreaId string `json:"area_id"`
	}{AreaId: areaId}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(send); err != nil {
		return
	}
	http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
}

func compareValue(action models.Actions, value float64) {
	switch action.ActionType {
	case 0:
		if value < float64(action.Value) {
			callMsgBrocker(action.AreaId)
		}
	case 1:
		if value > float64(action.Value) {
			callMsgBrocker(action.AreaId)
		}
	case 2:
		if value == float64(action.Value) {
			callMsgBrocker(action.AreaId)
		}
	}
}

func BackUpLocalDataCall() {
	db := utils.OpenDB(nil)
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(), `SELECT area_id, action_type, symbole, devise, value FROM "CryptoMoneyActions"`)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		action := models.Actions{}
		err := rows.Scan(&action.AreaId, &action.ActionType, &action.Symbole, &action.Devise, &action.Value)
		if err != nil {
			continue
		}

		currentValue, err := fetchCryptoPrice(action.Symbole, action.Devise)
		if err != nil {
			continue
		}

		compareValue(action, currentValue)
	}
}

func InitCronScheduler() *cron.Cron {
	c := cron.New()

	c.AddFunc("@every 00h00m10s", BackUpLocalDataCall)

	c.Start()
	return c
}

func main() {
	r := gin.Default()
	c := InitCronScheduler()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	routes.ApplyRoutes(r)
	r.Run(":8095")
	defer c.Stop()
}
