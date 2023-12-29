package sgew

import (
	"fmt"
	"os"
)

type TriggerCmd struct {
	ID  string `name:"id" help:"The ID of the Event Webhook you want to retrieve." required:""`
	URL string `name:"url" help:"The URL where you would like the test notification to be sent." required:""`
}

type triggerInput struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (c *TriggerCmd) Run() error {
	sg := newSendGrid(os.Getenv("SENDGRID_API_KEY"), "https://api.sendgrid.com")
	err := sg.sendEventNotificationForTest(c.ID, c.URL)
	if err != nil {
		return err
	}
	fmt.Println("Triggered events succesfully.")
	return nil
}
