package main

import (
	"fmt"
	"time"
)

type EventResponse struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		ID           int    `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		GravatarID   string `json:"gravatar_id"`
		URL          string `json:"url"`
		AvatarURL    string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		RepositoryID int    `json:"repository_id"`
		PushID       int64  `json:"push_id"`
		Size         int    `json:"size"`
		DistinctSize int    `json:"distinct_size"`
		Ref          string `json:"ref"`
		Head         string `json:"head"`
		Before       string `json:"before"`
		Commits      []struct {
			Sha    string `json:"sha"`
			Author struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"author"`
			Message  string `json:"message"`
			Distinct bool   `json:"distinct"`
			URL      string `json:"url"`
		} `json:"commits"`
	} `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}

type Event struct {
	Handler func(e EventResponse) string
}

func GetEvents() map[string]Event {
	return map[string]Event{
		"PushEvent": {
			Handler: HandlerPush,
		},
	}
}

func HandlerPush(e EventResponse) string {
	return fmt.Sprintf("Pushed %d commits to %s", len(e.Payload.Commits), e.Repo.Name)
}
