package parse

import (
	"testing"
)

func TestGetGithubActionPayloadURL(t *testing.T) {
	payload, err := GetGithubActionPayloadURL()
	if payload == "" {
		t.Error("GITHUB_ACTIONS is not set")
	}
	if err != nil {
		t.Error(err)
	}
}

func TestGenericEventHandler(t *testing.T) {
	payload := map[string]interface{}{
		"action": "opened",
		"issue": map[string]interface{}{
			"number":   1,
			"title":    "test",
			"html_url": "stiles.io",
		},
	}
	output, err := genericEventHandler(payload, "issue")
	if err != nil {
		t.Error(err)
	}
	if output != "opened issue 1 test" {
		t.Errorf("expected output to be 'opened issue 1 test', got %s", output)
	}
}
