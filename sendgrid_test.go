package sgew

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAllWebhookEvents(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v3/user/webhooks/event/settings/all", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"max_allowed": 5, "webhooks": [{"id": "test", "created_date": "2017-01-01", "enabled": true, "friendly_name": "test", "url": "https://example.com", "public_key": "test"}]}`)
	})

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	sg := newSendGrid("test-key", testServer.URL)
	got, err := sg.getAllEventWebhooks()
	if err != nil {
		t.Fatal(err)
	}

	expect := &GetAllEventWebhooksOutput{
		Webhooks: []struct {
			ID           string `json:"id"`
			CreatedDate  string `json:"created_date"`
			Enabled      bool   `json:"enabled"`
			FriendlyName string `json:"friendly_name"`
			URL          string `json:"url"`
		}{
			{
				ID:           "test",
				CreatedDate:  "2017-01-01",
				Enabled:      true,
				FriendlyName: "test",
				URL:          "https://example.com",
			},
		},
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %#v, got %#v", expect, got)
	}
}
func TestSendEventNotificationForTest(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v3/user/webhooks/event/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	sg := newSendGrid("test-key", testServer.URL)
	err := sg.sendEventNotificationForTest("test-id", "https://example.com")
	if err != nil {
		t.Error(err)
	}

}

func TestCreateEventWebhook(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v3/user/webhooks/event/settings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		fmt.Fprint(w, `{"id": "test", "created_date": "2017-01-01", "enabled": true, "friendly_name": "test", "url": "https://example.com", "public_key": "test"}`)
	})
	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	sg := newSendGrid("test-key", testServer.URL)
	outputs, err := sg.CreateEventWebhook("https://example.com")
	if err != nil {
		t.Fatal(err)
	}
	if outputs.ID != "test" {
		t.Errorf("expected %#v, got %#v", "test", outputs.ID)
	}
}

func TestDeleteEventWebhook(t *testing.T) {
	mux := http.NewServeMux()
	testID := "test-id"
	mux.HandleFunc("/v3/user/webhooks/event/settings/"+testID, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	sg := newSendGrid("test-key", testServer.URL)
	err := sg.deleteEventWebhook(testID)
	if err != nil {
		t.Error(err)
	}
}
