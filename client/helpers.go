package fhir

import (
	"strings"

	"github.com/gotidy/fhir-client/models"
	"github.com/tidwall/gjson"
)

func Path(parts ...string) string {
	return strings.Join(parts, "/")
}

func GetDataResourceType(data []byte) ResourceType {
	return ResourceType(gjson.GetBytes(data, "resourceType").String())
}

func ExpectedBundle(resp *FhirResponse, err error) (*models.Bundle, error) {
	if err != nil {
		return nil, err
	}
	if err := resp.MustBundle(); err != nil {
		return nil, err
	}
	return resp.Bundle, nil
}

func ExpectedOperationOutcome(resp *FhirResponse, err error) (*models.OperationOutcome, error) {
	if err != nil {
		return nil, err
	}
	if err := resp.MustOperationOutcome(); err != nil {
		return nil, err
	}
	return resp.OperationOutcome, nil
}

func ExpectedResource(resource ResourceType, resp *FhirResponse, err error) (ResourceData, error) {
	if err != nil {
		return nil, err
	}
	if err := resp.MustResourceData(resource); err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func EnumBundleResources(bundle *models.Bundle, resource ResourceType, f func(resource ResourceData) error) error {
	if bundle == nil {
		return nil
	}

	for _, entry := range bundle.Entry {
		if len(entry.Resource) != 0 && GetDataResourceType(entry.Resource) == resource {
			if err := f(ResourceData(entry.Resource)); err != nil {
				return err
			}
		}
	}

	return nil
}

func checkSingular(id string, resource ResourceType, count int) error {
	switch {
	case count == 0:
		return NewNotFoundError(id, "Account")
	case count > 1:
		return NewTooManyEntitiesError(id, "Account", count)
	}
	return nil
}

func StrPtrToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
