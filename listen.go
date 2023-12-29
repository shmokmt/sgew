package sgew

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

type ListenCmd struct {
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
	sg := newSendGrid(os.Getenv("SENDGRID_API_KEY"), "https://api.sendgrid.com")
	outputs, err := sg.CreateEventWebhook(tunnel.URL())
	if err != nil {
		return err
	}
	fmt.Printf("Event Webhook ID: %s\n", outputs.ID)
	go func() {
		trap := make(chan os.Signal, 1)
		signal.Notify(trap, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		<-trap
		// TODO: err check
		sg.deleteEventWebhook(outputs.ID)
		fmt.Println("\nDeleted an event webhook")
		os.Exit(0)
	}()
	return http.Serve(tunnel, methodValidationMiddleware(printRequestBodyMiddleware(okHandler)))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func printRequestBodyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bufOfRequestBody, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bufOfRequestBody))
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, bufOfRequestBody, "", "  ")
		fmt.Println(prettyJSON.String())
		fmt.Println()
		next.ServeHTTP(w, r)
	}
}

func methodValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "Method not allowed")
			return
		}
		next.ServeHTTP(w, r)
	}
}
