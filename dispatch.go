package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/go-playground/webhooks.v5/github"
)

func main() {
	githubEventPath := os.Getenv("GITHUB_EVENT_PATH")
	// githubToken := os.Getenv("GITHUB_TOKEN")
	eventFile, err := ioutil.ReadFile(githubEventPath)
	if err != nil {
		log.Fatal(err)
	}

	var event github.PullRequestPayload
	err = json.Unmarshal(eventFile, &event)
	if err != nil {
		log.Fatal(err)
	}

	if event.Action == "closed" && event.PullRequest.Merged {
		// call hecate api asking for an email
		log.Printf("Merged PR #%v.", event.PullRequest.Number)
	} else {
		log.Printf("No need to do anything for PR #%v. Action: %v, merged: %v", event.PullRequest.Number, event.Action, event.PullRequest.Merged)
		os.Exit(78)
	}
}
