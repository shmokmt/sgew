package sgew

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

type ListenCmd struct {
	URL string `name:"forward-to" help:"Webhook listener's port" required:"" type:"path"`
}

func (c *ListenCmd) Run() error {
	ctx := context.Background()
	tunnel, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}
	fmt.Println("Ingress established at:", tunnel.URL())
	return http.Serve(tunnel, HTTPMethodValidationMiddleware(PrintRequestBodyMiddleware(okHandler)))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func PrintRequestBodyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bufOfRequestBody, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bufOfRequestBody))
		fmt.Println(string(bufOfRequestBody))
		next.ServeHTTP(w, r)
	}
}

func HTTPMethodValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "Method not allowed")
			return
		}
		next.ServeHTTP(w, r)
	}
}
