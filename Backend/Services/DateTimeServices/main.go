package main

import (
	"bytes"
	"context"
	models "date-time-service/Models"
	"date-time-service/routes"
	"date-time-service/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func getDatabaseSlice(actionType int) []models.Database {
	var databaseSlice []models.Database = nil
	db := utils.OpenDB(nil)

	if db == nil {
		return nil
	}
	rows, err := db.Query(context.Background(), "SELECT * FROM \"TimeAction\" WHERE action_type = $1", actionType)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error on reading response of the query", err)
		return nil
	}
	defer rows.Close()
	defer db.Close(context.Background())

	for rows.Next() {
		var database models.Database
		err := rows.Scan(&database.Id, &database.AreaId, &database.ActionType, &database.Continent, &database.City, &database.Hour, &database.Minute)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		databaseSlice = append(databaseSlice, database)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	return databaseSlice
}

func GetEveryDayAt() {
	databaseSlice := getDatabaseSlice(0)
	fmt.Println(databaseSlice)
	if databaseSlice == nil {
		return
	}

	for _, slice := range databaseSlice {

		resp, err := http.Get(utils.GetEnvKey("DATE_TIME_API") + slice.Continent + "/" + slice.City)
		if err != nil {
			log.Fatal("Error while calling the API:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading the response body:", err)
			return
		}

		jsonBody := utils.BytesToJson(body)

		if jsonBody["hour"].(float64) == float64(slice.Hour) && jsonBody["minute"].(float64) == float64(slice.Minute) {
			send := models.TimeModelSendReaction{}
			send.ReactionId = slice.AreaId
			var buf bytes.Buffer
			if err := json.NewEncoder(&buf).Encode(send); err != nil {
				return
			}
			http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
		}
	}
}

func GetEveryHourAt() {
	databaseSlice := getDatabaseSlice(1)
	if databaseSlice == nil {
		return
	}

	for _, slice := range databaseSlice {

		resp, err := http.Get(utils.GetEnvKey("DATE_TIME_API") + slice.Continent + "/" + slice.City)
		if err != nil {
			log.Fatal("Error while calling the API:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading the response body:", err)
			return
		}

		jsonBody := utils.BytesToJson(body)

		if jsonBody["minute"].(float64) == float64(slice.Minute) {
			send := models.TimeModelSendReaction{}
			send.ReactionId = slice.AreaId
			var buf bytes.Buffer
			if err := json.NewEncoder(&buf).Encode(send); err != nil {
				return
			}
			http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
		}
	}
}

func GetHourIsEven() {
	databaseSlice := getDatabaseSlice(2)
	if databaseSlice == nil {
		return
	}

	for _, slice := range databaseSlice {

		resp, err := http.Get(utils.GetEnvKey("DATE_TIME_API") + slice.Continent + "/" + slice.City)
		if err != nil {
			log.Fatal("Error while calling the API:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading the response body:", err)
			return
		}

		jsonBody := utils.BytesToJson(body)

		if int(jsonBody["hour"].(float64))%2 == 0 && jsonBody["minute"].(float64) == float64(slice.Minute) {
			send := models.TimeModelSendReaction{}
			send.ReactionId = slice.AreaId
			var buf bytes.Buffer
			if err := json.NewEncoder(&buf).Encode(send); err != nil {
				return
			}
			http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
		}
	}
}

func GetHourIsOdd() {
	databaseSlice := getDatabaseSlice(3)
	if databaseSlice == nil {
		return
	}

	for _, slice := range databaseSlice {

		resp, err := http.Get(utils.GetEnvKey("DATE_TIME_API") + slice.Continent + "/" + slice.City)
		if err != nil {
			log.Fatal("Error while calling the API:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading the response body:", err)
			return
		}

		jsonBody := utils.BytesToJson(body)

		if int(jsonBody["hour"].(float64))%2 != 0 && jsonBody["minute"].(float64) == float64(slice.Minute) {
			send := models.TimeModelSendReaction{}
			send.ReactionId = slice.AreaId
			var buf bytes.Buffer
			if err := json.NewEncoder(&buf).Encode(send); err != nil {
				return
			}
			http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
		}
	}
}

func InitCronScheduler() *cron.Cron {
	c := cron.New()

	c.AddFunc("@every 00h01m00s", GetEveryDayAt)
	c.AddFunc("@every 00h01m00s", GetEveryHourAt)
	c.AddFunc("@every 00h01m00s", GetHourIsEven)
	c.AddFunc("@every 00h01m00s", GetHourIsOdd)

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

	r.Run(":8082")
	defer c.Stop()
}
