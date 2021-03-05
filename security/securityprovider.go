// Package securityprovider contains some default securityprovider
// implementations, which can be used as a RequestEditorFn of a
// client.
package security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrAPIKeyInvalidIn indicates a usage of an invalid In.
	// Should be cookie, header or query.
	ErrAPIKeyInvalidIn = errors.New("invalid 'in' specified for apiKey")
)

// BasicAuth provides a SecurityProvider, which can solve
// the BasicAuth challenge for api-calls.
func BasicAuth(username, password string) func(ctx context.Context, req *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}
}

// BearerToken provides a SecurityProvider, which can solve
// the Bearer Auth challende for api-calls.
func BearerToken(token string) func(ctx context.Context, req *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}
}

// ApiKey will attach a generic apiKey for a given name
// either to a cookie, header or as a query parameter.
func APIKey(in, name, apiKey string) (func(ctx context.Context, req *http.Request) error, error) {
	interceptors := map[string]func(ctx context.Context, req *http.Request) error{
		"cookie": func(ctx context.Context, req *http.Request) error {
			req.AddCookie(&http.Cookie{Name: name, Value: apiKey})
			return nil
		},
		"header": func(ctx context.Context, req *http.Request) error {
			req.Header.Add(name, apiKey)
			return nil
		},
		"query": func(ctx context.Context, req *http.Request) error {
			query := req.URL.Query()
			query.Add(name, apiKey)
			req.URL.RawQuery = query.Encode()
			return nil
		},
	}

	interceptor, ok := interceptors[in]
	if !ok {
		return nil, ErrAPIKeyInvalidIn
	}

	return interceptor, nil
}
