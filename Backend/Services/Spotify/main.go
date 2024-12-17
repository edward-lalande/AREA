package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	models "spotify/Models"
	"spotify/routes"
	"spotify/utils"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func getDatabaseSlice() []models.Database {
	var databaseSlice []models.Database = nil
	db := utils.OpenDB(nil)

	if db == nil {
		return nil
	}
	rows, err := db.Query(context.Background(), "SELECT * FROM \"SpotifyActions\"")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error on reading response of the query", err)
		return nil
	}
	defer rows.Close()
	defer db.Close(context.Background())

	for rows.Next() {
		var database models.Database
		err := rows.Scan(&database.Id, &database.AreaId, &database.ActionType, &database.AccessToken, &database.IsPlaying, &database.MusicName)
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

func CallMessageBrocker(AreaId string) {

}

func BackUpLocalDataCall() {
	databaseSlice := getDatabaseSlice()
	db := utils.OpenDB(nil)
	if db == nil || databaseSlice == nil {
		return
	}

	for _, slice := range databaseSlice {

		if slice.ActionType == 0 {

			client := &http.Client{}
			req, err := http.NewRequest("GET", utils.GetEnvKey("SPOTIFY_PLAYER_API"), nil)
			if err != nil {
				return
			}
			req.Header.Set("Authorization", "Bearer "+slice.AccessToken)
			req.Header.Set("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal("Error while calling the API:", err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusNoContent && slice.IsPlaying == 1 {
				db.Exec(context.Background(),
					`UPDATE "SpotifyActions" 
				 SET is_playing = CASE WHEN is_playing = 1 THEN 0 ELSE 1 END
				 WHERE area_id = $1 AND user_token = $2`, slice.AreaId, slice.AccessToken)
				continue
			}

			if resp.StatusCode == http.StatusNoContent {
				continue
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal("Error reading the response body:", err)
				return
			}
			jsonBody := utils.BytesToJson(body)
			if jsonBody != nil && slice.IsPlaying == 0 {
				db.Exec(context.Background(),
					`UPDATE "SpotifyActions" 
				 SET is_playing = CASE WHEN is_playing = 1 THEN 0 ELSE 1 END
				 WHERE area_id = $1 AND user_token = $2`, slice.AreaId, slice.AccessToken)
				send := struct {
					AreaId string `json:"area_id"`
				}{slice.AreaId}
				var buf bytes.Buffer
				if err := json.NewEncoder(&buf).Encode(send); err != nil {
					return
				}
				http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
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
	r.Run(":8091")
	defer c.Stop()
}
