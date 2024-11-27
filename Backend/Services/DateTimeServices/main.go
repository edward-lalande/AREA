package main

import (
	"date-time-service/routes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func BackUpLocalDataCall() {
	resp, err := http.Get("https://timeapi.io/api/time/current/zone?timeZone=Europe%2FParis")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while calling the API:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading the response body:", err)
		return
	}

	fmt.Println(string(body))
}

func InitCronScheduler() *cron.Cron {
	c := cron.New()

	c.AddFunc("@every 00h00m25s", BackUpLocalDataCall)

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

	r.Run(":8082")
	defer c.Stop()
}
