package sgew

import (
	"errors"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

type TriggerCmd struct {
	ID string `name:"id" help:"The ID of the Event Webhook you want to retrieve." required:""`
}

func (c *TriggerCmd) Run() error {
	endpoint := "/v3/user/webhooks/event/test"
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), endpoint, "https://api.sendgrid.com")
	request.Method = "POST"
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
