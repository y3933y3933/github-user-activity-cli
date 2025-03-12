package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const cliName = "github-activity"

func main() {
	if len(os.Args) < 3 || os.Args[1] != cliName {
		fmt.Printf("Usage: %s <username>\n", cliName)
		return
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Args[2]
	apiToken := os.Getenv("API_TOKEN")
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Generate request fail :%v", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("get response fail: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status code : %v", res.StatusCode)
	}

	var events []EventResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&events)
	if err != nil {
		log.Fatalf("decode json fail: %v", err)
	}

	fmt.Println("Output: ")
	for _, e := range events {
		event, exist := GetEvents()[e.Type]
		if !exist {
			continue
		}

		event.ShowMsg(e)

	}
}
