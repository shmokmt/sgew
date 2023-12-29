package sgew

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

type LsCmd struct {
}

type LsOutputs struct {
	MaxAllowed int `json:"max_allowed"`
	Webhooks   []struct {
		ID           string `json:"id"`
		CreatedDate  string `json:"created_date"`
		Enabled      bool   `json:"enabled"`
		FriendlyName string `json:"friendly_name"`
		URL          string `json:"url"`
		PublicKey    string `json:"public_key,omitempty"`
	} `json:"webhooks"`
}

func (c *LsCmd) Run() error {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/user/webhooks/event/settings/all", "https://api.sendgrid.com")
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New("expected status: 200, got status: " + fmt.Sprint(response.StatusCode))
	}
	outputs := &LsOutputs{}
	err = json.Unmarshal([]byte(response.Body), outputs)
	if err != nil {
		return err
	}
	fmt.Println("ID\tEnabled\tURL\tFriendlyName")
	for _, webhook := range outputs.Webhooks {
		fmt.Println(webhook.ID + "\t" + fmt.Sprint(webhook.Enabled) + "\t" + webhook.URL + "\t" + webhook.FriendlyName)
	}
	return nil
}
