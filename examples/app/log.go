package main

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ClientLogger struct {
	log    zerolog.Logger
	client HTTPClient
}

func Client(log zerolog.Logger, client HTTPClient) *ClientLogger {
	return &ClientLogger{
		log:    log,
		client: client,
	}
}

func (c ClientLogger) Do(r *http.Request) (*http.Response, error) {
	start := time.Now()
	res, err := c.client.Do(r)
	latency := time.Since(start)

	var event *zerolog.Event
	if err != nil {
		event = c.log.Error().Err(err)
	} else {
		event = c.log.Debug()
	}
	event = event.Str("method", r.Method).Str("url", r.URL.String()).Dur("latency", latency)

	if res != nil {
		event.Int("status", res.StatusCode).
			Int64("size", res.ContentLength)
	} else {
		event.Int("status", -1).
			Int64("size", -1)
	}

	event.Msg("sending a request")

	return res, err
}

func NewLogger(level string) zerolog.Logger {
	zl := zerolog.InfoLevel
	switch level {
	case "debug":
		zl = zerolog.DebugLevel
	case "info":
		zl = zerolog.InfoLevel
	case "warn":
		zl = zerolog.WarnLevel
	case "error":
		zl = zerolog.ErrorLevel
	case "fatal":
		zl = zerolog.FatalLevel
	case "panic":
		zl = zerolog.PanicLevel
	case "trace":
		zl = zerolog.TraceLevel
	}

	return zerolog.New(os.Stderr).Level(zl).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
}
