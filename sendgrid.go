package sgew

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
)

type sendGrid struct {
	apiKey string
	host   string
}

func newSendGrid(apiKey string, host string) *sendGrid {
	return &sendGrid{
		apiKey: apiKey,
		host:   host,
	}
}

type GetAllEventWebhooksOutput struct {
	Webhooks []struct {
		ID           string `json:"id"`
		CreatedDate  string `json:"created_date"`
		Enabled      bool   `json:"enabled"`
		FriendlyName string `json:"friendly_name"`
		URL          string `json:"url"`
	} `json:"webhooks"`
}

func (s *sendGrid) getAllEventWebhooks() (*GetAllEventWebhooksOutput, error) {
	endpoint := "/v3/user/webhooks/event/settings/all"
	request := sendgrid.GetRequest(s.apiKey, endpoint, s.host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(endpoint + ": expected status is 200, got status is " + fmt.Sprint(response.StatusCode))
	}
	output := &GetAllEventWebhooksOutput{}
	err = json.Unmarshal([]byte(response.Body), output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// ref. https://docs.sendgrid.com/api-reference/webhooks/test-event-notification-settings
func (s *sendGrid) sendEventNotificationForTest(id string, url string) error {
	endpoint := "/v3/user/webhooks/event/test"
	request := sendgrid.GetRequest(s.apiKey, endpoint, s.host)
	request.Method = "POST"
	b, err := json.Marshal(&triggerInput{
		ID:  id,
		URL: url,
	})
	if err != nil {
		return err
	}
	request.Body = b
	response, err := sendgrid.API(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return errors.New(endpoint + ": expected status is 204, got status is " + fmt.Sprint(response.StatusCode))
	}
	if err != nil {
		return err
	}
	return nil
}

type createWebhookInput struct {
	Enabled          bool   `json:"enabled"`
	URL              string `json:"url"`
	GroupReSubscribe bool   `json:"group_resubscribe"`
	Delivered        bool   `json:"delivered"`
	GroupUnsubscribe bool   `json:"group_unsubscribe"`
	SpamReport       bool   `json:"spam_report"`
	Bounce           bool   `json:"bounce"`
	Deferred         bool   `json:"deferred"`
	Unsubscribe      bool   `json:"unsubscribe"`
	Processed        bool   `json:"processed"`
	Open             bool   `json:"open"`
	Click            bool   `json:"click"`
	Dropped          bool   `json:"dropped"`
	FriendlyName     string `json:"friendly_name"`
}

type createWebhookOutput struct {
	ID string `json:"id"`
}

// ref. https://docs.sendgrid.com/api-reference/webhooks/create-an-event-webhook
func (s *sendGrid) CreateEventWebhook(url string) (*createWebhookOutput, error) {
	endpoint := "/v3/user/webhooks/event/settings"
	request := sendgrid.GetRequest(s.apiKey, endpoint, s.host)
	b, err := json.Marshal(&createWebhookInput{
		Enabled:          true,
		URL:              url,
		GroupReSubscribe: true,
		Delivered:        true,
		GroupUnsubscribe: true,
		SpamReport:       true,
		Bounce:           true,
		Deferred:         true,
		Unsubscribe:      true,
		Processed:        true,
		Open:             true,
		Click:            true,
		Dropped:          true,
		FriendlyName:     "created by sgew",
	})
	if err != nil {
		return nil, err
	}
	request.Method = "POST"
	request.Body = b
	response, err := sendgrid.API(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, errors.New(endpoint + ": expected status is 201, got status is " + fmt.Sprint(response.StatusCode))
	}
	output := &createWebhookOutput{}
	json.Unmarshal([]byte(response.Body), output)
	return output, nil
}

// ref. https://docs.sendgrid.com/api-reference/webhooks/delete-an-event-webhook
func (s *sendGrid) deleteEventWebhook(id string) error {
	endpoint := "/v3/user/webhooks/event/settings/" + id
	request := sendgrid.GetRequest(s.apiKey, endpoint, s.host)
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return errors.New(endpoint + ": expected status is 204, got status is " + fmt.Sprint(response.StatusCode))
	}
	return nil
}
