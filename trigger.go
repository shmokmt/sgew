package sgew

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

type TriggerCmd struct {
	ID  string `name:"id" help:"The ID of the Event Webhook you want to retrieve." required:""`
	URL string `name:"url" help:"The URL where you would like the test notification to be sent." required:""`
}

type TriggerInput struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (c *TriggerCmd) Run() error {
	endpoint := "/v3/user/webhooks/event/test"
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), endpoint, "https://api.sendgrid.com")
	request.Method = "POST"
	b, err := json.Marshal(&TriggerInput{
		ID:  c.ID,
		URL: c.URL,
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
	fmt.Println("Triggered events succesfully.")
	return nil
}
