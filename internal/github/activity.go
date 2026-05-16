package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchRecentActivity(client *http.Client, username string) ([]Event, error) {
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

	var events []Event

	err = json.Unmarshal(data, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func SummarizeActivity(events []Event) map[ActivitySummary]int {
	eventMap := make(map[ActivitySummary]int)
	for _, val := range events {
		eventMap[ActivitySummary{
			EventType: val.Type,
			RepoName:  val.Repo.Name,
		}]++
	}

	return eventMap
}

func FormatActivities(eventMap map[ActivitySummary]int) []string {
	var activities []string
	for summaryKey, count := range eventMap {
		switch summaryKey.EventType {
		case "PushEvent":
			if count > 1 {
				activities = append(activities, fmt.Sprintf("Pushed %d commits to %s", count, summaryKey.RepoName))
			} else {
				activities = append(activities, fmt.Sprintf("Pushed %d commit to %s", count, summaryKey.RepoName))
			}
		case "IssueCommentEvent":
			if count > 1 {
				activities = append(activities, fmt.Sprintf("Left %d comments on issues in %s", count, summaryKey.RepoName))
			} else {
				activities = append(activities, fmt.Sprintf("Left %d comment on issues in %s", count, summaryKey.RepoName))
			}
		case "CreateEvent":
			if count > 1 {
				activities = append(activities, fmt.Sprintf("Created %d branches or tags in %s", count, summaryKey.RepoName))
			} else {
				activities = append(activities, fmt.Sprintf("Created %d branch or tag in %s", count, summaryKey.RepoName))
			}
		case "DeleteEvent":
			if count > 1 {
				activities = append(activities, fmt.Sprintf("Deleted %d branches or tags in %s", count, summaryKey.RepoName))
			} else {
				activities = append(activities, fmt.Sprintf("Deleted %d branch or tag in %s", count, summaryKey.RepoName))
			}
		case "PullRequestReviewEvent":
			if count > 1 {
				activities = append(activities, fmt.Sprintf("Submitted %d pull requests reviews in %s", count, summaryKey.RepoName))
			} else {
				activities = append(activities, fmt.Sprintf("Submitted %d pull request review in %s", count, summaryKey.RepoName))
			}
		case "IssuesEvent":
			if count > 1 {
				activities = append(activities, fmt.Sprintf("Opened/closed %d issues in %s", count, summaryKey.RepoName))
			} else {
				activities = append(activities, fmt.Sprintf("Opened/closed %d issue in %s", count, summaryKey.RepoName))
			}
		case "PullRequestEvent":
			if count > 1 {
				activities = append(activities, fmt.Sprintf("Submitted %d pull requests in %s", count, summaryKey.RepoName))
			} else {
				activities = append(activities, fmt.Sprintf("Submitted %d pull request in %s", count, summaryKey.RepoName))
			}

		case "WatchEvent":
			activities = append(activities, fmt.Sprintf("Starred %s", summaryKey.RepoName))

		case "ForkEvent":
			activities = append(activities, fmt.Sprintf("Forked %s", summaryKey.RepoName))

		default:
			activities = append(activities, fmt.Sprintf("%s (%d events) on %s", summaryKey.EventType, count, summaryKey.RepoName))
		}
	}

	return activities
}
