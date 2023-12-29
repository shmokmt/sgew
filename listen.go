package sgew

import "errors"

type ListenCmd struct {
	URL string `name:"forward-to" help:"Webhook listener's port" required:"" type:"path"`
}

func (c *ListenCmd) Run() error {
	return errors.New("NotImplementation")
}
