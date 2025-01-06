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
	area "spotify/Area"
	models "spotify/Models"
	"spotify/routes"
	"spotify/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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
		err := rows.Scan(&database.Id, &database.AreaId, &database.ActionType, &database.AccessToken, &database.UserId, &database.IsPlaying, &database.NbPlaylists)
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

func listenActions(db *pgx.Conn, slice models.Database) int {
	client := &http.Client{}
	req, err := http.NewRequest("GET", utils.GetEnvKey("SPOTIFY_PLAYER_API")+"me/player", nil)

	if err != nil {
		return 1
	}

	req.Header.Set("Authorization", "Bearer "+slice.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Error while calling the API:", err)
		return 1
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent && slice.IsPlaying == 1 {
		db.Exec(context.Background(),
			`UPDATE "SpotifyActions" 
		 SET is_playing = CASE WHEN is_playing = 1 THEN 0 ELSE 1 END
		 WHERE area_id = $1 AND user_token = $2`, slice.AreaId, slice.AccessToken)
		return 1
	}

	if resp.StatusCode == http.StatusNoContent {
		return 1
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading the response body:", err)
		return 1
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
			return -1
		}
		http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
	}

	return 0
}

func addPlaylistAction(db *pgx.Conn, slice models.Database) int {
	actualNbPlaylists := area.GetNbPlaylists(slice.AccessToken, slice.UserId)

	if actualNbPlaylists > slice.NbPlaylists {
		var send models.TriggerModelGateway
		var buf bytes.Buffer
		send.AreaId = slice.AreaId
		if err := json.NewEncoder(&buf).Encode(send); err != nil {
			return -1
		}
		db.Exec(context.Background(),
			`UPDATE "SpotifyActions" 
			SET nb_playlists = $1 WHERE area_id = $2 AND user_token = $3
		`, slice.NbPlaylists+1, slice.AreaId, slice.AccessToken)
		http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
	}

	return 0
}

func removePlaylistAction(db *pgx.Conn, slice models.Database) int {
	actualNbPlaylists := area.GetNbPlaylists(slice.AccessToken, slice.UserId)

	if actualNbPlaylists < slice.NbPlaylists {
		var send models.TriggerModelGateway
		var buf bytes.Buffer
		send.AreaId = slice.AreaId
		if err := json.NewEncoder(&buf).Encode(send); err != nil {
			return -1
		}
		db.Exec(context.Background(),
			`UPDATE "SpotifyActions" 
			SET nb_playlists = $1 WHERE area_id = $2 AND user_token = $3
		`, slice.NbPlaylists-1, slice.AreaId, slice.AccessToken)
		http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
	}

	return 0
}

func BackUpLocalDataCall() {
	databaseSlice := getDatabaseSlice()
	db := utils.OpenDB(nil)
	if db == nil || databaseSlice == nil {
		return
	}

	for _, slice := range databaseSlice {

		if slice.ActionType == 0 {
			switch listenActions(db, slice) {
			case 1:
				continue
			case -1:
				return
			}
		}

		if slice.ActionType == 1 {
			switch addPlaylistAction(db, slice) {
			case 1:
				continue
			case -1:
				return
			}
		}

		if slice.ActionType == 2 {
			switch removePlaylistAction(db, slice) {
			case 1:
				continue
			case -1:
				return
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
	r.Run(":8091")
	defer c.Stop()
}
