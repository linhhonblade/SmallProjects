package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

const (
	requestsPerMinute = 1000
	url               = "http://localhost:3000/v1/products"
)

func makeRequest(client *http.Client) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	resp.Body.Close()
}

var loadTestCmd = &cobra.Command{
	Use:   "loadtest",
	Short: "Send 1000req/min to get list product curl --location 'localhost:3000/v1/products'",
	Run: func(cmd *cobra.Command, args []string) {
		rateLimit := rate.Every(time.Minute / requestsPerMinute)
		limiter := rate.NewLimiter(rateLimit, 1)
		client := &http.Client{}
		for {
			err := limiter.Wait(context.Background()) // Block until it's OK to proceed
			if err != nil {
				fmt.Println("Limiter error:", err)
				continue
			}
			go makeRequest(client)
		}
	},
}
