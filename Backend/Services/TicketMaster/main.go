package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	models "ticket-master/Models"
	"ticket-master/routes"
	"ticket-master/utils"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func FetchEvents(city string, genre string) int {
	url := fmt.Sprintf(
		"https://app.ticketmaster.com/discovery/v2/events.json?apikey=%s&city=%s&classificationName=%s",
		utils.GetEnvKey("API_KEY"), city, genre,
	)

	resp, err := http.Get(url)
	if err != nil {
		return -1
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return -1
	}

	eventsCount := 0.0
	if embedded, ok := data["_embedded"].(map[string]interface{}); ok {
		if eventsData, ok := embedded["events"].([]interface{}); ok {
			return len(eventsData)
		}
	}

	if page, ok := data["page"].(map[string]interface{}); ok {
		eventsCount = page["size"].(float64)
	}

	return int(eventsCount)
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

func updateDatabase(AreaID string, len int) {
	db := utils.OpenDB(nil)
	defer db.Close(context.Background())

	updateQuery := `UPDATE "TicketMasterActions" SET nb_events = $1 WHERE area_id = $2`
	result, err := db.Exec(context.Background(), updateQuery, len, AreaID)
	if err != nil {
		fmt.Printf("Error updating nb_events: %v\n", err)
	} else {
		rowsAffected := result.RowsAffected()
		fmt.Printf("Rows affected: %d\n", rowsAffected)
	}
}

func BackUpLocalDataCall() {
	db := utils.OpenDB(nil)
	defer db.Close(context.Background())

	if db == nil {
		return
	}
	query := `SELECT area_id, action_type, name, venue, city, nb_events FROM "TicketMasterActions"`
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var action models.Action
		if err := rows.Scan(&action.AreaID, &action.ActionType, &action.Name, &action.Venue, &action.City, &action.NbEvents); err != nil {
			continue
		}

		switch action.ActionType {
		case 0:
			len := FetchEvents(action.City, action.Venue)
			if action.NbEvents < len {
				updateDatabase(action.AreaID, len)
				callMsgBrocker(action.AreaID)
			}
		}
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

	r.Run(":8093")
	defer c.Stop()
}
