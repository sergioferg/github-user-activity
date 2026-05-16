package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sergioferg/github-user-activity/internal/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Please provide a username.")
		fmt.Fprintf(os.Stderr, "Usage: %s <username>\n", os.Args[0])
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "Error: Too many arguments")
		fmt.Fprintf(os.Stderr, "Usage: %s <username>\n", os.Args[0])
		os.Exit(1)
	}
	username := os.Args[1]

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	data, err := github.FetchRecentActivity(client, username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	myEventMap := github.SummarizeActivity(data)
	activities := github.FormatActivities(myEventMap)
	for _, activity := range activities {
		fmt.Println(activity)
	}
}
