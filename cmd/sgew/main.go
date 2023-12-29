package main

import (
	"github.com/alecthomas/kong"
	"github.com/shmokmt/sgew"
)

var CLI struct {
	Listen  sgew.ListenCmd  `cmd:"" help:"Listen for webhook events"`
	Trigger sgew.TriggerCmd `cmd:"" help:"trigger test webhook events"`
	Ls      sgew.LsCmd      `cmd:"" help:"List all event webhooks"`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("sgew"),
		kong.Description("SendGrid Event Webhook Debugger"),
	)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
