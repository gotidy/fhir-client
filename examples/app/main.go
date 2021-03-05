package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gotidy/fhir-client"
	"github.com/gotidy/fhir-client/security"
)

func main() {
	logLevel := os.Getenv("APP_LOGGING_LEVEL")
	clientID := os.Getenv("APP_CLIENT_ID")
	clientSecret := os.Getenv("APP_CLIENT_SECRET")
	server := os.Getenv("APP_SERVER")

	log := NewLogger(logLevel)

	client, err := fhir.New(
		server,
		fhir.WithRequestEditorFn(security.BasicAuth(clientID, clientSecret)),
		fhir.WithHTTPClient(Client(log, &http.Client{
			Timeout: time.Second * 10,
		})),
	)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	app := NewApplication(log, client)
	if err := app.Run(context.Background()); err != nil {
		log.Fatal().Err(err).Send()
	}
}
