package github

import (
	"time"
)

type Repo struct {
	Name string `json:"name"`
}

type Event struct {
	Repo Repo      `json:"repo"`
	Type string    `json:"type"`
	Date time.Time `json:"created_at"`
}

type ActivitySummary struct {
	EventType string
	RepoName  string
}

func (e Event) FormatForCLI() string {
	// Inside here, you can access e.Type, e.Repo.Name, etc.
	// and return a nicely formatted string.

	return "a"
}
