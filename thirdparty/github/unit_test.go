package github

import (
	"net/http"
	"testing"
)

func TestClient_ListPublicEventsForUser(t *testing.T) {
	client := NewClient(http.DefaultClient, nil)
	req := GetUserEventsRequest{
		Username: "jecosine",
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
