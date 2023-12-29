package sgew

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type LsCmd struct {
}

func (c *LsCmd) Run() error {
	sg := newSendGrid(os.Getenv("SENDGRID_API_KEY"), "https://api.sendgrid.com")
	outputs, err := sg.getAllEventWebhooks()
	if err != nil {
		return err
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Enabled", "URL", "Friendly Name", "Created Date"})
	for _, webhook := range outputs.Webhooks {
		t.AppendRow(table.Row{webhook.ID, fmt.Sprint(webhook.Enabled), webhook.URL, webhook.FriendlyName, webhook.CreatedDate})
		t.AppendSeparator()
	}
	t.Render()
	return nil
}
