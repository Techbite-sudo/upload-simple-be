package main

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	// Ping your application every 25 minutes
	c.AddFunc("*/25 * * * *", func() {
		fmt.Println("Pinging the application to keep it alive")
		// Replace the following URL with your application's URL
		_, err := http.Get("https://upload-simple-backend.onrender.com/query")
		if err != nil {
			fmt.Printf("Error pinging the application: %v\n", err)
		} else {
			fmt.Println("Application pinged successfully")
		}
	})

	c.Start()

	// Block the main goroutine to keep the application running
	select {}
}