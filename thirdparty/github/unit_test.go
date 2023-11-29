package github

import (
	"net/http"
	"os"
	"testing"
)

func TestClient_ListPublicEventsForUser(t *testing.T) {
	client := NewClient(http.DefaultClient, nil)
	userName := os.Getenv("GITHUB_USERNAME")
	req := GetUserEventsRequest{
		Username: userName,
	}
	if userName == "" {
		t.Skipf("username is empty")
	}
	resp, err := client.ListPublicEventsForUser(req)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
	t.Logf("resp number: %d", len(resp))
	for _, event := range resp {
		t.Logf("%s\n", event.Type)
	}
}
