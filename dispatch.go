package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/badoux/checkmail"
	"gopkg.in/go-playground/webhooks.v5/github"
)

type dispatchRequest struct {
	PrURL  string
	APIKey string
	Emails []string
}

var apiURL = "https://api.hecate.co/github_actions/merge"

func main() {
	githubEventPath := os.Getenv("GITHUB_EVENT_PATH")
	githubToken := os.Getenv("GITHUB_TOKEN")
	rawEmails := os.Getenv("EMAILS")

	log.Printf("EMAILS to send are %v", rawEmails)
	emails := []string{}
	for _, email := range strings.Split(rawEmails, ",") {
		email = strings.TrimSpace(email)
		err := checkmail.ValidateFormat(email)
		if err != nil {
			log.Fatalf("Error with email %v: %v", email, err)
		}
		emails = append(emails, strings.TrimSpace(email))
	}

	if len(emails) == 0 {
		log.Fatal("Must define EMAILS in ENV (comma separated)")
	}

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

		requestBody := dispatchRequest{
			PrURL:  event.PullRequest.URL,
			APIKey: githubToken,
			Emails: emails,
		}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			log.Fatal("Unable to serialise request to Hecate API")
		}
		resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonBody))
		if err != nil {
			log.Fatalf("Error talking to Hecate API %v", err)
		}
		if resp.StatusCode != http.StatusNoContent {
			log.Fatalf("Hecate API returned error code %v", resp.StatusCode)
		}
	} else {
		log.Printf("No need to do anything for PR #%v. Action: %v, merged: %v", event.PullRequest.Number, event.Action, event.PullRequest.Merged)
		os.Exit(78)
	}
}
