package fhir

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gotidy/fhir-client/models"
)

type ResponseError struct {
	Status      int
	ContentType string
	Message     string
}

func NewResponseError(resp *http.Response, message string) *ResponseError {
	return &ResponseError{
		Status:      resp.StatusCode,
		ContentType: resp.Header.Get("Content-Type"),
		Message:     message,
	}
}

func (e ResponseError) Error() string {
	return e.Message
}

type FhirError struct {
	Status           int
	OperationOutcome *models.OperationOutcome
	Message          string
}

func NewFhirError(resp *http.Response, operationOutcome *models.OperationOutcome, msg string) FhirError {
	return FhirError{
		Status:           resp.StatusCode,
		OperationOutcome: operationOutcome,
		Message:          msg,
	}
}

func (e FhirError) Error() string {
	switch {
	case e.OperationOutcome != nil && e.OperationOutcome.Text.Div != "":
		return e.OperationOutcome.Text.Div
	case e.Message != "":
		return e.Message
	default:
		return fmt.Sprintf("request error: %d (%s)", e.Status, http.StatusText(e.Status))
	}
}

func AsFhirError(err error) (FhirError, bool) {
	var e FhirError
	return e, errors.As(err, &e)
}

type UnmarshalError struct {
	Message  string
	Resource ResourceType
	Err      error
	Data     []byte
}

func NewUnmarshalError(msg string, resource ResourceType, data []byte, err error) UnmarshalError {
	return UnmarshalError{
		Message:  msg,
		Resource: resource,
		Err:      err,
		Data:     data,
	}
}

func (e UnmarshalError) Error() string {
	return fmt.Sprintf("%s (%s): %s", e.Message, e.Resource, e.Err)
}

func (e UnmarshalError) Unwrap() error {
	return e.Err
}

func AsUnmarshalError(err error) (UnmarshalError, bool) {
	var e UnmarshalError
	return e, errors.As(err, &e)
}

type NotFoundError struct {
	ID       string
	Resource string
}

func NewNotFoundError(id, resource string) NotFoundError {
	return NotFoundError{
		ID:       id,
		Resource: resource,
	}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("resource \"%s\" with ID \"%s\"", e.Resource, e.ID)
}

func IsNotFound(err error) bool {
	var e NotFoundError
	return errors.As(err, &e)
}

func NewTooManyEntitiesError(id, resource string, count int) error {
	return fmt.Errorf("too more \"%s\" entities for ID \"%s\" in the response: %d", resource, id, count)
}
