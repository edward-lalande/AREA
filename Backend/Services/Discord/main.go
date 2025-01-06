package main

import (
	"bytes"
	"context"
	models "discord-service/Models"
	"discord-service/routes"
	"discord-service/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func getDatabaseSlice() []models.Database {
	var databaseSlice []models.Database = nil
	db := utils.OpenDB(nil)

	if db == nil {
		return nil
	}
	rows, err := db.Query(context.Background(), "SELECT * FROM \"DiscordAction\"")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error on reading response of the query", err)
		return nil
	}
	defer rows.Close()
	defer db.Close(context.Background())

	for rows.Next() {
		var database models.Database
		err := rows.Scan(&database.Id, &database.Type, &database.AreaId, &database.ChannelId, &database.MessageId, &database.UserToken)
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

func BackUpLocalDataCall() {
	databaseSlice := getDatabaseSlice()

	if databaseSlice == nil {
		return
	}

	for _, slice := range databaseSlice {

		if slice.Type == 0 {

			url := fmt.Sprintf(
				"https://discord.com/api/v10/channels/%s/messages/%s/reactions/%s",
				slice.ChannelId,
				slice.MessageId,
				"✅",
			)

			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)

			if err != nil {
				return
			}

			req.Header.Set("Authorization", "Bot "+utils.GetEnvKey("BOT_TOKEN"))
			req.Header.Set("Content-Type", "application/json")

			response, err := client.Do(req)

			if err != nil {
				return
			}

			responseData, err := io.ReadAll(response.Body)

			if err != nil {
				return
			}

			var messages []any

			err = json.Unmarshal(responseData, &messages)

			if err != nil {
				return
			}

			if len(messages) > 0 {
				send := models.DiscordModelSendReaction{}
				send.ReactionId = slice.AreaId
				var buf bytes.Buffer
				if err := json.NewEncoder(&buf).Encode(send); err != nil {
					return
				}
				http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)

				url := fmt.Sprintf(
					"https://discord.com/api/v10/channels/%s/messages/%s/reactions/%s",
					slice.ChannelId,
					slice.MessageId,
					"✅",
				)

				client := &http.Client{}
				req, err := http.NewRequest("DELETE", url, nil)

				if err != nil {
					return
				}

				req.Header.Set("Authorization", "Bot "+utils.GetEnvKey("BOT_TOKEN"))
				req.Header.Set("Content-Type", "application/json")

				client.Do(req)
			}

		}

		if slice.Type == 1 {

			url := fmt.Sprintf(
				"https://discord.com/api/v10/channels/%s/messages/%s",
				slice.ChannelId,
				slice.MessageId,
			)

			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)

			if err != nil {
				return
			}

			req.Header.Set("Authorization", "Bot "+utils.GetEnvKey("BOT_TOKEN"))
			req.Header.Set("Content-Type", "application/json")

			response, err := client.Do(req)

			if err != nil {
				return
			}

			responseData, err := io.ReadAll(response.Body)

			if err != nil {
				return
			}

			var message models.Message

			err = json.Unmarshal(responseData, &message)

			if err != nil {
				return
			}

			if message.Pinned {
				send := models.DiscordModelSendReaction{}
				send.ReactionId = slice.AreaId
				var buf bytes.Buffer
				if err := json.NewEncoder(&buf).Encode(send); err != nil {
					return
				}
				http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)

				url := fmt.Sprintf(
					"https://discord.com/api/v10/channels/%s/pins/%s",
					slice.ChannelId,
					slice.MessageId,
				)

				client := &http.Client{}
				req, err := http.NewRequest("DELETE", url, nil)

				if err != nil {
					return
				}

				req.Header.Set("Authorization", "Bot "+utils.GetEnvKey("BOT_TOKEN"))
				req.Header.Set("Content-Type", "application/json")

				client.Do(req)
			}

		}

	}
}

func InitCronScheduler() *cron.Cron {
	c := cron.New()

	c.AddFunc("@every 00h00m01s", BackUpLocalDataCall)

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

	r.Run(":8083")
	defer c.Stop()
}
