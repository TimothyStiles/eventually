package parse

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

var eventHandlerMap = map[string]func(map[string]interface{}, string) (string, error){
	"push":                        genericEventHandler,
	"pull_request":                genericEventHandler,
	"issue":                       genericEventHandler, // in corner case map
	"issue_comment":               genericEventHandler,
	"pull_request_review":         genericEventHandler,
	"pull_request_review_comment": genericEventHandler,
}

var eventCornerCasesMap = map[string]string{
	// "release":       "release",
	"issues": "issue",
	// pull_request:  "pull_request",
	// issue_comment: "issue_comment",
	// discussion:    "discussion",
	// push:          "push",
}

// get Github action payload
func GetGithubActionPayloadURL() (string, error) {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		return "", errors.New("GITHUB_EVENT_PATH is not set")
	}

	// GITHUB_EVENT_NAME
	eventName := os.Getenv("GITHUB_EVENT_NAME")
	if eventName == "" {
		return "", errors.New("GITHUB_EVENT_NAME is not set")
	}

	// open and read the event file
	eventFile, err := os.Open(eventPath)
	if err != nil {
		return "", errors.Wrap(err, "failed to open event file")
	}

	// decode the event file
	var event map[string]interface{}
	if err := json.NewDecoder(eventFile).Decode(&event); err != nil {
		return "", errors.Wrap(err, "failed to decode event file")
	}
	if eventCornerCasesMap[eventName] != "" {
		eventName = eventCornerCasesMap[eventName]
	}

	handler := eventHandlerMap[eventName]
	if handler == nil {
		return "", errors.Errorf("no handler for event %s", eventName)
	}

	return handler(event, eventName)
}

func genericEventHandler(event map[string]interface{}, eventName string) (string, error) {
	payload := event[eventName]
	for key, value := range payload.(map[string]interface{}) {
		fmt.Println(key, value)
		switch value.(type) {
		case string:
			if key == "html_url" {
				return value.(string), nil
			}
			continue
		default:
			// return "", errors.Errorf("unexpected type %T", value)
		}
	}
	return "", nil
}
