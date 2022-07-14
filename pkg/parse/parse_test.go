package parse

import (
	"os"
	"testing"
)

// func TestGetGithubActionPayloadURL(t *testing.T) {
// 	payload, err := GetGithubActionPayloadURL()
// 	if payload == "" {
// 		t.Error("GITHUB_ACTIONS is not set")
// 	}
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

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
	if output != "stiles.io" {
		t.Errorf("expected output to be 'stiles.io', got %s", output)
	}
}

func TestGetGithubActionPayloadURL(t *testing.T) {
	tests := []struct {
		name      string
		eventName string
		eventPath string
		want      string
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name:      "GITHUB_EVENT_PATH is not set",
			eventName: "push",
			wantErr:   true,
		},
		{
			name:      "GITHUB_EVENT_NAME is not set",
			eventPath: "./data/issue.json",
			wantErr:   true,
		},
		{
			name:      "GITHUB_EVENT_NAME is issue",
			eventName: "issue",
			eventPath: "data/issues.json",
			want:      "https://github.com/Codertocat/Hello-World/issues/1",
		},
		// {
		// 	name:      "GITHUB_EVENT_NAME is push",
		// 	eventName: "push",
		// 	eventPath: "data/push.json",
		// 	want:      "https://github.com/Codertocat/Hello-World/commit/6113728f27ae82c7b1a177c8d03f9e96e0adf246",
		// },
		{
			name:      "GITHUB_EVENT_NAME is pull_request",
			eventName: "pull_request",
			eventPath: "data/pull_request.json",
			want:      "https://github.com/Codertocat/Hello-World/pull/2",
		},
		// {
		// 	name:      "GITHUB_EVENT_NAME is issue_comment",
		// 	eventName: "issue_comment",
		// 	eventPath: "data/issue_comment.json",
		// 	want:      "stiles.io",
		// },
		// {
		// 	name:      "GITHUB_EVENT_NAME is discussion",
		// 	eventName: "discussion",
		// 	eventPath: "data/discussion.json",
		// 	want:      "stiles.io",
		// },
		// {
		// 	name:      "GITHUB_EVENT_NAME is pull_request_review",
		// 	eventName: "pull_request_review",
		// 	eventPath: "data/pull_request_review.json",
		// 	want:      "stiles.io",
		// },
		// {
		// 	name:      "GITHUB_EVENT_NAME is pull_request_review_comment",
		// 	eventName: "pull_request_review_comment",
		// 	eventPath: "data/pull_request_review_comment.json",
		// 	want:      "stiles.io",
		// },
	}
	for _, tt := range tests {
		os.Setenv("GITHUB_EVENT_NAME", tt.eventName)
		os.Setenv("GITHUB_EVENT_PATH", tt.eventPath)
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGithubActionPayloadURL()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGithubActionPayloadURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetGithubActionPayloadURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
