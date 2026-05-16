package github

import (
	"fmt"
	"io"
	"net/http"
)

func FetchRecentActivity(client *http.Client, username string) ([]byte, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	recentActivityRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	recentActivityRequest.Header.Add("Accept", "application/vnd.github+json")
	recentActivityRequest.Header.Add("User-Agent", "sergioferg-github-cli-app")
	recentActivityRequest.Header.Add("X-GitHub-Api-Version", "2026-03-10")

	response, err := client.Do(recentActivityRequest)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected API status code: %d, couldn't handle request", response.StatusCode)
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
