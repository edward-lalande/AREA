package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	models "meteo/Models"
	"meteo/routes"
	"meteo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func checkTemperature(info models.MeteoDatabase) int {
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + info.Latitude + "&longitude=" + info.Longitude + "&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m"
	resp, err := http.Get(url)

	if err != nil {
		return -1
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return -1
	}

	hourly, ok := jsonMap["hourly"].(map[string]interface{})
	if !ok {
		return -1
	}

	tempArr, ok := hourly["temperature_2m"].([]interface{})
	if !ok {
		return -1
	}

	for _, temp := range tempArr {
		if tempFloat, ok := temp.(float64); ok {
			if info.Value == int(tempFloat) {
				return 1
			}
		}
	}

	return 0
}

func checkHumidty(info models.MeteoDatabase) int {
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + info.Latitude + "&longitude=" + info.Longitude + "&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m"
	resp, err := http.Get(url)

	if err != nil {
		return -1
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return -1
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return -1
	}

	hourly, ok := jsonMap["hourly"].(map[string]interface{})
	if !ok {
		return -1
	}

	tempArr, ok := hourly["relative_humidity_2m"].([]interface{})
	if !ok {
		return -1
	}

	for _, temp := range tempArr {
		if tempFloat, ok := temp.(float64); ok {
			if info.Value == int(tempFloat) {
				return 1
			}
		}
	}

	return 0
}

func checkWindSpeed(info models.MeteoDatabase) int {
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + info.Latitude + "&longitude=" + info.Longitude + "&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m"
	resp, err := http.Get(url)

	if err != nil {
		return -1
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return -1
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return -1
	}

	hourly, ok := jsonMap["hourly"].(map[string]interface{})
	if !ok {
		return -1
	}

	tempArr, ok := hourly["wind_speed_10m"].([]interface{})
	if !ok {
		return -1
	}

	for _, temp := range tempArr {
		if tempFloat, ok := temp.(float64); ok {
			if info.Value == int(tempFloat) {
				return 1
			}
		}
	}

	return 0
}

func sendTrigger(areaId string) {
	send := struct {
		AreaId string `json:"area_id"`
	}{areaId}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(send); err != nil {
		return
	}
	http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)

}

func BackUpLocalDataCall() {
	db := utils.OpenDB(nil)
	query := `SELECT * FROM "MeteoActions"`
	rows, _ := db.Query(context.Background(), query)

	for rows.Next() {
		dbSlice := models.MeteoDatabase{}
		err := rows.Scan(&dbSlice.Id, &dbSlice.AreaId, &dbSlice.ActionType, &dbSlice.Latitude, &dbSlice.Longitude, &dbSlice.Value)
		if err != nil {
			log.Fatal(err)
			continue
		}

		switch dbSlice.ActionType {
		case 0:
			value := checkTemperature(dbSlice)
			if value == -1 {
				continue
			}
			if value == 1 {
				sendTrigger(dbSlice.AreaId)
			}
		case 1:
			value := checkWindSpeed(dbSlice)
			if value == -1 {
				continue
			}
			if value == 1 {
				sendTrigger(dbSlice.AreaId)
			}
		case 2:
			value := checkHumidty(dbSlice)
			if value == -1 {
				continue
			}
			if value == 1 {
				sendTrigger(dbSlice.AreaId)
			}
		}
	}
	defer db.Close(context.Background())
}

func InitCronScheduler() *cron.Cron {
	c := cron.New()

	c.AddFunc("@every 24h00m00s", BackUpLocalDataCall)

	c.Start()
	return c
}

func main() {
	r := gin.Default()
	c := InitCronScheduler()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	routes.ApplyRoutes(r)
	defer c.Stop()

	r.Run(":8089")
}
