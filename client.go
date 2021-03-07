package fhir

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function.
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HTTPRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HTTPRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction.
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults.
func New(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HTTPRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	req = req.WithContext(ctx)
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) DoRequest(ctx context.Context, req *http.Request) (*FhirResponse, error) {
	if err := c.applyEditors(ctx, req, nil); err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	return NewFhirResponse(resp)
}

func BodyReader(body interface{}) (reader io.Reader, err error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return bodyReader, nil
}

func (c *Client) RequestWithBodyReader(ctx context.Context, method string, path string, params Parameters, body io.Reader) (*FhirResponse, error) {
	queryURL, err := url.Parse(c.Server)
	if err != nil {
		return nil, err
	}

	queryURL, err = queryURL.Parse(path)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryURL.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return c.DoRequest(ctx, req)
}

func (c *Client) RequestWithBody(ctx context.Context, method string, path string, params Parameters, body interface{}) (*FhirResponse, error) {
	bodyReader, err := BodyReader(body)
	if err != nil {
		return nil, err
	}
	return c.RequestWithBodyReader(ctx, method, path, params, bodyReader)
}

func (c *Client) Request(ctx context.Context, method string, path string, params Parameters) (*FhirResponse, error) {
	return c.RequestWithBodyReader(ctx, method, path, params, nil)
}

func (c *Client) Get(ctx context.Context, resource ResourceType, params Parameters) (*FhirResponse, error) {
	return c.Request(ctx, http.MethodGet, string(resource), params)
}

func (c *Client) GetByID(ctx context.Context, resource ResourceType, id string, params Parameters) (*FhirResponse, error) {
	return c.Request(ctx, http.MethodGet, Path(string(resource), id), params)
}

func (c *Client) Create(ctx context.Context, resource ResourceType, params Parameters, body interface{}) (*FhirResponse, error) {
	return c.RequestWithBody(ctx, http.MethodPost, string(resource), params, body)
}

func (c *Client) Update(ctx context.Context, resource ResourceType, params Parameters, body interface{}) (*FhirResponse, error) {
	return c.RequestWithBody(ctx, http.MethodPut, string(resource), params, body)
}

func (c *Client) UpdateByID(ctx context.Context, resource ResourceType, id string, params Parameters, body interface{}) (*FhirResponse, error) {
	return c.RequestWithBody(ctx, http.MethodPut, Path(string(resource), id), params, body)
}

func (c *Client) Patch(ctx context.Context, resource ResourceType, params Parameters, body interface{}) (*FhirResponse, error) {
	return c.RequestWithBody(ctx, http.MethodPatch, string(resource), params, body)
}

func (c *Client) PatchByID(ctx context.Context, resource ResourceType, id string, params Parameters, body interface{}) (*FhirResponse, error) {
	return c.RequestWithBody(ctx, http.MethodPatch, Path(string(resource), id), params, body)
}

func (c *Client) Delete(ctx context.Context, resource ResourceType, params Parameters) (*FhirResponse, error) {
	return c.Request(ctx, http.MethodDelete, string(resource), params)
}

func (c *Client) DeleteByID(ctx context.Context, resource ResourceType, id string, params Parameters) (*FhirResponse, error) {
	return c.Request(ctx, http.MethodDelete, Path(string(resource), id), params)
}
