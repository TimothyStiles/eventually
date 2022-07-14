package main

import (
	"os"

	"github.com/the-gophers/go-action/pkg/parse"
)

func main() {
	payloadURL, _ := parse.GetGithubActionPayloadURL()
	os.Setenv("GITHUB_PAYLOAD_URL", payloadURL)
}
