package main

import (
	"bytes"
	"context"
	"encoding/json"
	area "google/Area"
	models "google/Models"
	"google/routes"
	"google/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func Events(slice models.DatabaseActions, Compare func(a, b int) bool) int {
	db := utils.OpenDB(nil)

	nbEvents := area.GetNbEvents(slice.UserToken)

	if nbEvents == -1 {
		return -1
	}
	defer db.Close(context.Background())

	if Compare(nbEvents, slice.NbEvents) {
		db.Exec(context.Background(),
			`UPDATE "GoogleActions" 
		SET nb_events = $1 WHERE area_id = $2 AND user_token = $3`,
			nbEvents, slice.AreaId, slice.UserToken)
		return 1
	}
	return 0
}

func callMsgBrocker(areaId string) {
	send := models.MessageBrocker{}
	send.AreaId = areaId
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(send); err != nil {
		return
	}
	http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
}

func NewMessage(slice models.DatabaseActions) int {
	db := utils.OpenDB(nil)

	profile, err := area.GetGmailProfile(slice.UserToken)
	if err != nil {
		return -1
	}
	defer db.Close(context.Background())

	if profile.MessagesTotal > slice.NbMessage {
		db.Exec(context.Background(),
			`UPDATE "GoogleActions" 
		SET nb_message = $1 WHERE area_id = $2 AND user_token = $3`,
			profile.MessagesTotal, slice.AreaId, slice.UserToken)
		return 1
	}

	return 0
}

func BackUpLocalDataCall() {
	db := utils.OpenDB(nil)

	query := `SELECT * FROM "GoogleActions"`
	rows, _ := db.Query(context.Background(), query)

	defer db.Close(context.Background())
	for rows.Next() {
		dbSlice := models.DatabaseActions{}
		err := rows.Scan(&dbSlice.Id, &dbSlice.UserToken, &dbSlice.AreaId, &dbSlice.ActionType, &dbSlice.NbMessage, &dbSlice.NbEvents)
		if err != nil {
			log.Fatal(err)
			continue
		}

		switch dbSlice.ActionType {
		case 0:
			value := Events(dbSlice, func(a, b int) bool { return a > b })
			if value == -1 {
				continue
			}
			if value == 1 {
				callMsgBrocker(dbSlice.AreaId)
				return
			}
		case 1:
			value := Events(dbSlice, func(a, b int) bool { return a < b })
			if value == -1 {
				continue
			}
			if value == 1 {
				callMsgBrocker(dbSlice.AreaId)
				return
			}
		case 2:
			value := NewMessage(dbSlice)
			if value == -1 {
				continue
			}
			if value == 1 {
				callMsgBrocker(dbSlice.AreaId)
				return
			}
		}
	}
}

func InitCronScheduler() *cron.Cron {
	c := cron.New()

	c.AddFunc("@every 00h00m05s", BackUpLocalDataCall)

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
	defer c.Stop()
	r.Run(":8088")
}
