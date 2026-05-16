package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sergioferg/github-user-activity/internal/github"
)

func main() {
	username := os.Args[1]
	fmt.Println(username)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	data, err := github.FetchRecentActivity(client, username)
	if err != nil {
		//TODO: handle errors
		fmt.Println("ERROR: %e", err)
		return
	}
	fmt.Println(string(data))

}
