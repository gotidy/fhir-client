package fhir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gotidy/fhir-client/models"
	"github.com/tidwall/gjson"
)

type ResourceType string

func (r ResourceType) String() string {
	return string(r)
}

type ResourceData []byte

func (r ResourceData) String() string {
	return string(r)
}

func (r ResourceData) UnmarshalTo(v interface{}) error {
	return json.Unmarshal(r, v)
}

type Parameters interface {
	Encode() string
}

type FhirResponse struct {
	StatusCode       int
	Body             []byte
	Bundle           *models.Bundle
	OperationOutcome *models.OperationOutcome
	ResourceType     ResourceType
}

func NewFhirResponse(resp *http.Response) (*FhirResponse, error) {
	fresp := &FhirResponse{
		StatusCode: resp.StatusCode,
	}

	var err error
	fresp.Body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fresp, err
	}

	if contentType := resp.Header.Get("Content-Type"); !strings.Contains(contentType, "json") {
		return fresp, NewResponseError(resp, fmt.Sprintf("Content-Type is \"%s\" but expected \"application/json\"", contentType))
	}

	var message string

	fresp.ResourceType = GetDataResourceType(fresp.Body)
	switch fresp.ResourceType {
	case "Bundle":
		var dest models.Bundle
		if err := json.Unmarshal(fresp.Body, &dest); err != nil {
			return fresp, NewUnmarshalError("response parsing", fresp.ResourceType, fresp.Body, err)
		}
		fresp.Bundle = &dest
	case "OperationOutcome":
		var dest models.OperationOutcome
		if err := json.Unmarshal(fresp.Body, &dest); err != nil {
			return fresp, NewUnmarshalError("response parsing", fresp.ResourceType, fresp.Body, err)
		}
		fresp.OperationOutcome = &dest
	case "":
		message = gjson.GetBytes(fresp.Body, "message").String()
	}

	switch {
	case resp.StatusCode >= 200 && resp.StatusCode < 300:
		return fresp, nil
	case resp.StatusCode >= 400 && resp.StatusCode < 500:
		return fresp, NewFhirError(resp, fresp.OperationOutcome, message)
	default:
		return fresp, NewResponseError(resp, fmt.Sprintf("unexpected respose status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode)))
	}
}

func (r *FhirResponse) MustBundle() error {
	if r.Bundle == nil {
		return fmt.Errorf("expected Bundle data in the response, but have: %s", r.ResourceType)
	}
	return nil
}

func (r *FhirResponse) MustResourceData(resource ResourceType) error {
	if r.Body == nil && r.ResourceType != resource {
		return fmt.Errorf("expected resource data in the response")
	}
	return nil
}

func (r *FhirResponse) MustOperationOutcome() error {
	if r.Bundle == nil {
		return fmt.Errorf("expected OperationOutcome data in the response, but have: %s", r.ResourceType)
	}
	return nil
}
