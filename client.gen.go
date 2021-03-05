package fhir

import (
	"context"
	"encoding/json"

	"github.com/gotidy/fhir-client/models"
)

// ---------------------------------------------------------------------------------------------------------------------------
// Account
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToAccounts(bundle *models.Bundle) ([]*models.Account, error) {
	var entities []*models.Account
	err := EnumBundleResources(bundle, "Account", func(resource ResourceData) error {
		var entity models.Account
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToAccounts(resp *FhirResponse) ([]*models.Account, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToAccounts(resp.Bundle)
	case "Account":
		var entity models.Account
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Account{&entity}, nil
	}
	return nil, nil
}

func fhirRespToAccount(id string, resp *FhirResponse) (*models.Account, error) {
	entities, err := fhirRespToAccounts(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Account", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Account.
func (c *Client) GetAccount(ctx context.Context, params Parameters) ([]*models.Account, error) {
	resp, err := c.Get(ctx, "Account", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAccounts(resp)
}

// Get Account by ID.
func (c *Client) GetAccountByID(ctx context.Context, id string, params Parameters) (*models.Account, error) {
	resp, err := c.GetByID(ctx, "Account", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAccount(id, resp)
}

func (c *Client) CreateAccount(ctx context.Context, resource string, params Parameters, entity *models.Account) (*models.Account, error) {
	resp, err := c.Create(ctx, "Account", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAccount(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAccount(ctx context.Context, resource string, params Parameters, entity *models.Account) (*models.Account, error) {
	resp, err := c.Update(ctx, "Account", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAccount(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAccountByID(ctx context.Context, resource, id string, params Parameters, entity *models.Account) (*models.Account, error) {
	resp, err := c.UpdateByID(ctx, "Account", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAccount(id, resp)
}

func (c *Client) PatchAccount(ctx context.Context, resource string, params Parameters, entity *models.Account) ([]*models.Account, error) {
	resp, err := c.Patch(ctx, "Account", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAccounts(resp)
}

func (c *Client) PatchAccountByID(ctx context.Context, resource, id string, params Parameters, entity *models.Account) (*models.Account, error) {
	resp, err := c.PatchByID(ctx, "Account", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAccount(id, resp)
}

func (c *Client) DeleteAccount(ctx context.Context, resource string, params Parameters) ([]*models.Account, error) {
	resp, err := c.Delete(ctx, "Account", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAccounts(resp)
}

func (c *Client) DeleteAccountByID(ctx context.Context, resource, id string, params Parameters) (*models.Account, error) {
	resp, err := c.DeleteByID(ctx, "Account", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAccount(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ActivityDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToActivityDefinitions(bundle *models.Bundle) ([]*models.ActivityDefinition, error) {
	var entities []*models.ActivityDefinition
	err := EnumBundleResources(bundle, "ActivityDefinition", func(resource ResourceData) error {
		var entity models.ActivityDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToActivityDefinitions(resp *FhirResponse) ([]*models.ActivityDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToActivityDefinitions(resp.Bundle)
	case "ActivityDefinition":
		var entity models.ActivityDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ActivityDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToActivityDefinition(id string, resp *FhirResponse) (*models.ActivityDefinition, error) {
	entities, err := fhirRespToActivityDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ActivityDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ActivityDefinition.
func (c *Client) GetActivityDefinition(ctx context.Context, params Parameters) ([]*models.ActivityDefinition, error) {
	resp, err := c.Get(ctx, "ActivityDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToActivityDefinitions(resp)
}

// Get ActivityDefinition by ID.
func (c *Client) GetActivityDefinitionByID(ctx context.Context, id string, params Parameters) (*models.ActivityDefinition, error) {
	resp, err := c.GetByID(ctx, "ActivityDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToActivityDefinition(id, resp)
}

func (c *Client) CreateActivityDefinition(ctx context.Context, resource string, params Parameters, entity *models.ActivityDefinition) (*models.ActivityDefinition, error) {
	resp, err := c.Create(ctx, "ActivityDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToActivityDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateActivityDefinition(ctx context.Context, resource string, params Parameters, entity *models.ActivityDefinition) (*models.ActivityDefinition, error) {
	resp, err := c.Update(ctx, "ActivityDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToActivityDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateActivityDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ActivityDefinition) (*models.ActivityDefinition, error) {
	resp, err := c.UpdateByID(ctx, "ActivityDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToActivityDefinition(id, resp)
}

func (c *Client) PatchActivityDefinition(ctx context.Context, resource string, params Parameters, entity *models.ActivityDefinition) ([]*models.ActivityDefinition, error) {
	resp, err := c.Patch(ctx, "ActivityDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToActivityDefinitions(resp)
}

func (c *Client) PatchActivityDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ActivityDefinition) (*models.ActivityDefinition, error) {
	resp, err := c.PatchByID(ctx, "ActivityDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToActivityDefinition(id, resp)
}

func (c *Client) DeleteActivityDefinition(ctx context.Context, resource string, params Parameters) ([]*models.ActivityDefinition, error) {
	resp, err := c.Delete(ctx, "ActivityDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToActivityDefinitions(resp)
}

func (c *Client) DeleteActivityDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.ActivityDefinition, error) {
	resp, err := c.DeleteByID(ctx, "ActivityDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToActivityDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// AdverseEvent
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToAdverseEvents(bundle *models.Bundle) ([]*models.AdverseEvent, error) {
	var entities []*models.AdverseEvent
	err := EnumBundleResources(bundle, "AdverseEvent", func(resource ResourceData) error {
		var entity models.AdverseEvent
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToAdverseEvents(resp *FhirResponse) ([]*models.AdverseEvent, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToAdverseEvents(resp.Bundle)
	case "AdverseEvent":
		var entity models.AdverseEvent
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.AdverseEvent{&entity}, nil
	}
	return nil, nil
}

func fhirRespToAdverseEvent(id string, resp *FhirResponse) (*models.AdverseEvent, error) {
	entities, err := fhirRespToAdverseEvents(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "AdverseEvent", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get AdverseEvent.
func (c *Client) GetAdverseEvent(ctx context.Context, params Parameters) ([]*models.AdverseEvent, error) {
	resp, err := c.Get(ctx, "AdverseEvent", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAdverseEvents(resp)
}

// Get AdverseEvent by ID.
func (c *Client) GetAdverseEventByID(ctx context.Context, id string, params Parameters) (*models.AdverseEvent, error) {
	resp, err := c.GetByID(ctx, "AdverseEvent", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAdverseEvent(id, resp)
}

func (c *Client) CreateAdverseEvent(ctx context.Context, resource string, params Parameters, entity *models.AdverseEvent) (*models.AdverseEvent, error) {
	resp, err := c.Create(ctx, "AdverseEvent", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAdverseEvent(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAdverseEvent(ctx context.Context, resource string, params Parameters, entity *models.AdverseEvent) (*models.AdverseEvent, error) {
	resp, err := c.Update(ctx, "AdverseEvent", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAdverseEvent(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAdverseEventByID(ctx context.Context, resource, id string, params Parameters, entity *models.AdverseEvent) (*models.AdverseEvent, error) {
	resp, err := c.UpdateByID(ctx, "AdverseEvent", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAdverseEvent(id, resp)
}

func (c *Client) PatchAdverseEvent(ctx context.Context, resource string, params Parameters, entity *models.AdverseEvent) ([]*models.AdverseEvent, error) {
	resp, err := c.Patch(ctx, "AdverseEvent", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAdverseEvents(resp)
}

func (c *Client) PatchAdverseEventByID(ctx context.Context, resource, id string, params Parameters, entity *models.AdverseEvent) (*models.AdverseEvent, error) {
	resp, err := c.PatchByID(ctx, "AdverseEvent", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAdverseEvent(id, resp)
}

func (c *Client) DeleteAdverseEvent(ctx context.Context, resource string, params Parameters) ([]*models.AdverseEvent, error) {
	resp, err := c.Delete(ctx, "AdverseEvent", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAdverseEvents(resp)
}

func (c *Client) DeleteAdverseEventByID(ctx context.Context, resource, id string, params Parameters) (*models.AdverseEvent, error) {
	resp, err := c.DeleteByID(ctx, "AdverseEvent", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAdverseEvent(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// AllergyIntolerance
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToAllergyIntolerances(bundle *models.Bundle) ([]*models.AllergyIntolerance, error) {
	var entities []*models.AllergyIntolerance
	err := EnumBundleResources(bundle, "AllergyIntolerance", func(resource ResourceData) error {
		var entity models.AllergyIntolerance
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToAllergyIntolerances(resp *FhirResponse) ([]*models.AllergyIntolerance, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToAllergyIntolerances(resp.Bundle)
	case "AllergyIntolerance":
		var entity models.AllergyIntolerance
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.AllergyIntolerance{&entity}, nil
	}
	return nil, nil
}

func fhirRespToAllergyIntolerance(id string, resp *FhirResponse) (*models.AllergyIntolerance, error) {
	entities, err := fhirRespToAllergyIntolerances(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "AllergyIntolerance", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get AllergyIntolerance.
func (c *Client) GetAllergyIntolerance(ctx context.Context, params Parameters) ([]*models.AllergyIntolerance, error) {
	resp, err := c.Get(ctx, "AllergyIntolerance", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAllergyIntolerances(resp)
}

// Get AllergyIntolerance by ID.
func (c *Client) GetAllergyIntoleranceByID(ctx context.Context, id string, params Parameters) (*models.AllergyIntolerance, error) {
	resp, err := c.GetByID(ctx, "AllergyIntolerance", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAllergyIntolerance(id, resp)
}

func (c *Client) CreateAllergyIntolerance(ctx context.Context, resource string, params Parameters, entity *models.AllergyIntolerance) (*models.AllergyIntolerance, error) {
	resp, err := c.Create(ctx, "AllergyIntolerance", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAllergyIntolerance(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAllergyIntolerance(ctx context.Context, resource string, params Parameters, entity *models.AllergyIntolerance) (*models.AllergyIntolerance, error) {
	resp, err := c.Update(ctx, "AllergyIntolerance", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAllergyIntolerance(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAllergyIntoleranceByID(ctx context.Context, resource, id string, params Parameters, entity *models.AllergyIntolerance) (*models.AllergyIntolerance, error) {
	resp, err := c.UpdateByID(ctx, "AllergyIntolerance", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAllergyIntolerance(id, resp)
}

func (c *Client) PatchAllergyIntolerance(ctx context.Context, resource string, params Parameters, entity *models.AllergyIntolerance) ([]*models.AllergyIntolerance, error) {
	resp, err := c.Patch(ctx, "AllergyIntolerance", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAllergyIntolerances(resp)
}

func (c *Client) PatchAllergyIntoleranceByID(ctx context.Context, resource, id string, params Parameters, entity *models.AllergyIntolerance) (*models.AllergyIntolerance, error) {
	resp, err := c.PatchByID(ctx, "AllergyIntolerance", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAllergyIntolerance(id, resp)
}

func (c *Client) DeleteAllergyIntolerance(ctx context.Context, resource string, params Parameters) ([]*models.AllergyIntolerance, error) {
	resp, err := c.Delete(ctx, "AllergyIntolerance", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAllergyIntolerances(resp)
}

func (c *Client) DeleteAllergyIntoleranceByID(ctx context.Context, resource, id string, params Parameters) (*models.AllergyIntolerance, error) {
	resp, err := c.DeleteByID(ctx, "AllergyIntolerance", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAllergyIntolerance(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Appointment
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToAppointments(bundle *models.Bundle) ([]*models.Appointment, error) {
	var entities []*models.Appointment
	err := EnumBundleResources(bundle, "Appointment", func(resource ResourceData) error {
		var entity models.Appointment
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToAppointments(resp *FhirResponse) ([]*models.Appointment, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToAppointments(resp.Bundle)
	case "Appointment":
		var entity models.Appointment
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Appointment{&entity}, nil
	}
	return nil, nil
}

func fhirRespToAppointment(id string, resp *FhirResponse) (*models.Appointment, error) {
	entities, err := fhirRespToAppointments(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Appointment", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Appointment.
func (c *Client) GetAppointment(ctx context.Context, params Parameters) ([]*models.Appointment, error) {
	resp, err := c.Get(ctx, "Appointment", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointments(resp)
}

// Get Appointment by ID.
func (c *Client) GetAppointmentByID(ctx context.Context, id string, params Parameters) (*models.Appointment, error) {
	resp, err := c.GetByID(ctx, "Appointment", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointment(id, resp)
}

func (c *Client) CreateAppointment(ctx context.Context, resource string, params Parameters, entity *models.Appointment) (*models.Appointment, error) {
	resp, err := c.Create(ctx, "Appointment", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointment(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAppointment(ctx context.Context, resource string, params Parameters, entity *models.Appointment) (*models.Appointment, error) {
	resp, err := c.Update(ctx, "Appointment", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointment(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAppointmentByID(ctx context.Context, resource, id string, params Parameters, entity *models.Appointment) (*models.Appointment, error) {
	resp, err := c.UpdateByID(ctx, "Appointment", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointment(id, resp)
}

func (c *Client) PatchAppointment(ctx context.Context, resource string, params Parameters, entity *models.Appointment) ([]*models.Appointment, error) {
	resp, err := c.Patch(ctx, "Appointment", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointments(resp)
}

func (c *Client) PatchAppointmentByID(ctx context.Context, resource, id string, params Parameters, entity *models.Appointment) (*models.Appointment, error) {
	resp, err := c.PatchByID(ctx, "Appointment", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointment(id, resp)
}

func (c *Client) DeleteAppointment(ctx context.Context, resource string, params Parameters) ([]*models.Appointment, error) {
	resp, err := c.Delete(ctx, "Appointment", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointments(resp)
}

func (c *Client) DeleteAppointmentByID(ctx context.Context, resource, id string, params Parameters) (*models.Appointment, error) {
	resp, err := c.DeleteByID(ctx, "Appointment", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointment(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// AppointmentResponse
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToAppointmentResponses(bundle *models.Bundle) ([]*models.AppointmentResponse, error) {
	var entities []*models.AppointmentResponse
	err := EnumBundleResources(bundle, "AppointmentResponse", func(resource ResourceData) error {
		var entity models.AppointmentResponse
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToAppointmentResponses(resp *FhirResponse) ([]*models.AppointmentResponse, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToAppointmentResponses(resp.Bundle)
	case "AppointmentResponse":
		var entity models.AppointmentResponse
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.AppointmentResponse{&entity}, nil
	}
	return nil, nil
}

func fhirRespToAppointmentResponse(id string, resp *FhirResponse) (*models.AppointmentResponse, error) {
	entities, err := fhirRespToAppointmentResponses(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "AppointmentResponse", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get AppointmentResponse.
func (c *Client) GetAppointmentResponse(ctx context.Context, params Parameters) ([]*models.AppointmentResponse, error) {
	resp, err := c.Get(ctx, "AppointmentResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointmentResponses(resp)
}

// Get AppointmentResponse by ID.
func (c *Client) GetAppointmentResponseByID(ctx context.Context, id string, params Parameters) (*models.AppointmentResponse, error) {
	resp, err := c.GetByID(ctx, "AppointmentResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointmentResponse(id, resp)
}

func (c *Client) CreateAppointmentResponse(ctx context.Context, resource string, params Parameters, entity *models.AppointmentResponse) (*models.AppointmentResponse, error) {
	resp, err := c.Create(ctx, "AppointmentResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointmentResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAppointmentResponse(ctx context.Context, resource string, params Parameters, entity *models.AppointmentResponse) (*models.AppointmentResponse, error) {
	resp, err := c.Update(ctx, "AppointmentResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointmentResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAppointmentResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.AppointmentResponse) (*models.AppointmentResponse, error) {
	resp, err := c.UpdateByID(ctx, "AppointmentResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointmentResponse(id, resp)
}

func (c *Client) PatchAppointmentResponse(ctx context.Context, resource string, params Parameters, entity *models.AppointmentResponse) ([]*models.AppointmentResponse, error) {
	resp, err := c.Patch(ctx, "AppointmentResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointmentResponses(resp)
}

func (c *Client) PatchAppointmentResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.AppointmentResponse) (*models.AppointmentResponse, error) {
	resp, err := c.PatchByID(ctx, "AppointmentResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointmentResponse(id, resp)
}

func (c *Client) DeleteAppointmentResponse(ctx context.Context, resource string, params Parameters) ([]*models.AppointmentResponse, error) {
	resp, err := c.Delete(ctx, "AppointmentResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointmentResponses(resp)
}

func (c *Client) DeleteAppointmentResponseByID(ctx context.Context, resource, id string, params Parameters) (*models.AppointmentResponse, error) {
	resp, err := c.DeleteByID(ctx, "AppointmentResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAppointmentResponse(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// AuditEvent
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToAuditEvents(bundle *models.Bundle) ([]*models.AuditEvent, error) {
	var entities []*models.AuditEvent
	err := EnumBundleResources(bundle, "AuditEvent", func(resource ResourceData) error {
		var entity models.AuditEvent
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToAuditEvents(resp *FhirResponse) ([]*models.AuditEvent, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToAuditEvents(resp.Bundle)
	case "AuditEvent":
		var entity models.AuditEvent
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.AuditEvent{&entity}, nil
	}
	return nil, nil
}

func fhirRespToAuditEvent(id string, resp *FhirResponse) (*models.AuditEvent, error) {
	entities, err := fhirRespToAuditEvents(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "AuditEvent", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get AuditEvent.
func (c *Client) GetAuditEvent(ctx context.Context, params Parameters) ([]*models.AuditEvent, error) {
	resp, err := c.Get(ctx, "AuditEvent", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAuditEvents(resp)
}

// Get AuditEvent by ID.
func (c *Client) GetAuditEventByID(ctx context.Context, id string, params Parameters) (*models.AuditEvent, error) {
	resp, err := c.GetByID(ctx, "AuditEvent", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAuditEvent(id, resp)
}

func (c *Client) CreateAuditEvent(ctx context.Context, resource string, params Parameters, entity *models.AuditEvent) (*models.AuditEvent, error) {
	resp, err := c.Create(ctx, "AuditEvent", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAuditEvent(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAuditEvent(ctx context.Context, resource string, params Parameters, entity *models.AuditEvent) (*models.AuditEvent, error) {
	resp, err := c.Update(ctx, "AuditEvent", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAuditEvent(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateAuditEventByID(ctx context.Context, resource, id string, params Parameters, entity *models.AuditEvent) (*models.AuditEvent, error) {
	resp, err := c.UpdateByID(ctx, "AuditEvent", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAuditEvent(id, resp)
}

func (c *Client) PatchAuditEvent(ctx context.Context, resource string, params Parameters, entity *models.AuditEvent) ([]*models.AuditEvent, error) {
	resp, err := c.Patch(ctx, "AuditEvent", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAuditEvents(resp)
}

func (c *Client) PatchAuditEventByID(ctx context.Context, resource, id string, params Parameters, entity *models.AuditEvent) (*models.AuditEvent, error) {
	resp, err := c.PatchByID(ctx, "AuditEvent", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToAuditEvent(id, resp)
}

func (c *Client) DeleteAuditEvent(ctx context.Context, resource string, params Parameters) ([]*models.AuditEvent, error) {
	resp, err := c.Delete(ctx, "AuditEvent", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAuditEvents(resp)
}

func (c *Client) DeleteAuditEventByID(ctx context.Context, resource, id string, params Parameters) (*models.AuditEvent, error) {
	resp, err := c.DeleteByID(ctx, "AuditEvent", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToAuditEvent(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Basic
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToBasics(bundle *models.Bundle) ([]*models.Basic, error) {
	var entities []*models.Basic
	err := EnumBundleResources(bundle, "Basic", func(resource ResourceData) error {
		var entity models.Basic
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToBasics(resp *FhirResponse) ([]*models.Basic, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToBasics(resp.Bundle)
	case "Basic":
		var entity models.Basic
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Basic{&entity}, nil
	}
	return nil, nil
}

func fhirRespToBasic(id string, resp *FhirResponse) (*models.Basic, error) {
	entities, err := fhirRespToBasics(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Basic", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Basic.
func (c *Client) GetBasic(ctx context.Context, params Parameters) ([]*models.Basic, error) {
	resp, err := c.Get(ctx, "Basic", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBasics(resp)
}

// Get Basic by ID.
func (c *Client) GetBasicByID(ctx context.Context, id string, params Parameters) (*models.Basic, error) {
	resp, err := c.GetByID(ctx, "Basic", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBasic(id, resp)
}

func (c *Client) CreateBasic(ctx context.Context, resource string, params Parameters, entity *models.Basic) (*models.Basic, error) {
	resp, err := c.Create(ctx, "Basic", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBasic(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateBasic(ctx context.Context, resource string, params Parameters, entity *models.Basic) (*models.Basic, error) {
	resp, err := c.Update(ctx, "Basic", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBasic(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateBasicByID(ctx context.Context, resource, id string, params Parameters, entity *models.Basic) (*models.Basic, error) {
	resp, err := c.UpdateByID(ctx, "Basic", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBasic(id, resp)
}

func (c *Client) PatchBasic(ctx context.Context, resource string, params Parameters, entity *models.Basic) ([]*models.Basic, error) {
	resp, err := c.Patch(ctx, "Basic", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBasics(resp)
}

func (c *Client) PatchBasicByID(ctx context.Context, resource, id string, params Parameters, entity *models.Basic) (*models.Basic, error) {
	resp, err := c.PatchByID(ctx, "Basic", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBasic(id, resp)
}

func (c *Client) DeleteBasic(ctx context.Context, resource string, params Parameters) ([]*models.Basic, error) {
	resp, err := c.Delete(ctx, "Basic", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBasics(resp)
}

func (c *Client) DeleteBasicByID(ctx context.Context, resource, id string, params Parameters) (*models.Basic, error) {
	resp, err := c.DeleteByID(ctx, "Basic", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBasic(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Binary
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToBinarys(bundle *models.Bundle) ([]*models.Binary, error) {
	var entities []*models.Binary
	err := EnumBundleResources(bundle, "Binary", func(resource ResourceData) error {
		var entity models.Binary
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToBinarys(resp *FhirResponse) ([]*models.Binary, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToBinarys(resp.Bundle)
	case "Binary":
		var entity models.Binary
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Binary{&entity}, nil
	}
	return nil, nil
}

func fhirRespToBinary(id string, resp *FhirResponse) (*models.Binary, error) {
	entities, err := fhirRespToBinarys(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Binary", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Binary.
func (c *Client) GetBinary(ctx context.Context, params Parameters) ([]*models.Binary, error) {
	resp, err := c.Get(ctx, "Binary", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBinarys(resp)
}

// Get Binary by ID.
func (c *Client) GetBinaryByID(ctx context.Context, id string, params Parameters) (*models.Binary, error) {
	resp, err := c.GetByID(ctx, "Binary", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBinary(id, resp)
}

func (c *Client) CreateBinary(ctx context.Context, resource string, params Parameters, entity *models.Binary) (*models.Binary, error) {
	resp, err := c.Create(ctx, "Binary", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBinary(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateBinary(ctx context.Context, resource string, params Parameters, entity *models.Binary) (*models.Binary, error) {
	resp, err := c.Update(ctx, "Binary", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBinary(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateBinaryByID(ctx context.Context, resource, id string, params Parameters, entity *models.Binary) (*models.Binary, error) {
	resp, err := c.UpdateByID(ctx, "Binary", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBinary(id, resp)
}

func (c *Client) PatchBinary(ctx context.Context, resource string, params Parameters, entity *models.Binary) ([]*models.Binary, error) {
	resp, err := c.Patch(ctx, "Binary", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBinarys(resp)
}

func (c *Client) PatchBinaryByID(ctx context.Context, resource, id string, params Parameters, entity *models.Binary) (*models.Binary, error) {
	resp, err := c.PatchByID(ctx, "Binary", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBinary(id, resp)
}

func (c *Client) DeleteBinary(ctx context.Context, resource string, params Parameters) ([]*models.Binary, error) {
	resp, err := c.Delete(ctx, "Binary", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBinarys(resp)
}

func (c *Client) DeleteBinaryByID(ctx context.Context, resource, id string, params Parameters) (*models.Binary, error) {
	resp, err := c.DeleteByID(ctx, "Binary", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBinary(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// BiologicallyDerivedProduct
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToBiologicallyDerivedProducts(bundle *models.Bundle) ([]*models.BiologicallyDerivedProduct, error) {
	var entities []*models.BiologicallyDerivedProduct
	err := EnumBundleResources(bundle, "BiologicallyDerivedProduct", func(resource ResourceData) error {
		var entity models.BiologicallyDerivedProduct
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToBiologicallyDerivedProducts(resp *FhirResponse) ([]*models.BiologicallyDerivedProduct, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToBiologicallyDerivedProducts(resp.Bundle)
	case "BiologicallyDerivedProduct":
		var entity models.BiologicallyDerivedProduct
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.BiologicallyDerivedProduct{&entity}, nil
	}
	return nil, nil
}

func fhirRespToBiologicallyDerivedProduct(id string, resp *FhirResponse) (*models.BiologicallyDerivedProduct, error) {
	entities, err := fhirRespToBiologicallyDerivedProducts(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "BiologicallyDerivedProduct", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get BiologicallyDerivedProduct.
func (c *Client) GetBiologicallyDerivedProduct(ctx context.Context, params Parameters) ([]*models.BiologicallyDerivedProduct, error) {
	resp, err := c.Get(ctx, "BiologicallyDerivedProduct", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBiologicallyDerivedProducts(resp)
}

// Get BiologicallyDerivedProduct by ID.
func (c *Client) GetBiologicallyDerivedProductByID(ctx context.Context, id string, params Parameters) (*models.BiologicallyDerivedProduct, error) {
	resp, err := c.GetByID(ctx, "BiologicallyDerivedProduct", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBiologicallyDerivedProduct(id, resp)
}

func (c *Client) CreateBiologicallyDerivedProduct(ctx context.Context, resource string, params Parameters, entity *models.BiologicallyDerivedProduct) (*models.BiologicallyDerivedProduct, error) {
	resp, err := c.Create(ctx, "BiologicallyDerivedProduct", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBiologicallyDerivedProduct(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateBiologicallyDerivedProduct(ctx context.Context, resource string, params Parameters, entity *models.BiologicallyDerivedProduct) (*models.BiologicallyDerivedProduct, error) {
	resp, err := c.Update(ctx, "BiologicallyDerivedProduct", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBiologicallyDerivedProduct(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateBiologicallyDerivedProductByID(ctx context.Context, resource, id string, params Parameters, entity *models.BiologicallyDerivedProduct) (*models.BiologicallyDerivedProduct, error) {
	resp, err := c.UpdateByID(ctx, "BiologicallyDerivedProduct", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBiologicallyDerivedProduct(id, resp)
}

func (c *Client) PatchBiologicallyDerivedProduct(ctx context.Context, resource string, params Parameters, entity *models.BiologicallyDerivedProduct) ([]*models.BiologicallyDerivedProduct, error) {
	resp, err := c.Patch(ctx, "BiologicallyDerivedProduct", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBiologicallyDerivedProducts(resp)
}

func (c *Client) PatchBiologicallyDerivedProductByID(ctx context.Context, resource, id string, params Parameters, entity *models.BiologicallyDerivedProduct) (*models.BiologicallyDerivedProduct, error) {
	resp, err := c.PatchByID(ctx, "BiologicallyDerivedProduct", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBiologicallyDerivedProduct(id, resp)
}

func (c *Client) DeleteBiologicallyDerivedProduct(ctx context.Context, resource string, params Parameters) ([]*models.BiologicallyDerivedProduct, error) {
	resp, err := c.Delete(ctx, "BiologicallyDerivedProduct", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBiologicallyDerivedProducts(resp)
}

func (c *Client) DeleteBiologicallyDerivedProductByID(ctx context.Context, resource, id string, params Parameters) (*models.BiologicallyDerivedProduct, error) {
	resp, err := c.DeleteByID(ctx, "BiologicallyDerivedProduct", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBiologicallyDerivedProduct(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// BodyStructure
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToBodyStructures(bundle *models.Bundle) ([]*models.BodyStructure, error) {
	var entities []*models.BodyStructure
	err := EnumBundleResources(bundle, "BodyStructure", func(resource ResourceData) error {
		var entity models.BodyStructure
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToBodyStructures(resp *FhirResponse) ([]*models.BodyStructure, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToBodyStructures(resp.Bundle)
	case "BodyStructure":
		var entity models.BodyStructure
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.BodyStructure{&entity}, nil
	}
	return nil, nil
}

func fhirRespToBodyStructure(id string, resp *FhirResponse) (*models.BodyStructure, error) {
	entities, err := fhirRespToBodyStructures(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "BodyStructure", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get BodyStructure.
func (c *Client) GetBodyStructure(ctx context.Context, params Parameters) ([]*models.BodyStructure, error) {
	resp, err := c.Get(ctx, "BodyStructure", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBodyStructures(resp)
}

// Get BodyStructure by ID.
func (c *Client) GetBodyStructureByID(ctx context.Context, id string, params Parameters) (*models.BodyStructure, error) {
	resp, err := c.GetByID(ctx, "BodyStructure", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBodyStructure(id, resp)
}

func (c *Client) CreateBodyStructure(ctx context.Context, resource string, params Parameters, entity *models.BodyStructure) (*models.BodyStructure, error) {
	resp, err := c.Create(ctx, "BodyStructure", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBodyStructure(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateBodyStructure(ctx context.Context, resource string, params Parameters, entity *models.BodyStructure) (*models.BodyStructure, error) {
	resp, err := c.Update(ctx, "BodyStructure", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBodyStructure(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateBodyStructureByID(ctx context.Context, resource, id string, params Parameters, entity *models.BodyStructure) (*models.BodyStructure, error) {
	resp, err := c.UpdateByID(ctx, "BodyStructure", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBodyStructure(id, resp)
}

func (c *Client) PatchBodyStructure(ctx context.Context, resource string, params Parameters, entity *models.BodyStructure) ([]*models.BodyStructure, error) {
	resp, err := c.Patch(ctx, "BodyStructure", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBodyStructures(resp)
}

func (c *Client) PatchBodyStructureByID(ctx context.Context, resource, id string, params Parameters, entity *models.BodyStructure) (*models.BodyStructure, error) {
	resp, err := c.PatchByID(ctx, "BodyStructure", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToBodyStructure(id, resp)
}

func (c *Client) DeleteBodyStructure(ctx context.Context, resource string, params Parameters) ([]*models.BodyStructure, error) {
	resp, err := c.Delete(ctx, "BodyStructure", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBodyStructures(resp)
}

func (c *Client) DeleteBodyStructureByID(ctx context.Context, resource, id string, params Parameters) (*models.BodyStructure, error) {
	resp, err := c.DeleteByID(ctx, "BodyStructure", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToBodyStructure(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// CapabilityStatement
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCapabilityStatements(bundle *models.Bundle) ([]*models.CapabilityStatement, error) {
	var entities []*models.CapabilityStatement
	err := EnumBundleResources(bundle, "CapabilityStatement", func(resource ResourceData) error {
		var entity models.CapabilityStatement
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCapabilityStatements(resp *FhirResponse) ([]*models.CapabilityStatement, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCapabilityStatements(resp.Bundle)
	case "CapabilityStatement":
		var entity models.CapabilityStatement
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.CapabilityStatement{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCapabilityStatement(id string, resp *FhirResponse) (*models.CapabilityStatement, error) {
	entities, err := fhirRespToCapabilityStatements(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "CapabilityStatement", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get CapabilityStatement.
func (c *Client) GetCapabilityStatement(ctx context.Context, params Parameters) ([]*models.CapabilityStatement, error) {
	resp, err := c.Get(ctx, "CapabilityStatement", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCapabilityStatements(resp)
}

// Get CapabilityStatement by ID.
func (c *Client) GetCapabilityStatementByID(ctx context.Context, id string, params Parameters) (*models.CapabilityStatement, error) {
	resp, err := c.GetByID(ctx, "CapabilityStatement", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCapabilityStatement(id, resp)
}

func (c *Client) CreateCapabilityStatement(ctx context.Context, resource string, params Parameters, entity *models.CapabilityStatement) (*models.CapabilityStatement, error) {
	resp, err := c.Create(ctx, "CapabilityStatement", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCapabilityStatement(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCapabilityStatement(ctx context.Context, resource string, params Parameters, entity *models.CapabilityStatement) (*models.CapabilityStatement, error) {
	resp, err := c.Update(ctx, "CapabilityStatement", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCapabilityStatement(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCapabilityStatementByID(ctx context.Context, resource, id string, params Parameters, entity *models.CapabilityStatement) (*models.CapabilityStatement, error) {
	resp, err := c.UpdateByID(ctx, "CapabilityStatement", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCapabilityStatement(id, resp)
}

func (c *Client) PatchCapabilityStatement(ctx context.Context, resource string, params Parameters, entity *models.CapabilityStatement) ([]*models.CapabilityStatement, error) {
	resp, err := c.Patch(ctx, "CapabilityStatement", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCapabilityStatements(resp)
}

func (c *Client) PatchCapabilityStatementByID(ctx context.Context, resource, id string, params Parameters, entity *models.CapabilityStatement) (*models.CapabilityStatement, error) {
	resp, err := c.PatchByID(ctx, "CapabilityStatement", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCapabilityStatement(id, resp)
}

func (c *Client) DeleteCapabilityStatement(ctx context.Context, resource string, params Parameters) ([]*models.CapabilityStatement, error) {
	resp, err := c.Delete(ctx, "CapabilityStatement", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCapabilityStatements(resp)
}

func (c *Client) DeleteCapabilityStatementByID(ctx context.Context, resource, id string, params Parameters) (*models.CapabilityStatement, error) {
	resp, err := c.DeleteByID(ctx, "CapabilityStatement", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCapabilityStatement(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// CarePlan
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCarePlans(bundle *models.Bundle) ([]*models.CarePlan, error) {
	var entities []*models.CarePlan
	err := EnumBundleResources(bundle, "CarePlan", func(resource ResourceData) error {
		var entity models.CarePlan
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCarePlans(resp *FhirResponse) ([]*models.CarePlan, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCarePlans(resp.Bundle)
	case "CarePlan":
		var entity models.CarePlan
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.CarePlan{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCarePlan(id string, resp *FhirResponse) (*models.CarePlan, error) {
	entities, err := fhirRespToCarePlans(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "CarePlan", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get CarePlan.
func (c *Client) GetCarePlan(ctx context.Context, params Parameters) ([]*models.CarePlan, error) {
	resp, err := c.Get(ctx, "CarePlan", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCarePlans(resp)
}

// Get CarePlan by ID.
func (c *Client) GetCarePlanByID(ctx context.Context, id string, params Parameters) (*models.CarePlan, error) {
	resp, err := c.GetByID(ctx, "CarePlan", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCarePlan(id, resp)
}

func (c *Client) CreateCarePlan(ctx context.Context, resource string, params Parameters, entity *models.CarePlan) (*models.CarePlan, error) {
	resp, err := c.Create(ctx, "CarePlan", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCarePlan(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCarePlan(ctx context.Context, resource string, params Parameters, entity *models.CarePlan) (*models.CarePlan, error) {
	resp, err := c.Update(ctx, "CarePlan", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCarePlan(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCarePlanByID(ctx context.Context, resource, id string, params Parameters, entity *models.CarePlan) (*models.CarePlan, error) {
	resp, err := c.UpdateByID(ctx, "CarePlan", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCarePlan(id, resp)
}

func (c *Client) PatchCarePlan(ctx context.Context, resource string, params Parameters, entity *models.CarePlan) ([]*models.CarePlan, error) {
	resp, err := c.Patch(ctx, "CarePlan", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCarePlans(resp)
}

func (c *Client) PatchCarePlanByID(ctx context.Context, resource, id string, params Parameters, entity *models.CarePlan) (*models.CarePlan, error) {
	resp, err := c.PatchByID(ctx, "CarePlan", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCarePlan(id, resp)
}

func (c *Client) DeleteCarePlan(ctx context.Context, resource string, params Parameters) ([]*models.CarePlan, error) {
	resp, err := c.Delete(ctx, "CarePlan", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCarePlans(resp)
}

func (c *Client) DeleteCarePlanByID(ctx context.Context, resource, id string, params Parameters) (*models.CarePlan, error) {
	resp, err := c.DeleteByID(ctx, "CarePlan", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCarePlan(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// CareTeam
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCareTeams(bundle *models.Bundle) ([]*models.CareTeam, error) {
	var entities []*models.CareTeam
	err := EnumBundleResources(bundle, "CareTeam", func(resource ResourceData) error {
		var entity models.CareTeam
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCareTeams(resp *FhirResponse) ([]*models.CareTeam, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCareTeams(resp.Bundle)
	case "CareTeam":
		var entity models.CareTeam
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.CareTeam{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCareTeam(id string, resp *FhirResponse) (*models.CareTeam, error) {
	entities, err := fhirRespToCareTeams(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "CareTeam", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get CareTeam.
func (c *Client) GetCareTeam(ctx context.Context, params Parameters) ([]*models.CareTeam, error) {
	resp, err := c.Get(ctx, "CareTeam", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCareTeams(resp)
}

// Get CareTeam by ID.
func (c *Client) GetCareTeamByID(ctx context.Context, id string, params Parameters) (*models.CareTeam, error) {
	resp, err := c.GetByID(ctx, "CareTeam", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCareTeam(id, resp)
}

func (c *Client) CreateCareTeam(ctx context.Context, resource string, params Parameters, entity *models.CareTeam) (*models.CareTeam, error) {
	resp, err := c.Create(ctx, "CareTeam", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCareTeam(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCareTeam(ctx context.Context, resource string, params Parameters, entity *models.CareTeam) (*models.CareTeam, error) {
	resp, err := c.Update(ctx, "CareTeam", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCareTeam(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCareTeamByID(ctx context.Context, resource, id string, params Parameters, entity *models.CareTeam) (*models.CareTeam, error) {
	resp, err := c.UpdateByID(ctx, "CareTeam", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCareTeam(id, resp)
}

func (c *Client) PatchCareTeam(ctx context.Context, resource string, params Parameters, entity *models.CareTeam) ([]*models.CareTeam, error) {
	resp, err := c.Patch(ctx, "CareTeam", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCareTeams(resp)
}

func (c *Client) PatchCareTeamByID(ctx context.Context, resource, id string, params Parameters, entity *models.CareTeam) (*models.CareTeam, error) {
	resp, err := c.PatchByID(ctx, "CareTeam", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCareTeam(id, resp)
}

func (c *Client) DeleteCareTeam(ctx context.Context, resource string, params Parameters) ([]*models.CareTeam, error) {
	resp, err := c.Delete(ctx, "CareTeam", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCareTeams(resp)
}

func (c *Client) DeleteCareTeamByID(ctx context.Context, resource, id string, params Parameters) (*models.CareTeam, error) {
	resp, err := c.DeleteByID(ctx, "CareTeam", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCareTeam(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// CatalogEntry
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCatalogEntrys(bundle *models.Bundle) ([]*models.CatalogEntry, error) {
	var entities []*models.CatalogEntry
	err := EnumBundleResources(bundle, "CatalogEntry", func(resource ResourceData) error {
		var entity models.CatalogEntry
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCatalogEntrys(resp *FhirResponse) ([]*models.CatalogEntry, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCatalogEntrys(resp.Bundle)
	case "CatalogEntry":
		var entity models.CatalogEntry
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.CatalogEntry{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCatalogEntry(id string, resp *FhirResponse) (*models.CatalogEntry, error) {
	entities, err := fhirRespToCatalogEntrys(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "CatalogEntry", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get CatalogEntry.
func (c *Client) GetCatalogEntry(ctx context.Context, params Parameters) ([]*models.CatalogEntry, error) {
	resp, err := c.Get(ctx, "CatalogEntry", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCatalogEntrys(resp)
}

// Get CatalogEntry by ID.
func (c *Client) GetCatalogEntryByID(ctx context.Context, id string, params Parameters) (*models.CatalogEntry, error) {
	resp, err := c.GetByID(ctx, "CatalogEntry", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCatalogEntry(id, resp)
}

func (c *Client) CreateCatalogEntry(ctx context.Context, resource string, params Parameters, entity *models.CatalogEntry) (*models.CatalogEntry, error) {
	resp, err := c.Create(ctx, "CatalogEntry", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCatalogEntry(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCatalogEntry(ctx context.Context, resource string, params Parameters, entity *models.CatalogEntry) (*models.CatalogEntry, error) {
	resp, err := c.Update(ctx, "CatalogEntry", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCatalogEntry(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCatalogEntryByID(ctx context.Context, resource, id string, params Parameters, entity *models.CatalogEntry) (*models.CatalogEntry, error) {
	resp, err := c.UpdateByID(ctx, "CatalogEntry", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCatalogEntry(id, resp)
}

func (c *Client) PatchCatalogEntry(ctx context.Context, resource string, params Parameters, entity *models.CatalogEntry) ([]*models.CatalogEntry, error) {
	resp, err := c.Patch(ctx, "CatalogEntry", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCatalogEntrys(resp)
}

func (c *Client) PatchCatalogEntryByID(ctx context.Context, resource, id string, params Parameters, entity *models.CatalogEntry) (*models.CatalogEntry, error) {
	resp, err := c.PatchByID(ctx, "CatalogEntry", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCatalogEntry(id, resp)
}

func (c *Client) DeleteCatalogEntry(ctx context.Context, resource string, params Parameters) ([]*models.CatalogEntry, error) {
	resp, err := c.Delete(ctx, "CatalogEntry", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCatalogEntrys(resp)
}

func (c *Client) DeleteCatalogEntryByID(ctx context.Context, resource, id string, params Parameters) (*models.CatalogEntry, error) {
	resp, err := c.DeleteByID(ctx, "CatalogEntry", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCatalogEntry(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ChargeItem
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToChargeItems(bundle *models.Bundle) ([]*models.ChargeItem, error) {
	var entities []*models.ChargeItem
	err := EnumBundleResources(bundle, "ChargeItem", func(resource ResourceData) error {
		var entity models.ChargeItem
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToChargeItems(resp *FhirResponse) ([]*models.ChargeItem, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToChargeItems(resp.Bundle)
	case "ChargeItem":
		var entity models.ChargeItem
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ChargeItem{&entity}, nil
	}
	return nil, nil
}

func fhirRespToChargeItem(id string, resp *FhirResponse) (*models.ChargeItem, error) {
	entities, err := fhirRespToChargeItems(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ChargeItem", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ChargeItem.
func (c *Client) GetChargeItem(ctx context.Context, params Parameters) ([]*models.ChargeItem, error) {
	resp, err := c.Get(ctx, "ChargeItem", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItems(resp)
}

// Get ChargeItem by ID.
func (c *Client) GetChargeItemByID(ctx context.Context, id string, params Parameters) (*models.ChargeItem, error) {
	resp, err := c.GetByID(ctx, "ChargeItem", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItem(id, resp)
}

func (c *Client) CreateChargeItem(ctx context.Context, resource string, params Parameters, entity *models.ChargeItem) (*models.ChargeItem, error) {
	resp, err := c.Create(ctx, "ChargeItem", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItem(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateChargeItem(ctx context.Context, resource string, params Parameters, entity *models.ChargeItem) (*models.ChargeItem, error) {
	resp, err := c.Update(ctx, "ChargeItem", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItem(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateChargeItemByID(ctx context.Context, resource, id string, params Parameters, entity *models.ChargeItem) (*models.ChargeItem, error) {
	resp, err := c.UpdateByID(ctx, "ChargeItem", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItem(id, resp)
}

func (c *Client) PatchChargeItem(ctx context.Context, resource string, params Parameters, entity *models.ChargeItem) ([]*models.ChargeItem, error) {
	resp, err := c.Patch(ctx, "ChargeItem", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItems(resp)
}

func (c *Client) PatchChargeItemByID(ctx context.Context, resource, id string, params Parameters, entity *models.ChargeItem) (*models.ChargeItem, error) {
	resp, err := c.PatchByID(ctx, "ChargeItem", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItem(id, resp)
}

func (c *Client) DeleteChargeItem(ctx context.Context, resource string, params Parameters) ([]*models.ChargeItem, error) {
	resp, err := c.Delete(ctx, "ChargeItem", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItems(resp)
}

func (c *Client) DeleteChargeItemByID(ctx context.Context, resource, id string, params Parameters) (*models.ChargeItem, error) {
	resp, err := c.DeleteByID(ctx, "ChargeItem", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItem(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ChargeItemDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToChargeItemDefinitions(bundle *models.Bundle) ([]*models.ChargeItemDefinition, error) {
	var entities []*models.ChargeItemDefinition
	err := EnumBundleResources(bundle, "ChargeItemDefinition", func(resource ResourceData) error {
		var entity models.ChargeItemDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToChargeItemDefinitions(resp *FhirResponse) ([]*models.ChargeItemDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToChargeItemDefinitions(resp.Bundle)
	case "ChargeItemDefinition":
		var entity models.ChargeItemDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ChargeItemDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToChargeItemDefinition(id string, resp *FhirResponse) (*models.ChargeItemDefinition, error) {
	entities, err := fhirRespToChargeItemDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ChargeItemDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ChargeItemDefinition.
func (c *Client) GetChargeItemDefinition(ctx context.Context, params Parameters) ([]*models.ChargeItemDefinition, error) {
	resp, err := c.Get(ctx, "ChargeItemDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItemDefinitions(resp)
}

// Get ChargeItemDefinition by ID.
func (c *Client) GetChargeItemDefinitionByID(ctx context.Context, id string, params Parameters) (*models.ChargeItemDefinition, error) {
	resp, err := c.GetByID(ctx, "ChargeItemDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItemDefinition(id, resp)
}

func (c *Client) CreateChargeItemDefinition(ctx context.Context, resource string, params Parameters, entity *models.ChargeItemDefinition) (*models.ChargeItemDefinition, error) {
	resp, err := c.Create(ctx, "ChargeItemDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItemDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateChargeItemDefinition(ctx context.Context, resource string, params Parameters, entity *models.ChargeItemDefinition) (*models.ChargeItemDefinition, error) {
	resp, err := c.Update(ctx, "ChargeItemDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItemDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateChargeItemDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ChargeItemDefinition) (*models.ChargeItemDefinition, error) {
	resp, err := c.UpdateByID(ctx, "ChargeItemDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItemDefinition(id, resp)
}

func (c *Client) PatchChargeItemDefinition(ctx context.Context, resource string, params Parameters, entity *models.ChargeItemDefinition) ([]*models.ChargeItemDefinition, error) {
	resp, err := c.Patch(ctx, "ChargeItemDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItemDefinitions(resp)
}

func (c *Client) PatchChargeItemDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ChargeItemDefinition) (*models.ChargeItemDefinition, error) {
	resp, err := c.PatchByID(ctx, "ChargeItemDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItemDefinition(id, resp)
}

func (c *Client) DeleteChargeItemDefinition(ctx context.Context, resource string, params Parameters) ([]*models.ChargeItemDefinition, error) {
	resp, err := c.Delete(ctx, "ChargeItemDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItemDefinitions(resp)
}

func (c *Client) DeleteChargeItemDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.ChargeItemDefinition, error) {
	resp, err := c.DeleteByID(ctx, "ChargeItemDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToChargeItemDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Claim
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToClaims(bundle *models.Bundle) ([]*models.Claim, error) {
	var entities []*models.Claim
	err := EnumBundleResources(bundle, "Claim", func(resource ResourceData) error {
		var entity models.Claim
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToClaims(resp *FhirResponse) ([]*models.Claim, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToClaims(resp.Bundle)
	case "Claim":
		var entity models.Claim
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Claim{&entity}, nil
	}
	return nil, nil
}

func fhirRespToClaim(id string, resp *FhirResponse) (*models.Claim, error) {
	entities, err := fhirRespToClaims(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Claim", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Claim.
func (c *Client) GetClaim(ctx context.Context, params Parameters) ([]*models.Claim, error) {
	resp, err := c.Get(ctx, "Claim", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaims(resp)
}

// Get Claim by ID.
func (c *Client) GetClaimByID(ctx context.Context, id string, params Parameters) (*models.Claim, error) {
	resp, err := c.GetByID(ctx, "Claim", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaim(id, resp)
}

func (c *Client) CreateClaim(ctx context.Context, resource string, params Parameters, entity *models.Claim) (*models.Claim, error) {
	resp, err := c.Create(ctx, "Claim", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaim(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateClaim(ctx context.Context, resource string, params Parameters, entity *models.Claim) (*models.Claim, error) {
	resp, err := c.Update(ctx, "Claim", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaim(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateClaimByID(ctx context.Context, resource, id string, params Parameters, entity *models.Claim) (*models.Claim, error) {
	resp, err := c.UpdateByID(ctx, "Claim", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaim(id, resp)
}

func (c *Client) PatchClaim(ctx context.Context, resource string, params Parameters, entity *models.Claim) ([]*models.Claim, error) {
	resp, err := c.Patch(ctx, "Claim", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaims(resp)
}

func (c *Client) PatchClaimByID(ctx context.Context, resource, id string, params Parameters, entity *models.Claim) (*models.Claim, error) {
	resp, err := c.PatchByID(ctx, "Claim", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaim(id, resp)
}

func (c *Client) DeleteClaim(ctx context.Context, resource string, params Parameters) ([]*models.Claim, error) {
	resp, err := c.Delete(ctx, "Claim", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaims(resp)
}

func (c *Client) DeleteClaimByID(ctx context.Context, resource, id string, params Parameters) (*models.Claim, error) {
	resp, err := c.DeleteByID(ctx, "Claim", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaim(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ClaimResponse
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToClaimResponses(bundle *models.Bundle) ([]*models.ClaimResponse, error) {
	var entities []*models.ClaimResponse
	err := EnumBundleResources(bundle, "ClaimResponse", func(resource ResourceData) error {
		var entity models.ClaimResponse
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToClaimResponses(resp *FhirResponse) ([]*models.ClaimResponse, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToClaimResponses(resp.Bundle)
	case "ClaimResponse":
		var entity models.ClaimResponse
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ClaimResponse{&entity}, nil
	}
	return nil, nil
}

func fhirRespToClaimResponse(id string, resp *FhirResponse) (*models.ClaimResponse, error) {
	entities, err := fhirRespToClaimResponses(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ClaimResponse", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ClaimResponse.
func (c *Client) GetClaimResponse(ctx context.Context, params Parameters) ([]*models.ClaimResponse, error) {
	resp, err := c.Get(ctx, "ClaimResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaimResponses(resp)
}

// Get ClaimResponse by ID.
func (c *Client) GetClaimResponseByID(ctx context.Context, id string, params Parameters) (*models.ClaimResponse, error) {
	resp, err := c.GetByID(ctx, "ClaimResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaimResponse(id, resp)
}

func (c *Client) CreateClaimResponse(ctx context.Context, resource string, params Parameters, entity *models.ClaimResponse) (*models.ClaimResponse, error) {
	resp, err := c.Create(ctx, "ClaimResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaimResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateClaimResponse(ctx context.Context, resource string, params Parameters, entity *models.ClaimResponse) (*models.ClaimResponse, error) {
	resp, err := c.Update(ctx, "ClaimResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaimResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateClaimResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.ClaimResponse) (*models.ClaimResponse, error) {
	resp, err := c.UpdateByID(ctx, "ClaimResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaimResponse(id, resp)
}

func (c *Client) PatchClaimResponse(ctx context.Context, resource string, params Parameters, entity *models.ClaimResponse) ([]*models.ClaimResponse, error) {
	resp, err := c.Patch(ctx, "ClaimResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaimResponses(resp)
}

func (c *Client) PatchClaimResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.ClaimResponse) (*models.ClaimResponse, error) {
	resp, err := c.PatchByID(ctx, "ClaimResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaimResponse(id, resp)
}

func (c *Client) DeleteClaimResponse(ctx context.Context, resource string, params Parameters) ([]*models.ClaimResponse, error) {
	resp, err := c.Delete(ctx, "ClaimResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaimResponses(resp)
}

func (c *Client) DeleteClaimResponseByID(ctx context.Context, resource, id string, params Parameters) (*models.ClaimResponse, error) {
	resp, err := c.DeleteByID(ctx, "ClaimResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClaimResponse(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ClinicalImpression
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToClinicalImpressions(bundle *models.Bundle) ([]*models.ClinicalImpression, error) {
	var entities []*models.ClinicalImpression
	err := EnumBundleResources(bundle, "ClinicalImpression", func(resource ResourceData) error {
		var entity models.ClinicalImpression
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToClinicalImpressions(resp *FhirResponse) ([]*models.ClinicalImpression, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToClinicalImpressions(resp.Bundle)
	case "ClinicalImpression":
		var entity models.ClinicalImpression
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ClinicalImpression{&entity}, nil
	}
	return nil, nil
}

func fhirRespToClinicalImpression(id string, resp *FhirResponse) (*models.ClinicalImpression, error) {
	entities, err := fhirRespToClinicalImpressions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ClinicalImpression", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ClinicalImpression.
func (c *Client) GetClinicalImpression(ctx context.Context, params Parameters) ([]*models.ClinicalImpression, error) {
	resp, err := c.Get(ctx, "ClinicalImpression", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClinicalImpressions(resp)
}

// Get ClinicalImpression by ID.
func (c *Client) GetClinicalImpressionByID(ctx context.Context, id string, params Parameters) (*models.ClinicalImpression, error) {
	resp, err := c.GetByID(ctx, "ClinicalImpression", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClinicalImpression(id, resp)
}

func (c *Client) CreateClinicalImpression(ctx context.Context, resource string, params Parameters, entity *models.ClinicalImpression) (*models.ClinicalImpression, error) {
	resp, err := c.Create(ctx, "ClinicalImpression", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClinicalImpression(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateClinicalImpression(ctx context.Context, resource string, params Parameters, entity *models.ClinicalImpression) (*models.ClinicalImpression, error) {
	resp, err := c.Update(ctx, "ClinicalImpression", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClinicalImpression(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateClinicalImpressionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ClinicalImpression) (*models.ClinicalImpression, error) {
	resp, err := c.UpdateByID(ctx, "ClinicalImpression", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClinicalImpression(id, resp)
}

func (c *Client) PatchClinicalImpression(ctx context.Context, resource string, params Parameters, entity *models.ClinicalImpression) ([]*models.ClinicalImpression, error) {
	resp, err := c.Patch(ctx, "ClinicalImpression", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClinicalImpressions(resp)
}

func (c *Client) PatchClinicalImpressionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ClinicalImpression) (*models.ClinicalImpression, error) {
	resp, err := c.PatchByID(ctx, "ClinicalImpression", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToClinicalImpression(id, resp)
}

func (c *Client) DeleteClinicalImpression(ctx context.Context, resource string, params Parameters) ([]*models.ClinicalImpression, error) {
	resp, err := c.Delete(ctx, "ClinicalImpression", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClinicalImpressions(resp)
}

func (c *Client) DeleteClinicalImpressionByID(ctx context.Context, resource, id string, params Parameters) (*models.ClinicalImpression, error) {
	resp, err := c.DeleteByID(ctx, "ClinicalImpression", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToClinicalImpression(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// CodeSystem
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCodeSystems(bundle *models.Bundle) ([]*models.CodeSystem, error) {
	var entities []*models.CodeSystem
	err := EnumBundleResources(bundle, "CodeSystem", func(resource ResourceData) error {
		var entity models.CodeSystem
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCodeSystems(resp *FhirResponse) ([]*models.CodeSystem, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCodeSystems(resp.Bundle)
	case "CodeSystem":
		var entity models.CodeSystem
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.CodeSystem{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCodeSystem(id string, resp *FhirResponse) (*models.CodeSystem, error) {
	entities, err := fhirRespToCodeSystems(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "CodeSystem", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get CodeSystem.
func (c *Client) GetCodeSystem(ctx context.Context, params Parameters) ([]*models.CodeSystem, error) {
	resp, err := c.Get(ctx, "CodeSystem", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCodeSystems(resp)
}

// Get CodeSystem by ID.
func (c *Client) GetCodeSystemByID(ctx context.Context, id string, params Parameters) (*models.CodeSystem, error) {
	resp, err := c.GetByID(ctx, "CodeSystem", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCodeSystem(id, resp)
}

func (c *Client) CreateCodeSystem(ctx context.Context, resource string, params Parameters, entity *models.CodeSystem) (*models.CodeSystem, error) {
	resp, err := c.Create(ctx, "CodeSystem", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCodeSystem(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCodeSystem(ctx context.Context, resource string, params Parameters, entity *models.CodeSystem) (*models.CodeSystem, error) {
	resp, err := c.Update(ctx, "CodeSystem", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCodeSystem(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCodeSystemByID(ctx context.Context, resource, id string, params Parameters, entity *models.CodeSystem) (*models.CodeSystem, error) {
	resp, err := c.UpdateByID(ctx, "CodeSystem", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCodeSystem(id, resp)
}

func (c *Client) PatchCodeSystem(ctx context.Context, resource string, params Parameters, entity *models.CodeSystem) ([]*models.CodeSystem, error) {
	resp, err := c.Patch(ctx, "CodeSystem", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCodeSystems(resp)
}

func (c *Client) PatchCodeSystemByID(ctx context.Context, resource, id string, params Parameters, entity *models.CodeSystem) (*models.CodeSystem, error) {
	resp, err := c.PatchByID(ctx, "CodeSystem", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCodeSystem(id, resp)
}

func (c *Client) DeleteCodeSystem(ctx context.Context, resource string, params Parameters) ([]*models.CodeSystem, error) {
	resp, err := c.Delete(ctx, "CodeSystem", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCodeSystems(resp)
}

func (c *Client) DeleteCodeSystemByID(ctx context.Context, resource, id string, params Parameters) (*models.CodeSystem, error) {
	resp, err := c.DeleteByID(ctx, "CodeSystem", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCodeSystem(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Communication
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCommunications(bundle *models.Bundle) ([]*models.Communication, error) {
	var entities []*models.Communication
	err := EnumBundleResources(bundle, "Communication", func(resource ResourceData) error {
		var entity models.Communication
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCommunications(resp *FhirResponse) ([]*models.Communication, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCommunications(resp.Bundle)
	case "Communication":
		var entity models.Communication
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Communication{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCommunication(id string, resp *FhirResponse) (*models.Communication, error) {
	entities, err := fhirRespToCommunications(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Communication", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Communication.
func (c *Client) GetCommunication(ctx context.Context, params Parameters) ([]*models.Communication, error) {
	resp, err := c.Get(ctx, "Communication", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunications(resp)
}

// Get Communication by ID.
func (c *Client) GetCommunicationByID(ctx context.Context, id string, params Parameters) (*models.Communication, error) {
	resp, err := c.GetByID(ctx, "Communication", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunication(id, resp)
}

func (c *Client) CreateCommunication(ctx context.Context, resource string, params Parameters, entity *models.Communication) (*models.Communication, error) {
	resp, err := c.Create(ctx, "Communication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunication(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCommunication(ctx context.Context, resource string, params Parameters, entity *models.Communication) (*models.Communication, error) {
	resp, err := c.Update(ctx, "Communication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunication(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCommunicationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Communication) (*models.Communication, error) {
	resp, err := c.UpdateByID(ctx, "Communication", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunication(id, resp)
}

func (c *Client) PatchCommunication(ctx context.Context, resource string, params Parameters, entity *models.Communication) ([]*models.Communication, error) {
	resp, err := c.Patch(ctx, "Communication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunications(resp)
}

func (c *Client) PatchCommunicationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Communication) (*models.Communication, error) {
	resp, err := c.PatchByID(ctx, "Communication", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunication(id, resp)
}

func (c *Client) DeleteCommunication(ctx context.Context, resource string, params Parameters) ([]*models.Communication, error) {
	resp, err := c.Delete(ctx, "Communication", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunications(resp)
}

func (c *Client) DeleteCommunicationByID(ctx context.Context, resource, id string, params Parameters) (*models.Communication, error) {
	resp, err := c.DeleteByID(ctx, "Communication", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunication(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// CommunicationRequest
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCommunicationRequests(bundle *models.Bundle) ([]*models.CommunicationRequest, error) {
	var entities []*models.CommunicationRequest
	err := EnumBundleResources(bundle, "CommunicationRequest", func(resource ResourceData) error {
		var entity models.CommunicationRequest
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCommunicationRequests(resp *FhirResponse) ([]*models.CommunicationRequest, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCommunicationRequests(resp.Bundle)
	case "CommunicationRequest":
		var entity models.CommunicationRequest
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.CommunicationRequest{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCommunicationRequest(id string, resp *FhirResponse) (*models.CommunicationRequest, error) {
	entities, err := fhirRespToCommunicationRequests(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "CommunicationRequest", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get CommunicationRequest.
func (c *Client) GetCommunicationRequest(ctx context.Context, params Parameters) ([]*models.CommunicationRequest, error) {
	resp, err := c.Get(ctx, "CommunicationRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunicationRequests(resp)
}

// Get CommunicationRequest by ID.
func (c *Client) GetCommunicationRequestByID(ctx context.Context, id string, params Parameters) (*models.CommunicationRequest, error) {
	resp, err := c.GetByID(ctx, "CommunicationRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunicationRequest(id, resp)
}

func (c *Client) CreateCommunicationRequest(ctx context.Context, resource string, params Parameters, entity *models.CommunicationRequest) (*models.CommunicationRequest, error) {
	resp, err := c.Create(ctx, "CommunicationRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunicationRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCommunicationRequest(ctx context.Context, resource string, params Parameters, entity *models.CommunicationRequest) (*models.CommunicationRequest, error) {
	resp, err := c.Update(ctx, "CommunicationRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunicationRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCommunicationRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.CommunicationRequest) (*models.CommunicationRequest, error) {
	resp, err := c.UpdateByID(ctx, "CommunicationRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunicationRequest(id, resp)
}

func (c *Client) PatchCommunicationRequest(ctx context.Context, resource string, params Parameters, entity *models.CommunicationRequest) ([]*models.CommunicationRequest, error) {
	resp, err := c.Patch(ctx, "CommunicationRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunicationRequests(resp)
}

func (c *Client) PatchCommunicationRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.CommunicationRequest) (*models.CommunicationRequest, error) {
	resp, err := c.PatchByID(ctx, "CommunicationRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunicationRequest(id, resp)
}

func (c *Client) DeleteCommunicationRequest(ctx context.Context, resource string, params Parameters) ([]*models.CommunicationRequest, error) {
	resp, err := c.Delete(ctx, "CommunicationRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunicationRequests(resp)
}

func (c *Client) DeleteCommunicationRequestByID(ctx context.Context, resource, id string, params Parameters) (*models.CommunicationRequest, error) {
	resp, err := c.DeleteByID(ctx, "CommunicationRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCommunicationRequest(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// CompartmentDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCompartmentDefinitions(bundle *models.Bundle) ([]*models.CompartmentDefinition, error) {
	var entities []*models.CompartmentDefinition
	err := EnumBundleResources(bundle, "CompartmentDefinition", func(resource ResourceData) error {
		var entity models.CompartmentDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCompartmentDefinitions(resp *FhirResponse) ([]*models.CompartmentDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCompartmentDefinitions(resp.Bundle)
	case "CompartmentDefinition":
		var entity models.CompartmentDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.CompartmentDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCompartmentDefinition(id string, resp *FhirResponse) (*models.CompartmentDefinition, error) {
	entities, err := fhirRespToCompartmentDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "CompartmentDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get CompartmentDefinition.
func (c *Client) GetCompartmentDefinition(ctx context.Context, params Parameters) ([]*models.CompartmentDefinition, error) {
	resp, err := c.Get(ctx, "CompartmentDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompartmentDefinitions(resp)
}

// Get CompartmentDefinition by ID.
func (c *Client) GetCompartmentDefinitionByID(ctx context.Context, id string, params Parameters) (*models.CompartmentDefinition, error) {
	resp, err := c.GetByID(ctx, "CompartmentDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompartmentDefinition(id, resp)
}

func (c *Client) CreateCompartmentDefinition(ctx context.Context, resource string, params Parameters, entity *models.CompartmentDefinition) (*models.CompartmentDefinition, error) {
	resp, err := c.Create(ctx, "CompartmentDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompartmentDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCompartmentDefinition(ctx context.Context, resource string, params Parameters, entity *models.CompartmentDefinition) (*models.CompartmentDefinition, error) {
	resp, err := c.Update(ctx, "CompartmentDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompartmentDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCompartmentDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.CompartmentDefinition) (*models.CompartmentDefinition, error) {
	resp, err := c.UpdateByID(ctx, "CompartmentDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompartmentDefinition(id, resp)
}

func (c *Client) PatchCompartmentDefinition(ctx context.Context, resource string, params Parameters, entity *models.CompartmentDefinition) ([]*models.CompartmentDefinition, error) {
	resp, err := c.Patch(ctx, "CompartmentDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompartmentDefinitions(resp)
}

func (c *Client) PatchCompartmentDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.CompartmentDefinition) (*models.CompartmentDefinition, error) {
	resp, err := c.PatchByID(ctx, "CompartmentDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompartmentDefinition(id, resp)
}

func (c *Client) DeleteCompartmentDefinition(ctx context.Context, resource string, params Parameters) ([]*models.CompartmentDefinition, error) {
	resp, err := c.Delete(ctx, "CompartmentDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompartmentDefinitions(resp)
}

func (c *Client) DeleteCompartmentDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.CompartmentDefinition, error) {
	resp, err := c.DeleteByID(ctx, "CompartmentDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompartmentDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Composition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCompositions(bundle *models.Bundle) ([]*models.Composition, error) {
	var entities []*models.Composition
	err := EnumBundleResources(bundle, "Composition", func(resource ResourceData) error {
		var entity models.Composition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCompositions(resp *FhirResponse) ([]*models.Composition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCompositions(resp.Bundle)
	case "Composition":
		var entity models.Composition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Composition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToComposition(id string, resp *FhirResponse) (*models.Composition, error) {
	entities, err := fhirRespToCompositions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Composition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Composition.
func (c *Client) GetComposition(ctx context.Context, params Parameters) ([]*models.Composition, error) {
	resp, err := c.Get(ctx, "Composition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompositions(resp)
}

// Get Composition by ID.
func (c *Client) GetCompositionByID(ctx context.Context, id string, params Parameters) (*models.Composition, error) {
	resp, err := c.GetByID(ctx, "Composition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToComposition(id, resp)
}

func (c *Client) CreateComposition(ctx context.Context, resource string, params Parameters, entity *models.Composition) (*models.Composition, error) {
	resp, err := c.Create(ctx, "Composition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToComposition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateComposition(ctx context.Context, resource string, params Parameters, entity *models.Composition) (*models.Composition, error) {
	resp, err := c.Update(ctx, "Composition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToComposition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCompositionByID(ctx context.Context, resource, id string, params Parameters, entity *models.Composition) (*models.Composition, error) {
	resp, err := c.UpdateByID(ctx, "Composition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToComposition(id, resp)
}

func (c *Client) PatchComposition(ctx context.Context, resource string, params Parameters, entity *models.Composition) ([]*models.Composition, error) {
	resp, err := c.Patch(ctx, "Composition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompositions(resp)
}

func (c *Client) PatchCompositionByID(ctx context.Context, resource, id string, params Parameters, entity *models.Composition) (*models.Composition, error) {
	resp, err := c.PatchByID(ctx, "Composition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToComposition(id, resp)
}

func (c *Client) DeleteComposition(ctx context.Context, resource string, params Parameters) ([]*models.Composition, error) {
	resp, err := c.Delete(ctx, "Composition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCompositions(resp)
}

func (c *Client) DeleteCompositionByID(ctx context.Context, resource, id string, params Parameters) (*models.Composition, error) {
	resp, err := c.DeleteByID(ctx, "Composition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToComposition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ConceptMap
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToConceptMaps(bundle *models.Bundle) ([]*models.ConceptMap, error) {
	var entities []*models.ConceptMap
	err := EnumBundleResources(bundle, "ConceptMap", func(resource ResourceData) error {
		var entity models.ConceptMap
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToConceptMaps(resp *FhirResponse) ([]*models.ConceptMap, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToConceptMaps(resp.Bundle)
	case "ConceptMap":
		var entity models.ConceptMap
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ConceptMap{&entity}, nil
	}
	return nil, nil
}

func fhirRespToConceptMap(id string, resp *FhirResponse) (*models.ConceptMap, error) {
	entities, err := fhirRespToConceptMaps(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ConceptMap", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ConceptMap.
func (c *Client) GetConceptMap(ctx context.Context, params Parameters) ([]*models.ConceptMap, error) {
	resp, err := c.Get(ctx, "ConceptMap", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToConceptMaps(resp)
}

// Get ConceptMap by ID.
func (c *Client) GetConceptMapByID(ctx context.Context, id string, params Parameters) (*models.ConceptMap, error) {
	resp, err := c.GetByID(ctx, "ConceptMap", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToConceptMap(id, resp)
}

func (c *Client) CreateConceptMap(ctx context.Context, resource string, params Parameters, entity *models.ConceptMap) (*models.ConceptMap, error) {
	resp, err := c.Create(ctx, "ConceptMap", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConceptMap(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateConceptMap(ctx context.Context, resource string, params Parameters, entity *models.ConceptMap) (*models.ConceptMap, error) {
	resp, err := c.Update(ctx, "ConceptMap", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConceptMap(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateConceptMapByID(ctx context.Context, resource, id string, params Parameters, entity *models.ConceptMap) (*models.ConceptMap, error) {
	resp, err := c.UpdateByID(ctx, "ConceptMap", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConceptMap(id, resp)
}

func (c *Client) PatchConceptMap(ctx context.Context, resource string, params Parameters, entity *models.ConceptMap) ([]*models.ConceptMap, error) {
	resp, err := c.Patch(ctx, "ConceptMap", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConceptMaps(resp)
}

func (c *Client) PatchConceptMapByID(ctx context.Context, resource, id string, params Parameters, entity *models.ConceptMap) (*models.ConceptMap, error) {
	resp, err := c.PatchByID(ctx, "ConceptMap", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConceptMap(id, resp)
}

func (c *Client) DeleteConceptMap(ctx context.Context, resource string, params Parameters) ([]*models.ConceptMap, error) {
	resp, err := c.Delete(ctx, "ConceptMap", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToConceptMaps(resp)
}

func (c *Client) DeleteConceptMapByID(ctx context.Context, resource, id string, params Parameters) (*models.ConceptMap, error) {
	resp, err := c.DeleteByID(ctx, "ConceptMap", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToConceptMap(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Condition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToConditions(bundle *models.Bundle) ([]*models.Condition, error) {
	var entities []*models.Condition
	err := EnumBundleResources(bundle, "Condition", func(resource ResourceData) error {
		var entity models.Condition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToConditions(resp *FhirResponse) ([]*models.Condition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToConditions(resp.Bundle)
	case "Condition":
		var entity models.Condition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Condition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCondition(id string, resp *FhirResponse) (*models.Condition, error) {
	entities, err := fhirRespToConditions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Condition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Condition.
func (c *Client) GetCondition(ctx context.Context, params Parameters) ([]*models.Condition, error) {
	resp, err := c.Get(ctx, "Condition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToConditions(resp)
}

// Get Condition by ID.
func (c *Client) GetConditionByID(ctx context.Context, id string, params Parameters) (*models.Condition, error) {
	resp, err := c.GetByID(ctx, "Condition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCondition(id, resp)
}

func (c *Client) CreateCondition(ctx context.Context, resource string, params Parameters, entity *models.Condition) (*models.Condition, error) {
	resp, err := c.Create(ctx, "Condition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCondition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCondition(ctx context.Context, resource string, params Parameters, entity *models.Condition) (*models.Condition, error) {
	resp, err := c.Update(ctx, "Condition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCondition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateConditionByID(ctx context.Context, resource, id string, params Parameters, entity *models.Condition) (*models.Condition, error) {
	resp, err := c.UpdateByID(ctx, "Condition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCondition(id, resp)
}

func (c *Client) PatchCondition(ctx context.Context, resource string, params Parameters, entity *models.Condition) ([]*models.Condition, error) {
	resp, err := c.Patch(ctx, "Condition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConditions(resp)
}

func (c *Client) PatchConditionByID(ctx context.Context, resource, id string, params Parameters, entity *models.Condition) (*models.Condition, error) {
	resp, err := c.PatchByID(ctx, "Condition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCondition(id, resp)
}

func (c *Client) DeleteCondition(ctx context.Context, resource string, params Parameters) ([]*models.Condition, error) {
	resp, err := c.Delete(ctx, "Condition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToConditions(resp)
}

func (c *Client) DeleteConditionByID(ctx context.Context, resource, id string, params Parameters) (*models.Condition, error) {
	resp, err := c.DeleteByID(ctx, "Condition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCondition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Consent
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToConsents(bundle *models.Bundle) ([]*models.Consent, error) {
	var entities []*models.Consent
	err := EnumBundleResources(bundle, "Consent", func(resource ResourceData) error {
		var entity models.Consent
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToConsents(resp *FhirResponse) ([]*models.Consent, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToConsents(resp.Bundle)
	case "Consent":
		var entity models.Consent
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Consent{&entity}, nil
	}
	return nil, nil
}

func fhirRespToConsent(id string, resp *FhirResponse) (*models.Consent, error) {
	entities, err := fhirRespToConsents(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Consent", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Consent.
func (c *Client) GetConsent(ctx context.Context, params Parameters) ([]*models.Consent, error) {
	resp, err := c.Get(ctx, "Consent", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToConsents(resp)
}

// Get Consent by ID.
func (c *Client) GetConsentByID(ctx context.Context, id string, params Parameters) (*models.Consent, error) {
	resp, err := c.GetByID(ctx, "Consent", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToConsent(id, resp)
}

func (c *Client) CreateConsent(ctx context.Context, resource string, params Parameters, entity *models.Consent) (*models.Consent, error) {
	resp, err := c.Create(ctx, "Consent", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConsent(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateConsent(ctx context.Context, resource string, params Parameters, entity *models.Consent) (*models.Consent, error) {
	resp, err := c.Update(ctx, "Consent", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConsent(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateConsentByID(ctx context.Context, resource, id string, params Parameters, entity *models.Consent) (*models.Consent, error) {
	resp, err := c.UpdateByID(ctx, "Consent", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConsent(id, resp)
}

func (c *Client) PatchConsent(ctx context.Context, resource string, params Parameters, entity *models.Consent) ([]*models.Consent, error) {
	resp, err := c.Patch(ctx, "Consent", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConsents(resp)
}

func (c *Client) PatchConsentByID(ctx context.Context, resource, id string, params Parameters, entity *models.Consent) (*models.Consent, error) {
	resp, err := c.PatchByID(ctx, "Consent", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToConsent(id, resp)
}

func (c *Client) DeleteConsent(ctx context.Context, resource string, params Parameters) ([]*models.Consent, error) {
	resp, err := c.Delete(ctx, "Consent", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToConsents(resp)
}

func (c *Client) DeleteConsentByID(ctx context.Context, resource, id string, params Parameters) (*models.Consent, error) {
	resp, err := c.DeleteByID(ctx, "Consent", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToConsent(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Contract
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToContracts(bundle *models.Bundle) ([]*models.Contract, error) {
	var entities []*models.Contract
	err := EnumBundleResources(bundle, "Contract", func(resource ResourceData) error {
		var entity models.Contract
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToContracts(resp *FhirResponse) ([]*models.Contract, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToContracts(resp.Bundle)
	case "Contract":
		var entity models.Contract
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Contract{&entity}, nil
	}
	return nil, nil
}

func fhirRespToContract(id string, resp *FhirResponse) (*models.Contract, error) {
	entities, err := fhirRespToContracts(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Contract", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Contract.
func (c *Client) GetContract(ctx context.Context, params Parameters) ([]*models.Contract, error) {
	resp, err := c.Get(ctx, "Contract", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToContracts(resp)
}

// Get Contract by ID.
func (c *Client) GetContractByID(ctx context.Context, id string, params Parameters) (*models.Contract, error) {
	resp, err := c.GetByID(ctx, "Contract", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToContract(id, resp)
}

func (c *Client) CreateContract(ctx context.Context, resource string, params Parameters, entity *models.Contract) (*models.Contract, error) {
	resp, err := c.Create(ctx, "Contract", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToContract(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateContract(ctx context.Context, resource string, params Parameters, entity *models.Contract) (*models.Contract, error) {
	resp, err := c.Update(ctx, "Contract", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToContract(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateContractByID(ctx context.Context, resource, id string, params Parameters, entity *models.Contract) (*models.Contract, error) {
	resp, err := c.UpdateByID(ctx, "Contract", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToContract(id, resp)
}

func (c *Client) PatchContract(ctx context.Context, resource string, params Parameters, entity *models.Contract) ([]*models.Contract, error) {
	resp, err := c.Patch(ctx, "Contract", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToContracts(resp)
}

func (c *Client) PatchContractByID(ctx context.Context, resource, id string, params Parameters, entity *models.Contract) (*models.Contract, error) {
	resp, err := c.PatchByID(ctx, "Contract", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToContract(id, resp)
}

func (c *Client) DeleteContract(ctx context.Context, resource string, params Parameters) ([]*models.Contract, error) {
	resp, err := c.Delete(ctx, "Contract", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToContracts(resp)
}

func (c *Client) DeleteContractByID(ctx context.Context, resource, id string, params Parameters) (*models.Contract, error) {
	resp, err := c.DeleteByID(ctx, "Contract", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToContract(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Coverage
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCoverages(bundle *models.Bundle) ([]*models.Coverage, error) {
	var entities []*models.Coverage
	err := EnumBundleResources(bundle, "Coverage", func(resource ResourceData) error {
		var entity models.Coverage
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCoverages(resp *FhirResponse) ([]*models.Coverage, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCoverages(resp.Bundle)
	case "Coverage":
		var entity models.Coverage
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Coverage{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCoverage(id string, resp *FhirResponse) (*models.Coverage, error) {
	entities, err := fhirRespToCoverages(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Coverage", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Coverage.
func (c *Client) GetCoverage(ctx context.Context, params Parameters) ([]*models.Coverage, error) {
	resp, err := c.Get(ctx, "Coverage", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverages(resp)
}

// Get Coverage by ID.
func (c *Client) GetCoverageByID(ctx context.Context, id string, params Parameters) (*models.Coverage, error) {
	resp, err := c.GetByID(ctx, "Coverage", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverage(id, resp)
}

func (c *Client) CreateCoverage(ctx context.Context, resource string, params Parameters, entity *models.Coverage) (*models.Coverage, error) {
	resp, err := c.Create(ctx, "Coverage", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverage(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCoverage(ctx context.Context, resource string, params Parameters, entity *models.Coverage) (*models.Coverage, error) {
	resp, err := c.Update(ctx, "Coverage", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverage(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCoverageByID(ctx context.Context, resource, id string, params Parameters, entity *models.Coverage) (*models.Coverage, error) {
	resp, err := c.UpdateByID(ctx, "Coverage", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverage(id, resp)
}

func (c *Client) PatchCoverage(ctx context.Context, resource string, params Parameters, entity *models.Coverage) ([]*models.Coverage, error) {
	resp, err := c.Patch(ctx, "Coverage", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverages(resp)
}

func (c *Client) PatchCoverageByID(ctx context.Context, resource, id string, params Parameters, entity *models.Coverage) (*models.Coverage, error) {
	resp, err := c.PatchByID(ctx, "Coverage", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverage(id, resp)
}

func (c *Client) DeleteCoverage(ctx context.Context, resource string, params Parameters) ([]*models.Coverage, error) {
	resp, err := c.Delete(ctx, "Coverage", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverages(resp)
}

func (c *Client) DeleteCoverageByID(ctx context.Context, resource, id string, params Parameters) (*models.Coverage, error) {
	resp, err := c.DeleteByID(ctx, "Coverage", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverage(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// CoverageEligibilityRequest
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCoverageEligibilityRequests(bundle *models.Bundle) ([]*models.CoverageEligibilityRequest, error) {
	var entities []*models.CoverageEligibilityRequest
	err := EnumBundleResources(bundle, "CoverageEligibilityRequest", func(resource ResourceData) error {
		var entity models.CoverageEligibilityRequest
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCoverageEligibilityRequests(resp *FhirResponse) ([]*models.CoverageEligibilityRequest, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCoverageEligibilityRequests(resp.Bundle)
	case "CoverageEligibilityRequest":
		var entity models.CoverageEligibilityRequest
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.CoverageEligibilityRequest{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCoverageEligibilityRequest(id string, resp *FhirResponse) (*models.CoverageEligibilityRequest, error) {
	entities, err := fhirRespToCoverageEligibilityRequests(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "CoverageEligibilityRequest", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get CoverageEligibilityRequest.
func (c *Client) GetCoverageEligibilityRequest(ctx context.Context, params Parameters) ([]*models.CoverageEligibilityRequest, error) {
	resp, err := c.Get(ctx, "CoverageEligibilityRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityRequests(resp)
}

// Get CoverageEligibilityRequest by ID.
func (c *Client) GetCoverageEligibilityRequestByID(ctx context.Context, id string, params Parameters) (*models.CoverageEligibilityRequest, error) {
	resp, err := c.GetByID(ctx, "CoverageEligibilityRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityRequest(id, resp)
}

func (c *Client) CreateCoverageEligibilityRequest(ctx context.Context, resource string, params Parameters, entity *models.CoverageEligibilityRequest) (*models.CoverageEligibilityRequest, error) {
	resp, err := c.Create(ctx, "CoverageEligibilityRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCoverageEligibilityRequest(ctx context.Context, resource string, params Parameters, entity *models.CoverageEligibilityRequest) (*models.CoverageEligibilityRequest, error) {
	resp, err := c.Update(ctx, "CoverageEligibilityRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCoverageEligibilityRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.CoverageEligibilityRequest) (*models.CoverageEligibilityRequest, error) {
	resp, err := c.UpdateByID(ctx, "CoverageEligibilityRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityRequest(id, resp)
}

func (c *Client) PatchCoverageEligibilityRequest(ctx context.Context, resource string, params Parameters, entity *models.CoverageEligibilityRequest) ([]*models.CoverageEligibilityRequest, error) {
	resp, err := c.Patch(ctx, "CoverageEligibilityRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityRequests(resp)
}

func (c *Client) PatchCoverageEligibilityRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.CoverageEligibilityRequest) (*models.CoverageEligibilityRequest, error) {
	resp, err := c.PatchByID(ctx, "CoverageEligibilityRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityRequest(id, resp)
}

func (c *Client) DeleteCoverageEligibilityRequest(ctx context.Context, resource string, params Parameters) ([]*models.CoverageEligibilityRequest, error) {
	resp, err := c.Delete(ctx, "CoverageEligibilityRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityRequests(resp)
}

func (c *Client) DeleteCoverageEligibilityRequestByID(ctx context.Context, resource, id string, params Parameters) (*models.CoverageEligibilityRequest, error) {
	resp, err := c.DeleteByID(ctx, "CoverageEligibilityRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityRequest(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// CoverageEligibilityResponse
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToCoverageEligibilityResponses(bundle *models.Bundle) ([]*models.CoverageEligibilityResponse, error) {
	var entities []*models.CoverageEligibilityResponse
	err := EnumBundleResources(bundle, "CoverageEligibilityResponse", func(resource ResourceData) error {
		var entity models.CoverageEligibilityResponse
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToCoverageEligibilityResponses(resp *FhirResponse) ([]*models.CoverageEligibilityResponse, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToCoverageEligibilityResponses(resp.Bundle)
	case "CoverageEligibilityResponse":
		var entity models.CoverageEligibilityResponse
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.CoverageEligibilityResponse{&entity}, nil
	}
	return nil, nil
}

func fhirRespToCoverageEligibilityResponse(id string, resp *FhirResponse) (*models.CoverageEligibilityResponse, error) {
	entities, err := fhirRespToCoverageEligibilityResponses(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "CoverageEligibilityResponse", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get CoverageEligibilityResponse.
func (c *Client) GetCoverageEligibilityResponse(ctx context.Context, params Parameters) ([]*models.CoverageEligibilityResponse, error) {
	resp, err := c.Get(ctx, "CoverageEligibilityResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityResponses(resp)
}

// Get CoverageEligibilityResponse by ID.
func (c *Client) GetCoverageEligibilityResponseByID(ctx context.Context, id string, params Parameters) (*models.CoverageEligibilityResponse, error) {
	resp, err := c.GetByID(ctx, "CoverageEligibilityResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityResponse(id, resp)
}

func (c *Client) CreateCoverageEligibilityResponse(ctx context.Context, resource string, params Parameters, entity *models.CoverageEligibilityResponse) (*models.CoverageEligibilityResponse, error) {
	resp, err := c.Create(ctx, "CoverageEligibilityResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCoverageEligibilityResponse(ctx context.Context, resource string, params Parameters, entity *models.CoverageEligibilityResponse) (*models.CoverageEligibilityResponse, error) {
	resp, err := c.Update(ctx, "CoverageEligibilityResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateCoverageEligibilityResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.CoverageEligibilityResponse) (*models.CoverageEligibilityResponse, error) {
	resp, err := c.UpdateByID(ctx, "CoverageEligibilityResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityResponse(id, resp)
}

func (c *Client) PatchCoverageEligibilityResponse(ctx context.Context, resource string, params Parameters, entity *models.CoverageEligibilityResponse) ([]*models.CoverageEligibilityResponse, error) {
	resp, err := c.Patch(ctx, "CoverageEligibilityResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityResponses(resp)
}

func (c *Client) PatchCoverageEligibilityResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.CoverageEligibilityResponse) (*models.CoverageEligibilityResponse, error) {
	resp, err := c.PatchByID(ctx, "CoverageEligibilityResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityResponse(id, resp)
}

func (c *Client) DeleteCoverageEligibilityResponse(ctx context.Context, resource string, params Parameters) ([]*models.CoverageEligibilityResponse, error) {
	resp, err := c.Delete(ctx, "CoverageEligibilityResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityResponses(resp)
}

func (c *Client) DeleteCoverageEligibilityResponseByID(ctx context.Context, resource, id string, params Parameters) (*models.CoverageEligibilityResponse, error) {
	resp, err := c.DeleteByID(ctx, "CoverageEligibilityResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToCoverageEligibilityResponse(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// DetectedIssue
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToDetectedIssues(bundle *models.Bundle) ([]*models.DetectedIssue, error) {
	var entities []*models.DetectedIssue
	err := EnumBundleResources(bundle, "DetectedIssue", func(resource ResourceData) error {
		var entity models.DetectedIssue
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToDetectedIssues(resp *FhirResponse) ([]*models.DetectedIssue, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToDetectedIssues(resp.Bundle)
	case "DetectedIssue":
		var entity models.DetectedIssue
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.DetectedIssue{&entity}, nil
	}
	return nil, nil
}

func fhirRespToDetectedIssue(id string, resp *FhirResponse) (*models.DetectedIssue, error) {
	entities, err := fhirRespToDetectedIssues(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "DetectedIssue", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get DetectedIssue.
func (c *Client) GetDetectedIssue(ctx context.Context, params Parameters) ([]*models.DetectedIssue, error) {
	resp, err := c.Get(ctx, "DetectedIssue", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDetectedIssues(resp)
}

// Get DetectedIssue by ID.
func (c *Client) GetDetectedIssueByID(ctx context.Context, id string, params Parameters) (*models.DetectedIssue, error) {
	resp, err := c.GetByID(ctx, "DetectedIssue", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDetectedIssue(id, resp)
}

func (c *Client) CreateDetectedIssue(ctx context.Context, resource string, params Parameters, entity *models.DetectedIssue) (*models.DetectedIssue, error) {
	resp, err := c.Create(ctx, "DetectedIssue", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDetectedIssue(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDetectedIssue(ctx context.Context, resource string, params Parameters, entity *models.DetectedIssue) (*models.DetectedIssue, error) {
	resp, err := c.Update(ctx, "DetectedIssue", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDetectedIssue(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDetectedIssueByID(ctx context.Context, resource, id string, params Parameters, entity *models.DetectedIssue) (*models.DetectedIssue, error) {
	resp, err := c.UpdateByID(ctx, "DetectedIssue", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDetectedIssue(id, resp)
}

func (c *Client) PatchDetectedIssue(ctx context.Context, resource string, params Parameters, entity *models.DetectedIssue) ([]*models.DetectedIssue, error) {
	resp, err := c.Patch(ctx, "DetectedIssue", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDetectedIssues(resp)
}

func (c *Client) PatchDetectedIssueByID(ctx context.Context, resource, id string, params Parameters, entity *models.DetectedIssue) (*models.DetectedIssue, error) {
	resp, err := c.PatchByID(ctx, "DetectedIssue", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDetectedIssue(id, resp)
}

func (c *Client) DeleteDetectedIssue(ctx context.Context, resource string, params Parameters) ([]*models.DetectedIssue, error) {
	resp, err := c.Delete(ctx, "DetectedIssue", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDetectedIssues(resp)
}

func (c *Client) DeleteDetectedIssueByID(ctx context.Context, resource, id string, params Parameters) (*models.DetectedIssue, error) {
	resp, err := c.DeleteByID(ctx, "DetectedIssue", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDetectedIssue(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Device
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToDevices(bundle *models.Bundle) ([]*models.Device, error) {
	var entities []*models.Device
	err := EnumBundleResources(bundle, "Device", func(resource ResourceData) error {
		var entity models.Device
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToDevices(resp *FhirResponse) ([]*models.Device, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToDevices(resp.Bundle)
	case "Device":
		var entity models.Device
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Device{&entity}, nil
	}
	return nil, nil
}

func fhirRespToDevice(id string, resp *FhirResponse) (*models.Device, error) {
	entities, err := fhirRespToDevices(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Device", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Device.
func (c *Client) GetDevice(ctx context.Context, params Parameters) ([]*models.Device, error) {
	resp, err := c.Get(ctx, "Device", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDevices(resp)
}

// Get Device by ID.
func (c *Client) GetDeviceByID(ctx context.Context, id string, params Parameters) (*models.Device, error) {
	resp, err := c.GetByID(ctx, "Device", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDevice(id, resp)
}

func (c *Client) CreateDevice(ctx context.Context, resource string, params Parameters, entity *models.Device) (*models.Device, error) {
	resp, err := c.Create(ctx, "Device", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDevice(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDevice(ctx context.Context, resource string, params Parameters, entity *models.Device) (*models.Device, error) {
	resp, err := c.Update(ctx, "Device", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDevice(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDeviceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Device) (*models.Device, error) {
	resp, err := c.UpdateByID(ctx, "Device", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDevice(id, resp)
}

func (c *Client) PatchDevice(ctx context.Context, resource string, params Parameters, entity *models.Device) ([]*models.Device, error) {
	resp, err := c.Patch(ctx, "Device", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDevices(resp)
}

func (c *Client) PatchDeviceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Device) (*models.Device, error) {
	resp, err := c.PatchByID(ctx, "Device", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDevice(id, resp)
}

func (c *Client) DeleteDevice(ctx context.Context, resource string, params Parameters) ([]*models.Device, error) {
	resp, err := c.Delete(ctx, "Device", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDevices(resp)
}

func (c *Client) DeleteDeviceByID(ctx context.Context, resource, id string, params Parameters) (*models.Device, error) {
	resp, err := c.DeleteByID(ctx, "Device", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDevice(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// DeviceDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToDeviceDefinitions(bundle *models.Bundle) ([]*models.DeviceDefinition, error) {
	var entities []*models.DeviceDefinition
	err := EnumBundleResources(bundle, "DeviceDefinition", func(resource ResourceData) error {
		var entity models.DeviceDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToDeviceDefinitions(resp *FhirResponse) ([]*models.DeviceDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToDeviceDefinitions(resp.Bundle)
	case "DeviceDefinition":
		var entity models.DeviceDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.DeviceDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToDeviceDefinition(id string, resp *FhirResponse) (*models.DeviceDefinition, error) {
	entities, err := fhirRespToDeviceDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "DeviceDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get DeviceDefinition.
func (c *Client) GetDeviceDefinition(ctx context.Context, params Parameters) ([]*models.DeviceDefinition, error) {
	resp, err := c.Get(ctx, "DeviceDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceDefinitions(resp)
}

// Get DeviceDefinition by ID.
func (c *Client) GetDeviceDefinitionByID(ctx context.Context, id string, params Parameters) (*models.DeviceDefinition, error) {
	resp, err := c.GetByID(ctx, "DeviceDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceDefinition(id, resp)
}

func (c *Client) CreateDeviceDefinition(ctx context.Context, resource string, params Parameters, entity *models.DeviceDefinition) (*models.DeviceDefinition, error) {
	resp, err := c.Create(ctx, "DeviceDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDeviceDefinition(ctx context.Context, resource string, params Parameters, entity *models.DeviceDefinition) (*models.DeviceDefinition, error) {
	resp, err := c.Update(ctx, "DeviceDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDeviceDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.DeviceDefinition) (*models.DeviceDefinition, error) {
	resp, err := c.UpdateByID(ctx, "DeviceDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceDefinition(id, resp)
}

func (c *Client) PatchDeviceDefinition(ctx context.Context, resource string, params Parameters, entity *models.DeviceDefinition) ([]*models.DeviceDefinition, error) {
	resp, err := c.Patch(ctx, "DeviceDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceDefinitions(resp)
}

func (c *Client) PatchDeviceDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.DeviceDefinition) (*models.DeviceDefinition, error) {
	resp, err := c.PatchByID(ctx, "DeviceDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceDefinition(id, resp)
}

func (c *Client) DeleteDeviceDefinition(ctx context.Context, resource string, params Parameters) ([]*models.DeviceDefinition, error) {
	resp, err := c.Delete(ctx, "DeviceDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceDefinitions(resp)
}

func (c *Client) DeleteDeviceDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.DeviceDefinition, error) {
	resp, err := c.DeleteByID(ctx, "DeviceDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// DeviceMetric
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToDeviceMetrics(bundle *models.Bundle) ([]*models.DeviceMetric, error) {
	var entities []*models.DeviceMetric
	err := EnumBundleResources(bundle, "DeviceMetric", func(resource ResourceData) error {
		var entity models.DeviceMetric
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToDeviceMetrics(resp *FhirResponse) ([]*models.DeviceMetric, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToDeviceMetrics(resp.Bundle)
	case "DeviceMetric":
		var entity models.DeviceMetric
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.DeviceMetric{&entity}, nil
	}
	return nil, nil
}

func fhirRespToDeviceMetric(id string, resp *FhirResponse) (*models.DeviceMetric, error) {
	entities, err := fhirRespToDeviceMetrics(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "DeviceMetric", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get DeviceMetric.
func (c *Client) GetDeviceMetric(ctx context.Context, params Parameters) ([]*models.DeviceMetric, error) {
	resp, err := c.Get(ctx, "DeviceMetric", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceMetrics(resp)
}

// Get DeviceMetric by ID.
func (c *Client) GetDeviceMetricByID(ctx context.Context, id string, params Parameters) (*models.DeviceMetric, error) {
	resp, err := c.GetByID(ctx, "DeviceMetric", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceMetric(id, resp)
}

func (c *Client) CreateDeviceMetric(ctx context.Context, resource string, params Parameters, entity *models.DeviceMetric) (*models.DeviceMetric, error) {
	resp, err := c.Create(ctx, "DeviceMetric", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceMetric(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDeviceMetric(ctx context.Context, resource string, params Parameters, entity *models.DeviceMetric) (*models.DeviceMetric, error) {
	resp, err := c.Update(ctx, "DeviceMetric", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceMetric(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDeviceMetricByID(ctx context.Context, resource, id string, params Parameters, entity *models.DeviceMetric) (*models.DeviceMetric, error) {
	resp, err := c.UpdateByID(ctx, "DeviceMetric", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceMetric(id, resp)
}

func (c *Client) PatchDeviceMetric(ctx context.Context, resource string, params Parameters, entity *models.DeviceMetric) ([]*models.DeviceMetric, error) {
	resp, err := c.Patch(ctx, "DeviceMetric", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceMetrics(resp)
}

func (c *Client) PatchDeviceMetricByID(ctx context.Context, resource, id string, params Parameters, entity *models.DeviceMetric) (*models.DeviceMetric, error) {
	resp, err := c.PatchByID(ctx, "DeviceMetric", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceMetric(id, resp)
}

func (c *Client) DeleteDeviceMetric(ctx context.Context, resource string, params Parameters) ([]*models.DeviceMetric, error) {
	resp, err := c.Delete(ctx, "DeviceMetric", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceMetrics(resp)
}

func (c *Client) DeleteDeviceMetricByID(ctx context.Context, resource, id string, params Parameters) (*models.DeviceMetric, error) {
	resp, err := c.DeleteByID(ctx, "DeviceMetric", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceMetric(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// DeviceRequest
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToDeviceRequests(bundle *models.Bundle) ([]*models.DeviceRequest, error) {
	var entities []*models.DeviceRequest
	err := EnumBundleResources(bundle, "DeviceRequest", func(resource ResourceData) error {
		var entity models.DeviceRequest
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToDeviceRequests(resp *FhirResponse) ([]*models.DeviceRequest, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToDeviceRequests(resp.Bundle)
	case "DeviceRequest":
		var entity models.DeviceRequest
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.DeviceRequest{&entity}, nil
	}
	return nil, nil
}

func fhirRespToDeviceRequest(id string, resp *FhirResponse) (*models.DeviceRequest, error) {
	entities, err := fhirRespToDeviceRequests(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "DeviceRequest", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get DeviceRequest.
func (c *Client) GetDeviceRequest(ctx context.Context, params Parameters) ([]*models.DeviceRequest, error) {
	resp, err := c.Get(ctx, "DeviceRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceRequests(resp)
}

// Get DeviceRequest by ID.
func (c *Client) GetDeviceRequestByID(ctx context.Context, id string, params Parameters) (*models.DeviceRequest, error) {
	resp, err := c.GetByID(ctx, "DeviceRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceRequest(id, resp)
}

func (c *Client) CreateDeviceRequest(ctx context.Context, resource string, params Parameters, entity *models.DeviceRequest) (*models.DeviceRequest, error) {
	resp, err := c.Create(ctx, "DeviceRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDeviceRequest(ctx context.Context, resource string, params Parameters, entity *models.DeviceRequest) (*models.DeviceRequest, error) {
	resp, err := c.Update(ctx, "DeviceRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDeviceRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.DeviceRequest) (*models.DeviceRequest, error) {
	resp, err := c.UpdateByID(ctx, "DeviceRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceRequest(id, resp)
}

func (c *Client) PatchDeviceRequest(ctx context.Context, resource string, params Parameters, entity *models.DeviceRequest) ([]*models.DeviceRequest, error) {
	resp, err := c.Patch(ctx, "DeviceRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceRequests(resp)
}

func (c *Client) PatchDeviceRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.DeviceRequest) (*models.DeviceRequest, error) {
	resp, err := c.PatchByID(ctx, "DeviceRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceRequest(id, resp)
}

func (c *Client) DeleteDeviceRequest(ctx context.Context, resource string, params Parameters) ([]*models.DeviceRequest, error) {
	resp, err := c.Delete(ctx, "DeviceRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceRequests(resp)
}

func (c *Client) DeleteDeviceRequestByID(ctx context.Context, resource, id string, params Parameters) (*models.DeviceRequest, error) {
	resp, err := c.DeleteByID(ctx, "DeviceRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceRequest(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// DeviceUseStatement
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToDeviceUseStatements(bundle *models.Bundle) ([]*models.DeviceUseStatement, error) {
	var entities []*models.DeviceUseStatement
	err := EnumBundleResources(bundle, "DeviceUseStatement", func(resource ResourceData) error {
		var entity models.DeviceUseStatement
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToDeviceUseStatements(resp *FhirResponse) ([]*models.DeviceUseStatement, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToDeviceUseStatements(resp.Bundle)
	case "DeviceUseStatement":
		var entity models.DeviceUseStatement
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.DeviceUseStatement{&entity}, nil
	}
	return nil, nil
}

func fhirRespToDeviceUseStatement(id string, resp *FhirResponse) (*models.DeviceUseStatement, error) {
	entities, err := fhirRespToDeviceUseStatements(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "DeviceUseStatement", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get DeviceUseStatement.
func (c *Client) GetDeviceUseStatement(ctx context.Context, params Parameters) ([]*models.DeviceUseStatement, error) {
	resp, err := c.Get(ctx, "DeviceUseStatement", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceUseStatements(resp)
}

// Get DeviceUseStatement by ID.
func (c *Client) GetDeviceUseStatementByID(ctx context.Context, id string, params Parameters) (*models.DeviceUseStatement, error) {
	resp, err := c.GetByID(ctx, "DeviceUseStatement", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceUseStatement(id, resp)
}

func (c *Client) CreateDeviceUseStatement(ctx context.Context, resource string, params Parameters, entity *models.DeviceUseStatement) (*models.DeviceUseStatement, error) {
	resp, err := c.Create(ctx, "DeviceUseStatement", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceUseStatement(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDeviceUseStatement(ctx context.Context, resource string, params Parameters, entity *models.DeviceUseStatement) (*models.DeviceUseStatement, error) {
	resp, err := c.Update(ctx, "DeviceUseStatement", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceUseStatement(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDeviceUseStatementByID(ctx context.Context, resource, id string, params Parameters, entity *models.DeviceUseStatement) (*models.DeviceUseStatement, error) {
	resp, err := c.UpdateByID(ctx, "DeviceUseStatement", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceUseStatement(id, resp)
}

func (c *Client) PatchDeviceUseStatement(ctx context.Context, resource string, params Parameters, entity *models.DeviceUseStatement) ([]*models.DeviceUseStatement, error) {
	resp, err := c.Patch(ctx, "DeviceUseStatement", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceUseStatements(resp)
}

func (c *Client) PatchDeviceUseStatementByID(ctx context.Context, resource, id string, params Parameters, entity *models.DeviceUseStatement) (*models.DeviceUseStatement, error) {
	resp, err := c.PatchByID(ctx, "DeviceUseStatement", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceUseStatement(id, resp)
}

func (c *Client) DeleteDeviceUseStatement(ctx context.Context, resource string, params Parameters) ([]*models.DeviceUseStatement, error) {
	resp, err := c.Delete(ctx, "DeviceUseStatement", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceUseStatements(resp)
}

func (c *Client) DeleteDeviceUseStatementByID(ctx context.Context, resource, id string, params Parameters) (*models.DeviceUseStatement, error) {
	resp, err := c.DeleteByID(ctx, "DeviceUseStatement", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDeviceUseStatement(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// DiagnosticReport
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToDiagnosticReports(bundle *models.Bundle) ([]*models.DiagnosticReport, error) {
	var entities []*models.DiagnosticReport
	err := EnumBundleResources(bundle, "DiagnosticReport", func(resource ResourceData) error {
		var entity models.DiagnosticReport
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToDiagnosticReports(resp *FhirResponse) ([]*models.DiagnosticReport, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToDiagnosticReports(resp.Bundle)
	case "DiagnosticReport":
		var entity models.DiagnosticReport
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.DiagnosticReport{&entity}, nil
	}
	return nil, nil
}

func fhirRespToDiagnosticReport(id string, resp *FhirResponse) (*models.DiagnosticReport, error) {
	entities, err := fhirRespToDiagnosticReports(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "DiagnosticReport", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get DiagnosticReport.
func (c *Client) GetDiagnosticReport(ctx context.Context, params Parameters) ([]*models.DiagnosticReport, error) {
	resp, err := c.Get(ctx, "DiagnosticReport", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDiagnosticReports(resp)
}

// Get DiagnosticReport by ID.
func (c *Client) GetDiagnosticReportByID(ctx context.Context, id string, params Parameters) (*models.DiagnosticReport, error) {
	resp, err := c.GetByID(ctx, "DiagnosticReport", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDiagnosticReport(id, resp)
}

func (c *Client) CreateDiagnosticReport(ctx context.Context, resource string, params Parameters, entity *models.DiagnosticReport) (*models.DiagnosticReport, error) {
	resp, err := c.Create(ctx, "DiagnosticReport", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDiagnosticReport(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDiagnosticReport(ctx context.Context, resource string, params Parameters, entity *models.DiagnosticReport) (*models.DiagnosticReport, error) {
	resp, err := c.Update(ctx, "DiagnosticReport", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDiagnosticReport(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDiagnosticReportByID(ctx context.Context, resource, id string, params Parameters, entity *models.DiagnosticReport) (*models.DiagnosticReport, error) {
	resp, err := c.UpdateByID(ctx, "DiagnosticReport", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDiagnosticReport(id, resp)
}

func (c *Client) PatchDiagnosticReport(ctx context.Context, resource string, params Parameters, entity *models.DiagnosticReport) ([]*models.DiagnosticReport, error) {
	resp, err := c.Patch(ctx, "DiagnosticReport", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDiagnosticReports(resp)
}

func (c *Client) PatchDiagnosticReportByID(ctx context.Context, resource, id string, params Parameters, entity *models.DiagnosticReport) (*models.DiagnosticReport, error) {
	resp, err := c.PatchByID(ctx, "DiagnosticReport", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDiagnosticReport(id, resp)
}

func (c *Client) DeleteDiagnosticReport(ctx context.Context, resource string, params Parameters) ([]*models.DiagnosticReport, error) {
	resp, err := c.Delete(ctx, "DiagnosticReport", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDiagnosticReports(resp)
}

func (c *Client) DeleteDiagnosticReportByID(ctx context.Context, resource, id string, params Parameters) (*models.DiagnosticReport, error) {
	resp, err := c.DeleteByID(ctx, "DiagnosticReport", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDiagnosticReport(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// DocumentManifest
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToDocumentManifests(bundle *models.Bundle) ([]*models.DocumentManifest, error) {
	var entities []*models.DocumentManifest
	err := EnumBundleResources(bundle, "DocumentManifest", func(resource ResourceData) error {
		var entity models.DocumentManifest
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToDocumentManifests(resp *FhirResponse) ([]*models.DocumentManifest, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToDocumentManifests(resp.Bundle)
	case "DocumentManifest":
		var entity models.DocumentManifest
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.DocumentManifest{&entity}, nil
	}
	return nil, nil
}

func fhirRespToDocumentManifest(id string, resp *FhirResponse) (*models.DocumentManifest, error) {
	entities, err := fhirRespToDocumentManifests(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "DocumentManifest", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get DocumentManifest.
func (c *Client) GetDocumentManifest(ctx context.Context, params Parameters) ([]*models.DocumentManifest, error) {
	resp, err := c.Get(ctx, "DocumentManifest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentManifests(resp)
}

// Get DocumentManifest by ID.
func (c *Client) GetDocumentManifestByID(ctx context.Context, id string, params Parameters) (*models.DocumentManifest, error) {
	resp, err := c.GetByID(ctx, "DocumentManifest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentManifest(id, resp)
}

func (c *Client) CreateDocumentManifest(ctx context.Context, resource string, params Parameters, entity *models.DocumentManifest) (*models.DocumentManifest, error) {
	resp, err := c.Create(ctx, "DocumentManifest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentManifest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDocumentManifest(ctx context.Context, resource string, params Parameters, entity *models.DocumentManifest) (*models.DocumentManifest, error) {
	resp, err := c.Update(ctx, "DocumentManifest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentManifest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDocumentManifestByID(ctx context.Context, resource, id string, params Parameters, entity *models.DocumentManifest) (*models.DocumentManifest, error) {
	resp, err := c.UpdateByID(ctx, "DocumentManifest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentManifest(id, resp)
}

func (c *Client) PatchDocumentManifest(ctx context.Context, resource string, params Parameters, entity *models.DocumentManifest) ([]*models.DocumentManifest, error) {
	resp, err := c.Patch(ctx, "DocumentManifest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentManifests(resp)
}

func (c *Client) PatchDocumentManifestByID(ctx context.Context, resource, id string, params Parameters, entity *models.DocumentManifest) (*models.DocumentManifest, error) {
	resp, err := c.PatchByID(ctx, "DocumentManifest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentManifest(id, resp)
}

func (c *Client) DeleteDocumentManifest(ctx context.Context, resource string, params Parameters) ([]*models.DocumentManifest, error) {
	resp, err := c.Delete(ctx, "DocumentManifest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentManifests(resp)
}

func (c *Client) DeleteDocumentManifestByID(ctx context.Context, resource, id string, params Parameters) (*models.DocumentManifest, error) {
	resp, err := c.DeleteByID(ctx, "DocumentManifest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentManifest(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// DocumentReference
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToDocumentReferences(bundle *models.Bundle) ([]*models.DocumentReference, error) {
	var entities []*models.DocumentReference
	err := EnumBundleResources(bundle, "DocumentReference", func(resource ResourceData) error {
		var entity models.DocumentReference
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToDocumentReferences(resp *FhirResponse) ([]*models.DocumentReference, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToDocumentReferences(resp.Bundle)
	case "DocumentReference":
		var entity models.DocumentReference
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.DocumentReference{&entity}, nil
	}
	return nil, nil
}

func fhirRespToDocumentReference(id string, resp *FhirResponse) (*models.DocumentReference, error) {
	entities, err := fhirRespToDocumentReferences(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "DocumentReference", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get DocumentReference.
func (c *Client) GetDocumentReference(ctx context.Context, params Parameters) ([]*models.DocumentReference, error) {
	resp, err := c.Get(ctx, "DocumentReference", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentReferences(resp)
}

// Get DocumentReference by ID.
func (c *Client) GetDocumentReferenceByID(ctx context.Context, id string, params Parameters) (*models.DocumentReference, error) {
	resp, err := c.GetByID(ctx, "DocumentReference", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentReference(id, resp)
}

func (c *Client) CreateDocumentReference(ctx context.Context, resource string, params Parameters, entity *models.DocumentReference) (*models.DocumentReference, error) {
	resp, err := c.Create(ctx, "DocumentReference", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentReference(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDocumentReference(ctx context.Context, resource string, params Parameters, entity *models.DocumentReference) (*models.DocumentReference, error) {
	resp, err := c.Update(ctx, "DocumentReference", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentReference(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDocumentReferenceByID(ctx context.Context, resource, id string, params Parameters, entity *models.DocumentReference) (*models.DocumentReference, error) {
	resp, err := c.UpdateByID(ctx, "DocumentReference", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentReference(id, resp)
}

func (c *Client) PatchDocumentReference(ctx context.Context, resource string, params Parameters, entity *models.DocumentReference) ([]*models.DocumentReference, error) {
	resp, err := c.Patch(ctx, "DocumentReference", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentReferences(resp)
}

func (c *Client) PatchDocumentReferenceByID(ctx context.Context, resource, id string, params Parameters, entity *models.DocumentReference) (*models.DocumentReference, error) {
	resp, err := c.PatchByID(ctx, "DocumentReference", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentReference(id, resp)
}

func (c *Client) DeleteDocumentReference(ctx context.Context, resource string, params Parameters) ([]*models.DocumentReference, error) {
	resp, err := c.Delete(ctx, "DocumentReference", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentReferences(resp)
}

func (c *Client) DeleteDocumentReferenceByID(ctx context.Context, resource, id string, params Parameters) (*models.DocumentReference, error) {
	resp, err := c.DeleteByID(ctx, "DocumentReference", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDocumentReference(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// DomainResource
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToDomainResources(bundle *models.Bundle) ([]*models.DomainResource, error) {
	var entities []*models.DomainResource
	err := EnumBundleResources(bundle, "DomainResource", func(resource ResourceData) error {
		var entity models.DomainResource
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToDomainResources(resp *FhirResponse) ([]*models.DomainResource, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToDomainResources(resp.Bundle)
	case "DomainResource":
		var entity models.DomainResource
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.DomainResource{&entity}, nil
	}
	return nil, nil
}

func fhirRespToDomainResource(id string, resp *FhirResponse) (*models.DomainResource, error) {
	entities, err := fhirRespToDomainResources(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "DomainResource", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get DomainResource.
func (c *Client) GetDomainResource(ctx context.Context, params Parameters) ([]*models.DomainResource, error) {
	resp, err := c.Get(ctx, "DomainResource", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDomainResources(resp)
}

// Get DomainResource by ID.
func (c *Client) GetDomainResourceByID(ctx context.Context, id string, params Parameters) (*models.DomainResource, error) {
	resp, err := c.GetByID(ctx, "DomainResource", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDomainResource(id, resp)
}

func (c *Client) CreateDomainResource(ctx context.Context, resource string, params Parameters, entity *models.DomainResource) (*models.DomainResource, error) {
	resp, err := c.Create(ctx, "DomainResource", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDomainResource(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDomainResource(ctx context.Context, resource string, params Parameters, entity *models.DomainResource) (*models.DomainResource, error) {
	resp, err := c.Update(ctx, "DomainResource", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDomainResource(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateDomainResourceByID(ctx context.Context, resource, id string, params Parameters, entity *models.DomainResource) (*models.DomainResource, error) {
	resp, err := c.UpdateByID(ctx, "DomainResource", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDomainResource(id, resp)
}

func (c *Client) PatchDomainResource(ctx context.Context, resource string, params Parameters, entity *models.DomainResource) ([]*models.DomainResource, error) {
	resp, err := c.Patch(ctx, "DomainResource", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDomainResources(resp)
}

func (c *Client) PatchDomainResourceByID(ctx context.Context, resource, id string, params Parameters, entity *models.DomainResource) (*models.DomainResource, error) {
	resp, err := c.PatchByID(ctx, "DomainResource", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToDomainResource(id, resp)
}

func (c *Client) DeleteDomainResource(ctx context.Context, resource string, params Parameters) ([]*models.DomainResource, error) {
	resp, err := c.Delete(ctx, "DomainResource", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDomainResources(resp)
}

func (c *Client) DeleteDomainResourceByID(ctx context.Context, resource, id string, params Parameters) (*models.DomainResource, error) {
	resp, err := c.DeleteByID(ctx, "DomainResource", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToDomainResource(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// EffectEvidenceSynthesis
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToEffectEvidenceSynthesiss(bundle *models.Bundle) ([]*models.EffectEvidenceSynthesis, error) {
	var entities []*models.EffectEvidenceSynthesis
	err := EnumBundleResources(bundle, "EffectEvidenceSynthesis", func(resource ResourceData) error {
		var entity models.EffectEvidenceSynthesis
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToEffectEvidenceSynthesiss(resp *FhirResponse) ([]*models.EffectEvidenceSynthesis, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToEffectEvidenceSynthesiss(resp.Bundle)
	case "EffectEvidenceSynthesis":
		var entity models.EffectEvidenceSynthesis
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.EffectEvidenceSynthesis{&entity}, nil
	}
	return nil, nil
}

func fhirRespToEffectEvidenceSynthesis(id string, resp *FhirResponse) (*models.EffectEvidenceSynthesis, error) {
	entities, err := fhirRespToEffectEvidenceSynthesiss(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "EffectEvidenceSynthesis", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get EffectEvidenceSynthesis.
func (c *Client) GetEffectEvidenceSynthesis(ctx context.Context, params Parameters) ([]*models.EffectEvidenceSynthesis, error) {
	resp, err := c.Get(ctx, "EffectEvidenceSynthesis", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEffectEvidenceSynthesiss(resp)
}

// Get EffectEvidenceSynthesis by ID.
func (c *Client) GetEffectEvidenceSynthesisByID(ctx context.Context, id string, params Parameters) (*models.EffectEvidenceSynthesis, error) {
	resp, err := c.GetByID(ctx, "EffectEvidenceSynthesis", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEffectEvidenceSynthesis(id, resp)
}

func (c *Client) CreateEffectEvidenceSynthesis(ctx context.Context, resource string, params Parameters, entity *models.EffectEvidenceSynthesis) (*models.EffectEvidenceSynthesis, error) {
	resp, err := c.Create(ctx, "EffectEvidenceSynthesis", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEffectEvidenceSynthesis(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEffectEvidenceSynthesis(ctx context.Context, resource string, params Parameters, entity *models.EffectEvidenceSynthesis) (*models.EffectEvidenceSynthesis, error) {
	resp, err := c.Update(ctx, "EffectEvidenceSynthesis", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEffectEvidenceSynthesis(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEffectEvidenceSynthesisByID(ctx context.Context, resource, id string, params Parameters, entity *models.EffectEvidenceSynthesis) (*models.EffectEvidenceSynthesis, error) {
	resp, err := c.UpdateByID(ctx, "EffectEvidenceSynthesis", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEffectEvidenceSynthesis(id, resp)
}

func (c *Client) PatchEffectEvidenceSynthesis(ctx context.Context, resource string, params Parameters, entity *models.EffectEvidenceSynthesis) ([]*models.EffectEvidenceSynthesis, error) {
	resp, err := c.Patch(ctx, "EffectEvidenceSynthesis", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEffectEvidenceSynthesiss(resp)
}

func (c *Client) PatchEffectEvidenceSynthesisByID(ctx context.Context, resource, id string, params Parameters, entity *models.EffectEvidenceSynthesis) (*models.EffectEvidenceSynthesis, error) {
	resp, err := c.PatchByID(ctx, "EffectEvidenceSynthesis", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEffectEvidenceSynthesis(id, resp)
}

func (c *Client) DeleteEffectEvidenceSynthesis(ctx context.Context, resource string, params Parameters) ([]*models.EffectEvidenceSynthesis, error) {
	resp, err := c.Delete(ctx, "EffectEvidenceSynthesis", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEffectEvidenceSynthesiss(resp)
}

func (c *Client) DeleteEffectEvidenceSynthesisByID(ctx context.Context, resource, id string, params Parameters) (*models.EffectEvidenceSynthesis, error) {
	resp, err := c.DeleteByID(ctx, "EffectEvidenceSynthesis", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEffectEvidenceSynthesis(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Encounter
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToEncounters(bundle *models.Bundle) ([]*models.Encounter, error) {
	var entities []*models.Encounter
	err := EnumBundleResources(bundle, "Encounter", func(resource ResourceData) error {
		var entity models.Encounter
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToEncounters(resp *FhirResponse) ([]*models.Encounter, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToEncounters(resp.Bundle)
	case "Encounter":
		var entity models.Encounter
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Encounter{&entity}, nil
	}
	return nil, nil
}

func fhirRespToEncounter(id string, resp *FhirResponse) (*models.Encounter, error) {
	entities, err := fhirRespToEncounters(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Encounter", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Encounter.
func (c *Client) GetEncounter(ctx context.Context, params Parameters) ([]*models.Encounter, error) {
	resp, err := c.Get(ctx, "Encounter", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEncounters(resp)
}

// Get Encounter by ID.
func (c *Client) GetEncounterByID(ctx context.Context, id string, params Parameters) (*models.Encounter, error) {
	resp, err := c.GetByID(ctx, "Encounter", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEncounter(id, resp)
}

func (c *Client) CreateEncounter(ctx context.Context, resource string, params Parameters, entity *models.Encounter) (*models.Encounter, error) {
	resp, err := c.Create(ctx, "Encounter", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEncounter(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEncounter(ctx context.Context, resource string, params Parameters, entity *models.Encounter) (*models.Encounter, error) {
	resp, err := c.Update(ctx, "Encounter", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEncounter(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEncounterByID(ctx context.Context, resource, id string, params Parameters, entity *models.Encounter) (*models.Encounter, error) {
	resp, err := c.UpdateByID(ctx, "Encounter", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEncounter(id, resp)
}

func (c *Client) PatchEncounter(ctx context.Context, resource string, params Parameters, entity *models.Encounter) ([]*models.Encounter, error) {
	resp, err := c.Patch(ctx, "Encounter", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEncounters(resp)
}

func (c *Client) PatchEncounterByID(ctx context.Context, resource, id string, params Parameters, entity *models.Encounter) (*models.Encounter, error) {
	resp, err := c.PatchByID(ctx, "Encounter", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEncounter(id, resp)
}

func (c *Client) DeleteEncounter(ctx context.Context, resource string, params Parameters) ([]*models.Encounter, error) {
	resp, err := c.Delete(ctx, "Encounter", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEncounters(resp)
}

func (c *Client) DeleteEncounterByID(ctx context.Context, resource, id string, params Parameters) (*models.Encounter, error) {
	resp, err := c.DeleteByID(ctx, "Encounter", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEncounter(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Endpoint
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToEndpoints(bundle *models.Bundle) ([]*models.Endpoint, error) {
	var entities []*models.Endpoint
	err := EnumBundleResources(bundle, "Endpoint", func(resource ResourceData) error {
		var entity models.Endpoint
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToEndpoints(resp *FhirResponse) ([]*models.Endpoint, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToEndpoints(resp.Bundle)
	case "Endpoint":
		var entity models.Endpoint
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Endpoint{&entity}, nil
	}
	return nil, nil
}

func fhirRespToEndpoint(id string, resp *FhirResponse) (*models.Endpoint, error) {
	entities, err := fhirRespToEndpoints(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Endpoint", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Endpoint.
func (c *Client) GetEndpoint(ctx context.Context, params Parameters) ([]*models.Endpoint, error) {
	resp, err := c.Get(ctx, "Endpoint", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEndpoints(resp)
}

// Get Endpoint by ID.
func (c *Client) GetEndpointByID(ctx context.Context, id string, params Parameters) (*models.Endpoint, error) {
	resp, err := c.GetByID(ctx, "Endpoint", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEndpoint(id, resp)
}

func (c *Client) CreateEndpoint(ctx context.Context, resource string, params Parameters, entity *models.Endpoint) (*models.Endpoint, error) {
	resp, err := c.Create(ctx, "Endpoint", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEndpoint(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEndpoint(ctx context.Context, resource string, params Parameters, entity *models.Endpoint) (*models.Endpoint, error) {
	resp, err := c.Update(ctx, "Endpoint", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEndpoint(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEndpointByID(ctx context.Context, resource, id string, params Parameters, entity *models.Endpoint) (*models.Endpoint, error) {
	resp, err := c.UpdateByID(ctx, "Endpoint", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEndpoint(id, resp)
}

func (c *Client) PatchEndpoint(ctx context.Context, resource string, params Parameters, entity *models.Endpoint) ([]*models.Endpoint, error) {
	resp, err := c.Patch(ctx, "Endpoint", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEndpoints(resp)
}

func (c *Client) PatchEndpointByID(ctx context.Context, resource, id string, params Parameters, entity *models.Endpoint) (*models.Endpoint, error) {
	resp, err := c.PatchByID(ctx, "Endpoint", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEndpoint(id, resp)
}

func (c *Client) DeleteEndpoint(ctx context.Context, resource string, params Parameters) ([]*models.Endpoint, error) {
	resp, err := c.Delete(ctx, "Endpoint", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEndpoints(resp)
}

func (c *Client) DeleteEndpointByID(ctx context.Context, resource, id string, params Parameters) (*models.Endpoint, error) {
	resp, err := c.DeleteByID(ctx, "Endpoint", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEndpoint(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// EnrollmentRequest
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToEnrollmentRequests(bundle *models.Bundle) ([]*models.EnrollmentRequest, error) {
	var entities []*models.EnrollmentRequest
	err := EnumBundleResources(bundle, "EnrollmentRequest", func(resource ResourceData) error {
		var entity models.EnrollmentRequest
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToEnrollmentRequests(resp *FhirResponse) ([]*models.EnrollmentRequest, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToEnrollmentRequests(resp.Bundle)
	case "EnrollmentRequest":
		var entity models.EnrollmentRequest
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.EnrollmentRequest{&entity}, nil
	}
	return nil, nil
}

func fhirRespToEnrollmentRequest(id string, resp *FhirResponse) (*models.EnrollmentRequest, error) {
	entities, err := fhirRespToEnrollmentRequests(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "EnrollmentRequest", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get EnrollmentRequest.
func (c *Client) GetEnrollmentRequest(ctx context.Context, params Parameters) ([]*models.EnrollmentRequest, error) {
	resp, err := c.Get(ctx, "EnrollmentRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentRequests(resp)
}

// Get EnrollmentRequest by ID.
func (c *Client) GetEnrollmentRequestByID(ctx context.Context, id string, params Parameters) (*models.EnrollmentRequest, error) {
	resp, err := c.GetByID(ctx, "EnrollmentRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentRequest(id, resp)
}

func (c *Client) CreateEnrollmentRequest(ctx context.Context, resource string, params Parameters, entity *models.EnrollmentRequest) (*models.EnrollmentRequest, error) {
	resp, err := c.Create(ctx, "EnrollmentRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEnrollmentRequest(ctx context.Context, resource string, params Parameters, entity *models.EnrollmentRequest) (*models.EnrollmentRequest, error) {
	resp, err := c.Update(ctx, "EnrollmentRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEnrollmentRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.EnrollmentRequest) (*models.EnrollmentRequest, error) {
	resp, err := c.UpdateByID(ctx, "EnrollmentRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentRequest(id, resp)
}

func (c *Client) PatchEnrollmentRequest(ctx context.Context, resource string, params Parameters, entity *models.EnrollmentRequest) ([]*models.EnrollmentRequest, error) {
	resp, err := c.Patch(ctx, "EnrollmentRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentRequests(resp)
}

func (c *Client) PatchEnrollmentRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.EnrollmentRequest) (*models.EnrollmentRequest, error) {
	resp, err := c.PatchByID(ctx, "EnrollmentRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentRequest(id, resp)
}

func (c *Client) DeleteEnrollmentRequest(ctx context.Context, resource string, params Parameters) ([]*models.EnrollmentRequest, error) {
	resp, err := c.Delete(ctx, "EnrollmentRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentRequests(resp)
}

func (c *Client) DeleteEnrollmentRequestByID(ctx context.Context, resource, id string, params Parameters) (*models.EnrollmentRequest, error) {
	resp, err := c.DeleteByID(ctx, "EnrollmentRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentRequest(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// EnrollmentResponse
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToEnrollmentResponses(bundle *models.Bundle) ([]*models.EnrollmentResponse, error) {
	var entities []*models.EnrollmentResponse
	err := EnumBundleResources(bundle, "EnrollmentResponse", func(resource ResourceData) error {
		var entity models.EnrollmentResponse
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToEnrollmentResponses(resp *FhirResponse) ([]*models.EnrollmentResponse, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToEnrollmentResponses(resp.Bundle)
	case "EnrollmentResponse":
		var entity models.EnrollmentResponse
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.EnrollmentResponse{&entity}, nil
	}
	return nil, nil
}

func fhirRespToEnrollmentResponse(id string, resp *FhirResponse) (*models.EnrollmentResponse, error) {
	entities, err := fhirRespToEnrollmentResponses(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "EnrollmentResponse", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get EnrollmentResponse.
func (c *Client) GetEnrollmentResponse(ctx context.Context, params Parameters) ([]*models.EnrollmentResponse, error) {
	resp, err := c.Get(ctx, "EnrollmentResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentResponses(resp)
}

// Get EnrollmentResponse by ID.
func (c *Client) GetEnrollmentResponseByID(ctx context.Context, id string, params Parameters) (*models.EnrollmentResponse, error) {
	resp, err := c.GetByID(ctx, "EnrollmentResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentResponse(id, resp)
}

func (c *Client) CreateEnrollmentResponse(ctx context.Context, resource string, params Parameters, entity *models.EnrollmentResponse) (*models.EnrollmentResponse, error) {
	resp, err := c.Create(ctx, "EnrollmentResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEnrollmentResponse(ctx context.Context, resource string, params Parameters, entity *models.EnrollmentResponse) (*models.EnrollmentResponse, error) {
	resp, err := c.Update(ctx, "EnrollmentResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEnrollmentResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.EnrollmentResponse) (*models.EnrollmentResponse, error) {
	resp, err := c.UpdateByID(ctx, "EnrollmentResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentResponse(id, resp)
}

func (c *Client) PatchEnrollmentResponse(ctx context.Context, resource string, params Parameters, entity *models.EnrollmentResponse) ([]*models.EnrollmentResponse, error) {
	resp, err := c.Patch(ctx, "EnrollmentResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentResponses(resp)
}

func (c *Client) PatchEnrollmentResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.EnrollmentResponse) (*models.EnrollmentResponse, error) {
	resp, err := c.PatchByID(ctx, "EnrollmentResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentResponse(id, resp)
}

func (c *Client) DeleteEnrollmentResponse(ctx context.Context, resource string, params Parameters) ([]*models.EnrollmentResponse, error) {
	resp, err := c.Delete(ctx, "EnrollmentResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentResponses(resp)
}

func (c *Client) DeleteEnrollmentResponseByID(ctx context.Context, resource, id string, params Parameters) (*models.EnrollmentResponse, error) {
	resp, err := c.DeleteByID(ctx, "EnrollmentResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEnrollmentResponse(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// EpisodeOfCare
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToEpisodeOfCares(bundle *models.Bundle) ([]*models.EpisodeOfCare, error) {
	var entities []*models.EpisodeOfCare
	err := EnumBundleResources(bundle, "EpisodeOfCare", func(resource ResourceData) error {
		var entity models.EpisodeOfCare
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToEpisodeOfCares(resp *FhirResponse) ([]*models.EpisodeOfCare, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToEpisodeOfCares(resp.Bundle)
	case "EpisodeOfCare":
		var entity models.EpisodeOfCare
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.EpisodeOfCare{&entity}, nil
	}
	return nil, nil
}

func fhirRespToEpisodeOfCare(id string, resp *FhirResponse) (*models.EpisodeOfCare, error) {
	entities, err := fhirRespToEpisodeOfCares(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "EpisodeOfCare", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get EpisodeOfCare.
func (c *Client) GetEpisodeOfCare(ctx context.Context, params Parameters) ([]*models.EpisodeOfCare, error) {
	resp, err := c.Get(ctx, "EpisodeOfCare", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEpisodeOfCares(resp)
}

// Get EpisodeOfCare by ID.
func (c *Client) GetEpisodeOfCareByID(ctx context.Context, id string, params Parameters) (*models.EpisodeOfCare, error) {
	resp, err := c.GetByID(ctx, "EpisodeOfCare", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEpisodeOfCare(id, resp)
}

func (c *Client) CreateEpisodeOfCare(ctx context.Context, resource string, params Parameters, entity *models.EpisodeOfCare) (*models.EpisodeOfCare, error) {
	resp, err := c.Create(ctx, "EpisodeOfCare", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEpisodeOfCare(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEpisodeOfCare(ctx context.Context, resource string, params Parameters, entity *models.EpisodeOfCare) (*models.EpisodeOfCare, error) {
	resp, err := c.Update(ctx, "EpisodeOfCare", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEpisodeOfCare(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEpisodeOfCareByID(ctx context.Context, resource, id string, params Parameters, entity *models.EpisodeOfCare) (*models.EpisodeOfCare, error) {
	resp, err := c.UpdateByID(ctx, "EpisodeOfCare", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEpisodeOfCare(id, resp)
}

func (c *Client) PatchEpisodeOfCare(ctx context.Context, resource string, params Parameters, entity *models.EpisodeOfCare) ([]*models.EpisodeOfCare, error) {
	resp, err := c.Patch(ctx, "EpisodeOfCare", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEpisodeOfCares(resp)
}

func (c *Client) PatchEpisodeOfCareByID(ctx context.Context, resource, id string, params Parameters, entity *models.EpisodeOfCare) (*models.EpisodeOfCare, error) {
	resp, err := c.PatchByID(ctx, "EpisodeOfCare", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEpisodeOfCare(id, resp)
}

func (c *Client) DeleteEpisodeOfCare(ctx context.Context, resource string, params Parameters) ([]*models.EpisodeOfCare, error) {
	resp, err := c.Delete(ctx, "EpisodeOfCare", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEpisodeOfCares(resp)
}

func (c *Client) DeleteEpisodeOfCareByID(ctx context.Context, resource, id string, params Parameters) (*models.EpisodeOfCare, error) {
	resp, err := c.DeleteByID(ctx, "EpisodeOfCare", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEpisodeOfCare(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// EventDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToEventDefinitions(bundle *models.Bundle) ([]*models.EventDefinition, error) {
	var entities []*models.EventDefinition
	err := EnumBundleResources(bundle, "EventDefinition", func(resource ResourceData) error {
		var entity models.EventDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToEventDefinitions(resp *FhirResponse) ([]*models.EventDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToEventDefinitions(resp.Bundle)
	case "EventDefinition":
		var entity models.EventDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.EventDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToEventDefinition(id string, resp *FhirResponse) (*models.EventDefinition, error) {
	entities, err := fhirRespToEventDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "EventDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get EventDefinition.
func (c *Client) GetEventDefinition(ctx context.Context, params Parameters) ([]*models.EventDefinition, error) {
	resp, err := c.Get(ctx, "EventDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEventDefinitions(resp)
}

// Get EventDefinition by ID.
func (c *Client) GetEventDefinitionByID(ctx context.Context, id string, params Parameters) (*models.EventDefinition, error) {
	resp, err := c.GetByID(ctx, "EventDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEventDefinition(id, resp)
}

func (c *Client) CreateEventDefinition(ctx context.Context, resource string, params Parameters, entity *models.EventDefinition) (*models.EventDefinition, error) {
	resp, err := c.Create(ctx, "EventDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEventDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEventDefinition(ctx context.Context, resource string, params Parameters, entity *models.EventDefinition) (*models.EventDefinition, error) {
	resp, err := c.Update(ctx, "EventDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEventDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEventDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.EventDefinition) (*models.EventDefinition, error) {
	resp, err := c.UpdateByID(ctx, "EventDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEventDefinition(id, resp)
}

func (c *Client) PatchEventDefinition(ctx context.Context, resource string, params Parameters, entity *models.EventDefinition) ([]*models.EventDefinition, error) {
	resp, err := c.Patch(ctx, "EventDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEventDefinitions(resp)
}

func (c *Client) PatchEventDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.EventDefinition) (*models.EventDefinition, error) {
	resp, err := c.PatchByID(ctx, "EventDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEventDefinition(id, resp)
}

func (c *Client) DeleteEventDefinition(ctx context.Context, resource string, params Parameters) ([]*models.EventDefinition, error) {
	resp, err := c.Delete(ctx, "EventDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEventDefinitions(resp)
}

func (c *Client) DeleteEventDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.EventDefinition, error) {
	resp, err := c.DeleteByID(ctx, "EventDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEventDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Evidence
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToEvidences(bundle *models.Bundle) ([]*models.Evidence, error) {
	var entities []*models.Evidence
	err := EnumBundleResources(bundle, "Evidence", func(resource ResourceData) error {
		var entity models.Evidence
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToEvidences(resp *FhirResponse) ([]*models.Evidence, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToEvidences(resp.Bundle)
	case "Evidence":
		var entity models.Evidence
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Evidence{&entity}, nil
	}
	return nil, nil
}

func fhirRespToEvidence(id string, resp *FhirResponse) (*models.Evidence, error) {
	entities, err := fhirRespToEvidences(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Evidence", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Evidence.
func (c *Client) GetEvidence(ctx context.Context, params Parameters) ([]*models.Evidence, error) {
	resp, err := c.Get(ctx, "Evidence", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidences(resp)
}

// Get Evidence by ID.
func (c *Client) GetEvidenceByID(ctx context.Context, id string, params Parameters) (*models.Evidence, error) {
	resp, err := c.GetByID(ctx, "Evidence", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidence(id, resp)
}

func (c *Client) CreateEvidence(ctx context.Context, resource string, params Parameters, entity *models.Evidence) (*models.Evidence, error) {
	resp, err := c.Create(ctx, "Evidence", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidence(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEvidence(ctx context.Context, resource string, params Parameters, entity *models.Evidence) (*models.Evidence, error) {
	resp, err := c.Update(ctx, "Evidence", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidence(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEvidenceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Evidence) (*models.Evidence, error) {
	resp, err := c.UpdateByID(ctx, "Evidence", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidence(id, resp)
}

func (c *Client) PatchEvidence(ctx context.Context, resource string, params Parameters, entity *models.Evidence) ([]*models.Evidence, error) {
	resp, err := c.Patch(ctx, "Evidence", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidences(resp)
}

func (c *Client) PatchEvidenceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Evidence) (*models.Evidence, error) {
	resp, err := c.PatchByID(ctx, "Evidence", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidence(id, resp)
}

func (c *Client) DeleteEvidence(ctx context.Context, resource string, params Parameters) ([]*models.Evidence, error) {
	resp, err := c.Delete(ctx, "Evidence", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidences(resp)
}

func (c *Client) DeleteEvidenceByID(ctx context.Context, resource, id string, params Parameters) (*models.Evidence, error) {
	resp, err := c.DeleteByID(ctx, "Evidence", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidence(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// EvidenceVariable
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToEvidenceVariables(bundle *models.Bundle) ([]*models.EvidenceVariable, error) {
	var entities []*models.EvidenceVariable
	err := EnumBundleResources(bundle, "EvidenceVariable", func(resource ResourceData) error {
		var entity models.EvidenceVariable
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToEvidenceVariables(resp *FhirResponse) ([]*models.EvidenceVariable, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToEvidenceVariables(resp.Bundle)
	case "EvidenceVariable":
		var entity models.EvidenceVariable
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.EvidenceVariable{&entity}, nil
	}
	return nil, nil
}

func fhirRespToEvidenceVariable(id string, resp *FhirResponse) (*models.EvidenceVariable, error) {
	entities, err := fhirRespToEvidenceVariables(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "EvidenceVariable", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get EvidenceVariable.
func (c *Client) GetEvidenceVariable(ctx context.Context, params Parameters) ([]*models.EvidenceVariable, error) {
	resp, err := c.Get(ctx, "EvidenceVariable", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidenceVariables(resp)
}

// Get EvidenceVariable by ID.
func (c *Client) GetEvidenceVariableByID(ctx context.Context, id string, params Parameters) (*models.EvidenceVariable, error) {
	resp, err := c.GetByID(ctx, "EvidenceVariable", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidenceVariable(id, resp)
}

func (c *Client) CreateEvidenceVariable(ctx context.Context, resource string, params Parameters, entity *models.EvidenceVariable) (*models.EvidenceVariable, error) {
	resp, err := c.Create(ctx, "EvidenceVariable", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidenceVariable(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEvidenceVariable(ctx context.Context, resource string, params Parameters, entity *models.EvidenceVariable) (*models.EvidenceVariable, error) {
	resp, err := c.Update(ctx, "EvidenceVariable", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidenceVariable(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateEvidenceVariableByID(ctx context.Context, resource, id string, params Parameters, entity *models.EvidenceVariable) (*models.EvidenceVariable, error) {
	resp, err := c.UpdateByID(ctx, "EvidenceVariable", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidenceVariable(id, resp)
}

func (c *Client) PatchEvidenceVariable(ctx context.Context, resource string, params Parameters, entity *models.EvidenceVariable) ([]*models.EvidenceVariable, error) {
	resp, err := c.Patch(ctx, "EvidenceVariable", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidenceVariables(resp)
}

func (c *Client) PatchEvidenceVariableByID(ctx context.Context, resource, id string, params Parameters, entity *models.EvidenceVariable) (*models.EvidenceVariable, error) {
	resp, err := c.PatchByID(ctx, "EvidenceVariable", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidenceVariable(id, resp)
}

func (c *Client) DeleteEvidenceVariable(ctx context.Context, resource string, params Parameters) ([]*models.EvidenceVariable, error) {
	resp, err := c.Delete(ctx, "EvidenceVariable", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidenceVariables(resp)
}

func (c *Client) DeleteEvidenceVariableByID(ctx context.Context, resource, id string, params Parameters) (*models.EvidenceVariable, error) {
	resp, err := c.DeleteByID(ctx, "EvidenceVariable", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToEvidenceVariable(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ExampleScenario
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToExampleScenarios(bundle *models.Bundle) ([]*models.ExampleScenario, error) {
	var entities []*models.ExampleScenario
	err := EnumBundleResources(bundle, "ExampleScenario", func(resource ResourceData) error {
		var entity models.ExampleScenario
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToExampleScenarios(resp *FhirResponse) ([]*models.ExampleScenario, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToExampleScenarios(resp.Bundle)
	case "ExampleScenario":
		var entity models.ExampleScenario
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ExampleScenario{&entity}, nil
	}
	return nil, nil
}

func fhirRespToExampleScenario(id string, resp *FhirResponse) (*models.ExampleScenario, error) {
	entities, err := fhirRespToExampleScenarios(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ExampleScenario", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ExampleScenario.
func (c *Client) GetExampleScenario(ctx context.Context, params Parameters) ([]*models.ExampleScenario, error) {
	resp, err := c.Get(ctx, "ExampleScenario", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToExampleScenarios(resp)
}

// Get ExampleScenario by ID.
func (c *Client) GetExampleScenarioByID(ctx context.Context, id string, params Parameters) (*models.ExampleScenario, error) {
	resp, err := c.GetByID(ctx, "ExampleScenario", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToExampleScenario(id, resp)
}

func (c *Client) CreateExampleScenario(ctx context.Context, resource string, params Parameters, entity *models.ExampleScenario) (*models.ExampleScenario, error) {
	resp, err := c.Create(ctx, "ExampleScenario", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToExampleScenario(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateExampleScenario(ctx context.Context, resource string, params Parameters, entity *models.ExampleScenario) (*models.ExampleScenario, error) {
	resp, err := c.Update(ctx, "ExampleScenario", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToExampleScenario(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateExampleScenarioByID(ctx context.Context, resource, id string, params Parameters, entity *models.ExampleScenario) (*models.ExampleScenario, error) {
	resp, err := c.UpdateByID(ctx, "ExampleScenario", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToExampleScenario(id, resp)
}

func (c *Client) PatchExampleScenario(ctx context.Context, resource string, params Parameters, entity *models.ExampleScenario) ([]*models.ExampleScenario, error) {
	resp, err := c.Patch(ctx, "ExampleScenario", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToExampleScenarios(resp)
}

func (c *Client) PatchExampleScenarioByID(ctx context.Context, resource, id string, params Parameters, entity *models.ExampleScenario) (*models.ExampleScenario, error) {
	resp, err := c.PatchByID(ctx, "ExampleScenario", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToExampleScenario(id, resp)
}

func (c *Client) DeleteExampleScenario(ctx context.Context, resource string, params Parameters) ([]*models.ExampleScenario, error) {
	resp, err := c.Delete(ctx, "ExampleScenario", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToExampleScenarios(resp)
}

func (c *Client) DeleteExampleScenarioByID(ctx context.Context, resource, id string, params Parameters) (*models.ExampleScenario, error) {
	resp, err := c.DeleteByID(ctx, "ExampleScenario", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToExampleScenario(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ExplanationOfBenefit
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToExplanationOfBenefits(bundle *models.Bundle) ([]*models.ExplanationOfBenefit, error) {
	var entities []*models.ExplanationOfBenefit
	err := EnumBundleResources(bundle, "ExplanationOfBenefit", func(resource ResourceData) error {
		var entity models.ExplanationOfBenefit
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToExplanationOfBenefits(resp *FhirResponse) ([]*models.ExplanationOfBenefit, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToExplanationOfBenefits(resp.Bundle)
	case "ExplanationOfBenefit":
		var entity models.ExplanationOfBenefit
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ExplanationOfBenefit{&entity}, nil
	}
	return nil, nil
}

func fhirRespToExplanationOfBenefit(id string, resp *FhirResponse) (*models.ExplanationOfBenefit, error) {
	entities, err := fhirRespToExplanationOfBenefits(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ExplanationOfBenefit", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ExplanationOfBenefit.
func (c *Client) GetExplanationOfBenefit(ctx context.Context, params Parameters) ([]*models.ExplanationOfBenefit, error) {
	resp, err := c.Get(ctx, "ExplanationOfBenefit", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToExplanationOfBenefits(resp)
}

// Get ExplanationOfBenefit by ID.
func (c *Client) GetExplanationOfBenefitByID(ctx context.Context, id string, params Parameters) (*models.ExplanationOfBenefit, error) {
	resp, err := c.GetByID(ctx, "ExplanationOfBenefit", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToExplanationOfBenefit(id, resp)
}

func (c *Client) CreateExplanationOfBenefit(ctx context.Context, resource string, params Parameters, entity *models.ExplanationOfBenefit) (*models.ExplanationOfBenefit, error) {
	resp, err := c.Create(ctx, "ExplanationOfBenefit", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToExplanationOfBenefit(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateExplanationOfBenefit(ctx context.Context, resource string, params Parameters, entity *models.ExplanationOfBenefit) (*models.ExplanationOfBenefit, error) {
	resp, err := c.Update(ctx, "ExplanationOfBenefit", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToExplanationOfBenefit(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateExplanationOfBenefitByID(ctx context.Context, resource, id string, params Parameters, entity *models.ExplanationOfBenefit) (*models.ExplanationOfBenefit, error) {
	resp, err := c.UpdateByID(ctx, "ExplanationOfBenefit", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToExplanationOfBenefit(id, resp)
}

func (c *Client) PatchExplanationOfBenefit(ctx context.Context, resource string, params Parameters, entity *models.ExplanationOfBenefit) ([]*models.ExplanationOfBenefit, error) {
	resp, err := c.Patch(ctx, "ExplanationOfBenefit", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToExplanationOfBenefits(resp)
}

func (c *Client) PatchExplanationOfBenefitByID(ctx context.Context, resource, id string, params Parameters, entity *models.ExplanationOfBenefit) (*models.ExplanationOfBenefit, error) {
	resp, err := c.PatchByID(ctx, "ExplanationOfBenefit", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToExplanationOfBenefit(id, resp)
}

func (c *Client) DeleteExplanationOfBenefit(ctx context.Context, resource string, params Parameters) ([]*models.ExplanationOfBenefit, error) {
	resp, err := c.Delete(ctx, "ExplanationOfBenefit", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToExplanationOfBenefits(resp)
}

func (c *Client) DeleteExplanationOfBenefitByID(ctx context.Context, resource, id string, params Parameters) (*models.ExplanationOfBenefit, error) {
	resp, err := c.DeleteByID(ctx, "ExplanationOfBenefit", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToExplanationOfBenefit(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// FamilyMemberHistory
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToFamilyMemberHistorys(bundle *models.Bundle) ([]*models.FamilyMemberHistory, error) {
	var entities []*models.FamilyMemberHistory
	err := EnumBundleResources(bundle, "FamilyMemberHistory", func(resource ResourceData) error {
		var entity models.FamilyMemberHistory
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToFamilyMemberHistorys(resp *FhirResponse) ([]*models.FamilyMemberHistory, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToFamilyMemberHistorys(resp.Bundle)
	case "FamilyMemberHistory":
		var entity models.FamilyMemberHistory
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.FamilyMemberHistory{&entity}, nil
	}
	return nil, nil
}

func fhirRespToFamilyMemberHistory(id string, resp *FhirResponse) (*models.FamilyMemberHistory, error) {
	entities, err := fhirRespToFamilyMemberHistorys(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "FamilyMemberHistory", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get FamilyMemberHistory.
func (c *Client) GetFamilyMemberHistory(ctx context.Context, params Parameters) ([]*models.FamilyMemberHistory, error) {
	resp, err := c.Get(ctx, "FamilyMemberHistory", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToFamilyMemberHistorys(resp)
}

// Get FamilyMemberHistory by ID.
func (c *Client) GetFamilyMemberHistoryByID(ctx context.Context, id string, params Parameters) (*models.FamilyMemberHistory, error) {
	resp, err := c.GetByID(ctx, "FamilyMemberHistory", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToFamilyMemberHistory(id, resp)
}

func (c *Client) CreateFamilyMemberHistory(ctx context.Context, resource string, params Parameters, entity *models.FamilyMemberHistory) (*models.FamilyMemberHistory, error) {
	resp, err := c.Create(ctx, "FamilyMemberHistory", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToFamilyMemberHistory(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateFamilyMemberHistory(ctx context.Context, resource string, params Parameters, entity *models.FamilyMemberHistory) (*models.FamilyMemberHistory, error) {
	resp, err := c.Update(ctx, "FamilyMemberHistory", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToFamilyMemberHistory(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateFamilyMemberHistoryByID(ctx context.Context, resource, id string, params Parameters, entity *models.FamilyMemberHistory) (*models.FamilyMemberHistory, error) {
	resp, err := c.UpdateByID(ctx, "FamilyMemberHistory", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToFamilyMemberHistory(id, resp)
}

func (c *Client) PatchFamilyMemberHistory(ctx context.Context, resource string, params Parameters, entity *models.FamilyMemberHistory) ([]*models.FamilyMemberHistory, error) {
	resp, err := c.Patch(ctx, "FamilyMemberHistory", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToFamilyMemberHistorys(resp)
}

func (c *Client) PatchFamilyMemberHistoryByID(ctx context.Context, resource, id string, params Parameters, entity *models.FamilyMemberHistory) (*models.FamilyMemberHistory, error) {
	resp, err := c.PatchByID(ctx, "FamilyMemberHistory", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToFamilyMemberHistory(id, resp)
}

func (c *Client) DeleteFamilyMemberHistory(ctx context.Context, resource string, params Parameters) ([]*models.FamilyMemberHistory, error) {
	resp, err := c.Delete(ctx, "FamilyMemberHistory", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToFamilyMemberHistorys(resp)
}

func (c *Client) DeleteFamilyMemberHistoryByID(ctx context.Context, resource, id string, params Parameters) (*models.FamilyMemberHistory, error) {
	resp, err := c.DeleteByID(ctx, "FamilyMemberHistory", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToFamilyMemberHistory(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Flag
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToFlags(bundle *models.Bundle) ([]*models.Flag, error) {
	var entities []*models.Flag
	err := EnumBundleResources(bundle, "Flag", func(resource ResourceData) error {
		var entity models.Flag
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToFlags(resp *FhirResponse) ([]*models.Flag, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToFlags(resp.Bundle)
	case "Flag":
		var entity models.Flag
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Flag{&entity}, nil
	}
	return nil, nil
}

func fhirRespToFlag(id string, resp *FhirResponse) (*models.Flag, error) {
	entities, err := fhirRespToFlags(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Flag", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Flag.
func (c *Client) GetFlag(ctx context.Context, params Parameters) ([]*models.Flag, error) {
	resp, err := c.Get(ctx, "Flag", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToFlags(resp)
}

// Get Flag by ID.
func (c *Client) GetFlagByID(ctx context.Context, id string, params Parameters) (*models.Flag, error) {
	resp, err := c.GetByID(ctx, "Flag", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToFlag(id, resp)
}

func (c *Client) CreateFlag(ctx context.Context, resource string, params Parameters, entity *models.Flag) (*models.Flag, error) {
	resp, err := c.Create(ctx, "Flag", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToFlag(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateFlag(ctx context.Context, resource string, params Parameters, entity *models.Flag) (*models.Flag, error) {
	resp, err := c.Update(ctx, "Flag", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToFlag(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateFlagByID(ctx context.Context, resource, id string, params Parameters, entity *models.Flag) (*models.Flag, error) {
	resp, err := c.UpdateByID(ctx, "Flag", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToFlag(id, resp)
}

func (c *Client) PatchFlag(ctx context.Context, resource string, params Parameters, entity *models.Flag) ([]*models.Flag, error) {
	resp, err := c.Patch(ctx, "Flag", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToFlags(resp)
}

func (c *Client) PatchFlagByID(ctx context.Context, resource, id string, params Parameters, entity *models.Flag) (*models.Flag, error) {
	resp, err := c.PatchByID(ctx, "Flag", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToFlag(id, resp)
}

func (c *Client) DeleteFlag(ctx context.Context, resource string, params Parameters) ([]*models.Flag, error) {
	resp, err := c.Delete(ctx, "Flag", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToFlags(resp)
}

func (c *Client) DeleteFlagByID(ctx context.Context, resource, id string, params Parameters) (*models.Flag, error) {
	resp, err := c.DeleteByID(ctx, "Flag", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToFlag(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Goal
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToGoals(bundle *models.Bundle) ([]*models.Goal, error) {
	var entities []*models.Goal
	err := EnumBundleResources(bundle, "Goal", func(resource ResourceData) error {
		var entity models.Goal
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToGoals(resp *FhirResponse) ([]*models.Goal, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToGoals(resp.Bundle)
	case "Goal":
		var entity models.Goal
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Goal{&entity}, nil
	}
	return nil, nil
}

func fhirRespToGoal(id string, resp *FhirResponse) (*models.Goal, error) {
	entities, err := fhirRespToGoals(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Goal", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Goal.
func (c *Client) GetGoal(ctx context.Context, params Parameters) ([]*models.Goal, error) {
	resp, err := c.Get(ctx, "Goal", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGoals(resp)
}

// Get Goal by ID.
func (c *Client) GetGoalByID(ctx context.Context, id string, params Parameters) (*models.Goal, error) {
	resp, err := c.GetByID(ctx, "Goal", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGoal(id, resp)
}

func (c *Client) CreateGoal(ctx context.Context, resource string, params Parameters, entity *models.Goal) (*models.Goal, error) {
	resp, err := c.Create(ctx, "Goal", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGoal(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateGoal(ctx context.Context, resource string, params Parameters, entity *models.Goal) (*models.Goal, error) {
	resp, err := c.Update(ctx, "Goal", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGoal(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateGoalByID(ctx context.Context, resource, id string, params Parameters, entity *models.Goal) (*models.Goal, error) {
	resp, err := c.UpdateByID(ctx, "Goal", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGoal(id, resp)
}

func (c *Client) PatchGoal(ctx context.Context, resource string, params Parameters, entity *models.Goal) ([]*models.Goal, error) {
	resp, err := c.Patch(ctx, "Goal", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGoals(resp)
}

func (c *Client) PatchGoalByID(ctx context.Context, resource, id string, params Parameters, entity *models.Goal) (*models.Goal, error) {
	resp, err := c.PatchByID(ctx, "Goal", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGoal(id, resp)
}

func (c *Client) DeleteGoal(ctx context.Context, resource string, params Parameters) ([]*models.Goal, error) {
	resp, err := c.Delete(ctx, "Goal", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGoals(resp)
}

func (c *Client) DeleteGoalByID(ctx context.Context, resource, id string, params Parameters) (*models.Goal, error) {
	resp, err := c.DeleteByID(ctx, "Goal", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGoal(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// GraphDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToGraphDefinitions(bundle *models.Bundle) ([]*models.GraphDefinition, error) {
	var entities []*models.GraphDefinition
	err := EnumBundleResources(bundle, "GraphDefinition", func(resource ResourceData) error {
		var entity models.GraphDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToGraphDefinitions(resp *FhirResponse) ([]*models.GraphDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToGraphDefinitions(resp.Bundle)
	case "GraphDefinition":
		var entity models.GraphDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.GraphDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToGraphDefinition(id string, resp *FhirResponse) (*models.GraphDefinition, error) {
	entities, err := fhirRespToGraphDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "GraphDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get GraphDefinition.
func (c *Client) GetGraphDefinition(ctx context.Context, params Parameters) ([]*models.GraphDefinition, error) {
	resp, err := c.Get(ctx, "GraphDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGraphDefinitions(resp)
}

// Get GraphDefinition by ID.
func (c *Client) GetGraphDefinitionByID(ctx context.Context, id string, params Parameters) (*models.GraphDefinition, error) {
	resp, err := c.GetByID(ctx, "GraphDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGraphDefinition(id, resp)
}

func (c *Client) CreateGraphDefinition(ctx context.Context, resource string, params Parameters, entity *models.GraphDefinition) (*models.GraphDefinition, error) {
	resp, err := c.Create(ctx, "GraphDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGraphDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateGraphDefinition(ctx context.Context, resource string, params Parameters, entity *models.GraphDefinition) (*models.GraphDefinition, error) {
	resp, err := c.Update(ctx, "GraphDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGraphDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateGraphDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.GraphDefinition) (*models.GraphDefinition, error) {
	resp, err := c.UpdateByID(ctx, "GraphDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGraphDefinition(id, resp)
}

func (c *Client) PatchGraphDefinition(ctx context.Context, resource string, params Parameters, entity *models.GraphDefinition) ([]*models.GraphDefinition, error) {
	resp, err := c.Patch(ctx, "GraphDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGraphDefinitions(resp)
}

func (c *Client) PatchGraphDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.GraphDefinition) (*models.GraphDefinition, error) {
	resp, err := c.PatchByID(ctx, "GraphDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGraphDefinition(id, resp)
}

func (c *Client) DeleteGraphDefinition(ctx context.Context, resource string, params Parameters) ([]*models.GraphDefinition, error) {
	resp, err := c.Delete(ctx, "GraphDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGraphDefinitions(resp)
}

func (c *Client) DeleteGraphDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.GraphDefinition, error) {
	resp, err := c.DeleteByID(ctx, "GraphDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGraphDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Group
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToGroups(bundle *models.Bundle) ([]*models.Group, error) {
	var entities []*models.Group
	err := EnumBundleResources(bundle, "Group", func(resource ResourceData) error {
		var entity models.Group
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToGroups(resp *FhirResponse) ([]*models.Group, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToGroups(resp.Bundle)
	case "Group":
		var entity models.Group
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Group{&entity}, nil
	}
	return nil, nil
}

func fhirRespToGroup(id string, resp *FhirResponse) (*models.Group, error) {
	entities, err := fhirRespToGroups(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Group", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Group.
func (c *Client) GetGroup(ctx context.Context, params Parameters) ([]*models.Group, error) {
	resp, err := c.Get(ctx, "Group", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGroups(resp)
}

// Get Group by ID.
func (c *Client) GetGroupByID(ctx context.Context, id string, params Parameters) (*models.Group, error) {
	resp, err := c.GetByID(ctx, "Group", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGroup(id, resp)
}

func (c *Client) CreateGroup(ctx context.Context, resource string, params Parameters, entity *models.Group) (*models.Group, error) {
	resp, err := c.Create(ctx, "Group", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGroup(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateGroup(ctx context.Context, resource string, params Parameters, entity *models.Group) (*models.Group, error) {
	resp, err := c.Update(ctx, "Group", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGroup(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateGroupByID(ctx context.Context, resource, id string, params Parameters, entity *models.Group) (*models.Group, error) {
	resp, err := c.UpdateByID(ctx, "Group", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGroup(id, resp)
}

func (c *Client) PatchGroup(ctx context.Context, resource string, params Parameters, entity *models.Group) ([]*models.Group, error) {
	resp, err := c.Patch(ctx, "Group", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGroups(resp)
}

func (c *Client) PatchGroupByID(ctx context.Context, resource, id string, params Parameters, entity *models.Group) (*models.Group, error) {
	resp, err := c.PatchByID(ctx, "Group", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGroup(id, resp)
}

func (c *Client) DeleteGroup(ctx context.Context, resource string, params Parameters) ([]*models.Group, error) {
	resp, err := c.Delete(ctx, "Group", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGroups(resp)
}

func (c *Client) DeleteGroupByID(ctx context.Context, resource, id string, params Parameters) (*models.Group, error) {
	resp, err := c.DeleteByID(ctx, "Group", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGroup(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// GuidanceResponse
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToGuidanceResponses(bundle *models.Bundle) ([]*models.GuidanceResponse, error) {
	var entities []*models.GuidanceResponse
	err := EnumBundleResources(bundle, "GuidanceResponse", func(resource ResourceData) error {
		var entity models.GuidanceResponse
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToGuidanceResponses(resp *FhirResponse) ([]*models.GuidanceResponse, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToGuidanceResponses(resp.Bundle)
	case "GuidanceResponse":
		var entity models.GuidanceResponse
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.GuidanceResponse{&entity}, nil
	}
	return nil, nil
}

func fhirRespToGuidanceResponse(id string, resp *FhirResponse) (*models.GuidanceResponse, error) {
	entities, err := fhirRespToGuidanceResponses(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "GuidanceResponse", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get GuidanceResponse.
func (c *Client) GetGuidanceResponse(ctx context.Context, params Parameters) ([]*models.GuidanceResponse, error) {
	resp, err := c.Get(ctx, "GuidanceResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGuidanceResponses(resp)
}

// Get GuidanceResponse by ID.
func (c *Client) GetGuidanceResponseByID(ctx context.Context, id string, params Parameters) (*models.GuidanceResponse, error) {
	resp, err := c.GetByID(ctx, "GuidanceResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGuidanceResponse(id, resp)
}

func (c *Client) CreateGuidanceResponse(ctx context.Context, resource string, params Parameters, entity *models.GuidanceResponse) (*models.GuidanceResponse, error) {
	resp, err := c.Create(ctx, "GuidanceResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGuidanceResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateGuidanceResponse(ctx context.Context, resource string, params Parameters, entity *models.GuidanceResponse) (*models.GuidanceResponse, error) {
	resp, err := c.Update(ctx, "GuidanceResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGuidanceResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateGuidanceResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.GuidanceResponse) (*models.GuidanceResponse, error) {
	resp, err := c.UpdateByID(ctx, "GuidanceResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGuidanceResponse(id, resp)
}

func (c *Client) PatchGuidanceResponse(ctx context.Context, resource string, params Parameters, entity *models.GuidanceResponse) ([]*models.GuidanceResponse, error) {
	resp, err := c.Patch(ctx, "GuidanceResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGuidanceResponses(resp)
}

func (c *Client) PatchGuidanceResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.GuidanceResponse) (*models.GuidanceResponse, error) {
	resp, err := c.PatchByID(ctx, "GuidanceResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToGuidanceResponse(id, resp)
}

func (c *Client) DeleteGuidanceResponse(ctx context.Context, resource string, params Parameters) ([]*models.GuidanceResponse, error) {
	resp, err := c.Delete(ctx, "GuidanceResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGuidanceResponses(resp)
}

func (c *Client) DeleteGuidanceResponseByID(ctx context.Context, resource, id string, params Parameters) (*models.GuidanceResponse, error) {
	resp, err := c.DeleteByID(ctx, "GuidanceResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToGuidanceResponse(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// HealthcareService
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToHealthcareServices(bundle *models.Bundle) ([]*models.HealthcareService, error) {
	var entities []*models.HealthcareService
	err := EnumBundleResources(bundle, "HealthcareService", func(resource ResourceData) error {
		var entity models.HealthcareService
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToHealthcareServices(resp *FhirResponse) ([]*models.HealthcareService, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToHealthcareServices(resp.Bundle)
	case "HealthcareService":
		var entity models.HealthcareService
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.HealthcareService{&entity}, nil
	}
	return nil, nil
}

func fhirRespToHealthcareService(id string, resp *FhirResponse) (*models.HealthcareService, error) {
	entities, err := fhirRespToHealthcareServices(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "HealthcareService", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get HealthcareService.
func (c *Client) GetHealthcareService(ctx context.Context, params Parameters) ([]*models.HealthcareService, error) {
	resp, err := c.Get(ctx, "HealthcareService", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToHealthcareServices(resp)
}

// Get HealthcareService by ID.
func (c *Client) GetHealthcareServiceByID(ctx context.Context, id string, params Parameters) (*models.HealthcareService, error) {
	resp, err := c.GetByID(ctx, "HealthcareService", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToHealthcareService(id, resp)
}

func (c *Client) CreateHealthcareService(ctx context.Context, resource string, params Parameters, entity *models.HealthcareService) (*models.HealthcareService, error) {
	resp, err := c.Create(ctx, "HealthcareService", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToHealthcareService(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateHealthcareService(ctx context.Context, resource string, params Parameters, entity *models.HealthcareService) (*models.HealthcareService, error) {
	resp, err := c.Update(ctx, "HealthcareService", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToHealthcareService(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateHealthcareServiceByID(ctx context.Context, resource, id string, params Parameters, entity *models.HealthcareService) (*models.HealthcareService, error) {
	resp, err := c.UpdateByID(ctx, "HealthcareService", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToHealthcareService(id, resp)
}

func (c *Client) PatchHealthcareService(ctx context.Context, resource string, params Parameters, entity *models.HealthcareService) ([]*models.HealthcareService, error) {
	resp, err := c.Patch(ctx, "HealthcareService", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToHealthcareServices(resp)
}

func (c *Client) PatchHealthcareServiceByID(ctx context.Context, resource, id string, params Parameters, entity *models.HealthcareService) (*models.HealthcareService, error) {
	resp, err := c.PatchByID(ctx, "HealthcareService", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToHealthcareService(id, resp)
}

func (c *Client) DeleteHealthcareService(ctx context.Context, resource string, params Parameters) ([]*models.HealthcareService, error) {
	resp, err := c.Delete(ctx, "HealthcareService", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToHealthcareServices(resp)
}

func (c *Client) DeleteHealthcareServiceByID(ctx context.Context, resource, id string, params Parameters) (*models.HealthcareService, error) {
	resp, err := c.DeleteByID(ctx, "HealthcareService", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToHealthcareService(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ImagingStudy
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToImagingStudys(bundle *models.Bundle) ([]*models.ImagingStudy, error) {
	var entities []*models.ImagingStudy
	err := EnumBundleResources(bundle, "ImagingStudy", func(resource ResourceData) error {
		var entity models.ImagingStudy
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToImagingStudys(resp *FhirResponse) ([]*models.ImagingStudy, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToImagingStudys(resp.Bundle)
	case "ImagingStudy":
		var entity models.ImagingStudy
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ImagingStudy{&entity}, nil
	}
	return nil, nil
}

func fhirRespToImagingStudy(id string, resp *FhirResponse) (*models.ImagingStudy, error) {
	entities, err := fhirRespToImagingStudys(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ImagingStudy", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ImagingStudy.
func (c *Client) GetImagingStudy(ctx context.Context, params Parameters) ([]*models.ImagingStudy, error) {
	resp, err := c.Get(ctx, "ImagingStudy", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImagingStudys(resp)
}

// Get ImagingStudy by ID.
func (c *Client) GetImagingStudyByID(ctx context.Context, id string, params Parameters) (*models.ImagingStudy, error) {
	resp, err := c.GetByID(ctx, "ImagingStudy", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImagingStudy(id, resp)
}

func (c *Client) CreateImagingStudy(ctx context.Context, resource string, params Parameters, entity *models.ImagingStudy) (*models.ImagingStudy, error) {
	resp, err := c.Create(ctx, "ImagingStudy", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImagingStudy(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateImagingStudy(ctx context.Context, resource string, params Parameters, entity *models.ImagingStudy) (*models.ImagingStudy, error) {
	resp, err := c.Update(ctx, "ImagingStudy", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImagingStudy(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateImagingStudyByID(ctx context.Context, resource, id string, params Parameters, entity *models.ImagingStudy) (*models.ImagingStudy, error) {
	resp, err := c.UpdateByID(ctx, "ImagingStudy", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImagingStudy(id, resp)
}

func (c *Client) PatchImagingStudy(ctx context.Context, resource string, params Parameters, entity *models.ImagingStudy) ([]*models.ImagingStudy, error) {
	resp, err := c.Patch(ctx, "ImagingStudy", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImagingStudys(resp)
}

func (c *Client) PatchImagingStudyByID(ctx context.Context, resource, id string, params Parameters, entity *models.ImagingStudy) (*models.ImagingStudy, error) {
	resp, err := c.PatchByID(ctx, "ImagingStudy", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImagingStudy(id, resp)
}

func (c *Client) DeleteImagingStudy(ctx context.Context, resource string, params Parameters) ([]*models.ImagingStudy, error) {
	resp, err := c.Delete(ctx, "ImagingStudy", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImagingStudys(resp)
}

func (c *Client) DeleteImagingStudyByID(ctx context.Context, resource, id string, params Parameters) (*models.ImagingStudy, error) {
	resp, err := c.DeleteByID(ctx, "ImagingStudy", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImagingStudy(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Immunization
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToImmunizations(bundle *models.Bundle) ([]*models.Immunization, error) {
	var entities []*models.Immunization
	err := EnumBundleResources(bundle, "Immunization", func(resource ResourceData) error {
		var entity models.Immunization
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToImmunizations(resp *FhirResponse) ([]*models.Immunization, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToImmunizations(resp.Bundle)
	case "Immunization":
		var entity models.Immunization
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Immunization{&entity}, nil
	}
	return nil, nil
}

func fhirRespToImmunization(id string, resp *FhirResponse) (*models.Immunization, error) {
	entities, err := fhirRespToImmunizations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Immunization", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Immunization.
func (c *Client) GetImmunization(ctx context.Context, params Parameters) ([]*models.Immunization, error) {
	resp, err := c.Get(ctx, "Immunization", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizations(resp)
}

// Get Immunization by ID.
func (c *Client) GetImmunizationByID(ctx context.Context, id string, params Parameters) (*models.Immunization, error) {
	resp, err := c.GetByID(ctx, "Immunization", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunization(id, resp)
}

func (c *Client) CreateImmunization(ctx context.Context, resource string, params Parameters, entity *models.Immunization) (*models.Immunization, error) {
	resp, err := c.Create(ctx, "Immunization", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunization(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateImmunization(ctx context.Context, resource string, params Parameters, entity *models.Immunization) (*models.Immunization, error) {
	resp, err := c.Update(ctx, "Immunization", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunization(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateImmunizationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Immunization) (*models.Immunization, error) {
	resp, err := c.UpdateByID(ctx, "Immunization", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunization(id, resp)
}

func (c *Client) PatchImmunization(ctx context.Context, resource string, params Parameters, entity *models.Immunization) ([]*models.Immunization, error) {
	resp, err := c.Patch(ctx, "Immunization", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizations(resp)
}

func (c *Client) PatchImmunizationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Immunization) (*models.Immunization, error) {
	resp, err := c.PatchByID(ctx, "Immunization", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunization(id, resp)
}

func (c *Client) DeleteImmunization(ctx context.Context, resource string, params Parameters) ([]*models.Immunization, error) {
	resp, err := c.Delete(ctx, "Immunization", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizations(resp)
}

func (c *Client) DeleteImmunizationByID(ctx context.Context, resource, id string, params Parameters) (*models.Immunization, error) {
	resp, err := c.DeleteByID(ctx, "Immunization", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunization(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ImmunizationEvaluation
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToImmunizationEvaluations(bundle *models.Bundle) ([]*models.ImmunizationEvaluation, error) {
	var entities []*models.ImmunizationEvaluation
	err := EnumBundleResources(bundle, "ImmunizationEvaluation", func(resource ResourceData) error {
		var entity models.ImmunizationEvaluation
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToImmunizationEvaluations(resp *FhirResponse) ([]*models.ImmunizationEvaluation, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToImmunizationEvaluations(resp.Bundle)
	case "ImmunizationEvaluation":
		var entity models.ImmunizationEvaluation
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ImmunizationEvaluation{&entity}, nil
	}
	return nil, nil
}

func fhirRespToImmunizationEvaluation(id string, resp *FhirResponse) (*models.ImmunizationEvaluation, error) {
	entities, err := fhirRespToImmunizationEvaluations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ImmunizationEvaluation", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ImmunizationEvaluation.
func (c *Client) GetImmunizationEvaluation(ctx context.Context, params Parameters) ([]*models.ImmunizationEvaluation, error) {
	resp, err := c.Get(ctx, "ImmunizationEvaluation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationEvaluations(resp)
}

// Get ImmunizationEvaluation by ID.
func (c *Client) GetImmunizationEvaluationByID(ctx context.Context, id string, params Parameters) (*models.ImmunizationEvaluation, error) {
	resp, err := c.GetByID(ctx, "ImmunizationEvaluation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationEvaluation(id, resp)
}

func (c *Client) CreateImmunizationEvaluation(ctx context.Context, resource string, params Parameters, entity *models.ImmunizationEvaluation) (*models.ImmunizationEvaluation, error) {
	resp, err := c.Create(ctx, "ImmunizationEvaluation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationEvaluation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateImmunizationEvaluation(ctx context.Context, resource string, params Parameters, entity *models.ImmunizationEvaluation) (*models.ImmunizationEvaluation, error) {
	resp, err := c.Update(ctx, "ImmunizationEvaluation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationEvaluation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateImmunizationEvaluationByID(ctx context.Context, resource, id string, params Parameters, entity *models.ImmunizationEvaluation) (*models.ImmunizationEvaluation, error) {
	resp, err := c.UpdateByID(ctx, "ImmunizationEvaluation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationEvaluation(id, resp)
}

func (c *Client) PatchImmunizationEvaluation(ctx context.Context, resource string, params Parameters, entity *models.ImmunizationEvaluation) ([]*models.ImmunizationEvaluation, error) {
	resp, err := c.Patch(ctx, "ImmunizationEvaluation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationEvaluations(resp)
}

func (c *Client) PatchImmunizationEvaluationByID(ctx context.Context, resource, id string, params Parameters, entity *models.ImmunizationEvaluation) (*models.ImmunizationEvaluation, error) {
	resp, err := c.PatchByID(ctx, "ImmunizationEvaluation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationEvaluation(id, resp)
}

func (c *Client) DeleteImmunizationEvaluation(ctx context.Context, resource string, params Parameters) ([]*models.ImmunizationEvaluation, error) {
	resp, err := c.Delete(ctx, "ImmunizationEvaluation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationEvaluations(resp)
}

func (c *Client) DeleteImmunizationEvaluationByID(ctx context.Context, resource, id string, params Parameters) (*models.ImmunizationEvaluation, error) {
	resp, err := c.DeleteByID(ctx, "ImmunizationEvaluation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationEvaluation(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ImmunizationRecommendation
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToImmunizationRecommendations(bundle *models.Bundle) ([]*models.ImmunizationRecommendation, error) {
	var entities []*models.ImmunizationRecommendation
	err := EnumBundleResources(bundle, "ImmunizationRecommendation", func(resource ResourceData) error {
		var entity models.ImmunizationRecommendation
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToImmunizationRecommendations(resp *FhirResponse) ([]*models.ImmunizationRecommendation, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToImmunizationRecommendations(resp.Bundle)
	case "ImmunizationRecommendation":
		var entity models.ImmunizationRecommendation
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ImmunizationRecommendation{&entity}, nil
	}
	return nil, nil
}

func fhirRespToImmunizationRecommendation(id string, resp *FhirResponse) (*models.ImmunizationRecommendation, error) {
	entities, err := fhirRespToImmunizationRecommendations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ImmunizationRecommendation", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ImmunizationRecommendation.
func (c *Client) GetImmunizationRecommendation(ctx context.Context, params Parameters) ([]*models.ImmunizationRecommendation, error) {
	resp, err := c.Get(ctx, "ImmunizationRecommendation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationRecommendations(resp)
}

// Get ImmunizationRecommendation by ID.
func (c *Client) GetImmunizationRecommendationByID(ctx context.Context, id string, params Parameters) (*models.ImmunizationRecommendation, error) {
	resp, err := c.GetByID(ctx, "ImmunizationRecommendation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationRecommendation(id, resp)
}

func (c *Client) CreateImmunizationRecommendation(ctx context.Context, resource string, params Parameters, entity *models.ImmunizationRecommendation) (*models.ImmunizationRecommendation, error) {
	resp, err := c.Create(ctx, "ImmunizationRecommendation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationRecommendation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateImmunizationRecommendation(ctx context.Context, resource string, params Parameters, entity *models.ImmunizationRecommendation) (*models.ImmunizationRecommendation, error) {
	resp, err := c.Update(ctx, "ImmunizationRecommendation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationRecommendation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateImmunizationRecommendationByID(ctx context.Context, resource, id string, params Parameters, entity *models.ImmunizationRecommendation) (*models.ImmunizationRecommendation, error) {
	resp, err := c.UpdateByID(ctx, "ImmunizationRecommendation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationRecommendation(id, resp)
}

func (c *Client) PatchImmunizationRecommendation(ctx context.Context, resource string, params Parameters, entity *models.ImmunizationRecommendation) ([]*models.ImmunizationRecommendation, error) {
	resp, err := c.Patch(ctx, "ImmunizationRecommendation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationRecommendations(resp)
}

func (c *Client) PatchImmunizationRecommendationByID(ctx context.Context, resource, id string, params Parameters, entity *models.ImmunizationRecommendation) (*models.ImmunizationRecommendation, error) {
	resp, err := c.PatchByID(ctx, "ImmunizationRecommendation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationRecommendation(id, resp)
}

func (c *Client) DeleteImmunizationRecommendation(ctx context.Context, resource string, params Parameters) ([]*models.ImmunizationRecommendation, error) {
	resp, err := c.Delete(ctx, "ImmunizationRecommendation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationRecommendations(resp)
}

func (c *Client) DeleteImmunizationRecommendationByID(ctx context.Context, resource, id string, params Parameters) (*models.ImmunizationRecommendation, error) {
	resp, err := c.DeleteByID(ctx, "ImmunizationRecommendation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImmunizationRecommendation(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ImplementationGuide
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToImplementationGuides(bundle *models.Bundle) ([]*models.ImplementationGuide, error) {
	var entities []*models.ImplementationGuide
	err := EnumBundleResources(bundle, "ImplementationGuide", func(resource ResourceData) error {
		var entity models.ImplementationGuide
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToImplementationGuides(resp *FhirResponse) ([]*models.ImplementationGuide, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToImplementationGuides(resp.Bundle)
	case "ImplementationGuide":
		var entity models.ImplementationGuide
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ImplementationGuide{&entity}, nil
	}
	return nil, nil
}

func fhirRespToImplementationGuide(id string, resp *FhirResponse) (*models.ImplementationGuide, error) {
	entities, err := fhirRespToImplementationGuides(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ImplementationGuide", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ImplementationGuide.
func (c *Client) GetImplementationGuide(ctx context.Context, params Parameters) ([]*models.ImplementationGuide, error) {
	resp, err := c.Get(ctx, "ImplementationGuide", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImplementationGuides(resp)
}

// Get ImplementationGuide by ID.
func (c *Client) GetImplementationGuideByID(ctx context.Context, id string, params Parameters) (*models.ImplementationGuide, error) {
	resp, err := c.GetByID(ctx, "ImplementationGuide", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImplementationGuide(id, resp)
}

func (c *Client) CreateImplementationGuide(ctx context.Context, resource string, params Parameters, entity *models.ImplementationGuide) (*models.ImplementationGuide, error) {
	resp, err := c.Create(ctx, "ImplementationGuide", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImplementationGuide(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateImplementationGuide(ctx context.Context, resource string, params Parameters, entity *models.ImplementationGuide) (*models.ImplementationGuide, error) {
	resp, err := c.Update(ctx, "ImplementationGuide", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImplementationGuide(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateImplementationGuideByID(ctx context.Context, resource, id string, params Parameters, entity *models.ImplementationGuide) (*models.ImplementationGuide, error) {
	resp, err := c.UpdateByID(ctx, "ImplementationGuide", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImplementationGuide(id, resp)
}

func (c *Client) PatchImplementationGuide(ctx context.Context, resource string, params Parameters, entity *models.ImplementationGuide) ([]*models.ImplementationGuide, error) {
	resp, err := c.Patch(ctx, "ImplementationGuide", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImplementationGuides(resp)
}

func (c *Client) PatchImplementationGuideByID(ctx context.Context, resource, id string, params Parameters, entity *models.ImplementationGuide) (*models.ImplementationGuide, error) {
	resp, err := c.PatchByID(ctx, "ImplementationGuide", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToImplementationGuide(id, resp)
}

func (c *Client) DeleteImplementationGuide(ctx context.Context, resource string, params Parameters) ([]*models.ImplementationGuide, error) {
	resp, err := c.Delete(ctx, "ImplementationGuide", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImplementationGuides(resp)
}

func (c *Client) DeleteImplementationGuideByID(ctx context.Context, resource, id string, params Parameters) (*models.ImplementationGuide, error) {
	resp, err := c.DeleteByID(ctx, "ImplementationGuide", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToImplementationGuide(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// InsurancePlan
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToInsurancePlans(bundle *models.Bundle) ([]*models.InsurancePlan, error) {
	var entities []*models.InsurancePlan
	err := EnumBundleResources(bundle, "InsurancePlan", func(resource ResourceData) error {
		var entity models.InsurancePlan
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToInsurancePlans(resp *FhirResponse) ([]*models.InsurancePlan, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToInsurancePlans(resp.Bundle)
	case "InsurancePlan":
		var entity models.InsurancePlan
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.InsurancePlan{&entity}, nil
	}
	return nil, nil
}

func fhirRespToInsurancePlan(id string, resp *FhirResponse) (*models.InsurancePlan, error) {
	entities, err := fhirRespToInsurancePlans(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "InsurancePlan", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get InsurancePlan.
func (c *Client) GetInsurancePlan(ctx context.Context, params Parameters) ([]*models.InsurancePlan, error) {
	resp, err := c.Get(ctx, "InsurancePlan", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToInsurancePlans(resp)
}

// Get InsurancePlan by ID.
func (c *Client) GetInsurancePlanByID(ctx context.Context, id string, params Parameters) (*models.InsurancePlan, error) {
	resp, err := c.GetByID(ctx, "InsurancePlan", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToInsurancePlan(id, resp)
}

func (c *Client) CreateInsurancePlan(ctx context.Context, resource string, params Parameters, entity *models.InsurancePlan) (*models.InsurancePlan, error) {
	resp, err := c.Create(ctx, "InsurancePlan", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToInsurancePlan(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateInsurancePlan(ctx context.Context, resource string, params Parameters, entity *models.InsurancePlan) (*models.InsurancePlan, error) {
	resp, err := c.Update(ctx, "InsurancePlan", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToInsurancePlan(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateInsurancePlanByID(ctx context.Context, resource, id string, params Parameters, entity *models.InsurancePlan) (*models.InsurancePlan, error) {
	resp, err := c.UpdateByID(ctx, "InsurancePlan", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToInsurancePlan(id, resp)
}

func (c *Client) PatchInsurancePlan(ctx context.Context, resource string, params Parameters, entity *models.InsurancePlan) ([]*models.InsurancePlan, error) {
	resp, err := c.Patch(ctx, "InsurancePlan", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToInsurancePlans(resp)
}

func (c *Client) PatchInsurancePlanByID(ctx context.Context, resource, id string, params Parameters, entity *models.InsurancePlan) (*models.InsurancePlan, error) {
	resp, err := c.PatchByID(ctx, "InsurancePlan", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToInsurancePlan(id, resp)
}

func (c *Client) DeleteInsurancePlan(ctx context.Context, resource string, params Parameters) ([]*models.InsurancePlan, error) {
	resp, err := c.Delete(ctx, "InsurancePlan", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToInsurancePlans(resp)
}

func (c *Client) DeleteInsurancePlanByID(ctx context.Context, resource, id string, params Parameters) (*models.InsurancePlan, error) {
	resp, err := c.DeleteByID(ctx, "InsurancePlan", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToInsurancePlan(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Invoice
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToInvoices(bundle *models.Bundle) ([]*models.Invoice, error) {
	var entities []*models.Invoice
	err := EnumBundleResources(bundle, "Invoice", func(resource ResourceData) error {
		var entity models.Invoice
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToInvoices(resp *FhirResponse) ([]*models.Invoice, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToInvoices(resp.Bundle)
	case "Invoice":
		var entity models.Invoice
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Invoice{&entity}, nil
	}
	return nil, nil
}

func fhirRespToInvoice(id string, resp *FhirResponse) (*models.Invoice, error) {
	entities, err := fhirRespToInvoices(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Invoice", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Invoice.
func (c *Client) GetInvoice(ctx context.Context, params Parameters) ([]*models.Invoice, error) {
	resp, err := c.Get(ctx, "Invoice", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToInvoices(resp)
}

// Get Invoice by ID.
func (c *Client) GetInvoiceByID(ctx context.Context, id string, params Parameters) (*models.Invoice, error) {
	resp, err := c.GetByID(ctx, "Invoice", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToInvoice(id, resp)
}

func (c *Client) CreateInvoice(ctx context.Context, resource string, params Parameters, entity *models.Invoice) (*models.Invoice, error) {
	resp, err := c.Create(ctx, "Invoice", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToInvoice(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateInvoice(ctx context.Context, resource string, params Parameters, entity *models.Invoice) (*models.Invoice, error) {
	resp, err := c.Update(ctx, "Invoice", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToInvoice(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateInvoiceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Invoice) (*models.Invoice, error) {
	resp, err := c.UpdateByID(ctx, "Invoice", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToInvoice(id, resp)
}

func (c *Client) PatchInvoice(ctx context.Context, resource string, params Parameters, entity *models.Invoice) ([]*models.Invoice, error) {
	resp, err := c.Patch(ctx, "Invoice", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToInvoices(resp)
}

func (c *Client) PatchInvoiceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Invoice) (*models.Invoice, error) {
	resp, err := c.PatchByID(ctx, "Invoice", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToInvoice(id, resp)
}

func (c *Client) DeleteInvoice(ctx context.Context, resource string, params Parameters) ([]*models.Invoice, error) {
	resp, err := c.Delete(ctx, "Invoice", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToInvoices(resp)
}

func (c *Client) DeleteInvoiceByID(ctx context.Context, resource, id string, params Parameters) (*models.Invoice, error) {
	resp, err := c.DeleteByID(ctx, "Invoice", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToInvoice(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Library
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToLibrarys(bundle *models.Bundle) ([]*models.Library, error) {
	var entities []*models.Library
	err := EnumBundleResources(bundle, "Library", func(resource ResourceData) error {
		var entity models.Library
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToLibrarys(resp *FhirResponse) ([]*models.Library, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToLibrarys(resp.Bundle)
	case "Library":
		var entity models.Library
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Library{&entity}, nil
	}
	return nil, nil
}

func fhirRespToLibrary(id string, resp *FhirResponse) (*models.Library, error) {
	entities, err := fhirRespToLibrarys(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Library", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Library.
func (c *Client) GetLibrary(ctx context.Context, params Parameters) ([]*models.Library, error) {
	resp, err := c.Get(ctx, "Library", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLibrarys(resp)
}

// Get Library by ID.
func (c *Client) GetLibraryByID(ctx context.Context, id string, params Parameters) (*models.Library, error) {
	resp, err := c.GetByID(ctx, "Library", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLibrary(id, resp)
}

func (c *Client) CreateLibrary(ctx context.Context, resource string, params Parameters, entity *models.Library) (*models.Library, error) {
	resp, err := c.Create(ctx, "Library", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLibrary(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateLibrary(ctx context.Context, resource string, params Parameters, entity *models.Library) (*models.Library, error) {
	resp, err := c.Update(ctx, "Library", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLibrary(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateLibraryByID(ctx context.Context, resource, id string, params Parameters, entity *models.Library) (*models.Library, error) {
	resp, err := c.UpdateByID(ctx, "Library", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLibrary(id, resp)
}

func (c *Client) PatchLibrary(ctx context.Context, resource string, params Parameters, entity *models.Library) ([]*models.Library, error) {
	resp, err := c.Patch(ctx, "Library", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLibrarys(resp)
}

func (c *Client) PatchLibraryByID(ctx context.Context, resource, id string, params Parameters, entity *models.Library) (*models.Library, error) {
	resp, err := c.PatchByID(ctx, "Library", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLibrary(id, resp)
}

func (c *Client) DeleteLibrary(ctx context.Context, resource string, params Parameters) ([]*models.Library, error) {
	resp, err := c.Delete(ctx, "Library", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLibrarys(resp)
}

func (c *Client) DeleteLibraryByID(ctx context.Context, resource, id string, params Parameters) (*models.Library, error) {
	resp, err := c.DeleteByID(ctx, "Library", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLibrary(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Linkage
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToLinkages(bundle *models.Bundle) ([]*models.Linkage, error) {
	var entities []*models.Linkage
	err := EnumBundleResources(bundle, "Linkage", func(resource ResourceData) error {
		var entity models.Linkage
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToLinkages(resp *FhirResponse) ([]*models.Linkage, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToLinkages(resp.Bundle)
	case "Linkage":
		var entity models.Linkage
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Linkage{&entity}, nil
	}
	return nil, nil
}

func fhirRespToLinkage(id string, resp *FhirResponse) (*models.Linkage, error) {
	entities, err := fhirRespToLinkages(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Linkage", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Linkage.
func (c *Client) GetLinkage(ctx context.Context, params Parameters) ([]*models.Linkage, error) {
	resp, err := c.Get(ctx, "Linkage", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLinkages(resp)
}

// Get Linkage by ID.
func (c *Client) GetLinkageByID(ctx context.Context, id string, params Parameters) (*models.Linkage, error) {
	resp, err := c.GetByID(ctx, "Linkage", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLinkage(id, resp)
}

func (c *Client) CreateLinkage(ctx context.Context, resource string, params Parameters, entity *models.Linkage) (*models.Linkage, error) {
	resp, err := c.Create(ctx, "Linkage", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLinkage(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateLinkage(ctx context.Context, resource string, params Parameters, entity *models.Linkage) (*models.Linkage, error) {
	resp, err := c.Update(ctx, "Linkage", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLinkage(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateLinkageByID(ctx context.Context, resource, id string, params Parameters, entity *models.Linkage) (*models.Linkage, error) {
	resp, err := c.UpdateByID(ctx, "Linkage", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLinkage(id, resp)
}

func (c *Client) PatchLinkage(ctx context.Context, resource string, params Parameters, entity *models.Linkage) ([]*models.Linkage, error) {
	resp, err := c.Patch(ctx, "Linkage", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLinkages(resp)
}

func (c *Client) PatchLinkageByID(ctx context.Context, resource, id string, params Parameters, entity *models.Linkage) (*models.Linkage, error) {
	resp, err := c.PatchByID(ctx, "Linkage", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLinkage(id, resp)
}

func (c *Client) DeleteLinkage(ctx context.Context, resource string, params Parameters) ([]*models.Linkage, error) {
	resp, err := c.Delete(ctx, "Linkage", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLinkages(resp)
}

func (c *Client) DeleteLinkageByID(ctx context.Context, resource, id string, params Parameters) (*models.Linkage, error) {
	resp, err := c.DeleteByID(ctx, "Linkage", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLinkage(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// List
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToLists(bundle *models.Bundle) ([]*models.List, error) {
	var entities []*models.List
	err := EnumBundleResources(bundle, "List", func(resource ResourceData) error {
		var entity models.List
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToLists(resp *FhirResponse) ([]*models.List, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToLists(resp.Bundle)
	case "List":
		var entity models.List
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.List{&entity}, nil
	}
	return nil, nil
}

func fhirRespToList(id string, resp *FhirResponse) (*models.List, error) {
	entities, err := fhirRespToLists(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "List", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get List.
func (c *Client) GetList(ctx context.Context, params Parameters) ([]*models.List, error) {
	resp, err := c.Get(ctx, "List", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLists(resp)
}

// Get List by ID.
func (c *Client) GetListByID(ctx context.Context, id string, params Parameters) (*models.List, error) {
	resp, err := c.GetByID(ctx, "List", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToList(id, resp)
}

func (c *Client) CreateList(ctx context.Context, resource string, params Parameters, entity *models.List) (*models.List, error) {
	resp, err := c.Create(ctx, "List", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToList(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateList(ctx context.Context, resource string, params Parameters, entity *models.List) (*models.List, error) {
	resp, err := c.Update(ctx, "List", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToList(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateListByID(ctx context.Context, resource, id string, params Parameters, entity *models.List) (*models.List, error) {
	resp, err := c.UpdateByID(ctx, "List", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToList(id, resp)
}

func (c *Client) PatchList(ctx context.Context, resource string, params Parameters, entity *models.List) ([]*models.List, error) {
	resp, err := c.Patch(ctx, "List", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLists(resp)
}

func (c *Client) PatchListByID(ctx context.Context, resource, id string, params Parameters, entity *models.List) (*models.List, error) {
	resp, err := c.PatchByID(ctx, "List", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToList(id, resp)
}

func (c *Client) DeleteList(ctx context.Context, resource string, params Parameters) ([]*models.List, error) {
	resp, err := c.Delete(ctx, "List", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLists(resp)
}

func (c *Client) DeleteListByID(ctx context.Context, resource, id string, params Parameters) (*models.List, error) {
	resp, err := c.DeleteByID(ctx, "List", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToList(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Location
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToLocations(bundle *models.Bundle) ([]*models.Location, error) {
	var entities []*models.Location
	err := EnumBundleResources(bundle, "Location", func(resource ResourceData) error {
		var entity models.Location
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToLocations(resp *FhirResponse) ([]*models.Location, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToLocations(resp.Bundle)
	case "Location":
		var entity models.Location
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Location{&entity}, nil
	}
	return nil, nil
}

func fhirRespToLocation(id string, resp *FhirResponse) (*models.Location, error) {
	entities, err := fhirRespToLocations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Location", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Location.
func (c *Client) GetLocation(ctx context.Context, params Parameters) ([]*models.Location, error) {
	resp, err := c.Get(ctx, "Location", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLocations(resp)
}

// Get Location by ID.
func (c *Client) GetLocationByID(ctx context.Context, id string, params Parameters) (*models.Location, error) {
	resp, err := c.GetByID(ctx, "Location", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLocation(id, resp)
}

func (c *Client) CreateLocation(ctx context.Context, resource string, params Parameters, entity *models.Location) (*models.Location, error) {
	resp, err := c.Create(ctx, "Location", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLocation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateLocation(ctx context.Context, resource string, params Parameters, entity *models.Location) (*models.Location, error) {
	resp, err := c.Update(ctx, "Location", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLocation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateLocationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Location) (*models.Location, error) {
	resp, err := c.UpdateByID(ctx, "Location", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLocation(id, resp)
}

func (c *Client) PatchLocation(ctx context.Context, resource string, params Parameters, entity *models.Location) ([]*models.Location, error) {
	resp, err := c.Patch(ctx, "Location", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLocations(resp)
}

func (c *Client) PatchLocationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Location) (*models.Location, error) {
	resp, err := c.PatchByID(ctx, "Location", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToLocation(id, resp)
}

func (c *Client) DeleteLocation(ctx context.Context, resource string, params Parameters) ([]*models.Location, error) {
	resp, err := c.Delete(ctx, "Location", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLocations(resp)
}

func (c *Client) DeleteLocationByID(ctx context.Context, resource, id string, params Parameters) (*models.Location, error) {
	resp, err := c.DeleteByID(ctx, "Location", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToLocation(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Measure
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMeasures(bundle *models.Bundle) ([]*models.Measure, error) {
	var entities []*models.Measure
	err := EnumBundleResources(bundle, "Measure", func(resource ResourceData) error {
		var entity models.Measure
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMeasures(resp *FhirResponse) ([]*models.Measure, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMeasures(resp.Bundle)
	case "Measure":
		var entity models.Measure
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Measure{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMeasure(id string, resp *FhirResponse) (*models.Measure, error) {
	entities, err := fhirRespToMeasures(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Measure", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Measure.
func (c *Client) GetMeasure(ctx context.Context, params Parameters) ([]*models.Measure, error) {
	resp, err := c.Get(ctx, "Measure", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasures(resp)
}

// Get Measure by ID.
func (c *Client) GetMeasureByID(ctx context.Context, id string, params Parameters) (*models.Measure, error) {
	resp, err := c.GetByID(ctx, "Measure", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasure(id, resp)
}

func (c *Client) CreateMeasure(ctx context.Context, resource string, params Parameters, entity *models.Measure) (*models.Measure, error) {
	resp, err := c.Create(ctx, "Measure", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasure(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMeasure(ctx context.Context, resource string, params Parameters, entity *models.Measure) (*models.Measure, error) {
	resp, err := c.Update(ctx, "Measure", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasure(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMeasureByID(ctx context.Context, resource, id string, params Parameters, entity *models.Measure) (*models.Measure, error) {
	resp, err := c.UpdateByID(ctx, "Measure", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasure(id, resp)
}

func (c *Client) PatchMeasure(ctx context.Context, resource string, params Parameters, entity *models.Measure) ([]*models.Measure, error) {
	resp, err := c.Patch(ctx, "Measure", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasures(resp)
}

func (c *Client) PatchMeasureByID(ctx context.Context, resource, id string, params Parameters, entity *models.Measure) (*models.Measure, error) {
	resp, err := c.PatchByID(ctx, "Measure", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasure(id, resp)
}

func (c *Client) DeleteMeasure(ctx context.Context, resource string, params Parameters) ([]*models.Measure, error) {
	resp, err := c.Delete(ctx, "Measure", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasures(resp)
}

func (c *Client) DeleteMeasureByID(ctx context.Context, resource, id string, params Parameters) (*models.Measure, error) {
	resp, err := c.DeleteByID(ctx, "Measure", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasure(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MeasureReport
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMeasureReports(bundle *models.Bundle) ([]*models.MeasureReport, error) {
	var entities []*models.MeasureReport
	err := EnumBundleResources(bundle, "MeasureReport", func(resource ResourceData) error {
		var entity models.MeasureReport
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMeasureReports(resp *FhirResponse) ([]*models.MeasureReport, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMeasureReports(resp.Bundle)
	case "MeasureReport":
		var entity models.MeasureReport
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MeasureReport{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMeasureReport(id string, resp *FhirResponse) (*models.MeasureReport, error) {
	entities, err := fhirRespToMeasureReports(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MeasureReport", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MeasureReport.
func (c *Client) GetMeasureReport(ctx context.Context, params Parameters) ([]*models.MeasureReport, error) {
	resp, err := c.Get(ctx, "MeasureReport", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasureReports(resp)
}

// Get MeasureReport by ID.
func (c *Client) GetMeasureReportByID(ctx context.Context, id string, params Parameters) (*models.MeasureReport, error) {
	resp, err := c.GetByID(ctx, "MeasureReport", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasureReport(id, resp)
}

func (c *Client) CreateMeasureReport(ctx context.Context, resource string, params Parameters, entity *models.MeasureReport) (*models.MeasureReport, error) {
	resp, err := c.Create(ctx, "MeasureReport", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasureReport(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMeasureReport(ctx context.Context, resource string, params Parameters, entity *models.MeasureReport) (*models.MeasureReport, error) {
	resp, err := c.Update(ctx, "MeasureReport", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasureReport(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMeasureReportByID(ctx context.Context, resource, id string, params Parameters, entity *models.MeasureReport) (*models.MeasureReport, error) {
	resp, err := c.UpdateByID(ctx, "MeasureReport", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasureReport(id, resp)
}

func (c *Client) PatchMeasureReport(ctx context.Context, resource string, params Parameters, entity *models.MeasureReport) ([]*models.MeasureReport, error) {
	resp, err := c.Patch(ctx, "MeasureReport", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasureReports(resp)
}

func (c *Client) PatchMeasureReportByID(ctx context.Context, resource, id string, params Parameters, entity *models.MeasureReport) (*models.MeasureReport, error) {
	resp, err := c.PatchByID(ctx, "MeasureReport", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasureReport(id, resp)
}

func (c *Client) DeleteMeasureReport(ctx context.Context, resource string, params Parameters) ([]*models.MeasureReport, error) {
	resp, err := c.Delete(ctx, "MeasureReport", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasureReports(resp)
}

func (c *Client) DeleteMeasureReportByID(ctx context.Context, resource, id string, params Parameters) (*models.MeasureReport, error) {
	resp, err := c.DeleteByID(ctx, "MeasureReport", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMeasureReport(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Media
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedias(bundle *models.Bundle) ([]*models.Media, error) {
	var entities []*models.Media
	err := EnumBundleResources(bundle, "Media", func(resource ResourceData) error {
		var entity models.Media
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedias(resp *FhirResponse) ([]*models.Media, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedias(resp.Bundle)
	case "Media":
		var entity models.Media
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Media{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedia(id string, resp *FhirResponse) (*models.Media, error) {
	entities, err := fhirRespToMedias(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Media", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Media.
func (c *Client) GetMedia(ctx context.Context, params Parameters) ([]*models.Media, error) {
	resp, err := c.Get(ctx, "Media", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedias(resp)
}

// Get Media by ID.
func (c *Client) GetMediaByID(ctx context.Context, id string, params Parameters) (*models.Media, error) {
	resp, err := c.GetByID(ctx, "Media", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedia(id, resp)
}

func (c *Client) CreateMedia(ctx context.Context, resource string, params Parameters, entity *models.Media) (*models.Media, error) {
	resp, err := c.Create(ctx, "Media", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedia(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedia(ctx context.Context, resource string, params Parameters, entity *models.Media) (*models.Media, error) {
	resp, err := c.Update(ctx, "Media", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedia(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMediaByID(ctx context.Context, resource, id string, params Parameters, entity *models.Media) (*models.Media, error) {
	resp, err := c.UpdateByID(ctx, "Media", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedia(id, resp)
}

func (c *Client) PatchMedia(ctx context.Context, resource string, params Parameters, entity *models.Media) ([]*models.Media, error) {
	resp, err := c.Patch(ctx, "Media", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedias(resp)
}

func (c *Client) PatchMediaByID(ctx context.Context, resource, id string, params Parameters, entity *models.Media) (*models.Media, error) {
	resp, err := c.PatchByID(ctx, "Media", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedia(id, resp)
}

func (c *Client) DeleteMedia(ctx context.Context, resource string, params Parameters) ([]*models.Media, error) {
	resp, err := c.Delete(ctx, "Media", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedias(resp)
}

func (c *Client) DeleteMediaByID(ctx context.Context, resource, id string, params Parameters) (*models.Media, error) {
	resp, err := c.DeleteByID(ctx, "Media", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedia(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Medication
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedications(bundle *models.Bundle) ([]*models.Medication, error) {
	var entities []*models.Medication
	err := EnumBundleResources(bundle, "Medication", func(resource ResourceData) error {
		var entity models.Medication
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedications(resp *FhirResponse) ([]*models.Medication, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedications(resp.Bundle)
	case "Medication":
		var entity models.Medication
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Medication{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedication(id string, resp *FhirResponse) (*models.Medication, error) {
	entities, err := fhirRespToMedications(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Medication", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Medication.
func (c *Client) GetMedication(ctx context.Context, params Parameters) ([]*models.Medication, error) {
	resp, err := c.Get(ctx, "Medication", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedications(resp)
}

// Get Medication by ID.
func (c *Client) GetMedicationByID(ctx context.Context, id string, params Parameters) (*models.Medication, error) {
	resp, err := c.GetByID(ctx, "Medication", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedication(id, resp)
}

func (c *Client) CreateMedication(ctx context.Context, resource string, params Parameters, entity *models.Medication) (*models.Medication, error) {
	resp, err := c.Create(ctx, "Medication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedication(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedication(ctx context.Context, resource string, params Parameters, entity *models.Medication) (*models.Medication, error) {
	resp, err := c.Update(ctx, "Medication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedication(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Medication) (*models.Medication, error) {
	resp, err := c.UpdateByID(ctx, "Medication", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedication(id, resp)
}

func (c *Client) PatchMedication(ctx context.Context, resource string, params Parameters, entity *models.Medication) ([]*models.Medication, error) {
	resp, err := c.Patch(ctx, "Medication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedications(resp)
}

func (c *Client) PatchMedicationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Medication) (*models.Medication, error) {
	resp, err := c.PatchByID(ctx, "Medication", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedication(id, resp)
}

func (c *Client) DeleteMedication(ctx context.Context, resource string, params Parameters) ([]*models.Medication, error) {
	resp, err := c.Delete(ctx, "Medication", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedications(resp)
}

func (c *Client) DeleteMedicationByID(ctx context.Context, resource, id string, params Parameters) (*models.Medication, error) {
	resp, err := c.DeleteByID(ctx, "Medication", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedication(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicationAdministration
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicationAdministrations(bundle *models.Bundle) ([]*models.MedicationAdministration, error) {
	var entities []*models.MedicationAdministration
	err := EnumBundleResources(bundle, "MedicationAdministration", func(resource ResourceData) error {
		var entity models.MedicationAdministration
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicationAdministrations(resp *FhirResponse) ([]*models.MedicationAdministration, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicationAdministrations(resp.Bundle)
	case "MedicationAdministration":
		var entity models.MedicationAdministration
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicationAdministration{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicationAdministration(id string, resp *FhirResponse) (*models.MedicationAdministration, error) {
	entities, err := fhirRespToMedicationAdministrations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicationAdministration", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicationAdministration.
func (c *Client) GetMedicationAdministration(ctx context.Context, params Parameters) ([]*models.MedicationAdministration, error) {
	resp, err := c.Get(ctx, "MedicationAdministration", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationAdministrations(resp)
}

// Get MedicationAdministration by ID.
func (c *Client) GetMedicationAdministrationByID(ctx context.Context, id string, params Parameters) (*models.MedicationAdministration, error) {
	resp, err := c.GetByID(ctx, "MedicationAdministration", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationAdministration(id, resp)
}

func (c *Client) CreateMedicationAdministration(ctx context.Context, resource string, params Parameters, entity *models.MedicationAdministration) (*models.MedicationAdministration, error) {
	resp, err := c.Create(ctx, "MedicationAdministration", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationAdministration(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationAdministration(ctx context.Context, resource string, params Parameters, entity *models.MedicationAdministration) (*models.MedicationAdministration, error) {
	resp, err := c.Update(ctx, "MedicationAdministration", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationAdministration(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationAdministrationByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicationAdministration) (*models.MedicationAdministration, error) {
	resp, err := c.UpdateByID(ctx, "MedicationAdministration", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationAdministration(id, resp)
}

func (c *Client) PatchMedicationAdministration(ctx context.Context, resource string, params Parameters, entity *models.MedicationAdministration) ([]*models.MedicationAdministration, error) {
	resp, err := c.Patch(ctx, "MedicationAdministration", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationAdministrations(resp)
}

func (c *Client) PatchMedicationAdministrationByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicationAdministration) (*models.MedicationAdministration, error) {
	resp, err := c.PatchByID(ctx, "MedicationAdministration", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationAdministration(id, resp)
}

func (c *Client) DeleteMedicationAdministration(ctx context.Context, resource string, params Parameters) ([]*models.MedicationAdministration, error) {
	resp, err := c.Delete(ctx, "MedicationAdministration", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationAdministrations(resp)
}

func (c *Client) DeleteMedicationAdministrationByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicationAdministration, error) {
	resp, err := c.DeleteByID(ctx, "MedicationAdministration", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationAdministration(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicationDispense
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicationDispenses(bundle *models.Bundle) ([]*models.MedicationDispense, error) {
	var entities []*models.MedicationDispense
	err := EnumBundleResources(bundle, "MedicationDispense", func(resource ResourceData) error {
		var entity models.MedicationDispense
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicationDispenses(resp *FhirResponse) ([]*models.MedicationDispense, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicationDispenses(resp.Bundle)
	case "MedicationDispense":
		var entity models.MedicationDispense
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicationDispense{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicationDispense(id string, resp *FhirResponse) (*models.MedicationDispense, error) {
	entities, err := fhirRespToMedicationDispenses(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicationDispense", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicationDispense.
func (c *Client) GetMedicationDispense(ctx context.Context, params Parameters) ([]*models.MedicationDispense, error) {
	resp, err := c.Get(ctx, "MedicationDispense", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationDispenses(resp)
}

// Get MedicationDispense by ID.
func (c *Client) GetMedicationDispenseByID(ctx context.Context, id string, params Parameters) (*models.MedicationDispense, error) {
	resp, err := c.GetByID(ctx, "MedicationDispense", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationDispense(id, resp)
}

func (c *Client) CreateMedicationDispense(ctx context.Context, resource string, params Parameters, entity *models.MedicationDispense) (*models.MedicationDispense, error) {
	resp, err := c.Create(ctx, "MedicationDispense", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationDispense(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationDispense(ctx context.Context, resource string, params Parameters, entity *models.MedicationDispense) (*models.MedicationDispense, error) {
	resp, err := c.Update(ctx, "MedicationDispense", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationDispense(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationDispenseByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicationDispense) (*models.MedicationDispense, error) {
	resp, err := c.UpdateByID(ctx, "MedicationDispense", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationDispense(id, resp)
}

func (c *Client) PatchMedicationDispense(ctx context.Context, resource string, params Parameters, entity *models.MedicationDispense) ([]*models.MedicationDispense, error) {
	resp, err := c.Patch(ctx, "MedicationDispense", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationDispenses(resp)
}

func (c *Client) PatchMedicationDispenseByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicationDispense) (*models.MedicationDispense, error) {
	resp, err := c.PatchByID(ctx, "MedicationDispense", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationDispense(id, resp)
}

func (c *Client) DeleteMedicationDispense(ctx context.Context, resource string, params Parameters) ([]*models.MedicationDispense, error) {
	resp, err := c.Delete(ctx, "MedicationDispense", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationDispenses(resp)
}

func (c *Client) DeleteMedicationDispenseByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicationDispense, error) {
	resp, err := c.DeleteByID(ctx, "MedicationDispense", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationDispense(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicationKnowledge
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicationKnowledges(bundle *models.Bundle) ([]*models.MedicationKnowledge, error) {
	var entities []*models.MedicationKnowledge
	err := EnumBundleResources(bundle, "MedicationKnowledge", func(resource ResourceData) error {
		var entity models.MedicationKnowledge
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicationKnowledges(resp *FhirResponse) ([]*models.MedicationKnowledge, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicationKnowledges(resp.Bundle)
	case "MedicationKnowledge":
		var entity models.MedicationKnowledge
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicationKnowledge{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicationKnowledge(id string, resp *FhirResponse) (*models.MedicationKnowledge, error) {
	entities, err := fhirRespToMedicationKnowledges(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicationKnowledge", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicationKnowledge.
func (c *Client) GetMedicationKnowledge(ctx context.Context, params Parameters) ([]*models.MedicationKnowledge, error) {
	resp, err := c.Get(ctx, "MedicationKnowledge", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationKnowledges(resp)
}

// Get MedicationKnowledge by ID.
func (c *Client) GetMedicationKnowledgeByID(ctx context.Context, id string, params Parameters) (*models.MedicationKnowledge, error) {
	resp, err := c.GetByID(ctx, "MedicationKnowledge", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationKnowledge(id, resp)
}

func (c *Client) CreateMedicationKnowledge(ctx context.Context, resource string, params Parameters, entity *models.MedicationKnowledge) (*models.MedicationKnowledge, error) {
	resp, err := c.Create(ctx, "MedicationKnowledge", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationKnowledge(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationKnowledge(ctx context.Context, resource string, params Parameters, entity *models.MedicationKnowledge) (*models.MedicationKnowledge, error) {
	resp, err := c.Update(ctx, "MedicationKnowledge", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationKnowledge(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationKnowledgeByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicationKnowledge) (*models.MedicationKnowledge, error) {
	resp, err := c.UpdateByID(ctx, "MedicationKnowledge", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationKnowledge(id, resp)
}

func (c *Client) PatchMedicationKnowledge(ctx context.Context, resource string, params Parameters, entity *models.MedicationKnowledge) ([]*models.MedicationKnowledge, error) {
	resp, err := c.Patch(ctx, "MedicationKnowledge", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationKnowledges(resp)
}

func (c *Client) PatchMedicationKnowledgeByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicationKnowledge) (*models.MedicationKnowledge, error) {
	resp, err := c.PatchByID(ctx, "MedicationKnowledge", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationKnowledge(id, resp)
}

func (c *Client) DeleteMedicationKnowledge(ctx context.Context, resource string, params Parameters) ([]*models.MedicationKnowledge, error) {
	resp, err := c.Delete(ctx, "MedicationKnowledge", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationKnowledges(resp)
}

func (c *Client) DeleteMedicationKnowledgeByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicationKnowledge, error) {
	resp, err := c.DeleteByID(ctx, "MedicationKnowledge", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationKnowledge(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicationRequest
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicationRequests(bundle *models.Bundle) ([]*models.MedicationRequest, error) {
	var entities []*models.MedicationRequest
	err := EnumBundleResources(bundle, "MedicationRequest", func(resource ResourceData) error {
		var entity models.MedicationRequest
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicationRequests(resp *FhirResponse) ([]*models.MedicationRequest, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicationRequests(resp.Bundle)
	case "MedicationRequest":
		var entity models.MedicationRequest
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicationRequest{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicationRequest(id string, resp *FhirResponse) (*models.MedicationRequest, error) {
	entities, err := fhirRespToMedicationRequests(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicationRequest", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicationRequest.
func (c *Client) GetMedicationRequest(ctx context.Context, params Parameters) ([]*models.MedicationRequest, error) {
	resp, err := c.Get(ctx, "MedicationRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationRequests(resp)
}

// Get MedicationRequest by ID.
func (c *Client) GetMedicationRequestByID(ctx context.Context, id string, params Parameters) (*models.MedicationRequest, error) {
	resp, err := c.GetByID(ctx, "MedicationRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationRequest(id, resp)
}

func (c *Client) CreateMedicationRequest(ctx context.Context, resource string, params Parameters, entity *models.MedicationRequest) (*models.MedicationRequest, error) {
	resp, err := c.Create(ctx, "MedicationRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationRequest(ctx context.Context, resource string, params Parameters, entity *models.MedicationRequest) (*models.MedicationRequest, error) {
	resp, err := c.Update(ctx, "MedicationRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicationRequest) (*models.MedicationRequest, error) {
	resp, err := c.UpdateByID(ctx, "MedicationRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationRequest(id, resp)
}

func (c *Client) PatchMedicationRequest(ctx context.Context, resource string, params Parameters, entity *models.MedicationRequest) ([]*models.MedicationRequest, error) {
	resp, err := c.Patch(ctx, "MedicationRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationRequests(resp)
}

func (c *Client) PatchMedicationRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicationRequest) (*models.MedicationRequest, error) {
	resp, err := c.PatchByID(ctx, "MedicationRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationRequest(id, resp)
}

func (c *Client) DeleteMedicationRequest(ctx context.Context, resource string, params Parameters) ([]*models.MedicationRequest, error) {
	resp, err := c.Delete(ctx, "MedicationRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationRequests(resp)
}

func (c *Client) DeleteMedicationRequestByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicationRequest, error) {
	resp, err := c.DeleteByID(ctx, "MedicationRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationRequest(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicationStatement
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicationStatements(bundle *models.Bundle) ([]*models.MedicationStatement, error) {
	var entities []*models.MedicationStatement
	err := EnumBundleResources(bundle, "MedicationStatement", func(resource ResourceData) error {
		var entity models.MedicationStatement
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicationStatements(resp *FhirResponse) ([]*models.MedicationStatement, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicationStatements(resp.Bundle)
	case "MedicationStatement":
		var entity models.MedicationStatement
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicationStatement{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicationStatement(id string, resp *FhirResponse) (*models.MedicationStatement, error) {
	entities, err := fhirRespToMedicationStatements(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicationStatement", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicationStatement.
func (c *Client) GetMedicationStatement(ctx context.Context, params Parameters) ([]*models.MedicationStatement, error) {
	resp, err := c.Get(ctx, "MedicationStatement", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationStatements(resp)
}

// Get MedicationStatement by ID.
func (c *Client) GetMedicationStatementByID(ctx context.Context, id string, params Parameters) (*models.MedicationStatement, error) {
	resp, err := c.GetByID(ctx, "MedicationStatement", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationStatement(id, resp)
}

func (c *Client) CreateMedicationStatement(ctx context.Context, resource string, params Parameters, entity *models.MedicationStatement) (*models.MedicationStatement, error) {
	resp, err := c.Create(ctx, "MedicationStatement", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationStatement(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationStatement(ctx context.Context, resource string, params Parameters, entity *models.MedicationStatement) (*models.MedicationStatement, error) {
	resp, err := c.Update(ctx, "MedicationStatement", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationStatement(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicationStatementByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicationStatement) (*models.MedicationStatement, error) {
	resp, err := c.UpdateByID(ctx, "MedicationStatement", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationStatement(id, resp)
}

func (c *Client) PatchMedicationStatement(ctx context.Context, resource string, params Parameters, entity *models.MedicationStatement) ([]*models.MedicationStatement, error) {
	resp, err := c.Patch(ctx, "MedicationStatement", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationStatements(resp)
}

func (c *Client) PatchMedicationStatementByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicationStatement) (*models.MedicationStatement, error) {
	resp, err := c.PatchByID(ctx, "MedicationStatement", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationStatement(id, resp)
}

func (c *Client) DeleteMedicationStatement(ctx context.Context, resource string, params Parameters) ([]*models.MedicationStatement, error) {
	resp, err := c.Delete(ctx, "MedicationStatement", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationStatements(resp)
}

func (c *Client) DeleteMedicationStatementByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicationStatement, error) {
	resp, err := c.DeleteByID(ctx, "MedicationStatement", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicationStatement(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicinalProduct
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicinalProducts(bundle *models.Bundle) ([]*models.MedicinalProduct, error) {
	var entities []*models.MedicinalProduct
	err := EnumBundleResources(bundle, "MedicinalProduct", func(resource ResourceData) error {
		var entity models.MedicinalProduct
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicinalProducts(resp *FhirResponse) ([]*models.MedicinalProduct, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicinalProducts(resp.Bundle)
	case "MedicinalProduct":
		var entity models.MedicinalProduct
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicinalProduct{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicinalProduct(id string, resp *FhirResponse) (*models.MedicinalProduct, error) {
	entities, err := fhirRespToMedicinalProducts(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicinalProduct", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicinalProduct.
func (c *Client) GetMedicinalProduct(ctx context.Context, params Parameters) ([]*models.MedicinalProduct, error) {
	resp, err := c.Get(ctx, "MedicinalProduct", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProducts(resp)
}

// Get MedicinalProduct by ID.
func (c *Client) GetMedicinalProductByID(ctx context.Context, id string, params Parameters) (*models.MedicinalProduct, error) {
	resp, err := c.GetByID(ctx, "MedicinalProduct", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProduct(id, resp)
}

func (c *Client) CreateMedicinalProduct(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProduct) (*models.MedicinalProduct, error) {
	resp, err := c.Create(ctx, "MedicinalProduct", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProduct(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProduct(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProduct) (*models.MedicinalProduct, error) {
	resp, err := c.Update(ctx, "MedicinalProduct", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProduct(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProduct) (*models.MedicinalProduct, error) {
	resp, err := c.UpdateByID(ctx, "MedicinalProduct", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProduct(id, resp)
}

func (c *Client) PatchMedicinalProduct(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProduct) ([]*models.MedicinalProduct, error) {
	resp, err := c.Patch(ctx, "MedicinalProduct", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProducts(resp)
}

func (c *Client) PatchMedicinalProductByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProduct) (*models.MedicinalProduct, error) {
	resp, err := c.PatchByID(ctx, "MedicinalProduct", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProduct(id, resp)
}

func (c *Client) DeleteMedicinalProduct(ctx context.Context, resource string, params Parameters) ([]*models.MedicinalProduct, error) {
	resp, err := c.Delete(ctx, "MedicinalProduct", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProducts(resp)
}

func (c *Client) DeleteMedicinalProductByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicinalProduct, error) {
	resp, err := c.DeleteByID(ctx, "MedicinalProduct", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProduct(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicinalProductAuthorization
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicinalProductAuthorizations(bundle *models.Bundle) ([]*models.MedicinalProductAuthorization, error) {
	var entities []*models.MedicinalProductAuthorization
	err := EnumBundleResources(bundle, "MedicinalProductAuthorization", func(resource ResourceData) error {
		var entity models.MedicinalProductAuthorization
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicinalProductAuthorizations(resp *FhirResponse) ([]*models.MedicinalProductAuthorization, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicinalProductAuthorizations(resp.Bundle)
	case "MedicinalProductAuthorization":
		var entity models.MedicinalProductAuthorization
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicinalProductAuthorization{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicinalProductAuthorization(id string, resp *FhirResponse) (*models.MedicinalProductAuthorization, error) {
	entities, err := fhirRespToMedicinalProductAuthorizations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicinalProductAuthorization", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicinalProductAuthorization.
func (c *Client) GetMedicinalProductAuthorization(ctx context.Context, params Parameters) ([]*models.MedicinalProductAuthorization, error) {
	resp, err := c.Get(ctx, "MedicinalProductAuthorization", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductAuthorizations(resp)
}

// Get MedicinalProductAuthorization by ID.
func (c *Client) GetMedicinalProductAuthorizationByID(ctx context.Context, id string, params Parameters) (*models.MedicinalProductAuthorization, error) {
	resp, err := c.GetByID(ctx, "MedicinalProductAuthorization", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductAuthorization(id, resp)
}

func (c *Client) CreateMedicinalProductAuthorization(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductAuthorization) (*models.MedicinalProductAuthorization, error) {
	resp, err := c.Create(ctx, "MedicinalProductAuthorization", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductAuthorization(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductAuthorization(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductAuthorization) (*models.MedicinalProductAuthorization, error) {
	resp, err := c.Update(ctx, "MedicinalProductAuthorization", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductAuthorization(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductAuthorizationByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductAuthorization) (*models.MedicinalProductAuthorization, error) {
	resp, err := c.UpdateByID(ctx, "MedicinalProductAuthorization", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductAuthorization(id, resp)
}

func (c *Client) PatchMedicinalProductAuthorization(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductAuthorization) ([]*models.MedicinalProductAuthorization, error) {
	resp, err := c.Patch(ctx, "MedicinalProductAuthorization", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductAuthorizations(resp)
}

func (c *Client) PatchMedicinalProductAuthorizationByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductAuthorization) (*models.MedicinalProductAuthorization, error) {
	resp, err := c.PatchByID(ctx, "MedicinalProductAuthorization", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductAuthorization(id, resp)
}

func (c *Client) DeleteMedicinalProductAuthorization(ctx context.Context, resource string, params Parameters) ([]*models.MedicinalProductAuthorization, error) {
	resp, err := c.Delete(ctx, "MedicinalProductAuthorization", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductAuthorizations(resp)
}

func (c *Client) DeleteMedicinalProductAuthorizationByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicinalProductAuthorization, error) {
	resp, err := c.DeleteByID(ctx, "MedicinalProductAuthorization", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductAuthorization(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicinalProductContraindication
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicinalProductContraindications(bundle *models.Bundle) ([]*models.MedicinalProductContraindication, error) {
	var entities []*models.MedicinalProductContraindication
	err := EnumBundleResources(bundle, "MedicinalProductContraindication", func(resource ResourceData) error {
		var entity models.MedicinalProductContraindication
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicinalProductContraindications(resp *FhirResponse) ([]*models.MedicinalProductContraindication, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicinalProductContraindications(resp.Bundle)
	case "MedicinalProductContraindication":
		var entity models.MedicinalProductContraindication
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicinalProductContraindication{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicinalProductContraindication(id string, resp *FhirResponse) (*models.MedicinalProductContraindication, error) {
	entities, err := fhirRespToMedicinalProductContraindications(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicinalProductContraindication", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicinalProductContraindication.
func (c *Client) GetMedicinalProductContraindication(ctx context.Context, params Parameters) ([]*models.MedicinalProductContraindication, error) {
	resp, err := c.Get(ctx, "MedicinalProductContraindication", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductContraindications(resp)
}

// Get MedicinalProductContraindication by ID.
func (c *Client) GetMedicinalProductContraindicationByID(ctx context.Context, id string, params Parameters) (*models.MedicinalProductContraindication, error) {
	resp, err := c.GetByID(ctx, "MedicinalProductContraindication", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductContraindication(id, resp)
}

func (c *Client) CreateMedicinalProductContraindication(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductContraindication) (*models.MedicinalProductContraindication, error) {
	resp, err := c.Create(ctx, "MedicinalProductContraindication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductContraindication(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductContraindication(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductContraindication) (*models.MedicinalProductContraindication, error) {
	resp, err := c.Update(ctx, "MedicinalProductContraindication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductContraindication(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductContraindicationByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductContraindication) (*models.MedicinalProductContraindication, error) {
	resp, err := c.UpdateByID(ctx, "MedicinalProductContraindication", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductContraindication(id, resp)
}

func (c *Client) PatchMedicinalProductContraindication(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductContraindication) ([]*models.MedicinalProductContraindication, error) {
	resp, err := c.Patch(ctx, "MedicinalProductContraindication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductContraindications(resp)
}

func (c *Client) PatchMedicinalProductContraindicationByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductContraindication) (*models.MedicinalProductContraindication, error) {
	resp, err := c.PatchByID(ctx, "MedicinalProductContraindication", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductContraindication(id, resp)
}

func (c *Client) DeleteMedicinalProductContraindication(ctx context.Context, resource string, params Parameters) ([]*models.MedicinalProductContraindication, error) {
	resp, err := c.Delete(ctx, "MedicinalProductContraindication", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductContraindications(resp)
}

func (c *Client) DeleteMedicinalProductContraindicationByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicinalProductContraindication, error) {
	resp, err := c.DeleteByID(ctx, "MedicinalProductContraindication", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductContraindication(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicinalProductIndication
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicinalProductIndications(bundle *models.Bundle) ([]*models.MedicinalProductIndication, error) {
	var entities []*models.MedicinalProductIndication
	err := EnumBundleResources(bundle, "MedicinalProductIndication", func(resource ResourceData) error {
		var entity models.MedicinalProductIndication
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicinalProductIndications(resp *FhirResponse) ([]*models.MedicinalProductIndication, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicinalProductIndications(resp.Bundle)
	case "MedicinalProductIndication":
		var entity models.MedicinalProductIndication
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicinalProductIndication{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicinalProductIndication(id string, resp *FhirResponse) (*models.MedicinalProductIndication, error) {
	entities, err := fhirRespToMedicinalProductIndications(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicinalProductIndication", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicinalProductIndication.
func (c *Client) GetMedicinalProductIndication(ctx context.Context, params Parameters) ([]*models.MedicinalProductIndication, error) {
	resp, err := c.Get(ctx, "MedicinalProductIndication", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIndications(resp)
}

// Get MedicinalProductIndication by ID.
func (c *Client) GetMedicinalProductIndicationByID(ctx context.Context, id string, params Parameters) (*models.MedicinalProductIndication, error) {
	resp, err := c.GetByID(ctx, "MedicinalProductIndication", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIndication(id, resp)
}

func (c *Client) CreateMedicinalProductIndication(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductIndication) (*models.MedicinalProductIndication, error) {
	resp, err := c.Create(ctx, "MedicinalProductIndication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIndication(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductIndication(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductIndication) (*models.MedicinalProductIndication, error) {
	resp, err := c.Update(ctx, "MedicinalProductIndication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIndication(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductIndicationByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductIndication) (*models.MedicinalProductIndication, error) {
	resp, err := c.UpdateByID(ctx, "MedicinalProductIndication", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIndication(id, resp)
}

func (c *Client) PatchMedicinalProductIndication(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductIndication) ([]*models.MedicinalProductIndication, error) {
	resp, err := c.Patch(ctx, "MedicinalProductIndication", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIndications(resp)
}

func (c *Client) PatchMedicinalProductIndicationByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductIndication) (*models.MedicinalProductIndication, error) {
	resp, err := c.PatchByID(ctx, "MedicinalProductIndication", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIndication(id, resp)
}

func (c *Client) DeleteMedicinalProductIndication(ctx context.Context, resource string, params Parameters) ([]*models.MedicinalProductIndication, error) {
	resp, err := c.Delete(ctx, "MedicinalProductIndication", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIndications(resp)
}

func (c *Client) DeleteMedicinalProductIndicationByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicinalProductIndication, error) {
	resp, err := c.DeleteByID(ctx, "MedicinalProductIndication", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIndication(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicinalProductIngredient
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicinalProductIngredients(bundle *models.Bundle) ([]*models.MedicinalProductIngredient, error) {
	var entities []*models.MedicinalProductIngredient
	err := EnumBundleResources(bundle, "MedicinalProductIngredient", func(resource ResourceData) error {
		var entity models.MedicinalProductIngredient
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicinalProductIngredients(resp *FhirResponse) ([]*models.MedicinalProductIngredient, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicinalProductIngredients(resp.Bundle)
	case "MedicinalProductIngredient":
		var entity models.MedicinalProductIngredient
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicinalProductIngredient{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicinalProductIngredient(id string, resp *FhirResponse) (*models.MedicinalProductIngredient, error) {
	entities, err := fhirRespToMedicinalProductIngredients(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicinalProductIngredient", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicinalProductIngredient.
func (c *Client) GetMedicinalProductIngredient(ctx context.Context, params Parameters) ([]*models.MedicinalProductIngredient, error) {
	resp, err := c.Get(ctx, "MedicinalProductIngredient", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIngredients(resp)
}

// Get MedicinalProductIngredient by ID.
func (c *Client) GetMedicinalProductIngredientByID(ctx context.Context, id string, params Parameters) (*models.MedicinalProductIngredient, error) {
	resp, err := c.GetByID(ctx, "MedicinalProductIngredient", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIngredient(id, resp)
}

func (c *Client) CreateMedicinalProductIngredient(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductIngredient) (*models.MedicinalProductIngredient, error) {
	resp, err := c.Create(ctx, "MedicinalProductIngredient", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIngredient(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductIngredient(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductIngredient) (*models.MedicinalProductIngredient, error) {
	resp, err := c.Update(ctx, "MedicinalProductIngredient", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIngredient(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductIngredientByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductIngredient) (*models.MedicinalProductIngredient, error) {
	resp, err := c.UpdateByID(ctx, "MedicinalProductIngredient", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIngredient(id, resp)
}

func (c *Client) PatchMedicinalProductIngredient(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductIngredient) ([]*models.MedicinalProductIngredient, error) {
	resp, err := c.Patch(ctx, "MedicinalProductIngredient", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIngredients(resp)
}

func (c *Client) PatchMedicinalProductIngredientByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductIngredient) (*models.MedicinalProductIngredient, error) {
	resp, err := c.PatchByID(ctx, "MedicinalProductIngredient", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIngredient(id, resp)
}

func (c *Client) DeleteMedicinalProductIngredient(ctx context.Context, resource string, params Parameters) ([]*models.MedicinalProductIngredient, error) {
	resp, err := c.Delete(ctx, "MedicinalProductIngredient", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIngredients(resp)
}

func (c *Client) DeleteMedicinalProductIngredientByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicinalProductIngredient, error) {
	resp, err := c.DeleteByID(ctx, "MedicinalProductIngredient", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductIngredient(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicinalProductInteraction
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicinalProductInteractions(bundle *models.Bundle) ([]*models.MedicinalProductInteraction, error) {
	var entities []*models.MedicinalProductInteraction
	err := EnumBundleResources(bundle, "MedicinalProductInteraction", func(resource ResourceData) error {
		var entity models.MedicinalProductInteraction
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicinalProductInteractions(resp *FhirResponse) ([]*models.MedicinalProductInteraction, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicinalProductInteractions(resp.Bundle)
	case "MedicinalProductInteraction":
		var entity models.MedicinalProductInteraction
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicinalProductInteraction{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicinalProductInteraction(id string, resp *FhirResponse) (*models.MedicinalProductInteraction, error) {
	entities, err := fhirRespToMedicinalProductInteractions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicinalProductInteraction", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicinalProductInteraction.
func (c *Client) GetMedicinalProductInteraction(ctx context.Context, params Parameters) ([]*models.MedicinalProductInteraction, error) {
	resp, err := c.Get(ctx, "MedicinalProductInteraction", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductInteractions(resp)
}

// Get MedicinalProductInteraction by ID.
func (c *Client) GetMedicinalProductInteractionByID(ctx context.Context, id string, params Parameters) (*models.MedicinalProductInteraction, error) {
	resp, err := c.GetByID(ctx, "MedicinalProductInteraction", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductInteraction(id, resp)
}

func (c *Client) CreateMedicinalProductInteraction(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductInteraction) (*models.MedicinalProductInteraction, error) {
	resp, err := c.Create(ctx, "MedicinalProductInteraction", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductInteraction(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductInteraction(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductInteraction) (*models.MedicinalProductInteraction, error) {
	resp, err := c.Update(ctx, "MedicinalProductInteraction", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductInteraction(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductInteractionByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductInteraction) (*models.MedicinalProductInteraction, error) {
	resp, err := c.UpdateByID(ctx, "MedicinalProductInteraction", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductInteraction(id, resp)
}

func (c *Client) PatchMedicinalProductInteraction(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductInteraction) ([]*models.MedicinalProductInteraction, error) {
	resp, err := c.Patch(ctx, "MedicinalProductInteraction", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductInteractions(resp)
}

func (c *Client) PatchMedicinalProductInteractionByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductInteraction) (*models.MedicinalProductInteraction, error) {
	resp, err := c.PatchByID(ctx, "MedicinalProductInteraction", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductInteraction(id, resp)
}

func (c *Client) DeleteMedicinalProductInteraction(ctx context.Context, resource string, params Parameters) ([]*models.MedicinalProductInteraction, error) {
	resp, err := c.Delete(ctx, "MedicinalProductInteraction", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductInteractions(resp)
}

func (c *Client) DeleteMedicinalProductInteractionByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicinalProductInteraction, error) {
	resp, err := c.DeleteByID(ctx, "MedicinalProductInteraction", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductInteraction(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicinalProductManufactured
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicinalProductManufactureds(bundle *models.Bundle) ([]*models.MedicinalProductManufactured, error) {
	var entities []*models.MedicinalProductManufactured
	err := EnumBundleResources(bundle, "MedicinalProductManufactured", func(resource ResourceData) error {
		var entity models.MedicinalProductManufactured
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicinalProductManufactureds(resp *FhirResponse) ([]*models.MedicinalProductManufactured, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicinalProductManufactureds(resp.Bundle)
	case "MedicinalProductManufactured":
		var entity models.MedicinalProductManufactured
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicinalProductManufactured{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicinalProductManufactured(id string, resp *FhirResponse) (*models.MedicinalProductManufactured, error) {
	entities, err := fhirRespToMedicinalProductManufactureds(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicinalProductManufactured", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicinalProductManufactured.
func (c *Client) GetMedicinalProductManufactured(ctx context.Context, params Parameters) ([]*models.MedicinalProductManufactured, error) {
	resp, err := c.Get(ctx, "MedicinalProductManufactured", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductManufactureds(resp)
}

// Get MedicinalProductManufactured by ID.
func (c *Client) GetMedicinalProductManufacturedByID(ctx context.Context, id string, params Parameters) (*models.MedicinalProductManufactured, error) {
	resp, err := c.GetByID(ctx, "MedicinalProductManufactured", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductManufactured(id, resp)
}

func (c *Client) CreateMedicinalProductManufactured(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductManufactured) (*models.MedicinalProductManufactured, error) {
	resp, err := c.Create(ctx, "MedicinalProductManufactured", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductManufactured(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductManufactured(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductManufactured) (*models.MedicinalProductManufactured, error) {
	resp, err := c.Update(ctx, "MedicinalProductManufactured", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductManufactured(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductManufacturedByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductManufactured) (*models.MedicinalProductManufactured, error) {
	resp, err := c.UpdateByID(ctx, "MedicinalProductManufactured", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductManufactured(id, resp)
}

func (c *Client) PatchMedicinalProductManufactured(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductManufactured) ([]*models.MedicinalProductManufactured, error) {
	resp, err := c.Patch(ctx, "MedicinalProductManufactured", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductManufactureds(resp)
}

func (c *Client) PatchMedicinalProductManufacturedByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductManufactured) (*models.MedicinalProductManufactured, error) {
	resp, err := c.PatchByID(ctx, "MedicinalProductManufactured", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductManufactured(id, resp)
}

func (c *Client) DeleteMedicinalProductManufactured(ctx context.Context, resource string, params Parameters) ([]*models.MedicinalProductManufactured, error) {
	resp, err := c.Delete(ctx, "MedicinalProductManufactured", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductManufactureds(resp)
}

func (c *Client) DeleteMedicinalProductManufacturedByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicinalProductManufactured, error) {
	resp, err := c.DeleteByID(ctx, "MedicinalProductManufactured", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductManufactured(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicinalProductPackaged
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicinalProductPackageds(bundle *models.Bundle) ([]*models.MedicinalProductPackaged, error) {
	var entities []*models.MedicinalProductPackaged
	err := EnumBundleResources(bundle, "MedicinalProductPackaged", func(resource ResourceData) error {
		var entity models.MedicinalProductPackaged
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicinalProductPackageds(resp *FhirResponse) ([]*models.MedicinalProductPackaged, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicinalProductPackageds(resp.Bundle)
	case "MedicinalProductPackaged":
		var entity models.MedicinalProductPackaged
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicinalProductPackaged{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicinalProductPackaged(id string, resp *FhirResponse) (*models.MedicinalProductPackaged, error) {
	entities, err := fhirRespToMedicinalProductPackageds(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicinalProductPackaged", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicinalProductPackaged.
func (c *Client) GetMedicinalProductPackaged(ctx context.Context, params Parameters) ([]*models.MedicinalProductPackaged, error) {
	resp, err := c.Get(ctx, "MedicinalProductPackaged", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPackageds(resp)
}

// Get MedicinalProductPackaged by ID.
func (c *Client) GetMedicinalProductPackagedByID(ctx context.Context, id string, params Parameters) (*models.MedicinalProductPackaged, error) {
	resp, err := c.GetByID(ctx, "MedicinalProductPackaged", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPackaged(id, resp)
}

func (c *Client) CreateMedicinalProductPackaged(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductPackaged) (*models.MedicinalProductPackaged, error) {
	resp, err := c.Create(ctx, "MedicinalProductPackaged", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPackaged(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductPackaged(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductPackaged) (*models.MedicinalProductPackaged, error) {
	resp, err := c.Update(ctx, "MedicinalProductPackaged", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPackaged(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductPackagedByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductPackaged) (*models.MedicinalProductPackaged, error) {
	resp, err := c.UpdateByID(ctx, "MedicinalProductPackaged", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPackaged(id, resp)
}

func (c *Client) PatchMedicinalProductPackaged(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductPackaged) ([]*models.MedicinalProductPackaged, error) {
	resp, err := c.Patch(ctx, "MedicinalProductPackaged", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPackageds(resp)
}

func (c *Client) PatchMedicinalProductPackagedByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductPackaged) (*models.MedicinalProductPackaged, error) {
	resp, err := c.PatchByID(ctx, "MedicinalProductPackaged", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPackaged(id, resp)
}

func (c *Client) DeleteMedicinalProductPackaged(ctx context.Context, resource string, params Parameters) ([]*models.MedicinalProductPackaged, error) {
	resp, err := c.Delete(ctx, "MedicinalProductPackaged", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPackageds(resp)
}

func (c *Client) DeleteMedicinalProductPackagedByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicinalProductPackaged, error) {
	resp, err := c.DeleteByID(ctx, "MedicinalProductPackaged", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPackaged(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicinalProductPharmaceutical
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicinalProductPharmaceuticals(bundle *models.Bundle) ([]*models.MedicinalProductPharmaceutical, error) {
	var entities []*models.MedicinalProductPharmaceutical
	err := EnumBundleResources(bundle, "MedicinalProductPharmaceutical", func(resource ResourceData) error {
		var entity models.MedicinalProductPharmaceutical
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicinalProductPharmaceuticals(resp *FhirResponse) ([]*models.MedicinalProductPharmaceutical, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicinalProductPharmaceuticals(resp.Bundle)
	case "MedicinalProductPharmaceutical":
		var entity models.MedicinalProductPharmaceutical
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicinalProductPharmaceutical{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicinalProductPharmaceutical(id string, resp *FhirResponse) (*models.MedicinalProductPharmaceutical, error) {
	entities, err := fhirRespToMedicinalProductPharmaceuticals(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicinalProductPharmaceutical", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicinalProductPharmaceutical.
func (c *Client) GetMedicinalProductPharmaceutical(ctx context.Context, params Parameters) ([]*models.MedicinalProductPharmaceutical, error) {
	resp, err := c.Get(ctx, "MedicinalProductPharmaceutical", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPharmaceuticals(resp)
}

// Get MedicinalProductPharmaceutical by ID.
func (c *Client) GetMedicinalProductPharmaceuticalByID(ctx context.Context, id string, params Parameters) (*models.MedicinalProductPharmaceutical, error) {
	resp, err := c.GetByID(ctx, "MedicinalProductPharmaceutical", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPharmaceutical(id, resp)
}

func (c *Client) CreateMedicinalProductPharmaceutical(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductPharmaceutical) (*models.MedicinalProductPharmaceutical, error) {
	resp, err := c.Create(ctx, "MedicinalProductPharmaceutical", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPharmaceutical(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductPharmaceutical(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductPharmaceutical) (*models.MedicinalProductPharmaceutical, error) {
	resp, err := c.Update(ctx, "MedicinalProductPharmaceutical", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPharmaceutical(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductPharmaceuticalByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductPharmaceutical) (*models.MedicinalProductPharmaceutical, error) {
	resp, err := c.UpdateByID(ctx, "MedicinalProductPharmaceutical", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPharmaceutical(id, resp)
}

func (c *Client) PatchMedicinalProductPharmaceutical(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductPharmaceutical) ([]*models.MedicinalProductPharmaceutical, error) {
	resp, err := c.Patch(ctx, "MedicinalProductPharmaceutical", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPharmaceuticals(resp)
}

func (c *Client) PatchMedicinalProductPharmaceuticalByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductPharmaceutical) (*models.MedicinalProductPharmaceutical, error) {
	resp, err := c.PatchByID(ctx, "MedicinalProductPharmaceutical", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPharmaceutical(id, resp)
}

func (c *Client) DeleteMedicinalProductPharmaceutical(ctx context.Context, resource string, params Parameters) ([]*models.MedicinalProductPharmaceutical, error) {
	resp, err := c.Delete(ctx, "MedicinalProductPharmaceutical", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPharmaceuticals(resp)
}

func (c *Client) DeleteMedicinalProductPharmaceuticalByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicinalProductPharmaceutical, error) {
	resp, err := c.DeleteByID(ctx, "MedicinalProductPharmaceutical", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductPharmaceutical(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MedicinalProductUndesirableEffect
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMedicinalProductUndesirableEffects(bundle *models.Bundle) ([]*models.MedicinalProductUndesirableEffect, error) {
	var entities []*models.MedicinalProductUndesirableEffect
	err := EnumBundleResources(bundle, "MedicinalProductUndesirableEffect", func(resource ResourceData) error {
		var entity models.MedicinalProductUndesirableEffect
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMedicinalProductUndesirableEffects(resp *FhirResponse) ([]*models.MedicinalProductUndesirableEffect, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMedicinalProductUndesirableEffects(resp.Bundle)
	case "MedicinalProductUndesirableEffect":
		var entity models.MedicinalProductUndesirableEffect
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MedicinalProductUndesirableEffect{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMedicinalProductUndesirableEffect(id string, resp *FhirResponse) (*models.MedicinalProductUndesirableEffect, error) {
	entities, err := fhirRespToMedicinalProductUndesirableEffects(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MedicinalProductUndesirableEffect", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MedicinalProductUndesirableEffect.
func (c *Client) GetMedicinalProductUndesirableEffect(ctx context.Context, params Parameters) ([]*models.MedicinalProductUndesirableEffect, error) {
	resp, err := c.Get(ctx, "MedicinalProductUndesirableEffect", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductUndesirableEffects(resp)
}

// Get MedicinalProductUndesirableEffect by ID.
func (c *Client) GetMedicinalProductUndesirableEffectByID(ctx context.Context, id string, params Parameters) (*models.MedicinalProductUndesirableEffect, error) {
	resp, err := c.GetByID(ctx, "MedicinalProductUndesirableEffect", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductUndesirableEffect(id, resp)
}

func (c *Client) CreateMedicinalProductUndesirableEffect(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductUndesirableEffect) (*models.MedicinalProductUndesirableEffect, error) {
	resp, err := c.Create(ctx, "MedicinalProductUndesirableEffect", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductUndesirableEffect(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductUndesirableEffect(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductUndesirableEffect) (*models.MedicinalProductUndesirableEffect, error) {
	resp, err := c.Update(ctx, "MedicinalProductUndesirableEffect", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductUndesirableEffect(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMedicinalProductUndesirableEffectByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductUndesirableEffect) (*models.MedicinalProductUndesirableEffect, error) {
	resp, err := c.UpdateByID(ctx, "MedicinalProductUndesirableEffect", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductUndesirableEffect(id, resp)
}

func (c *Client) PatchMedicinalProductUndesirableEffect(ctx context.Context, resource string, params Parameters, entity *models.MedicinalProductUndesirableEffect) ([]*models.MedicinalProductUndesirableEffect, error) {
	resp, err := c.Patch(ctx, "MedicinalProductUndesirableEffect", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductUndesirableEffects(resp)
}

func (c *Client) PatchMedicinalProductUndesirableEffectByID(ctx context.Context, resource, id string, params Parameters, entity *models.MedicinalProductUndesirableEffect) (*models.MedicinalProductUndesirableEffect, error) {
	resp, err := c.PatchByID(ctx, "MedicinalProductUndesirableEffect", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductUndesirableEffect(id, resp)
}

func (c *Client) DeleteMedicinalProductUndesirableEffect(ctx context.Context, resource string, params Parameters) ([]*models.MedicinalProductUndesirableEffect, error) {
	resp, err := c.Delete(ctx, "MedicinalProductUndesirableEffect", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductUndesirableEffects(resp)
}

func (c *Client) DeleteMedicinalProductUndesirableEffectByID(ctx context.Context, resource, id string, params Parameters) (*models.MedicinalProductUndesirableEffect, error) {
	resp, err := c.DeleteByID(ctx, "MedicinalProductUndesirableEffect", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMedicinalProductUndesirableEffect(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MessageDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMessageDefinitions(bundle *models.Bundle) ([]*models.MessageDefinition, error) {
	var entities []*models.MessageDefinition
	err := EnumBundleResources(bundle, "MessageDefinition", func(resource ResourceData) error {
		var entity models.MessageDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMessageDefinitions(resp *FhirResponse) ([]*models.MessageDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMessageDefinitions(resp.Bundle)
	case "MessageDefinition":
		var entity models.MessageDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MessageDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMessageDefinition(id string, resp *FhirResponse) (*models.MessageDefinition, error) {
	entities, err := fhirRespToMessageDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MessageDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MessageDefinition.
func (c *Client) GetMessageDefinition(ctx context.Context, params Parameters) ([]*models.MessageDefinition, error) {
	resp, err := c.Get(ctx, "MessageDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageDefinitions(resp)
}

// Get MessageDefinition by ID.
func (c *Client) GetMessageDefinitionByID(ctx context.Context, id string, params Parameters) (*models.MessageDefinition, error) {
	resp, err := c.GetByID(ctx, "MessageDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageDefinition(id, resp)
}

func (c *Client) CreateMessageDefinition(ctx context.Context, resource string, params Parameters, entity *models.MessageDefinition) (*models.MessageDefinition, error) {
	resp, err := c.Create(ctx, "MessageDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMessageDefinition(ctx context.Context, resource string, params Parameters, entity *models.MessageDefinition) (*models.MessageDefinition, error) {
	resp, err := c.Update(ctx, "MessageDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMessageDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.MessageDefinition) (*models.MessageDefinition, error) {
	resp, err := c.UpdateByID(ctx, "MessageDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageDefinition(id, resp)
}

func (c *Client) PatchMessageDefinition(ctx context.Context, resource string, params Parameters, entity *models.MessageDefinition) ([]*models.MessageDefinition, error) {
	resp, err := c.Patch(ctx, "MessageDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageDefinitions(resp)
}

func (c *Client) PatchMessageDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.MessageDefinition) (*models.MessageDefinition, error) {
	resp, err := c.PatchByID(ctx, "MessageDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageDefinition(id, resp)
}

func (c *Client) DeleteMessageDefinition(ctx context.Context, resource string, params Parameters) ([]*models.MessageDefinition, error) {
	resp, err := c.Delete(ctx, "MessageDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageDefinitions(resp)
}

func (c *Client) DeleteMessageDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.MessageDefinition, error) {
	resp, err := c.DeleteByID(ctx, "MessageDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MessageHeader
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMessageHeaders(bundle *models.Bundle) ([]*models.MessageHeader, error) {
	var entities []*models.MessageHeader
	err := EnumBundleResources(bundle, "MessageHeader", func(resource ResourceData) error {
		var entity models.MessageHeader
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMessageHeaders(resp *FhirResponse) ([]*models.MessageHeader, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMessageHeaders(resp.Bundle)
	case "MessageHeader":
		var entity models.MessageHeader
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MessageHeader{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMessageHeader(id string, resp *FhirResponse) (*models.MessageHeader, error) {
	entities, err := fhirRespToMessageHeaders(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MessageHeader", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MessageHeader.
func (c *Client) GetMessageHeader(ctx context.Context, params Parameters) ([]*models.MessageHeader, error) {
	resp, err := c.Get(ctx, "MessageHeader", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageHeaders(resp)
}

// Get MessageHeader by ID.
func (c *Client) GetMessageHeaderByID(ctx context.Context, id string, params Parameters) (*models.MessageHeader, error) {
	resp, err := c.GetByID(ctx, "MessageHeader", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageHeader(id, resp)
}

func (c *Client) CreateMessageHeader(ctx context.Context, resource string, params Parameters, entity *models.MessageHeader) (*models.MessageHeader, error) {
	resp, err := c.Create(ctx, "MessageHeader", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageHeader(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMessageHeader(ctx context.Context, resource string, params Parameters, entity *models.MessageHeader) (*models.MessageHeader, error) {
	resp, err := c.Update(ctx, "MessageHeader", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageHeader(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMessageHeaderByID(ctx context.Context, resource, id string, params Parameters, entity *models.MessageHeader) (*models.MessageHeader, error) {
	resp, err := c.UpdateByID(ctx, "MessageHeader", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageHeader(id, resp)
}

func (c *Client) PatchMessageHeader(ctx context.Context, resource string, params Parameters, entity *models.MessageHeader) ([]*models.MessageHeader, error) {
	resp, err := c.Patch(ctx, "MessageHeader", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageHeaders(resp)
}

func (c *Client) PatchMessageHeaderByID(ctx context.Context, resource, id string, params Parameters, entity *models.MessageHeader) (*models.MessageHeader, error) {
	resp, err := c.PatchByID(ctx, "MessageHeader", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageHeader(id, resp)
}

func (c *Client) DeleteMessageHeader(ctx context.Context, resource string, params Parameters) ([]*models.MessageHeader, error) {
	resp, err := c.Delete(ctx, "MessageHeader", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageHeaders(resp)
}

func (c *Client) DeleteMessageHeaderByID(ctx context.Context, resource, id string, params Parameters) (*models.MessageHeader, error) {
	resp, err := c.DeleteByID(ctx, "MessageHeader", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMessageHeader(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// MolecularSequence
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToMolecularSequences(bundle *models.Bundle) ([]*models.MolecularSequence, error) {
	var entities []*models.MolecularSequence
	err := EnumBundleResources(bundle, "MolecularSequence", func(resource ResourceData) error {
		var entity models.MolecularSequence
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToMolecularSequences(resp *FhirResponse) ([]*models.MolecularSequence, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToMolecularSequences(resp.Bundle)
	case "MolecularSequence":
		var entity models.MolecularSequence
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.MolecularSequence{&entity}, nil
	}
	return nil, nil
}

func fhirRespToMolecularSequence(id string, resp *FhirResponse) (*models.MolecularSequence, error) {
	entities, err := fhirRespToMolecularSequences(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "MolecularSequence", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get MolecularSequence.
func (c *Client) GetMolecularSequence(ctx context.Context, params Parameters) ([]*models.MolecularSequence, error) {
	resp, err := c.Get(ctx, "MolecularSequence", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMolecularSequences(resp)
}

// Get MolecularSequence by ID.
func (c *Client) GetMolecularSequenceByID(ctx context.Context, id string, params Parameters) (*models.MolecularSequence, error) {
	resp, err := c.GetByID(ctx, "MolecularSequence", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMolecularSequence(id, resp)
}

func (c *Client) CreateMolecularSequence(ctx context.Context, resource string, params Parameters, entity *models.MolecularSequence) (*models.MolecularSequence, error) {
	resp, err := c.Create(ctx, "MolecularSequence", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMolecularSequence(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMolecularSequence(ctx context.Context, resource string, params Parameters, entity *models.MolecularSequence) (*models.MolecularSequence, error) {
	resp, err := c.Update(ctx, "MolecularSequence", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMolecularSequence(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateMolecularSequenceByID(ctx context.Context, resource, id string, params Parameters, entity *models.MolecularSequence) (*models.MolecularSequence, error) {
	resp, err := c.UpdateByID(ctx, "MolecularSequence", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMolecularSequence(id, resp)
}

func (c *Client) PatchMolecularSequence(ctx context.Context, resource string, params Parameters, entity *models.MolecularSequence) ([]*models.MolecularSequence, error) {
	resp, err := c.Patch(ctx, "MolecularSequence", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMolecularSequences(resp)
}

func (c *Client) PatchMolecularSequenceByID(ctx context.Context, resource, id string, params Parameters, entity *models.MolecularSequence) (*models.MolecularSequence, error) {
	resp, err := c.PatchByID(ctx, "MolecularSequence", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToMolecularSequence(id, resp)
}

func (c *Client) DeleteMolecularSequence(ctx context.Context, resource string, params Parameters) ([]*models.MolecularSequence, error) {
	resp, err := c.Delete(ctx, "MolecularSequence", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMolecularSequences(resp)
}

func (c *Client) DeleteMolecularSequenceByID(ctx context.Context, resource, id string, params Parameters) (*models.MolecularSequence, error) {
	resp, err := c.DeleteByID(ctx, "MolecularSequence", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToMolecularSequence(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// NamingSystem
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToNamingSystems(bundle *models.Bundle) ([]*models.NamingSystem, error) {
	var entities []*models.NamingSystem
	err := EnumBundleResources(bundle, "NamingSystem", func(resource ResourceData) error {
		var entity models.NamingSystem
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToNamingSystems(resp *FhirResponse) ([]*models.NamingSystem, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToNamingSystems(resp.Bundle)
	case "NamingSystem":
		var entity models.NamingSystem
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.NamingSystem{&entity}, nil
	}
	return nil, nil
}

func fhirRespToNamingSystem(id string, resp *FhirResponse) (*models.NamingSystem, error) {
	entities, err := fhirRespToNamingSystems(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "NamingSystem", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get NamingSystem.
func (c *Client) GetNamingSystem(ctx context.Context, params Parameters) ([]*models.NamingSystem, error) {
	resp, err := c.Get(ctx, "NamingSystem", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToNamingSystems(resp)
}

// Get NamingSystem by ID.
func (c *Client) GetNamingSystemByID(ctx context.Context, id string, params Parameters) (*models.NamingSystem, error) {
	resp, err := c.GetByID(ctx, "NamingSystem", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToNamingSystem(id, resp)
}

func (c *Client) CreateNamingSystem(ctx context.Context, resource string, params Parameters, entity *models.NamingSystem) (*models.NamingSystem, error) {
	resp, err := c.Create(ctx, "NamingSystem", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToNamingSystem(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateNamingSystem(ctx context.Context, resource string, params Parameters, entity *models.NamingSystem) (*models.NamingSystem, error) {
	resp, err := c.Update(ctx, "NamingSystem", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToNamingSystem(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateNamingSystemByID(ctx context.Context, resource, id string, params Parameters, entity *models.NamingSystem) (*models.NamingSystem, error) {
	resp, err := c.UpdateByID(ctx, "NamingSystem", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToNamingSystem(id, resp)
}

func (c *Client) PatchNamingSystem(ctx context.Context, resource string, params Parameters, entity *models.NamingSystem) ([]*models.NamingSystem, error) {
	resp, err := c.Patch(ctx, "NamingSystem", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToNamingSystems(resp)
}

func (c *Client) PatchNamingSystemByID(ctx context.Context, resource, id string, params Parameters, entity *models.NamingSystem) (*models.NamingSystem, error) {
	resp, err := c.PatchByID(ctx, "NamingSystem", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToNamingSystem(id, resp)
}

func (c *Client) DeleteNamingSystem(ctx context.Context, resource string, params Parameters) ([]*models.NamingSystem, error) {
	resp, err := c.Delete(ctx, "NamingSystem", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToNamingSystems(resp)
}

func (c *Client) DeleteNamingSystemByID(ctx context.Context, resource, id string, params Parameters) (*models.NamingSystem, error) {
	resp, err := c.DeleteByID(ctx, "NamingSystem", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToNamingSystem(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// NutritionOrder
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToNutritionOrders(bundle *models.Bundle) ([]*models.NutritionOrder, error) {
	var entities []*models.NutritionOrder
	err := EnumBundleResources(bundle, "NutritionOrder", func(resource ResourceData) error {
		var entity models.NutritionOrder
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToNutritionOrders(resp *FhirResponse) ([]*models.NutritionOrder, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToNutritionOrders(resp.Bundle)
	case "NutritionOrder":
		var entity models.NutritionOrder
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.NutritionOrder{&entity}, nil
	}
	return nil, nil
}

func fhirRespToNutritionOrder(id string, resp *FhirResponse) (*models.NutritionOrder, error) {
	entities, err := fhirRespToNutritionOrders(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "NutritionOrder", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get NutritionOrder.
func (c *Client) GetNutritionOrder(ctx context.Context, params Parameters) ([]*models.NutritionOrder, error) {
	resp, err := c.Get(ctx, "NutritionOrder", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToNutritionOrders(resp)
}

// Get NutritionOrder by ID.
func (c *Client) GetNutritionOrderByID(ctx context.Context, id string, params Parameters) (*models.NutritionOrder, error) {
	resp, err := c.GetByID(ctx, "NutritionOrder", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToNutritionOrder(id, resp)
}

func (c *Client) CreateNutritionOrder(ctx context.Context, resource string, params Parameters, entity *models.NutritionOrder) (*models.NutritionOrder, error) {
	resp, err := c.Create(ctx, "NutritionOrder", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToNutritionOrder(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateNutritionOrder(ctx context.Context, resource string, params Parameters, entity *models.NutritionOrder) (*models.NutritionOrder, error) {
	resp, err := c.Update(ctx, "NutritionOrder", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToNutritionOrder(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateNutritionOrderByID(ctx context.Context, resource, id string, params Parameters, entity *models.NutritionOrder) (*models.NutritionOrder, error) {
	resp, err := c.UpdateByID(ctx, "NutritionOrder", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToNutritionOrder(id, resp)
}

func (c *Client) PatchNutritionOrder(ctx context.Context, resource string, params Parameters, entity *models.NutritionOrder) ([]*models.NutritionOrder, error) {
	resp, err := c.Patch(ctx, "NutritionOrder", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToNutritionOrders(resp)
}

func (c *Client) PatchNutritionOrderByID(ctx context.Context, resource, id string, params Parameters, entity *models.NutritionOrder) (*models.NutritionOrder, error) {
	resp, err := c.PatchByID(ctx, "NutritionOrder", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToNutritionOrder(id, resp)
}

func (c *Client) DeleteNutritionOrder(ctx context.Context, resource string, params Parameters) ([]*models.NutritionOrder, error) {
	resp, err := c.Delete(ctx, "NutritionOrder", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToNutritionOrders(resp)
}

func (c *Client) DeleteNutritionOrderByID(ctx context.Context, resource, id string, params Parameters) (*models.NutritionOrder, error) {
	resp, err := c.DeleteByID(ctx, "NutritionOrder", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToNutritionOrder(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Observation
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToObservations(bundle *models.Bundle) ([]*models.Observation, error) {
	var entities []*models.Observation
	err := EnumBundleResources(bundle, "Observation", func(resource ResourceData) error {
		var entity models.Observation
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToObservations(resp *FhirResponse) ([]*models.Observation, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToObservations(resp.Bundle)
	case "Observation":
		var entity models.Observation
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Observation{&entity}, nil
	}
	return nil, nil
}

func fhirRespToObservation(id string, resp *FhirResponse) (*models.Observation, error) {
	entities, err := fhirRespToObservations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Observation", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Observation.
func (c *Client) GetObservation(ctx context.Context, params Parameters) ([]*models.Observation, error) {
	resp, err := c.Get(ctx, "Observation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservations(resp)
}

// Get Observation by ID.
func (c *Client) GetObservationByID(ctx context.Context, id string, params Parameters) (*models.Observation, error) {
	resp, err := c.GetByID(ctx, "Observation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservation(id, resp)
}

func (c *Client) CreateObservation(ctx context.Context, resource string, params Parameters, entity *models.Observation) (*models.Observation, error) {
	resp, err := c.Create(ctx, "Observation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateObservation(ctx context.Context, resource string, params Parameters, entity *models.Observation) (*models.Observation, error) {
	resp, err := c.Update(ctx, "Observation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateObservationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Observation) (*models.Observation, error) {
	resp, err := c.UpdateByID(ctx, "Observation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservation(id, resp)
}

func (c *Client) PatchObservation(ctx context.Context, resource string, params Parameters, entity *models.Observation) ([]*models.Observation, error) {
	resp, err := c.Patch(ctx, "Observation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservations(resp)
}

func (c *Client) PatchObservationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Observation) (*models.Observation, error) {
	resp, err := c.PatchByID(ctx, "Observation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservation(id, resp)
}

func (c *Client) DeleteObservation(ctx context.Context, resource string, params Parameters) ([]*models.Observation, error) {
	resp, err := c.Delete(ctx, "Observation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservations(resp)
}

func (c *Client) DeleteObservationByID(ctx context.Context, resource, id string, params Parameters) (*models.Observation, error) {
	resp, err := c.DeleteByID(ctx, "Observation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservation(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ObservationDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToObservationDefinitions(bundle *models.Bundle) ([]*models.ObservationDefinition, error) {
	var entities []*models.ObservationDefinition
	err := EnumBundleResources(bundle, "ObservationDefinition", func(resource ResourceData) error {
		var entity models.ObservationDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToObservationDefinitions(resp *FhirResponse) ([]*models.ObservationDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToObservationDefinitions(resp.Bundle)
	case "ObservationDefinition":
		var entity models.ObservationDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ObservationDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToObservationDefinition(id string, resp *FhirResponse) (*models.ObservationDefinition, error) {
	entities, err := fhirRespToObservationDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ObservationDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ObservationDefinition.
func (c *Client) GetObservationDefinition(ctx context.Context, params Parameters) ([]*models.ObservationDefinition, error) {
	resp, err := c.Get(ctx, "ObservationDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservationDefinitions(resp)
}

// Get ObservationDefinition by ID.
func (c *Client) GetObservationDefinitionByID(ctx context.Context, id string, params Parameters) (*models.ObservationDefinition, error) {
	resp, err := c.GetByID(ctx, "ObservationDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservationDefinition(id, resp)
}

func (c *Client) CreateObservationDefinition(ctx context.Context, resource string, params Parameters, entity *models.ObservationDefinition) (*models.ObservationDefinition, error) {
	resp, err := c.Create(ctx, "ObservationDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservationDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateObservationDefinition(ctx context.Context, resource string, params Parameters, entity *models.ObservationDefinition) (*models.ObservationDefinition, error) {
	resp, err := c.Update(ctx, "ObservationDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservationDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateObservationDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ObservationDefinition) (*models.ObservationDefinition, error) {
	resp, err := c.UpdateByID(ctx, "ObservationDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservationDefinition(id, resp)
}

func (c *Client) PatchObservationDefinition(ctx context.Context, resource string, params Parameters, entity *models.ObservationDefinition) ([]*models.ObservationDefinition, error) {
	resp, err := c.Patch(ctx, "ObservationDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservationDefinitions(resp)
}

func (c *Client) PatchObservationDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ObservationDefinition) (*models.ObservationDefinition, error) {
	resp, err := c.PatchByID(ctx, "ObservationDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservationDefinition(id, resp)
}

func (c *Client) DeleteObservationDefinition(ctx context.Context, resource string, params Parameters) ([]*models.ObservationDefinition, error) {
	resp, err := c.Delete(ctx, "ObservationDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservationDefinitions(resp)
}

func (c *Client) DeleteObservationDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.ObservationDefinition, error) {
	resp, err := c.DeleteByID(ctx, "ObservationDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToObservationDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// OperationDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToOperationDefinitions(bundle *models.Bundle) ([]*models.OperationDefinition, error) {
	var entities []*models.OperationDefinition
	err := EnumBundleResources(bundle, "OperationDefinition", func(resource ResourceData) error {
		var entity models.OperationDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToOperationDefinitions(resp *FhirResponse) ([]*models.OperationDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToOperationDefinitions(resp.Bundle)
	case "OperationDefinition":
		var entity models.OperationDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.OperationDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToOperationDefinition(id string, resp *FhirResponse) (*models.OperationDefinition, error) {
	entities, err := fhirRespToOperationDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "OperationDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get OperationDefinition.
func (c *Client) GetOperationDefinition(ctx context.Context, params Parameters) ([]*models.OperationDefinition, error) {
	resp, err := c.Get(ctx, "OperationDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationDefinitions(resp)
}

// Get OperationDefinition by ID.
func (c *Client) GetOperationDefinitionByID(ctx context.Context, id string, params Parameters) (*models.OperationDefinition, error) {
	resp, err := c.GetByID(ctx, "OperationDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationDefinition(id, resp)
}

func (c *Client) CreateOperationDefinition(ctx context.Context, resource string, params Parameters, entity *models.OperationDefinition) (*models.OperationDefinition, error) {
	resp, err := c.Create(ctx, "OperationDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateOperationDefinition(ctx context.Context, resource string, params Parameters, entity *models.OperationDefinition) (*models.OperationDefinition, error) {
	resp, err := c.Update(ctx, "OperationDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateOperationDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.OperationDefinition) (*models.OperationDefinition, error) {
	resp, err := c.UpdateByID(ctx, "OperationDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationDefinition(id, resp)
}

func (c *Client) PatchOperationDefinition(ctx context.Context, resource string, params Parameters, entity *models.OperationDefinition) ([]*models.OperationDefinition, error) {
	resp, err := c.Patch(ctx, "OperationDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationDefinitions(resp)
}

func (c *Client) PatchOperationDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.OperationDefinition) (*models.OperationDefinition, error) {
	resp, err := c.PatchByID(ctx, "OperationDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationDefinition(id, resp)
}

func (c *Client) DeleteOperationDefinition(ctx context.Context, resource string, params Parameters) ([]*models.OperationDefinition, error) {
	resp, err := c.Delete(ctx, "OperationDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationDefinitions(resp)
}

func (c *Client) DeleteOperationDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.OperationDefinition, error) {
	resp, err := c.DeleteByID(ctx, "OperationDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// OperationOutcome
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToOperationOutcomes(bundle *models.Bundle) ([]*models.OperationOutcome, error) {
	var entities []*models.OperationOutcome
	err := EnumBundleResources(bundle, "OperationOutcome", func(resource ResourceData) error {
		var entity models.OperationOutcome
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToOperationOutcomes(resp *FhirResponse) ([]*models.OperationOutcome, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToOperationOutcomes(resp.Bundle)
	case "OperationOutcome":
		var entity models.OperationOutcome
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.OperationOutcome{&entity}, nil
	}
	return nil, nil
}

func fhirRespToOperationOutcome(id string, resp *FhirResponse) (*models.OperationOutcome, error) {
	entities, err := fhirRespToOperationOutcomes(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "OperationOutcome", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get OperationOutcome.
func (c *Client) GetOperationOutcome(ctx context.Context, params Parameters) ([]*models.OperationOutcome, error) {
	resp, err := c.Get(ctx, "OperationOutcome", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationOutcomes(resp)
}

// Get OperationOutcome by ID.
func (c *Client) GetOperationOutcomeByID(ctx context.Context, id string, params Parameters) (*models.OperationOutcome, error) {
	resp, err := c.GetByID(ctx, "OperationOutcome", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationOutcome(id, resp)
}

func (c *Client) CreateOperationOutcome(ctx context.Context, resource string, params Parameters, entity *models.OperationOutcome) (*models.OperationOutcome, error) {
	resp, err := c.Create(ctx, "OperationOutcome", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationOutcome(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateOperationOutcome(ctx context.Context, resource string, params Parameters, entity *models.OperationOutcome) (*models.OperationOutcome, error) {
	resp, err := c.Update(ctx, "OperationOutcome", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationOutcome(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateOperationOutcomeByID(ctx context.Context, resource, id string, params Parameters, entity *models.OperationOutcome) (*models.OperationOutcome, error) {
	resp, err := c.UpdateByID(ctx, "OperationOutcome", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationOutcome(id, resp)
}

func (c *Client) PatchOperationOutcome(ctx context.Context, resource string, params Parameters, entity *models.OperationOutcome) ([]*models.OperationOutcome, error) {
	resp, err := c.Patch(ctx, "OperationOutcome", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationOutcomes(resp)
}

func (c *Client) PatchOperationOutcomeByID(ctx context.Context, resource, id string, params Parameters, entity *models.OperationOutcome) (*models.OperationOutcome, error) {
	resp, err := c.PatchByID(ctx, "OperationOutcome", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationOutcome(id, resp)
}

func (c *Client) DeleteOperationOutcome(ctx context.Context, resource string, params Parameters) ([]*models.OperationOutcome, error) {
	resp, err := c.Delete(ctx, "OperationOutcome", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationOutcomes(resp)
}

func (c *Client) DeleteOperationOutcomeByID(ctx context.Context, resource, id string, params Parameters) (*models.OperationOutcome, error) {
	resp, err := c.DeleteByID(ctx, "OperationOutcome", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOperationOutcome(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Organization
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToOrganizations(bundle *models.Bundle) ([]*models.Organization, error) {
	var entities []*models.Organization
	err := EnumBundleResources(bundle, "Organization", func(resource ResourceData) error {
		var entity models.Organization
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToOrganizations(resp *FhirResponse) ([]*models.Organization, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToOrganizations(resp.Bundle)
	case "Organization":
		var entity models.Organization
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Organization{&entity}, nil
	}
	return nil, nil
}

func fhirRespToOrganization(id string, resp *FhirResponse) (*models.Organization, error) {
	entities, err := fhirRespToOrganizations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Organization", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Organization.
func (c *Client) GetOrganization(ctx context.Context, params Parameters) ([]*models.Organization, error) {
	resp, err := c.Get(ctx, "Organization", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizations(resp)
}

// Get Organization by ID.
func (c *Client) GetOrganizationByID(ctx context.Context, id string, params Parameters) (*models.Organization, error) {
	resp, err := c.GetByID(ctx, "Organization", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganization(id, resp)
}

func (c *Client) CreateOrganization(ctx context.Context, resource string, params Parameters, entity *models.Organization) (*models.Organization, error) {
	resp, err := c.Create(ctx, "Organization", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganization(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateOrganization(ctx context.Context, resource string, params Parameters, entity *models.Organization) (*models.Organization, error) {
	resp, err := c.Update(ctx, "Organization", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganization(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateOrganizationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Organization) (*models.Organization, error) {
	resp, err := c.UpdateByID(ctx, "Organization", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganization(id, resp)
}

func (c *Client) PatchOrganization(ctx context.Context, resource string, params Parameters, entity *models.Organization) ([]*models.Organization, error) {
	resp, err := c.Patch(ctx, "Organization", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizations(resp)
}

func (c *Client) PatchOrganizationByID(ctx context.Context, resource, id string, params Parameters, entity *models.Organization) (*models.Organization, error) {
	resp, err := c.PatchByID(ctx, "Organization", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganization(id, resp)
}

func (c *Client) DeleteOrganization(ctx context.Context, resource string, params Parameters) ([]*models.Organization, error) {
	resp, err := c.Delete(ctx, "Organization", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizations(resp)
}

func (c *Client) DeleteOrganizationByID(ctx context.Context, resource, id string, params Parameters) (*models.Organization, error) {
	resp, err := c.DeleteByID(ctx, "Organization", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganization(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// OrganizationAffiliation
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToOrganizationAffiliations(bundle *models.Bundle) ([]*models.OrganizationAffiliation, error) {
	var entities []*models.OrganizationAffiliation
	err := EnumBundleResources(bundle, "OrganizationAffiliation", func(resource ResourceData) error {
		var entity models.OrganizationAffiliation
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToOrganizationAffiliations(resp *FhirResponse) ([]*models.OrganizationAffiliation, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToOrganizationAffiliations(resp.Bundle)
	case "OrganizationAffiliation":
		var entity models.OrganizationAffiliation
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.OrganizationAffiliation{&entity}, nil
	}
	return nil, nil
}

func fhirRespToOrganizationAffiliation(id string, resp *FhirResponse) (*models.OrganizationAffiliation, error) {
	entities, err := fhirRespToOrganizationAffiliations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "OrganizationAffiliation", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get OrganizationAffiliation.
func (c *Client) GetOrganizationAffiliation(ctx context.Context, params Parameters) ([]*models.OrganizationAffiliation, error) {
	resp, err := c.Get(ctx, "OrganizationAffiliation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizationAffiliations(resp)
}

// Get OrganizationAffiliation by ID.
func (c *Client) GetOrganizationAffiliationByID(ctx context.Context, id string, params Parameters) (*models.OrganizationAffiliation, error) {
	resp, err := c.GetByID(ctx, "OrganizationAffiliation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizationAffiliation(id, resp)
}

func (c *Client) CreateOrganizationAffiliation(ctx context.Context, resource string, params Parameters, entity *models.OrganizationAffiliation) (*models.OrganizationAffiliation, error) {
	resp, err := c.Create(ctx, "OrganizationAffiliation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizationAffiliation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateOrganizationAffiliation(ctx context.Context, resource string, params Parameters, entity *models.OrganizationAffiliation) (*models.OrganizationAffiliation, error) {
	resp, err := c.Update(ctx, "OrganizationAffiliation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizationAffiliation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateOrganizationAffiliationByID(ctx context.Context, resource, id string, params Parameters, entity *models.OrganizationAffiliation) (*models.OrganizationAffiliation, error) {
	resp, err := c.UpdateByID(ctx, "OrganizationAffiliation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizationAffiliation(id, resp)
}

func (c *Client) PatchOrganizationAffiliation(ctx context.Context, resource string, params Parameters, entity *models.OrganizationAffiliation) ([]*models.OrganizationAffiliation, error) {
	resp, err := c.Patch(ctx, "OrganizationAffiliation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizationAffiliations(resp)
}

func (c *Client) PatchOrganizationAffiliationByID(ctx context.Context, resource, id string, params Parameters, entity *models.OrganizationAffiliation) (*models.OrganizationAffiliation, error) {
	resp, err := c.PatchByID(ctx, "OrganizationAffiliation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizationAffiliation(id, resp)
}

func (c *Client) DeleteOrganizationAffiliation(ctx context.Context, resource string, params Parameters) ([]*models.OrganizationAffiliation, error) {
	resp, err := c.Delete(ctx, "OrganizationAffiliation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizationAffiliations(resp)
}

func (c *Client) DeleteOrganizationAffiliationByID(ctx context.Context, resource, id string, params Parameters) (*models.OrganizationAffiliation, error) {
	resp, err := c.DeleteByID(ctx, "OrganizationAffiliation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToOrganizationAffiliation(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Parameters
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToParameterss(bundle *models.Bundle) ([]*models.Parameters, error) {
	var entities []*models.Parameters
	err := EnumBundleResources(bundle, "Parameters", func(resource ResourceData) error {
		var entity models.Parameters
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToParameterss(resp *FhirResponse) ([]*models.Parameters, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToParameterss(resp.Bundle)
	case "Parameters":
		var entity models.Parameters
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Parameters{&entity}, nil
	}
	return nil, nil
}

func fhirRespToParameters(id string, resp *FhirResponse) (*models.Parameters, error) {
	entities, err := fhirRespToParameterss(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Parameters", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Parameters.
func (c *Client) GetParameters(ctx context.Context, params Parameters) ([]*models.Parameters, error) {
	resp, err := c.Get(ctx, "Parameters", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToParameterss(resp)
}

// Get Parameters by ID.
func (c *Client) GetParametersByID(ctx context.Context, id string, params Parameters) (*models.Parameters, error) {
	resp, err := c.GetByID(ctx, "Parameters", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToParameters(id, resp)
}

func (c *Client) CreateParameters(ctx context.Context, resource string, params Parameters, entity *models.Parameters) (*models.Parameters, error) {
	resp, err := c.Create(ctx, "Parameters", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToParameters(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateParameters(ctx context.Context, resource string, params Parameters, entity *models.Parameters) (*models.Parameters, error) {
	resp, err := c.Update(ctx, "Parameters", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToParameters(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateParametersByID(ctx context.Context, resource, id string, params Parameters, entity *models.Parameters) (*models.Parameters, error) {
	resp, err := c.UpdateByID(ctx, "Parameters", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToParameters(id, resp)
}

func (c *Client) PatchParameters(ctx context.Context, resource string, params Parameters, entity *models.Parameters) ([]*models.Parameters, error) {
	resp, err := c.Patch(ctx, "Parameters", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToParameterss(resp)
}

func (c *Client) PatchParametersByID(ctx context.Context, resource, id string, params Parameters, entity *models.Parameters) (*models.Parameters, error) {
	resp, err := c.PatchByID(ctx, "Parameters", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToParameters(id, resp)
}

func (c *Client) DeleteParameters(ctx context.Context, resource string, params Parameters) ([]*models.Parameters, error) {
	resp, err := c.Delete(ctx, "Parameters", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToParameterss(resp)
}

func (c *Client) DeleteParametersByID(ctx context.Context, resource, id string, params Parameters) (*models.Parameters, error) {
	resp, err := c.DeleteByID(ctx, "Parameters", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToParameters(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Patient
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToPatients(bundle *models.Bundle) ([]*models.Patient, error) {
	var entities []*models.Patient
	err := EnumBundleResources(bundle, "Patient", func(resource ResourceData) error {
		var entity models.Patient
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToPatients(resp *FhirResponse) ([]*models.Patient, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToPatients(resp.Bundle)
	case "Patient":
		var entity models.Patient
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Patient{&entity}, nil
	}
	return nil, nil
}

func fhirRespToPatient(id string, resp *FhirResponse) (*models.Patient, error) {
	entities, err := fhirRespToPatients(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Patient", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Patient.
func (c *Client) GetPatient(ctx context.Context, params Parameters) ([]*models.Patient, error) {
	resp, err := c.Get(ctx, "Patient", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPatients(resp)
}

// Get Patient by ID.
func (c *Client) GetPatientByID(ctx context.Context, id string, params Parameters) (*models.Patient, error) {
	resp, err := c.GetByID(ctx, "Patient", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPatient(id, resp)
}

func (c *Client) CreatePatient(ctx context.Context, resource string, params Parameters, entity *models.Patient) (*models.Patient, error) {
	resp, err := c.Create(ctx, "Patient", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPatient(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePatient(ctx context.Context, resource string, params Parameters, entity *models.Patient) (*models.Patient, error) {
	resp, err := c.Update(ctx, "Patient", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPatient(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePatientByID(ctx context.Context, resource, id string, params Parameters, entity *models.Patient) (*models.Patient, error) {
	resp, err := c.UpdateByID(ctx, "Patient", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPatient(id, resp)
}

func (c *Client) PatchPatient(ctx context.Context, resource string, params Parameters, entity *models.Patient) ([]*models.Patient, error) {
	resp, err := c.Patch(ctx, "Patient", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPatients(resp)
}

func (c *Client) PatchPatientByID(ctx context.Context, resource, id string, params Parameters, entity *models.Patient) (*models.Patient, error) {
	resp, err := c.PatchByID(ctx, "Patient", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPatient(id, resp)
}

func (c *Client) DeletePatient(ctx context.Context, resource string, params Parameters) ([]*models.Patient, error) {
	resp, err := c.Delete(ctx, "Patient", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPatients(resp)
}

func (c *Client) DeletePatientByID(ctx context.Context, resource, id string, params Parameters) (*models.Patient, error) {
	resp, err := c.DeleteByID(ctx, "Patient", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPatient(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// PaymentNotice
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToPaymentNotices(bundle *models.Bundle) ([]*models.PaymentNotice, error) {
	var entities []*models.PaymentNotice
	err := EnumBundleResources(bundle, "PaymentNotice", func(resource ResourceData) error {
		var entity models.PaymentNotice
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToPaymentNotices(resp *FhirResponse) ([]*models.PaymentNotice, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToPaymentNotices(resp.Bundle)
	case "PaymentNotice":
		var entity models.PaymentNotice
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.PaymentNotice{&entity}, nil
	}
	return nil, nil
}

func fhirRespToPaymentNotice(id string, resp *FhirResponse) (*models.PaymentNotice, error) {
	entities, err := fhirRespToPaymentNotices(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "PaymentNotice", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get PaymentNotice.
func (c *Client) GetPaymentNotice(ctx context.Context, params Parameters) ([]*models.PaymentNotice, error) {
	resp, err := c.Get(ctx, "PaymentNotice", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentNotices(resp)
}

// Get PaymentNotice by ID.
func (c *Client) GetPaymentNoticeByID(ctx context.Context, id string, params Parameters) (*models.PaymentNotice, error) {
	resp, err := c.GetByID(ctx, "PaymentNotice", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentNotice(id, resp)
}

func (c *Client) CreatePaymentNotice(ctx context.Context, resource string, params Parameters, entity *models.PaymentNotice) (*models.PaymentNotice, error) {
	resp, err := c.Create(ctx, "PaymentNotice", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentNotice(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePaymentNotice(ctx context.Context, resource string, params Parameters, entity *models.PaymentNotice) (*models.PaymentNotice, error) {
	resp, err := c.Update(ctx, "PaymentNotice", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentNotice(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePaymentNoticeByID(ctx context.Context, resource, id string, params Parameters, entity *models.PaymentNotice) (*models.PaymentNotice, error) {
	resp, err := c.UpdateByID(ctx, "PaymentNotice", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentNotice(id, resp)
}

func (c *Client) PatchPaymentNotice(ctx context.Context, resource string, params Parameters, entity *models.PaymentNotice) ([]*models.PaymentNotice, error) {
	resp, err := c.Patch(ctx, "PaymentNotice", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentNotices(resp)
}

func (c *Client) PatchPaymentNoticeByID(ctx context.Context, resource, id string, params Parameters, entity *models.PaymentNotice) (*models.PaymentNotice, error) {
	resp, err := c.PatchByID(ctx, "PaymentNotice", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentNotice(id, resp)
}

func (c *Client) DeletePaymentNotice(ctx context.Context, resource string, params Parameters) ([]*models.PaymentNotice, error) {
	resp, err := c.Delete(ctx, "PaymentNotice", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentNotices(resp)
}

func (c *Client) DeletePaymentNoticeByID(ctx context.Context, resource, id string, params Parameters) (*models.PaymentNotice, error) {
	resp, err := c.DeleteByID(ctx, "PaymentNotice", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentNotice(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// PaymentReconciliation
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToPaymentReconciliations(bundle *models.Bundle) ([]*models.PaymentReconciliation, error) {
	var entities []*models.PaymentReconciliation
	err := EnumBundleResources(bundle, "PaymentReconciliation", func(resource ResourceData) error {
		var entity models.PaymentReconciliation
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToPaymentReconciliations(resp *FhirResponse) ([]*models.PaymentReconciliation, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToPaymentReconciliations(resp.Bundle)
	case "PaymentReconciliation":
		var entity models.PaymentReconciliation
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.PaymentReconciliation{&entity}, nil
	}
	return nil, nil
}

func fhirRespToPaymentReconciliation(id string, resp *FhirResponse) (*models.PaymentReconciliation, error) {
	entities, err := fhirRespToPaymentReconciliations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "PaymentReconciliation", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get PaymentReconciliation.
func (c *Client) GetPaymentReconciliation(ctx context.Context, params Parameters) ([]*models.PaymentReconciliation, error) {
	resp, err := c.Get(ctx, "PaymentReconciliation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentReconciliations(resp)
}

// Get PaymentReconciliation by ID.
func (c *Client) GetPaymentReconciliationByID(ctx context.Context, id string, params Parameters) (*models.PaymentReconciliation, error) {
	resp, err := c.GetByID(ctx, "PaymentReconciliation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentReconciliation(id, resp)
}

func (c *Client) CreatePaymentReconciliation(ctx context.Context, resource string, params Parameters, entity *models.PaymentReconciliation) (*models.PaymentReconciliation, error) {
	resp, err := c.Create(ctx, "PaymentReconciliation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentReconciliation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePaymentReconciliation(ctx context.Context, resource string, params Parameters, entity *models.PaymentReconciliation) (*models.PaymentReconciliation, error) {
	resp, err := c.Update(ctx, "PaymentReconciliation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentReconciliation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePaymentReconciliationByID(ctx context.Context, resource, id string, params Parameters, entity *models.PaymentReconciliation) (*models.PaymentReconciliation, error) {
	resp, err := c.UpdateByID(ctx, "PaymentReconciliation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentReconciliation(id, resp)
}

func (c *Client) PatchPaymentReconciliation(ctx context.Context, resource string, params Parameters, entity *models.PaymentReconciliation) ([]*models.PaymentReconciliation, error) {
	resp, err := c.Patch(ctx, "PaymentReconciliation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentReconciliations(resp)
}

func (c *Client) PatchPaymentReconciliationByID(ctx context.Context, resource, id string, params Parameters, entity *models.PaymentReconciliation) (*models.PaymentReconciliation, error) {
	resp, err := c.PatchByID(ctx, "PaymentReconciliation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentReconciliation(id, resp)
}

func (c *Client) DeletePaymentReconciliation(ctx context.Context, resource string, params Parameters) ([]*models.PaymentReconciliation, error) {
	resp, err := c.Delete(ctx, "PaymentReconciliation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentReconciliations(resp)
}

func (c *Client) DeletePaymentReconciliationByID(ctx context.Context, resource, id string, params Parameters) (*models.PaymentReconciliation, error) {
	resp, err := c.DeleteByID(ctx, "PaymentReconciliation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPaymentReconciliation(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Person
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToPersons(bundle *models.Bundle) ([]*models.Person, error) {
	var entities []*models.Person
	err := EnumBundleResources(bundle, "Person", func(resource ResourceData) error {
		var entity models.Person
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToPersons(resp *FhirResponse) ([]*models.Person, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToPersons(resp.Bundle)
	case "Person":
		var entity models.Person
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Person{&entity}, nil
	}
	return nil, nil
}

func fhirRespToPerson(id string, resp *FhirResponse) (*models.Person, error) {
	entities, err := fhirRespToPersons(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Person", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Person.
func (c *Client) GetPerson(ctx context.Context, params Parameters) ([]*models.Person, error) {
	resp, err := c.Get(ctx, "Person", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPersons(resp)
}

// Get Person by ID.
func (c *Client) GetPersonByID(ctx context.Context, id string, params Parameters) (*models.Person, error) {
	resp, err := c.GetByID(ctx, "Person", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPerson(id, resp)
}

func (c *Client) CreatePerson(ctx context.Context, resource string, params Parameters, entity *models.Person) (*models.Person, error) {
	resp, err := c.Create(ctx, "Person", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPerson(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePerson(ctx context.Context, resource string, params Parameters, entity *models.Person) (*models.Person, error) {
	resp, err := c.Update(ctx, "Person", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPerson(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePersonByID(ctx context.Context, resource, id string, params Parameters, entity *models.Person) (*models.Person, error) {
	resp, err := c.UpdateByID(ctx, "Person", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPerson(id, resp)
}

func (c *Client) PatchPerson(ctx context.Context, resource string, params Parameters, entity *models.Person) ([]*models.Person, error) {
	resp, err := c.Patch(ctx, "Person", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPersons(resp)
}

func (c *Client) PatchPersonByID(ctx context.Context, resource, id string, params Parameters, entity *models.Person) (*models.Person, error) {
	resp, err := c.PatchByID(ctx, "Person", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPerson(id, resp)
}

func (c *Client) DeletePerson(ctx context.Context, resource string, params Parameters) ([]*models.Person, error) {
	resp, err := c.Delete(ctx, "Person", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPersons(resp)
}

func (c *Client) DeletePersonByID(ctx context.Context, resource, id string, params Parameters) (*models.Person, error) {
	resp, err := c.DeleteByID(ctx, "Person", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPerson(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// PlanDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToPlanDefinitions(bundle *models.Bundle) ([]*models.PlanDefinition, error) {
	var entities []*models.PlanDefinition
	err := EnumBundleResources(bundle, "PlanDefinition", func(resource ResourceData) error {
		var entity models.PlanDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToPlanDefinitions(resp *FhirResponse) ([]*models.PlanDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToPlanDefinitions(resp.Bundle)
	case "PlanDefinition":
		var entity models.PlanDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.PlanDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToPlanDefinition(id string, resp *FhirResponse) (*models.PlanDefinition, error) {
	entities, err := fhirRespToPlanDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "PlanDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get PlanDefinition.
func (c *Client) GetPlanDefinition(ctx context.Context, params Parameters) ([]*models.PlanDefinition, error) {
	resp, err := c.Get(ctx, "PlanDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPlanDefinitions(resp)
}

// Get PlanDefinition by ID.
func (c *Client) GetPlanDefinitionByID(ctx context.Context, id string, params Parameters) (*models.PlanDefinition, error) {
	resp, err := c.GetByID(ctx, "PlanDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPlanDefinition(id, resp)
}

func (c *Client) CreatePlanDefinition(ctx context.Context, resource string, params Parameters, entity *models.PlanDefinition) (*models.PlanDefinition, error) {
	resp, err := c.Create(ctx, "PlanDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPlanDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePlanDefinition(ctx context.Context, resource string, params Parameters, entity *models.PlanDefinition) (*models.PlanDefinition, error) {
	resp, err := c.Update(ctx, "PlanDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPlanDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePlanDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.PlanDefinition) (*models.PlanDefinition, error) {
	resp, err := c.UpdateByID(ctx, "PlanDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPlanDefinition(id, resp)
}

func (c *Client) PatchPlanDefinition(ctx context.Context, resource string, params Parameters, entity *models.PlanDefinition) ([]*models.PlanDefinition, error) {
	resp, err := c.Patch(ctx, "PlanDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPlanDefinitions(resp)
}

func (c *Client) PatchPlanDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.PlanDefinition) (*models.PlanDefinition, error) {
	resp, err := c.PatchByID(ctx, "PlanDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPlanDefinition(id, resp)
}

func (c *Client) DeletePlanDefinition(ctx context.Context, resource string, params Parameters) ([]*models.PlanDefinition, error) {
	resp, err := c.Delete(ctx, "PlanDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPlanDefinitions(resp)
}

func (c *Client) DeletePlanDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.PlanDefinition, error) {
	resp, err := c.DeleteByID(ctx, "PlanDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPlanDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Practitioner
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToPractitioners(bundle *models.Bundle) ([]*models.Practitioner, error) {
	var entities []*models.Practitioner
	err := EnumBundleResources(bundle, "Practitioner", func(resource ResourceData) error {
		var entity models.Practitioner
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToPractitioners(resp *FhirResponse) ([]*models.Practitioner, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToPractitioners(resp.Bundle)
	case "Practitioner":
		var entity models.Practitioner
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Practitioner{&entity}, nil
	}
	return nil, nil
}

func fhirRespToPractitioner(id string, resp *FhirResponse) (*models.Practitioner, error) {
	entities, err := fhirRespToPractitioners(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Practitioner", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Practitioner.
func (c *Client) GetPractitioner(ctx context.Context, params Parameters) ([]*models.Practitioner, error) {
	resp, err := c.Get(ctx, "Practitioner", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitioners(resp)
}

// Get Practitioner by ID.
func (c *Client) GetPractitionerByID(ctx context.Context, id string, params Parameters) (*models.Practitioner, error) {
	resp, err := c.GetByID(ctx, "Practitioner", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitioner(id, resp)
}

func (c *Client) CreatePractitioner(ctx context.Context, resource string, params Parameters, entity *models.Practitioner) (*models.Practitioner, error) {
	resp, err := c.Create(ctx, "Practitioner", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitioner(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePractitioner(ctx context.Context, resource string, params Parameters, entity *models.Practitioner) (*models.Practitioner, error) {
	resp, err := c.Update(ctx, "Practitioner", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitioner(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePractitionerByID(ctx context.Context, resource, id string, params Parameters, entity *models.Practitioner) (*models.Practitioner, error) {
	resp, err := c.UpdateByID(ctx, "Practitioner", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitioner(id, resp)
}

func (c *Client) PatchPractitioner(ctx context.Context, resource string, params Parameters, entity *models.Practitioner) ([]*models.Practitioner, error) {
	resp, err := c.Patch(ctx, "Practitioner", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitioners(resp)
}

func (c *Client) PatchPractitionerByID(ctx context.Context, resource, id string, params Parameters, entity *models.Practitioner) (*models.Practitioner, error) {
	resp, err := c.PatchByID(ctx, "Practitioner", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitioner(id, resp)
}

func (c *Client) DeletePractitioner(ctx context.Context, resource string, params Parameters) ([]*models.Practitioner, error) {
	resp, err := c.Delete(ctx, "Practitioner", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitioners(resp)
}

func (c *Client) DeletePractitionerByID(ctx context.Context, resource, id string, params Parameters) (*models.Practitioner, error) {
	resp, err := c.DeleteByID(ctx, "Practitioner", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitioner(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// PractitionerRole
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToPractitionerRoles(bundle *models.Bundle) ([]*models.PractitionerRole, error) {
	var entities []*models.PractitionerRole
	err := EnumBundleResources(bundle, "PractitionerRole", func(resource ResourceData) error {
		var entity models.PractitionerRole
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToPractitionerRoles(resp *FhirResponse) ([]*models.PractitionerRole, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToPractitionerRoles(resp.Bundle)
	case "PractitionerRole":
		var entity models.PractitionerRole
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.PractitionerRole{&entity}, nil
	}
	return nil, nil
}

func fhirRespToPractitionerRole(id string, resp *FhirResponse) (*models.PractitionerRole, error) {
	entities, err := fhirRespToPractitionerRoles(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "PractitionerRole", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get PractitionerRole.
func (c *Client) GetPractitionerRole(ctx context.Context, params Parameters) ([]*models.PractitionerRole, error) {
	resp, err := c.Get(ctx, "PractitionerRole", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitionerRoles(resp)
}

// Get PractitionerRole by ID.
func (c *Client) GetPractitionerRoleByID(ctx context.Context, id string, params Parameters) (*models.PractitionerRole, error) {
	resp, err := c.GetByID(ctx, "PractitionerRole", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitionerRole(id, resp)
}

func (c *Client) CreatePractitionerRole(ctx context.Context, resource string, params Parameters, entity *models.PractitionerRole) (*models.PractitionerRole, error) {
	resp, err := c.Create(ctx, "PractitionerRole", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitionerRole(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePractitionerRole(ctx context.Context, resource string, params Parameters, entity *models.PractitionerRole) (*models.PractitionerRole, error) {
	resp, err := c.Update(ctx, "PractitionerRole", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitionerRole(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdatePractitionerRoleByID(ctx context.Context, resource, id string, params Parameters, entity *models.PractitionerRole) (*models.PractitionerRole, error) {
	resp, err := c.UpdateByID(ctx, "PractitionerRole", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitionerRole(id, resp)
}

func (c *Client) PatchPractitionerRole(ctx context.Context, resource string, params Parameters, entity *models.PractitionerRole) ([]*models.PractitionerRole, error) {
	resp, err := c.Patch(ctx, "PractitionerRole", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitionerRoles(resp)
}

func (c *Client) PatchPractitionerRoleByID(ctx context.Context, resource, id string, params Parameters, entity *models.PractitionerRole) (*models.PractitionerRole, error) {
	resp, err := c.PatchByID(ctx, "PractitionerRole", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitionerRole(id, resp)
}

func (c *Client) DeletePractitionerRole(ctx context.Context, resource string, params Parameters) ([]*models.PractitionerRole, error) {
	resp, err := c.Delete(ctx, "PractitionerRole", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitionerRoles(resp)
}

func (c *Client) DeletePractitionerRoleByID(ctx context.Context, resource, id string, params Parameters) (*models.PractitionerRole, error) {
	resp, err := c.DeleteByID(ctx, "PractitionerRole", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToPractitionerRole(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Procedure
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToProcedures(bundle *models.Bundle) ([]*models.Procedure, error) {
	var entities []*models.Procedure
	err := EnumBundleResources(bundle, "Procedure", func(resource ResourceData) error {
		var entity models.Procedure
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToProcedures(resp *FhirResponse) ([]*models.Procedure, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToProcedures(resp.Bundle)
	case "Procedure":
		var entity models.Procedure
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Procedure{&entity}, nil
	}
	return nil, nil
}

func fhirRespToProcedure(id string, resp *FhirResponse) (*models.Procedure, error) {
	entities, err := fhirRespToProcedures(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Procedure", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Procedure.
func (c *Client) GetProcedure(ctx context.Context, params Parameters) ([]*models.Procedure, error) {
	resp, err := c.Get(ctx, "Procedure", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToProcedures(resp)
}

// Get Procedure by ID.
func (c *Client) GetProcedureByID(ctx context.Context, id string, params Parameters) (*models.Procedure, error) {
	resp, err := c.GetByID(ctx, "Procedure", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToProcedure(id, resp)
}

func (c *Client) CreateProcedure(ctx context.Context, resource string, params Parameters, entity *models.Procedure) (*models.Procedure, error) {
	resp, err := c.Create(ctx, "Procedure", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToProcedure(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateProcedure(ctx context.Context, resource string, params Parameters, entity *models.Procedure) (*models.Procedure, error) {
	resp, err := c.Update(ctx, "Procedure", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToProcedure(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateProcedureByID(ctx context.Context, resource, id string, params Parameters, entity *models.Procedure) (*models.Procedure, error) {
	resp, err := c.UpdateByID(ctx, "Procedure", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToProcedure(id, resp)
}

func (c *Client) PatchProcedure(ctx context.Context, resource string, params Parameters, entity *models.Procedure) ([]*models.Procedure, error) {
	resp, err := c.Patch(ctx, "Procedure", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToProcedures(resp)
}

func (c *Client) PatchProcedureByID(ctx context.Context, resource, id string, params Parameters, entity *models.Procedure) (*models.Procedure, error) {
	resp, err := c.PatchByID(ctx, "Procedure", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToProcedure(id, resp)
}

func (c *Client) DeleteProcedure(ctx context.Context, resource string, params Parameters) ([]*models.Procedure, error) {
	resp, err := c.Delete(ctx, "Procedure", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToProcedures(resp)
}

func (c *Client) DeleteProcedureByID(ctx context.Context, resource, id string, params Parameters) (*models.Procedure, error) {
	resp, err := c.DeleteByID(ctx, "Procedure", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToProcedure(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Provenance
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToProvenances(bundle *models.Bundle) ([]*models.Provenance, error) {
	var entities []*models.Provenance
	err := EnumBundleResources(bundle, "Provenance", func(resource ResourceData) error {
		var entity models.Provenance
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToProvenances(resp *FhirResponse) ([]*models.Provenance, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToProvenances(resp.Bundle)
	case "Provenance":
		var entity models.Provenance
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Provenance{&entity}, nil
	}
	return nil, nil
}

func fhirRespToProvenance(id string, resp *FhirResponse) (*models.Provenance, error) {
	entities, err := fhirRespToProvenances(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Provenance", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Provenance.
func (c *Client) GetProvenance(ctx context.Context, params Parameters) ([]*models.Provenance, error) {
	resp, err := c.Get(ctx, "Provenance", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToProvenances(resp)
}

// Get Provenance by ID.
func (c *Client) GetProvenanceByID(ctx context.Context, id string, params Parameters) (*models.Provenance, error) {
	resp, err := c.GetByID(ctx, "Provenance", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToProvenance(id, resp)
}

func (c *Client) CreateProvenance(ctx context.Context, resource string, params Parameters, entity *models.Provenance) (*models.Provenance, error) {
	resp, err := c.Create(ctx, "Provenance", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToProvenance(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateProvenance(ctx context.Context, resource string, params Parameters, entity *models.Provenance) (*models.Provenance, error) {
	resp, err := c.Update(ctx, "Provenance", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToProvenance(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateProvenanceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Provenance) (*models.Provenance, error) {
	resp, err := c.UpdateByID(ctx, "Provenance", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToProvenance(id, resp)
}

func (c *Client) PatchProvenance(ctx context.Context, resource string, params Parameters, entity *models.Provenance) ([]*models.Provenance, error) {
	resp, err := c.Patch(ctx, "Provenance", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToProvenances(resp)
}

func (c *Client) PatchProvenanceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Provenance) (*models.Provenance, error) {
	resp, err := c.PatchByID(ctx, "Provenance", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToProvenance(id, resp)
}

func (c *Client) DeleteProvenance(ctx context.Context, resource string, params Parameters) ([]*models.Provenance, error) {
	resp, err := c.Delete(ctx, "Provenance", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToProvenances(resp)
}

func (c *Client) DeleteProvenanceByID(ctx context.Context, resource, id string, params Parameters) (*models.Provenance, error) {
	resp, err := c.DeleteByID(ctx, "Provenance", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToProvenance(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Questionnaire
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToQuestionnaires(bundle *models.Bundle) ([]*models.Questionnaire, error) {
	var entities []*models.Questionnaire
	err := EnumBundleResources(bundle, "Questionnaire", func(resource ResourceData) error {
		var entity models.Questionnaire
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToQuestionnaires(resp *FhirResponse) ([]*models.Questionnaire, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToQuestionnaires(resp.Bundle)
	case "Questionnaire":
		var entity models.Questionnaire
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Questionnaire{&entity}, nil
	}
	return nil, nil
}

func fhirRespToQuestionnaire(id string, resp *FhirResponse) (*models.Questionnaire, error) {
	entities, err := fhirRespToQuestionnaires(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Questionnaire", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Questionnaire.
func (c *Client) GetQuestionnaire(ctx context.Context, params Parameters) ([]*models.Questionnaire, error) {
	resp, err := c.Get(ctx, "Questionnaire", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaires(resp)
}

// Get Questionnaire by ID.
func (c *Client) GetQuestionnaireByID(ctx context.Context, id string, params Parameters) (*models.Questionnaire, error) {
	resp, err := c.GetByID(ctx, "Questionnaire", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaire(id, resp)
}

func (c *Client) CreateQuestionnaire(ctx context.Context, resource string, params Parameters, entity *models.Questionnaire) (*models.Questionnaire, error) {
	resp, err := c.Create(ctx, "Questionnaire", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaire(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateQuestionnaire(ctx context.Context, resource string, params Parameters, entity *models.Questionnaire) (*models.Questionnaire, error) {
	resp, err := c.Update(ctx, "Questionnaire", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaire(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateQuestionnaireByID(ctx context.Context, resource, id string, params Parameters, entity *models.Questionnaire) (*models.Questionnaire, error) {
	resp, err := c.UpdateByID(ctx, "Questionnaire", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaire(id, resp)
}

func (c *Client) PatchQuestionnaire(ctx context.Context, resource string, params Parameters, entity *models.Questionnaire) ([]*models.Questionnaire, error) {
	resp, err := c.Patch(ctx, "Questionnaire", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaires(resp)
}

func (c *Client) PatchQuestionnaireByID(ctx context.Context, resource, id string, params Parameters, entity *models.Questionnaire) (*models.Questionnaire, error) {
	resp, err := c.PatchByID(ctx, "Questionnaire", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaire(id, resp)
}

func (c *Client) DeleteQuestionnaire(ctx context.Context, resource string, params Parameters) ([]*models.Questionnaire, error) {
	resp, err := c.Delete(ctx, "Questionnaire", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaires(resp)
}

func (c *Client) DeleteQuestionnaireByID(ctx context.Context, resource, id string, params Parameters) (*models.Questionnaire, error) {
	resp, err := c.DeleteByID(ctx, "Questionnaire", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaire(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// QuestionnaireResponse
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToQuestionnaireResponses(bundle *models.Bundle) ([]*models.QuestionnaireResponse, error) {
	var entities []*models.QuestionnaireResponse
	err := EnumBundleResources(bundle, "QuestionnaireResponse", func(resource ResourceData) error {
		var entity models.QuestionnaireResponse
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToQuestionnaireResponses(resp *FhirResponse) ([]*models.QuestionnaireResponse, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToQuestionnaireResponses(resp.Bundle)
	case "QuestionnaireResponse":
		var entity models.QuestionnaireResponse
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.QuestionnaireResponse{&entity}, nil
	}
	return nil, nil
}

func fhirRespToQuestionnaireResponse(id string, resp *FhirResponse) (*models.QuestionnaireResponse, error) {
	entities, err := fhirRespToQuestionnaireResponses(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "QuestionnaireResponse", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get QuestionnaireResponse.
func (c *Client) GetQuestionnaireResponse(ctx context.Context, params Parameters) ([]*models.QuestionnaireResponse, error) {
	resp, err := c.Get(ctx, "QuestionnaireResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaireResponses(resp)
}

// Get QuestionnaireResponse by ID.
func (c *Client) GetQuestionnaireResponseByID(ctx context.Context, id string, params Parameters) (*models.QuestionnaireResponse, error) {
	resp, err := c.GetByID(ctx, "QuestionnaireResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaireResponse(id, resp)
}

func (c *Client) CreateQuestionnaireResponse(ctx context.Context, resource string, params Parameters, entity *models.QuestionnaireResponse) (*models.QuestionnaireResponse, error) {
	resp, err := c.Create(ctx, "QuestionnaireResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaireResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateQuestionnaireResponse(ctx context.Context, resource string, params Parameters, entity *models.QuestionnaireResponse) (*models.QuestionnaireResponse, error) {
	resp, err := c.Update(ctx, "QuestionnaireResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaireResponse(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateQuestionnaireResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.QuestionnaireResponse) (*models.QuestionnaireResponse, error) {
	resp, err := c.UpdateByID(ctx, "QuestionnaireResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaireResponse(id, resp)
}

func (c *Client) PatchQuestionnaireResponse(ctx context.Context, resource string, params Parameters, entity *models.QuestionnaireResponse) ([]*models.QuestionnaireResponse, error) {
	resp, err := c.Patch(ctx, "QuestionnaireResponse", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaireResponses(resp)
}

func (c *Client) PatchQuestionnaireResponseByID(ctx context.Context, resource, id string, params Parameters, entity *models.QuestionnaireResponse) (*models.QuestionnaireResponse, error) {
	resp, err := c.PatchByID(ctx, "QuestionnaireResponse", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaireResponse(id, resp)
}

func (c *Client) DeleteQuestionnaireResponse(ctx context.Context, resource string, params Parameters) ([]*models.QuestionnaireResponse, error) {
	resp, err := c.Delete(ctx, "QuestionnaireResponse", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaireResponses(resp)
}

func (c *Client) DeleteQuestionnaireResponseByID(ctx context.Context, resource, id string, params Parameters) (*models.QuestionnaireResponse, error) {
	resp, err := c.DeleteByID(ctx, "QuestionnaireResponse", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToQuestionnaireResponse(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// RelatedPerson
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToRelatedPersons(bundle *models.Bundle) ([]*models.RelatedPerson, error) {
	var entities []*models.RelatedPerson
	err := EnumBundleResources(bundle, "RelatedPerson", func(resource ResourceData) error {
		var entity models.RelatedPerson
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToRelatedPersons(resp *FhirResponse) ([]*models.RelatedPerson, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToRelatedPersons(resp.Bundle)
	case "RelatedPerson":
		var entity models.RelatedPerson
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.RelatedPerson{&entity}, nil
	}
	return nil, nil
}

func fhirRespToRelatedPerson(id string, resp *FhirResponse) (*models.RelatedPerson, error) {
	entities, err := fhirRespToRelatedPersons(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "RelatedPerson", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get RelatedPerson.
func (c *Client) GetRelatedPerson(ctx context.Context, params Parameters) ([]*models.RelatedPerson, error) {
	resp, err := c.Get(ctx, "RelatedPerson", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRelatedPersons(resp)
}

// Get RelatedPerson by ID.
func (c *Client) GetRelatedPersonByID(ctx context.Context, id string, params Parameters) (*models.RelatedPerson, error) {
	resp, err := c.GetByID(ctx, "RelatedPerson", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRelatedPerson(id, resp)
}

func (c *Client) CreateRelatedPerson(ctx context.Context, resource string, params Parameters, entity *models.RelatedPerson) (*models.RelatedPerson, error) {
	resp, err := c.Create(ctx, "RelatedPerson", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRelatedPerson(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateRelatedPerson(ctx context.Context, resource string, params Parameters, entity *models.RelatedPerson) (*models.RelatedPerson, error) {
	resp, err := c.Update(ctx, "RelatedPerson", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRelatedPerson(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateRelatedPersonByID(ctx context.Context, resource, id string, params Parameters, entity *models.RelatedPerson) (*models.RelatedPerson, error) {
	resp, err := c.UpdateByID(ctx, "RelatedPerson", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRelatedPerson(id, resp)
}

func (c *Client) PatchRelatedPerson(ctx context.Context, resource string, params Parameters, entity *models.RelatedPerson) ([]*models.RelatedPerson, error) {
	resp, err := c.Patch(ctx, "RelatedPerson", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRelatedPersons(resp)
}

func (c *Client) PatchRelatedPersonByID(ctx context.Context, resource, id string, params Parameters, entity *models.RelatedPerson) (*models.RelatedPerson, error) {
	resp, err := c.PatchByID(ctx, "RelatedPerson", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRelatedPerson(id, resp)
}

func (c *Client) DeleteRelatedPerson(ctx context.Context, resource string, params Parameters) ([]*models.RelatedPerson, error) {
	resp, err := c.Delete(ctx, "RelatedPerson", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRelatedPersons(resp)
}

func (c *Client) DeleteRelatedPersonByID(ctx context.Context, resource, id string, params Parameters) (*models.RelatedPerson, error) {
	resp, err := c.DeleteByID(ctx, "RelatedPerson", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRelatedPerson(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// RequestGroup
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToRequestGroups(bundle *models.Bundle) ([]*models.RequestGroup, error) {
	var entities []*models.RequestGroup
	err := EnumBundleResources(bundle, "RequestGroup", func(resource ResourceData) error {
		var entity models.RequestGroup
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToRequestGroups(resp *FhirResponse) ([]*models.RequestGroup, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToRequestGroups(resp.Bundle)
	case "RequestGroup":
		var entity models.RequestGroup
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.RequestGroup{&entity}, nil
	}
	return nil, nil
}

func fhirRespToRequestGroup(id string, resp *FhirResponse) (*models.RequestGroup, error) {
	entities, err := fhirRespToRequestGroups(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "RequestGroup", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get RequestGroup.
func (c *Client) GetRequestGroup(ctx context.Context, params Parameters) ([]*models.RequestGroup, error) {
	resp, err := c.Get(ctx, "RequestGroup", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRequestGroups(resp)
}

// Get RequestGroup by ID.
func (c *Client) GetRequestGroupByID(ctx context.Context, id string, params Parameters) (*models.RequestGroup, error) {
	resp, err := c.GetByID(ctx, "RequestGroup", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRequestGroup(id, resp)
}

func (c *Client) CreateRequestGroup(ctx context.Context, resource string, params Parameters, entity *models.RequestGroup) (*models.RequestGroup, error) {
	resp, err := c.Create(ctx, "RequestGroup", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRequestGroup(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateRequestGroup(ctx context.Context, resource string, params Parameters, entity *models.RequestGroup) (*models.RequestGroup, error) {
	resp, err := c.Update(ctx, "RequestGroup", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRequestGroup(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateRequestGroupByID(ctx context.Context, resource, id string, params Parameters, entity *models.RequestGroup) (*models.RequestGroup, error) {
	resp, err := c.UpdateByID(ctx, "RequestGroup", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRequestGroup(id, resp)
}

func (c *Client) PatchRequestGroup(ctx context.Context, resource string, params Parameters, entity *models.RequestGroup) ([]*models.RequestGroup, error) {
	resp, err := c.Patch(ctx, "RequestGroup", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRequestGroups(resp)
}

func (c *Client) PatchRequestGroupByID(ctx context.Context, resource, id string, params Parameters, entity *models.RequestGroup) (*models.RequestGroup, error) {
	resp, err := c.PatchByID(ctx, "RequestGroup", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRequestGroup(id, resp)
}

func (c *Client) DeleteRequestGroup(ctx context.Context, resource string, params Parameters) ([]*models.RequestGroup, error) {
	resp, err := c.Delete(ctx, "RequestGroup", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRequestGroups(resp)
}

func (c *Client) DeleteRequestGroupByID(ctx context.Context, resource, id string, params Parameters) (*models.RequestGroup, error) {
	resp, err := c.DeleteByID(ctx, "RequestGroup", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRequestGroup(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ResearchDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToResearchDefinitions(bundle *models.Bundle) ([]*models.ResearchDefinition, error) {
	var entities []*models.ResearchDefinition
	err := EnumBundleResources(bundle, "ResearchDefinition", func(resource ResourceData) error {
		var entity models.ResearchDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToResearchDefinitions(resp *FhirResponse) ([]*models.ResearchDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToResearchDefinitions(resp.Bundle)
	case "ResearchDefinition":
		var entity models.ResearchDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ResearchDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToResearchDefinition(id string, resp *FhirResponse) (*models.ResearchDefinition, error) {
	entities, err := fhirRespToResearchDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ResearchDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ResearchDefinition.
func (c *Client) GetResearchDefinition(ctx context.Context, params Parameters) ([]*models.ResearchDefinition, error) {
	resp, err := c.Get(ctx, "ResearchDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchDefinitions(resp)
}

// Get ResearchDefinition by ID.
func (c *Client) GetResearchDefinitionByID(ctx context.Context, id string, params Parameters) (*models.ResearchDefinition, error) {
	resp, err := c.GetByID(ctx, "ResearchDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchDefinition(id, resp)
}

func (c *Client) CreateResearchDefinition(ctx context.Context, resource string, params Parameters, entity *models.ResearchDefinition) (*models.ResearchDefinition, error) {
	resp, err := c.Create(ctx, "ResearchDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateResearchDefinition(ctx context.Context, resource string, params Parameters, entity *models.ResearchDefinition) (*models.ResearchDefinition, error) {
	resp, err := c.Update(ctx, "ResearchDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateResearchDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ResearchDefinition) (*models.ResearchDefinition, error) {
	resp, err := c.UpdateByID(ctx, "ResearchDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchDefinition(id, resp)
}

func (c *Client) PatchResearchDefinition(ctx context.Context, resource string, params Parameters, entity *models.ResearchDefinition) ([]*models.ResearchDefinition, error) {
	resp, err := c.Patch(ctx, "ResearchDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchDefinitions(resp)
}

func (c *Client) PatchResearchDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ResearchDefinition) (*models.ResearchDefinition, error) {
	resp, err := c.PatchByID(ctx, "ResearchDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchDefinition(id, resp)
}

func (c *Client) DeleteResearchDefinition(ctx context.Context, resource string, params Parameters) ([]*models.ResearchDefinition, error) {
	resp, err := c.Delete(ctx, "ResearchDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchDefinitions(resp)
}

func (c *Client) DeleteResearchDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.ResearchDefinition, error) {
	resp, err := c.DeleteByID(ctx, "ResearchDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ResearchElementDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToResearchElementDefinitions(bundle *models.Bundle) ([]*models.ResearchElementDefinition, error) {
	var entities []*models.ResearchElementDefinition
	err := EnumBundleResources(bundle, "ResearchElementDefinition", func(resource ResourceData) error {
		var entity models.ResearchElementDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToResearchElementDefinitions(resp *FhirResponse) ([]*models.ResearchElementDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToResearchElementDefinitions(resp.Bundle)
	case "ResearchElementDefinition":
		var entity models.ResearchElementDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ResearchElementDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToResearchElementDefinition(id string, resp *FhirResponse) (*models.ResearchElementDefinition, error) {
	entities, err := fhirRespToResearchElementDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ResearchElementDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ResearchElementDefinition.
func (c *Client) GetResearchElementDefinition(ctx context.Context, params Parameters) ([]*models.ResearchElementDefinition, error) {
	resp, err := c.Get(ctx, "ResearchElementDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchElementDefinitions(resp)
}

// Get ResearchElementDefinition by ID.
func (c *Client) GetResearchElementDefinitionByID(ctx context.Context, id string, params Parameters) (*models.ResearchElementDefinition, error) {
	resp, err := c.GetByID(ctx, "ResearchElementDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchElementDefinition(id, resp)
}

func (c *Client) CreateResearchElementDefinition(ctx context.Context, resource string, params Parameters, entity *models.ResearchElementDefinition) (*models.ResearchElementDefinition, error) {
	resp, err := c.Create(ctx, "ResearchElementDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchElementDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateResearchElementDefinition(ctx context.Context, resource string, params Parameters, entity *models.ResearchElementDefinition) (*models.ResearchElementDefinition, error) {
	resp, err := c.Update(ctx, "ResearchElementDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchElementDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateResearchElementDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ResearchElementDefinition) (*models.ResearchElementDefinition, error) {
	resp, err := c.UpdateByID(ctx, "ResearchElementDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchElementDefinition(id, resp)
}

func (c *Client) PatchResearchElementDefinition(ctx context.Context, resource string, params Parameters, entity *models.ResearchElementDefinition) ([]*models.ResearchElementDefinition, error) {
	resp, err := c.Patch(ctx, "ResearchElementDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchElementDefinitions(resp)
}

func (c *Client) PatchResearchElementDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.ResearchElementDefinition) (*models.ResearchElementDefinition, error) {
	resp, err := c.PatchByID(ctx, "ResearchElementDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchElementDefinition(id, resp)
}

func (c *Client) DeleteResearchElementDefinition(ctx context.Context, resource string, params Parameters) ([]*models.ResearchElementDefinition, error) {
	resp, err := c.Delete(ctx, "ResearchElementDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchElementDefinitions(resp)
}

func (c *Client) DeleteResearchElementDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.ResearchElementDefinition, error) {
	resp, err := c.DeleteByID(ctx, "ResearchElementDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchElementDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ResearchStudy
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToResearchStudys(bundle *models.Bundle) ([]*models.ResearchStudy, error) {
	var entities []*models.ResearchStudy
	err := EnumBundleResources(bundle, "ResearchStudy", func(resource ResourceData) error {
		var entity models.ResearchStudy
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToResearchStudys(resp *FhirResponse) ([]*models.ResearchStudy, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToResearchStudys(resp.Bundle)
	case "ResearchStudy":
		var entity models.ResearchStudy
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ResearchStudy{&entity}, nil
	}
	return nil, nil
}

func fhirRespToResearchStudy(id string, resp *FhirResponse) (*models.ResearchStudy, error) {
	entities, err := fhirRespToResearchStudys(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ResearchStudy", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ResearchStudy.
func (c *Client) GetResearchStudy(ctx context.Context, params Parameters) ([]*models.ResearchStudy, error) {
	resp, err := c.Get(ctx, "ResearchStudy", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchStudys(resp)
}

// Get ResearchStudy by ID.
func (c *Client) GetResearchStudyByID(ctx context.Context, id string, params Parameters) (*models.ResearchStudy, error) {
	resp, err := c.GetByID(ctx, "ResearchStudy", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchStudy(id, resp)
}

func (c *Client) CreateResearchStudy(ctx context.Context, resource string, params Parameters, entity *models.ResearchStudy) (*models.ResearchStudy, error) {
	resp, err := c.Create(ctx, "ResearchStudy", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchStudy(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateResearchStudy(ctx context.Context, resource string, params Parameters, entity *models.ResearchStudy) (*models.ResearchStudy, error) {
	resp, err := c.Update(ctx, "ResearchStudy", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchStudy(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateResearchStudyByID(ctx context.Context, resource, id string, params Parameters, entity *models.ResearchStudy) (*models.ResearchStudy, error) {
	resp, err := c.UpdateByID(ctx, "ResearchStudy", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchStudy(id, resp)
}

func (c *Client) PatchResearchStudy(ctx context.Context, resource string, params Parameters, entity *models.ResearchStudy) ([]*models.ResearchStudy, error) {
	resp, err := c.Patch(ctx, "ResearchStudy", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchStudys(resp)
}

func (c *Client) PatchResearchStudyByID(ctx context.Context, resource, id string, params Parameters, entity *models.ResearchStudy) (*models.ResearchStudy, error) {
	resp, err := c.PatchByID(ctx, "ResearchStudy", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchStudy(id, resp)
}

func (c *Client) DeleteResearchStudy(ctx context.Context, resource string, params Parameters) ([]*models.ResearchStudy, error) {
	resp, err := c.Delete(ctx, "ResearchStudy", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchStudys(resp)
}

func (c *Client) DeleteResearchStudyByID(ctx context.Context, resource, id string, params Parameters) (*models.ResearchStudy, error) {
	resp, err := c.DeleteByID(ctx, "ResearchStudy", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchStudy(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ResearchSubject
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToResearchSubjects(bundle *models.Bundle) ([]*models.ResearchSubject, error) {
	var entities []*models.ResearchSubject
	err := EnumBundleResources(bundle, "ResearchSubject", func(resource ResourceData) error {
		var entity models.ResearchSubject
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToResearchSubjects(resp *FhirResponse) ([]*models.ResearchSubject, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToResearchSubjects(resp.Bundle)
	case "ResearchSubject":
		var entity models.ResearchSubject
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ResearchSubject{&entity}, nil
	}
	return nil, nil
}

func fhirRespToResearchSubject(id string, resp *FhirResponse) (*models.ResearchSubject, error) {
	entities, err := fhirRespToResearchSubjects(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ResearchSubject", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ResearchSubject.
func (c *Client) GetResearchSubject(ctx context.Context, params Parameters) ([]*models.ResearchSubject, error) {
	resp, err := c.Get(ctx, "ResearchSubject", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchSubjects(resp)
}

// Get ResearchSubject by ID.
func (c *Client) GetResearchSubjectByID(ctx context.Context, id string, params Parameters) (*models.ResearchSubject, error) {
	resp, err := c.GetByID(ctx, "ResearchSubject", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchSubject(id, resp)
}

func (c *Client) CreateResearchSubject(ctx context.Context, resource string, params Parameters, entity *models.ResearchSubject) (*models.ResearchSubject, error) {
	resp, err := c.Create(ctx, "ResearchSubject", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchSubject(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateResearchSubject(ctx context.Context, resource string, params Parameters, entity *models.ResearchSubject) (*models.ResearchSubject, error) {
	resp, err := c.Update(ctx, "ResearchSubject", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchSubject(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateResearchSubjectByID(ctx context.Context, resource, id string, params Parameters, entity *models.ResearchSubject) (*models.ResearchSubject, error) {
	resp, err := c.UpdateByID(ctx, "ResearchSubject", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchSubject(id, resp)
}

func (c *Client) PatchResearchSubject(ctx context.Context, resource string, params Parameters, entity *models.ResearchSubject) ([]*models.ResearchSubject, error) {
	resp, err := c.Patch(ctx, "ResearchSubject", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchSubjects(resp)
}

func (c *Client) PatchResearchSubjectByID(ctx context.Context, resource, id string, params Parameters, entity *models.ResearchSubject) (*models.ResearchSubject, error) {
	resp, err := c.PatchByID(ctx, "ResearchSubject", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchSubject(id, resp)
}

func (c *Client) DeleteResearchSubject(ctx context.Context, resource string, params Parameters) ([]*models.ResearchSubject, error) {
	resp, err := c.Delete(ctx, "ResearchSubject", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchSubjects(resp)
}

func (c *Client) DeleteResearchSubjectByID(ctx context.Context, resource, id string, params Parameters) (*models.ResearchSubject, error) {
	resp, err := c.DeleteByID(ctx, "ResearchSubject", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResearchSubject(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Resource
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToResources(bundle *models.Bundle) ([]*models.Resource, error) {
	var entities []*models.Resource
	err := EnumBundleResources(bundle, "Resource", func(resource ResourceData) error {
		var entity models.Resource
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToResources(resp *FhirResponse) ([]*models.Resource, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToResources(resp.Bundle)
	case "Resource":
		var entity models.Resource
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Resource{&entity}, nil
	}
	return nil, nil
}

func fhirRespToResource(id string, resp *FhirResponse) (*models.Resource, error) {
	entities, err := fhirRespToResources(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Resource", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Resource.
func (c *Client) GetResource(ctx context.Context, params Parameters) ([]*models.Resource, error) {
	resp, err := c.Get(ctx, "Resource", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResources(resp)
}

// Get Resource by ID.
func (c *Client) GetResourceByID(ctx context.Context, id string, params Parameters) (*models.Resource, error) {
	resp, err := c.GetByID(ctx, "Resource", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResource(id, resp)
}

func (c *Client) CreateResource(ctx context.Context, resource string, params Parameters, entity *models.Resource) (*models.Resource, error) {
	resp, err := c.Create(ctx, "Resource", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResource(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateResource(ctx context.Context, resource string, params Parameters, entity *models.Resource) (*models.Resource, error) {
	resp, err := c.Update(ctx, "Resource", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResource(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateResourceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Resource) (*models.Resource, error) {
	resp, err := c.UpdateByID(ctx, "Resource", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResource(id, resp)
}

func (c *Client) PatchResource(ctx context.Context, resource string, params Parameters, entity *models.Resource) ([]*models.Resource, error) {
	resp, err := c.Patch(ctx, "Resource", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResources(resp)
}

func (c *Client) PatchResourceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Resource) (*models.Resource, error) {
	resp, err := c.PatchByID(ctx, "Resource", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToResource(id, resp)
}

func (c *Client) DeleteResource(ctx context.Context, resource string, params Parameters) ([]*models.Resource, error) {
	resp, err := c.Delete(ctx, "Resource", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResources(resp)
}

func (c *Client) DeleteResourceByID(ctx context.Context, resource, id string, params Parameters) (*models.Resource, error) {
	resp, err := c.DeleteByID(ctx, "Resource", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToResource(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// RiskAssessment
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToRiskAssessments(bundle *models.Bundle) ([]*models.RiskAssessment, error) {
	var entities []*models.RiskAssessment
	err := EnumBundleResources(bundle, "RiskAssessment", func(resource ResourceData) error {
		var entity models.RiskAssessment
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToRiskAssessments(resp *FhirResponse) ([]*models.RiskAssessment, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToRiskAssessments(resp.Bundle)
	case "RiskAssessment":
		var entity models.RiskAssessment
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.RiskAssessment{&entity}, nil
	}
	return nil, nil
}

func fhirRespToRiskAssessment(id string, resp *FhirResponse) (*models.RiskAssessment, error) {
	entities, err := fhirRespToRiskAssessments(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "RiskAssessment", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get RiskAssessment.
func (c *Client) GetRiskAssessment(ctx context.Context, params Parameters) ([]*models.RiskAssessment, error) {
	resp, err := c.Get(ctx, "RiskAssessment", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskAssessments(resp)
}

// Get RiskAssessment by ID.
func (c *Client) GetRiskAssessmentByID(ctx context.Context, id string, params Parameters) (*models.RiskAssessment, error) {
	resp, err := c.GetByID(ctx, "RiskAssessment", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskAssessment(id, resp)
}

func (c *Client) CreateRiskAssessment(ctx context.Context, resource string, params Parameters, entity *models.RiskAssessment) (*models.RiskAssessment, error) {
	resp, err := c.Create(ctx, "RiskAssessment", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskAssessment(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateRiskAssessment(ctx context.Context, resource string, params Parameters, entity *models.RiskAssessment) (*models.RiskAssessment, error) {
	resp, err := c.Update(ctx, "RiskAssessment", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskAssessment(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateRiskAssessmentByID(ctx context.Context, resource, id string, params Parameters, entity *models.RiskAssessment) (*models.RiskAssessment, error) {
	resp, err := c.UpdateByID(ctx, "RiskAssessment", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskAssessment(id, resp)
}

func (c *Client) PatchRiskAssessment(ctx context.Context, resource string, params Parameters, entity *models.RiskAssessment) ([]*models.RiskAssessment, error) {
	resp, err := c.Patch(ctx, "RiskAssessment", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskAssessments(resp)
}

func (c *Client) PatchRiskAssessmentByID(ctx context.Context, resource, id string, params Parameters, entity *models.RiskAssessment) (*models.RiskAssessment, error) {
	resp, err := c.PatchByID(ctx, "RiskAssessment", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskAssessment(id, resp)
}

func (c *Client) DeleteRiskAssessment(ctx context.Context, resource string, params Parameters) ([]*models.RiskAssessment, error) {
	resp, err := c.Delete(ctx, "RiskAssessment", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskAssessments(resp)
}

func (c *Client) DeleteRiskAssessmentByID(ctx context.Context, resource, id string, params Parameters) (*models.RiskAssessment, error) {
	resp, err := c.DeleteByID(ctx, "RiskAssessment", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskAssessment(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// RiskEvidenceSynthesis
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToRiskEvidenceSynthesiss(bundle *models.Bundle) ([]*models.RiskEvidenceSynthesis, error) {
	var entities []*models.RiskEvidenceSynthesis
	err := EnumBundleResources(bundle, "RiskEvidenceSynthesis", func(resource ResourceData) error {
		var entity models.RiskEvidenceSynthesis
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToRiskEvidenceSynthesiss(resp *FhirResponse) ([]*models.RiskEvidenceSynthesis, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToRiskEvidenceSynthesiss(resp.Bundle)
	case "RiskEvidenceSynthesis":
		var entity models.RiskEvidenceSynthesis
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.RiskEvidenceSynthesis{&entity}, nil
	}
	return nil, nil
}

func fhirRespToRiskEvidenceSynthesis(id string, resp *FhirResponse) (*models.RiskEvidenceSynthesis, error) {
	entities, err := fhirRespToRiskEvidenceSynthesiss(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "RiskEvidenceSynthesis", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get RiskEvidenceSynthesis.
func (c *Client) GetRiskEvidenceSynthesis(ctx context.Context, params Parameters) ([]*models.RiskEvidenceSynthesis, error) {
	resp, err := c.Get(ctx, "RiskEvidenceSynthesis", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskEvidenceSynthesiss(resp)
}

// Get RiskEvidenceSynthesis by ID.
func (c *Client) GetRiskEvidenceSynthesisByID(ctx context.Context, id string, params Parameters) (*models.RiskEvidenceSynthesis, error) {
	resp, err := c.GetByID(ctx, "RiskEvidenceSynthesis", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskEvidenceSynthesis(id, resp)
}

func (c *Client) CreateRiskEvidenceSynthesis(ctx context.Context, resource string, params Parameters, entity *models.RiskEvidenceSynthesis) (*models.RiskEvidenceSynthesis, error) {
	resp, err := c.Create(ctx, "RiskEvidenceSynthesis", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskEvidenceSynthesis(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateRiskEvidenceSynthesis(ctx context.Context, resource string, params Parameters, entity *models.RiskEvidenceSynthesis) (*models.RiskEvidenceSynthesis, error) {
	resp, err := c.Update(ctx, "RiskEvidenceSynthesis", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskEvidenceSynthesis(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateRiskEvidenceSynthesisByID(ctx context.Context, resource, id string, params Parameters, entity *models.RiskEvidenceSynthesis) (*models.RiskEvidenceSynthesis, error) {
	resp, err := c.UpdateByID(ctx, "RiskEvidenceSynthesis", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskEvidenceSynthesis(id, resp)
}

func (c *Client) PatchRiskEvidenceSynthesis(ctx context.Context, resource string, params Parameters, entity *models.RiskEvidenceSynthesis) ([]*models.RiskEvidenceSynthesis, error) {
	resp, err := c.Patch(ctx, "RiskEvidenceSynthesis", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskEvidenceSynthesiss(resp)
}

func (c *Client) PatchRiskEvidenceSynthesisByID(ctx context.Context, resource, id string, params Parameters, entity *models.RiskEvidenceSynthesis) (*models.RiskEvidenceSynthesis, error) {
	resp, err := c.PatchByID(ctx, "RiskEvidenceSynthesis", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskEvidenceSynthesis(id, resp)
}

func (c *Client) DeleteRiskEvidenceSynthesis(ctx context.Context, resource string, params Parameters) ([]*models.RiskEvidenceSynthesis, error) {
	resp, err := c.Delete(ctx, "RiskEvidenceSynthesis", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskEvidenceSynthesiss(resp)
}

func (c *Client) DeleteRiskEvidenceSynthesisByID(ctx context.Context, resource, id string, params Parameters) (*models.RiskEvidenceSynthesis, error) {
	resp, err := c.DeleteByID(ctx, "RiskEvidenceSynthesis", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToRiskEvidenceSynthesis(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Schedule
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSchedules(bundle *models.Bundle) ([]*models.Schedule, error) {
	var entities []*models.Schedule
	err := EnumBundleResources(bundle, "Schedule", func(resource ResourceData) error {
		var entity models.Schedule
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSchedules(resp *FhirResponse) ([]*models.Schedule, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSchedules(resp.Bundle)
	case "Schedule":
		var entity models.Schedule
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Schedule{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSchedule(id string, resp *FhirResponse) (*models.Schedule, error) {
	entities, err := fhirRespToSchedules(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Schedule", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Schedule.
func (c *Client) GetSchedule(ctx context.Context, params Parameters) ([]*models.Schedule, error) {
	resp, err := c.Get(ctx, "Schedule", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSchedules(resp)
}

// Get Schedule by ID.
func (c *Client) GetScheduleByID(ctx context.Context, id string, params Parameters) (*models.Schedule, error) {
	resp, err := c.GetByID(ctx, "Schedule", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSchedule(id, resp)
}

func (c *Client) CreateSchedule(ctx context.Context, resource string, params Parameters, entity *models.Schedule) (*models.Schedule, error) {
	resp, err := c.Create(ctx, "Schedule", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSchedule(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSchedule(ctx context.Context, resource string, params Parameters, entity *models.Schedule) (*models.Schedule, error) {
	resp, err := c.Update(ctx, "Schedule", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSchedule(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateScheduleByID(ctx context.Context, resource, id string, params Parameters, entity *models.Schedule) (*models.Schedule, error) {
	resp, err := c.UpdateByID(ctx, "Schedule", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSchedule(id, resp)
}

func (c *Client) PatchSchedule(ctx context.Context, resource string, params Parameters, entity *models.Schedule) ([]*models.Schedule, error) {
	resp, err := c.Patch(ctx, "Schedule", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSchedules(resp)
}

func (c *Client) PatchScheduleByID(ctx context.Context, resource, id string, params Parameters, entity *models.Schedule) (*models.Schedule, error) {
	resp, err := c.PatchByID(ctx, "Schedule", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSchedule(id, resp)
}

func (c *Client) DeleteSchedule(ctx context.Context, resource string, params Parameters) ([]*models.Schedule, error) {
	resp, err := c.Delete(ctx, "Schedule", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSchedules(resp)
}

func (c *Client) DeleteScheduleByID(ctx context.Context, resource, id string, params Parameters) (*models.Schedule, error) {
	resp, err := c.DeleteByID(ctx, "Schedule", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSchedule(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// SearchParameter
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSearchParameters(bundle *models.Bundle) ([]*models.SearchParameter, error) {
	var entities []*models.SearchParameter
	err := EnumBundleResources(bundle, "SearchParameter", func(resource ResourceData) error {
		var entity models.SearchParameter
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSearchParameters(resp *FhirResponse) ([]*models.SearchParameter, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSearchParameters(resp.Bundle)
	case "SearchParameter":
		var entity models.SearchParameter
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.SearchParameter{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSearchParameter(id string, resp *FhirResponse) (*models.SearchParameter, error) {
	entities, err := fhirRespToSearchParameters(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "SearchParameter", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get SearchParameter.
func (c *Client) GetSearchParameter(ctx context.Context, params Parameters) ([]*models.SearchParameter, error) {
	resp, err := c.Get(ctx, "SearchParameter", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSearchParameters(resp)
}

// Get SearchParameter by ID.
func (c *Client) GetSearchParameterByID(ctx context.Context, id string, params Parameters) (*models.SearchParameter, error) {
	resp, err := c.GetByID(ctx, "SearchParameter", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSearchParameter(id, resp)
}

func (c *Client) CreateSearchParameter(ctx context.Context, resource string, params Parameters, entity *models.SearchParameter) (*models.SearchParameter, error) {
	resp, err := c.Create(ctx, "SearchParameter", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSearchParameter(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSearchParameter(ctx context.Context, resource string, params Parameters, entity *models.SearchParameter) (*models.SearchParameter, error) {
	resp, err := c.Update(ctx, "SearchParameter", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSearchParameter(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSearchParameterByID(ctx context.Context, resource, id string, params Parameters, entity *models.SearchParameter) (*models.SearchParameter, error) {
	resp, err := c.UpdateByID(ctx, "SearchParameter", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSearchParameter(id, resp)
}

func (c *Client) PatchSearchParameter(ctx context.Context, resource string, params Parameters, entity *models.SearchParameter) ([]*models.SearchParameter, error) {
	resp, err := c.Patch(ctx, "SearchParameter", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSearchParameters(resp)
}

func (c *Client) PatchSearchParameterByID(ctx context.Context, resource, id string, params Parameters, entity *models.SearchParameter) (*models.SearchParameter, error) {
	resp, err := c.PatchByID(ctx, "SearchParameter", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSearchParameter(id, resp)
}

func (c *Client) DeleteSearchParameter(ctx context.Context, resource string, params Parameters) ([]*models.SearchParameter, error) {
	resp, err := c.Delete(ctx, "SearchParameter", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSearchParameters(resp)
}

func (c *Client) DeleteSearchParameterByID(ctx context.Context, resource, id string, params Parameters) (*models.SearchParameter, error) {
	resp, err := c.DeleteByID(ctx, "SearchParameter", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSearchParameter(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ServiceRequest
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToServiceRequests(bundle *models.Bundle) ([]*models.ServiceRequest, error) {
	var entities []*models.ServiceRequest
	err := EnumBundleResources(bundle, "ServiceRequest", func(resource ResourceData) error {
		var entity models.ServiceRequest
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToServiceRequests(resp *FhirResponse) ([]*models.ServiceRequest, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToServiceRequests(resp.Bundle)
	case "ServiceRequest":
		var entity models.ServiceRequest
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ServiceRequest{&entity}, nil
	}
	return nil, nil
}

func fhirRespToServiceRequest(id string, resp *FhirResponse) (*models.ServiceRequest, error) {
	entities, err := fhirRespToServiceRequests(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ServiceRequest", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ServiceRequest.
func (c *Client) GetServiceRequest(ctx context.Context, params Parameters) ([]*models.ServiceRequest, error) {
	resp, err := c.Get(ctx, "ServiceRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToServiceRequests(resp)
}

// Get ServiceRequest by ID.
func (c *Client) GetServiceRequestByID(ctx context.Context, id string, params Parameters) (*models.ServiceRequest, error) {
	resp, err := c.GetByID(ctx, "ServiceRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToServiceRequest(id, resp)
}

func (c *Client) CreateServiceRequest(ctx context.Context, resource string, params Parameters, entity *models.ServiceRequest) (*models.ServiceRequest, error) {
	resp, err := c.Create(ctx, "ServiceRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToServiceRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateServiceRequest(ctx context.Context, resource string, params Parameters, entity *models.ServiceRequest) (*models.ServiceRequest, error) {
	resp, err := c.Update(ctx, "ServiceRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToServiceRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateServiceRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.ServiceRequest) (*models.ServiceRequest, error) {
	resp, err := c.UpdateByID(ctx, "ServiceRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToServiceRequest(id, resp)
}

func (c *Client) PatchServiceRequest(ctx context.Context, resource string, params Parameters, entity *models.ServiceRequest) ([]*models.ServiceRequest, error) {
	resp, err := c.Patch(ctx, "ServiceRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToServiceRequests(resp)
}

func (c *Client) PatchServiceRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.ServiceRequest) (*models.ServiceRequest, error) {
	resp, err := c.PatchByID(ctx, "ServiceRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToServiceRequest(id, resp)
}

func (c *Client) DeleteServiceRequest(ctx context.Context, resource string, params Parameters) ([]*models.ServiceRequest, error) {
	resp, err := c.Delete(ctx, "ServiceRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToServiceRequests(resp)
}

func (c *Client) DeleteServiceRequestByID(ctx context.Context, resource, id string, params Parameters) (*models.ServiceRequest, error) {
	resp, err := c.DeleteByID(ctx, "ServiceRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToServiceRequest(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Slot
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSlots(bundle *models.Bundle) ([]*models.Slot, error) {
	var entities []*models.Slot
	err := EnumBundleResources(bundle, "Slot", func(resource ResourceData) error {
		var entity models.Slot
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSlots(resp *FhirResponse) ([]*models.Slot, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSlots(resp.Bundle)
	case "Slot":
		var entity models.Slot
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Slot{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSlot(id string, resp *FhirResponse) (*models.Slot, error) {
	entities, err := fhirRespToSlots(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Slot", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Slot.
func (c *Client) GetSlot(ctx context.Context, params Parameters) ([]*models.Slot, error) {
	resp, err := c.Get(ctx, "Slot", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSlots(resp)
}

// Get Slot by ID.
func (c *Client) GetSlotByID(ctx context.Context, id string, params Parameters) (*models.Slot, error) {
	resp, err := c.GetByID(ctx, "Slot", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSlot(id, resp)
}

func (c *Client) CreateSlot(ctx context.Context, resource string, params Parameters, entity *models.Slot) (*models.Slot, error) {
	resp, err := c.Create(ctx, "Slot", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSlot(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSlot(ctx context.Context, resource string, params Parameters, entity *models.Slot) (*models.Slot, error) {
	resp, err := c.Update(ctx, "Slot", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSlot(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSlotByID(ctx context.Context, resource, id string, params Parameters, entity *models.Slot) (*models.Slot, error) {
	resp, err := c.UpdateByID(ctx, "Slot", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSlot(id, resp)
}

func (c *Client) PatchSlot(ctx context.Context, resource string, params Parameters, entity *models.Slot) ([]*models.Slot, error) {
	resp, err := c.Patch(ctx, "Slot", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSlots(resp)
}

func (c *Client) PatchSlotByID(ctx context.Context, resource, id string, params Parameters, entity *models.Slot) (*models.Slot, error) {
	resp, err := c.PatchByID(ctx, "Slot", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSlot(id, resp)
}

func (c *Client) DeleteSlot(ctx context.Context, resource string, params Parameters) ([]*models.Slot, error) {
	resp, err := c.Delete(ctx, "Slot", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSlots(resp)
}

func (c *Client) DeleteSlotByID(ctx context.Context, resource, id string, params Parameters) (*models.Slot, error) {
	resp, err := c.DeleteByID(ctx, "Slot", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSlot(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Specimen
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSpecimens(bundle *models.Bundle) ([]*models.Specimen, error) {
	var entities []*models.Specimen
	err := EnumBundleResources(bundle, "Specimen", func(resource ResourceData) error {
		var entity models.Specimen
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSpecimens(resp *FhirResponse) ([]*models.Specimen, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSpecimens(resp.Bundle)
	case "Specimen":
		var entity models.Specimen
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Specimen{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSpecimen(id string, resp *FhirResponse) (*models.Specimen, error) {
	entities, err := fhirRespToSpecimens(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Specimen", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Specimen.
func (c *Client) GetSpecimen(ctx context.Context, params Parameters) ([]*models.Specimen, error) {
	resp, err := c.Get(ctx, "Specimen", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimens(resp)
}

// Get Specimen by ID.
func (c *Client) GetSpecimenByID(ctx context.Context, id string, params Parameters) (*models.Specimen, error) {
	resp, err := c.GetByID(ctx, "Specimen", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimen(id, resp)
}

func (c *Client) CreateSpecimen(ctx context.Context, resource string, params Parameters, entity *models.Specimen) (*models.Specimen, error) {
	resp, err := c.Create(ctx, "Specimen", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimen(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSpecimen(ctx context.Context, resource string, params Parameters, entity *models.Specimen) (*models.Specimen, error) {
	resp, err := c.Update(ctx, "Specimen", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimen(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSpecimenByID(ctx context.Context, resource, id string, params Parameters, entity *models.Specimen) (*models.Specimen, error) {
	resp, err := c.UpdateByID(ctx, "Specimen", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimen(id, resp)
}

func (c *Client) PatchSpecimen(ctx context.Context, resource string, params Parameters, entity *models.Specimen) ([]*models.Specimen, error) {
	resp, err := c.Patch(ctx, "Specimen", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimens(resp)
}

func (c *Client) PatchSpecimenByID(ctx context.Context, resource, id string, params Parameters, entity *models.Specimen) (*models.Specimen, error) {
	resp, err := c.PatchByID(ctx, "Specimen", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimen(id, resp)
}

func (c *Client) DeleteSpecimen(ctx context.Context, resource string, params Parameters) ([]*models.Specimen, error) {
	resp, err := c.Delete(ctx, "Specimen", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimens(resp)
}

func (c *Client) DeleteSpecimenByID(ctx context.Context, resource, id string, params Parameters) (*models.Specimen, error) {
	resp, err := c.DeleteByID(ctx, "Specimen", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimen(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// SpecimenDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSpecimenDefinitions(bundle *models.Bundle) ([]*models.SpecimenDefinition, error) {
	var entities []*models.SpecimenDefinition
	err := EnumBundleResources(bundle, "SpecimenDefinition", func(resource ResourceData) error {
		var entity models.SpecimenDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSpecimenDefinitions(resp *FhirResponse) ([]*models.SpecimenDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSpecimenDefinitions(resp.Bundle)
	case "SpecimenDefinition":
		var entity models.SpecimenDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.SpecimenDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSpecimenDefinition(id string, resp *FhirResponse) (*models.SpecimenDefinition, error) {
	entities, err := fhirRespToSpecimenDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "SpecimenDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get SpecimenDefinition.
func (c *Client) GetSpecimenDefinition(ctx context.Context, params Parameters) ([]*models.SpecimenDefinition, error) {
	resp, err := c.Get(ctx, "SpecimenDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimenDefinitions(resp)
}

// Get SpecimenDefinition by ID.
func (c *Client) GetSpecimenDefinitionByID(ctx context.Context, id string, params Parameters) (*models.SpecimenDefinition, error) {
	resp, err := c.GetByID(ctx, "SpecimenDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimenDefinition(id, resp)
}

func (c *Client) CreateSpecimenDefinition(ctx context.Context, resource string, params Parameters, entity *models.SpecimenDefinition) (*models.SpecimenDefinition, error) {
	resp, err := c.Create(ctx, "SpecimenDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimenDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSpecimenDefinition(ctx context.Context, resource string, params Parameters, entity *models.SpecimenDefinition) (*models.SpecimenDefinition, error) {
	resp, err := c.Update(ctx, "SpecimenDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimenDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSpecimenDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.SpecimenDefinition) (*models.SpecimenDefinition, error) {
	resp, err := c.UpdateByID(ctx, "SpecimenDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimenDefinition(id, resp)
}

func (c *Client) PatchSpecimenDefinition(ctx context.Context, resource string, params Parameters, entity *models.SpecimenDefinition) ([]*models.SpecimenDefinition, error) {
	resp, err := c.Patch(ctx, "SpecimenDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimenDefinitions(resp)
}

func (c *Client) PatchSpecimenDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.SpecimenDefinition) (*models.SpecimenDefinition, error) {
	resp, err := c.PatchByID(ctx, "SpecimenDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimenDefinition(id, resp)
}

func (c *Client) DeleteSpecimenDefinition(ctx context.Context, resource string, params Parameters) ([]*models.SpecimenDefinition, error) {
	resp, err := c.Delete(ctx, "SpecimenDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimenDefinitions(resp)
}

func (c *Client) DeleteSpecimenDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.SpecimenDefinition, error) {
	resp, err := c.DeleteByID(ctx, "SpecimenDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSpecimenDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// StructureDefinition
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToStructureDefinitions(bundle *models.Bundle) ([]*models.StructureDefinition, error) {
	var entities []*models.StructureDefinition
	err := EnumBundleResources(bundle, "StructureDefinition", func(resource ResourceData) error {
		var entity models.StructureDefinition
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToStructureDefinitions(resp *FhirResponse) ([]*models.StructureDefinition, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToStructureDefinitions(resp.Bundle)
	case "StructureDefinition":
		var entity models.StructureDefinition
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.StructureDefinition{&entity}, nil
	}
	return nil, nil
}

func fhirRespToStructureDefinition(id string, resp *FhirResponse) (*models.StructureDefinition, error) {
	entities, err := fhirRespToStructureDefinitions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "StructureDefinition", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get StructureDefinition.
func (c *Client) GetStructureDefinition(ctx context.Context, params Parameters) ([]*models.StructureDefinition, error) {
	resp, err := c.Get(ctx, "StructureDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureDefinitions(resp)
}

// Get StructureDefinition by ID.
func (c *Client) GetStructureDefinitionByID(ctx context.Context, id string, params Parameters) (*models.StructureDefinition, error) {
	resp, err := c.GetByID(ctx, "StructureDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureDefinition(id, resp)
}

func (c *Client) CreateStructureDefinition(ctx context.Context, resource string, params Parameters, entity *models.StructureDefinition) (*models.StructureDefinition, error) {
	resp, err := c.Create(ctx, "StructureDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateStructureDefinition(ctx context.Context, resource string, params Parameters, entity *models.StructureDefinition) (*models.StructureDefinition, error) {
	resp, err := c.Update(ctx, "StructureDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureDefinition(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateStructureDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.StructureDefinition) (*models.StructureDefinition, error) {
	resp, err := c.UpdateByID(ctx, "StructureDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureDefinition(id, resp)
}

func (c *Client) PatchStructureDefinition(ctx context.Context, resource string, params Parameters, entity *models.StructureDefinition) ([]*models.StructureDefinition, error) {
	resp, err := c.Patch(ctx, "StructureDefinition", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureDefinitions(resp)
}

func (c *Client) PatchStructureDefinitionByID(ctx context.Context, resource, id string, params Parameters, entity *models.StructureDefinition) (*models.StructureDefinition, error) {
	resp, err := c.PatchByID(ctx, "StructureDefinition", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureDefinition(id, resp)
}

func (c *Client) DeleteStructureDefinition(ctx context.Context, resource string, params Parameters) ([]*models.StructureDefinition, error) {
	resp, err := c.Delete(ctx, "StructureDefinition", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureDefinitions(resp)
}

func (c *Client) DeleteStructureDefinitionByID(ctx context.Context, resource, id string, params Parameters) (*models.StructureDefinition, error) {
	resp, err := c.DeleteByID(ctx, "StructureDefinition", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureDefinition(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// StructureMap
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToStructureMaps(bundle *models.Bundle) ([]*models.StructureMap, error) {
	var entities []*models.StructureMap
	err := EnumBundleResources(bundle, "StructureMap", func(resource ResourceData) error {
		var entity models.StructureMap
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToStructureMaps(resp *FhirResponse) ([]*models.StructureMap, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToStructureMaps(resp.Bundle)
	case "StructureMap":
		var entity models.StructureMap
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.StructureMap{&entity}, nil
	}
	return nil, nil
}

func fhirRespToStructureMap(id string, resp *FhirResponse) (*models.StructureMap, error) {
	entities, err := fhirRespToStructureMaps(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "StructureMap", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get StructureMap.
func (c *Client) GetStructureMap(ctx context.Context, params Parameters) ([]*models.StructureMap, error) {
	resp, err := c.Get(ctx, "StructureMap", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureMaps(resp)
}

// Get StructureMap by ID.
func (c *Client) GetStructureMapByID(ctx context.Context, id string, params Parameters) (*models.StructureMap, error) {
	resp, err := c.GetByID(ctx, "StructureMap", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureMap(id, resp)
}

func (c *Client) CreateStructureMap(ctx context.Context, resource string, params Parameters, entity *models.StructureMap) (*models.StructureMap, error) {
	resp, err := c.Create(ctx, "StructureMap", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureMap(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateStructureMap(ctx context.Context, resource string, params Parameters, entity *models.StructureMap) (*models.StructureMap, error) {
	resp, err := c.Update(ctx, "StructureMap", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureMap(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateStructureMapByID(ctx context.Context, resource, id string, params Parameters, entity *models.StructureMap) (*models.StructureMap, error) {
	resp, err := c.UpdateByID(ctx, "StructureMap", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureMap(id, resp)
}

func (c *Client) PatchStructureMap(ctx context.Context, resource string, params Parameters, entity *models.StructureMap) ([]*models.StructureMap, error) {
	resp, err := c.Patch(ctx, "StructureMap", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureMaps(resp)
}

func (c *Client) PatchStructureMapByID(ctx context.Context, resource, id string, params Parameters, entity *models.StructureMap) (*models.StructureMap, error) {
	resp, err := c.PatchByID(ctx, "StructureMap", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureMap(id, resp)
}

func (c *Client) DeleteStructureMap(ctx context.Context, resource string, params Parameters) ([]*models.StructureMap, error) {
	resp, err := c.Delete(ctx, "StructureMap", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureMaps(resp)
}

func (c *Client) DeleteStructureMapByID(ctx context.Context, resource, id string, params Parameters) (*models.StructureMap, error) {
	resp, err := c.DeleteByID(ctx, "StructureMap", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToStructureMap(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Subscription
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSubscriptions(bundle *models.Bundle) ([]*models.Subscription, error) {
	var entities []*models.Subscription
	err := EnumBundleResources(bundle, "Subscription", func(resource ResourceData) error {
		var entity models.Subscription
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSubscriptions(resp *FhirResponse) ([]*models.Subscription, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSubscriptions(resp.Bundle)
	case "Subscription":
		var entity models.Subscription
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Subscription{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSubscription(id string, resp *FhirResponse) (*models.Subscription, error) {
	entities, err := fhirRespToSubscriptions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Subscription", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Subscription.
func (c *Client) GetSubscription(ctx context.Context, params Parameters) ([]*models.Subscription, error) {
	resp, err := c.Get(ctx, "Subscription", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubscriptions(resp)
}

// Get Subscription by ID.
func (c *Client) GetSubscriptionByID(ctx context.Context, id string, params Parameters) (*models.Subscription, error) {
	resp, err := c.GetByID(ctx, "Subscription", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubscription(id, resp)
}

func (c *Client) CreateSubscription(ctx context.Context, resource string, params Parameters, entity *models.Subscription) (*models.Subscription, error) {
	resp, err := c.Create(ctx, "Subscription", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubscription(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubscription(ctx context.Context, resource string, params Parameters, entity *models.Subscription) (*models.Subscription, error) {
	resp, err := c.Update(ctx, "Subscription", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubscription(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubscriptionByID(ctx context.Context, resource, id string, params Parameters, entity *models.Subscription) (*models.Subscription, error) {
	resp, err := c.UpdateByID(ctx, "Subscription", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubscription(id, resp)
}

func (c *Client) PatchSubscription(ctx context.Context, resource string, params Parameters, entity *models.Subscription) ([]*models.Subscription, error) {
	resp, err := c.Patch(ctx, "Subscription", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubscriptions(resp)
}

func (c *Client) PatchSubscriptionByID(ctx context.Context, resource, id string, params Parameters, entity *models.Subscription) (*models.Subscription, error) {
	resp, err := c.PatchByID(ctx, "Subscription", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubscription(id, resp)
}

func (c *Client) DeleteSubscription(ctx context.Context, resource string, params Parameters) ([]*models.Subscription, error) {
	resp, err := c.Delete(ctx, "Subscription", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubscriptions(resp)
}

func (c *Client) DeleteSubscriptionByID(ctx context.Context, resource, id string, params Parameters) (*models.Subscription, error) {
	resp, err := c.DeleteByID(ctx, "Subscription", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubscription(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Substance
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSubstances(bundle *models.Bundle) ([]*models.Substance, error) {
	var entities []*models.Substance
	err := EnumBundleResources(bundle, "Substance", func(resource ResourceData) error {
		var entity models.Substance
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSubstances(resp *FhirResponse) ([]*models.Substance, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSubstances(resp.Bundle)
	case "Substance":
		var entity models.Substance
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Substance{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSubstance(id string, resp *FhirResponse) (*models.Substance, error) {
	entities, err := fhirRespToSubstances(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Substance", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Substance.
func (c *Client) GetSubstance(ctx context.Context, params Parameters) ([]*models.Substance, error) {
	resp, err := c.Get(ctx, "Substance", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstances(resp)
}

// Get Substance by ID.
func (c *Client) GetSubstanceByID(ctx context.Context, id string, params Parameters) (*models.Substance, error) {
	resp, err := c.GetByID(ctx, "Substance", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstance(id, resp)
}

func (c *Client) CreateSubstance(ctx context.Context, resource string, params Parameters, entity *models.Substance) (*models.Substance, error) {
	resp, err := c.Create(ctx, "Substance", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstance(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstance(ctx context.Context, resource string, params Parameters, entity *models.Substance) (*models.Substance, error) {
	resp, err := c.Update(ctx, "Substance", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstance(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Substance) (*models.Substance, error) {
	resp, err := c.UpdateByID(ctx, "Substance", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstance(id, resp)
}

func (c *Client) PatchSubstance(ctx context.Context, resource string, params Parameters, entity *models.Substance) ([]*models.Substance, error) {
	resp, err := c.Patch(ctx, "Substance", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstances(resp)
}

func (c *Client) PatchSubstanceByID(ctx context.Context, resource, id string, params Parameters, entity *models.Substance) (*models.Substance, error) {
	resp, err := c.PatchByID(ctx, "Substance", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstance(id, resp)
}

func (c *Client) DeleteSubstance(ctx context.Context, resource string, params Parameters) ([]*models.Substance, error) {
	resp, err := c.Delete(ctx, "Substance", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstances(resp)
}

func (c *Client) DeleteSubstanceByID(ctx context.Context, resource, id string, params Parameters) (*models.Substance, error) {
	resp, err := c.DeleteByID(ctx, "Substance", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstance(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// SubstanceNucleicAcid
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSubstanceNucleicAcids(bundle *models.Bundle) ([]*models.SubstanceNucleicAcid, error) {
	var entities []*models.SubstanceNucleicAcid
	err := EnumBundleResources(bundle, "SubstanceNucleicAcid", func(resource ResourceData) error {
		var entity models.SubstanceNucleicAcid
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSubstanceNucleicAcids(resp *FhirResponse) ([]*models.SubstanceNucleicAcid, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSubstanceNucleicAcids(resp.Bundle)
	case "SubstanceNucleicAcid":
		var entity models.SubstanceNucleicAcid
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.SubstanceNucleicAcid{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSubstanceNucleicAcid(id string, resp *FhirResponse) (*models.SubstanceNucleicAcid, error) {
	entities, err := fhirRespToSubstanceNucleicAcids(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "SubstanceNucleicAcid", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get SubstanceNucleicAcid.
func (c *Client) GetSubstanceNucleicAcid(ctx context.Context, params Parameters) ([]*models.SubstanceNucleicAcid, error) {
	resp, err := c.Get(ctx, "SubstanceNucleicAcid", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceNucleicAcids(resp)
}

// Get SubstanceNucleicAcid by ID.
func (c *Client) GetSubstanceNucleicAcidByID(ctx context.Context, id string, params Parameters) (*models.SubstanceNucleicAcid, error) {
	resp, err := c.GetByID(ctx, "SubstanceNucleicAcid", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceNucleicAcid(id, resp)
}

func (c *Client) CreateSubstanceNucleicAcid(ctx context.Context, resource string, params Parameters, entity *models.SubstanceNucleicAcid) (*models.SubstanceNucleicAcid, error) {
	resp, err := c.Create(ctx, "SubstanceNucleicAcid", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceNucleicAcid(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceNucleicAcid(ctx context.Context, resource string, params Parameters, entity *models.SubstanceNucleicAcid) (*models.SubstanceNucleicAcid, error) {
	resp, err := c.Update(ctx, "SubstanceNucleicAcid", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceNucleicAcid(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceNucleicAcidByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstanceNucleicAcid) (*models.SubstanceNucleicAcid, error) {
	resp, err := c.UpdateByID(ctx, "SubstanceNucleicAcid", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceNucleicAcid(id, resp)
}

func (c *Client) PatchSubstanceNucleicAcid(ctx context.Context, resource string, params Parameters, entity *models.SubstanceNucleicAcid) ([]*models.SubstanceNucleicAcid, error) {
	resp, err := c.Patch(ctx, "SubstanceNucleicAcid", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceNucleicAcids(resp)
}

func (c *Client) PatchSubstanceNucleicAcidByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstanceNucleicAcid) (*models.SubstanceNucleicAcid, error) {
	resp, err := c.PatchByID(ctx, "SubstanceNucleicAcid", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceNucleicAcid(id, resp)
}

func (c *Client) DeleteSubstanceNucleicAcid(ctx context.Context, resource string, params Parameters) ([]*models.SubstanceNucleicAcid, error) {
	resp, err := c.Delete(ctx, "SubstanceNucleicAcid", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceNucleicAcids(resp)
}

func (c *Client) DeleteSubstanceNucleicAcidByID(ctx context.Context, resource, id string, params Parameters) (*models.SubstanceNucleicAcid, error) {
	resp, err := c.DeleteByID(ctx, "SubstanceNucleicAcid", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceNucleicAcid(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// SubstancePolymer
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSubstancePolymers(bundle *models.Bundle) ([]*models.SubstancePolymer, error) {
	var entities []*models.SubstancePolymer
	err := EnumBundleResources(bundle, "SubstancePolymer", func(resource ResourceData) error {
		var entity models.SubstancePolymer
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSubstancePolymers(resp *FhirResponse) ([]*models.SubstancePolymer, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSubstancePolymers(resp.Bundle)
	case "SubstancePolymer":
		var entity models.SubstancePolymer
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.SubstancePolymer{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSubstancePolymer(id string, resp *FhirResponse) (*models.SubstancePolymer, error) {
	entities, err := fhirRespToSubstancePolymers(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "SubstancePolymer", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get SubstancePolymer.
func (c *Client) GetSubstancePolymer(ctx context.Context, params Parameters) ([]*models.SubstancePolymer, error) {
	resp, err := c.Get(ctx, "SubstancePolymer", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstancePolymers(resp)
}

// Get SubstancePolymer by ID.
func (c *Client) GetSubstancePolymerByID(ctx context.Context, id string, params Parameters) (*models.SubstancePolymer, error) {
	resp, err := c.GetByID(ctx, "SubstancePolymer", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstancePolymer(id, resp)
}

func (c *Client) CreateSubstancePolymer(ctx context.Context, resource string, params Parameters, entity *models.SubstancePolymer) (*models.SubstancePolymer, error) {
	resp, err := c.Create(ctx, "SubstancePolymer", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstancePolymer(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstancePolymer(ctx context.Context, resource string, params Parameters, entity *models.SubstancePolymer) (*models.SubstancePolymer, error) {
	resp, err := c.Update(ctx, "SubstancePolymer", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstancePolymer(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstancePolymerByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstancePolymer) (*models.SubstancePolymer, error) {
	resp, err := c.UpdateByID(ctx, "SubstancePolymer", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstancePolymer(id, resp)
}

func (c *Client) PatchSubstancePolymer(ctx context.Context, resource string, params Parameters, entity *models.SubstancePolymer) ([]*models.SubstancePolymer, error) {
	resp, err := c.Patch(ctx, "SubstancePolymer", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstancePolymers(resp)
}

func (c *Client) PatchSubstancePolymerByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstancePolymer) (*models.SubstancePolymer, error) {
	resp, err := c.PatchByID(ctx, "SubstancePolymer", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstancePolymer(id, resp)
}

func (c *Client) DeleteSubstancePolymer(ctx context.Context, resource string, params Parameters) ([]*models.SubstancePolymer, error) {
	resp, err := c.Delete(ctx, "SubstancePolymer", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstancePolymers(resp)
}

func (c *Client) DeleteSubstancePolymerByID(ctx context.Context, resource, id string, params Parameters) (*models.SubstancePolymer, error) {
	resp, err := c.DeleteByID(ctx, "SubstancePolymer", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstancePolymer(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// SubstanceProtein
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSubstanceProteins(bundle *models.Bundle) ([]*models.SubstanceProtein, error) {
	var entities []*models.SubstanceProtein
	err := EnumBundleResources(bundle, "SubstanceProtein", func(resource ResourceData) error {
		var entity models.SubstanceProtein
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSubstanceProteins(resp *FhirResponse) ([]*models.SubstanceProtein, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSubstanceProteins(resp.Bundle)
	case "SubstanceProtein":
		var entity models.SubstanceProtein
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.SubstanceProtein{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSubstanceProtein(id string, resp *FhirResponse) (*models.SubstanceProtein, error) {
	entities, err := fhirRespToSubstanceProteins(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "SubstanceProtein", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get SubstanceProtein.
func (c *Client) GetSubstanceProtein(ctx context.Context, params Parameters) ([]*models.SubstanceProtein, error) {
	resp, err := c.Get(ctx, "SubstanceProtein", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceProteins(resp)
}

// Get SubstanceProtein by ID.
func (c *Client) GetSubstanceProteinByID(ctx context.Context, id string, params Parameters) (*models.SubstanceProtein, error) {
	resp, err := c.GetByID(ctx, "SubstanceProtein", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceProtein(id, resp)
}

func (c *Client) CreateSubstanceProtein(ctx context.Context, resource string, params Parameters, entity *models.SubstanceProtein) (*models.SubstanceProtein, error) {
	resp, err := c.Create(ctx, "SubstanceProtein", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceProtein(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceProtein(ctx context.Context, resource string, params Parameters, entity *models.SubstanceProtein) (*models.SubstanceProtein, error) {
	resp, err := c.Update(ctx, "SubstanceProtein", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceProtein(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceProteinByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstanceProtein) (*models.SubstanceProtein, error) {
	resp, err := c.UpdateByID(ctx, "SubstanceProtein", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceProtein(id, resp)
}

func (c *Client) PatchSubstanceProtein(ctx context.Context, resource string, params Parameters, entity *models.SubstanceProtein) ([]*models.SubstanceProtein, error) {
	resp, err := c.Patch(ctx, "SubstanceProtein", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceProteins(resp)
}

func (c *Client) PatchSubstanceProteinByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstanceProtein) (*models.SubstanceProtein, error) {
	resp, err := c.PatchByID(ctx, "SubstanceProtein", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceProtein(id, resp)
}

func (c *Client) DeleteSubstanceProtein(ctx context.Context, resource string, params Parameters) ([]*models.SubstanceProtein, error) {
	resp, err := c.Delete(ctx, "SubstanceProtein", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceProteins(resp)
}

func (c *Client) DeleteSubstanceProteinByID(ctx context.Context, resource, id string, params Parameters) (*models.SubstanceProtein, error) {
	resp, err := c.DeleteByID(ctx, "SubstanceProtein", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceProtein(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// SubstanceReferenceInformation
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSubstanceReferenceInformations(bundle *models.Bundle) ([]*models.SubstanceReferenceInformation, error) {
	var entities []*models.SubstanceReferenceInformation
	err := EnumBundleResources(bundle, "SubstanceReferenceInformation", func(resource ResourceData) error {
		var entity models.SubstanceReferenceInformation
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSubstanceReferenceInformations(resp *FhirResponse) ([]*models.SubstanceReferenceInformation, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSubstanceReferenceInformations(resp.Bundle)
	case "SubstanceReferenceInformation":
		var entity models.SubstanceReferenceInformation
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.SubstanceReferenceInformation{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSubstanceReferenceInformation(id string, resp *FhirResponse) (*models.SubstanceReferenceInformation, error) {
	entities, err := fhirRespToSubstanceReferenceInformations(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "SubstanceReferenceInformation", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get SubstanceReferenceInformation.
func (c *Client) GetSubstanceReferenceInformation(ctx context.Context, params Parameters) ([]*models.SubstanceReferenceInformation, error) {
	resp, err := c.Get(ctx, "SubstanceReferenceInformation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceReferenceInformations(resp)
}

// Get SubstanceReferenceInformation by ID.
func (c *Client) GetSubstanceReferenceInformationByID(ctx context.Context, id string, params Parameters) (*models.SubstanceReferenceInformation, error) {
	resp, err := c.GetByID(ctx, "SubstanceReferenceInformation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceReferenceInformation(id, resp)
}

func (c *Client) CreateSubstanceReferenceInformation(ctx context.Context, resource string, params Parameters, entity *models.SubstanceReferenceInformation) (*models.SubstanceReferenceInformation, error) {
	resp, err := c.Create(ctx, "SubstanceReferenceInformation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceReferenceInformation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceReferenceInformation(ctx context.Context, resource string, params Parameters, entity *models.SubstanceReferenceInformation) (*models.SubstanceReferenceInformation, error) {
	resp, err := c.Update(ctx, "SubstanceReferenceInformation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceReferenceInformation(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceReferenceInformationByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstanceReferenceInformation) (*models.SubstanceReferenceInformation, error) {
	resp, err := c.UpdateByID(ctx, "SubstanceReferenceInformation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceReferenceInformation(id, resp)
}

func (c *Client) PatchSubstanceReferenceInformation(ctx context.Context, resource string, params Parameters, entity *models.SubstanceReferenceInformation) ([]*models.SubstanceReferenceInformation, error) {
	resp, err := c.Patch(ctx, "SubstanceReferenceInformation", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceReferenceInformations(resp)
}

func (c *Client) PatchSubstanceReferenceInformationByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstanceReferenceInformation) (*models.SubstanceReferenceInformation, error) {
	resp, err := c.PatchByID(ctx, "SubstanceReferenceInformation", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceReferenceInformation(id, resp)
}

func (c *Client) DeleteSubstanceReferenceInformation(ctx context.Context, resource string, params Parameters) ([]*models.SubstanceReferenceInformation, error) {
	resp, err := c.Delete(ctx, "SubstanceReferenceInformation", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceReferenceInformations(resp)
}

func (c *Client) DeleteSubstanceReferenceInformationByID(ctx context.Context, resource, id string, params Parameters) (*models.SubstanceReferenceInformation, error) {
	resp, err := c.DeleteByID(ctx, "SubstanceReferenceInformation", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceReferenceInformation(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// SubstanceSourceMaterial
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSubstanceSourceMaterials(bundle *models.Bundle) ([]*models.SubstanceSourceMaterial, error) {
	var entities []*models.SubstanceSourceMaterial
	err := EnumBundleResources(bundle, "SubstanceSourceMaterial", func(resource ResourceData) error {
		var entity models.SubstanceSourceMaterial
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSubstanceSourceMaterials(resp *FhirResponse) ([]*models.SubstanceSourceMaterial, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSubstanceSourceMaterials(resp.Bundle)
	case "SubstanceSourceMaterial":
		var entity models.SubstanceSourceMaterial
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.SubstanceSourceMaterial{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSubstanceSourceMaterial(id string, resp *FhirResponse) (*models.SubstanceSourceMaterial, error) {
	entities, err := fhirRespToSubstanceSourceMaterials(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "SubstanceSourceMaterial", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get SubstanceSourceMaterial.
func (c *Client) GetSubstanceSourceMaterial(ctx context.Context, params Parameters) ([]*models.SubstanceSourceMaterial, error) {
	resp, err := c.Get(ctx, "SubstanceSourceMaterial", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSourceMaterials(resp)
}

// Get SubstanceSourceMaterial by ID.
func (c *Client) GetSubstanceSourceMaterialByID(ctx context.Context, id string, params Parameters) (*models.SubstanceSourceMaterial, error) {
	resp, err := c.GetByID(ctx, "SubstanceSourceMaterial", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSourceMaterial(id, resp)
}

func (c *Client) CreateSubstanceSourceMaterial(ctx context.Context, resource string, params Parameters, entity *models.SubstanceSourceMaterial) (*models.SubstanceSourceMaterial, error) {
	resp, err := c.Create(ctx, "SubstanceSourceMaterial", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSourceMaterial(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceSourceMaterial(ctx context.Context, resource string, params Parameters, entity *models.SubstanceSourceMaterial) (*models.SubstanceSourceMaterial, error) {
	resp, err := c.Update(ctx, "SubstanceSourceMaterial", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSourceMaterial(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceSourceMaterialByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstanceSourceMaterial) (*models.SubstanceSourceMaterial, error) {
	resp, err := c.UpdateByID(ctx, "SubstanceSourceMaterial", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSourceMaterial(id, resp)
}

func (c *Client) PatchSubstanceSourceMaterial(ctx context.Context, resource string, params Parameters, entity *models.SubstanceSourceMaterial) ([]*models.SubstanceSourceMaterial, error) {
	resp, err := c.Patch(ctx, "SubstanceSourceMaterial", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSourceMaterials(resp)
}

func (c *Client) PatchSubstanceSourceMaterialByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstanceSourceMaterial) (*models.SubstanceSourceMaterial, error) {
	resp, err := c.PatchByID(ctx, "SubstanceSourceMaterial", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSourceMaterial(id, resp)
}

func (c *Client) DeleteSubstanceSourceMaterial(ctx context.Context, resource string, params Parameters) ([]*models.SubstanceSourceMaterial, error) {
	resp, err := c.Delete(ctx, "SubstanceSourceMaterial", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSourceMaterials(resp)
}

func (c *Client) DeleteSubstanceSourceMaterialByID(ctx context.Context, resource, id string, params Parameters) (*models.SubstanceSourceMaterial, error) {
	resp, err := c.DeleteByID(ctx, "SubstanceSourceMaterial", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSourceMaterial(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// SubstanceSpecification
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSubstanceSpecifications(bundle *models.Bundle) ([]*models.SubstanceSpecification, error) {
	var entities []*models.SubstanceSpecification
	err := EnumBundleResources(bundle, "SubstanceSpecification", func(resource ResourceData) error {
		var entity models.SubstanceSpecification
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSubstanceSpecifications(resp *FhirResponse) ([]*models.SubstanceSpecification, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSubstanceSpecifications(resp.Bundle)
	case "SubstanceSpecification":
		var entity models.SubstanceSpecification
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.SubstanceSpecification{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSubstanceSpecification(id string, resp *FhirResponse) (*models.SubstanceSpecification, error) {
	entities, err := fhirRespToSubstanceSpecifications(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "SubstanceSpecification", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get SubstanceSpecification.
func (c *Client) GetSubstanceSpecification(ctx context.Context, params Parameters) ([]*models.SubstanceSpecification, error) {
	resp, err := c.Get(ctx, "SubstanceSpecification", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSpecifications(resp)
}

// Get SubstanceSpecification by ID.
func (c *Client) GetSubstanceSpecificationByID(ctx context.Context, id string, params Parameters) (*models.SubstanceSpecification, error) {
	resp, err := c.GetByID(ctx, "SubstanceSpecification", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSpecification(id, resp)
}

func (c *Client) CreateSubstanceSpecification(ctx context.Context, resource string, params Parameters, entity *models.SubstanceSpecification) (*models.SubstanceSpecification, error) {
	resp, err := c.Create(ctx, "SubstanceSpecification", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSpecification(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceSpecification(ctx context.Context, resource string, params Parameters, entity *models.SubstanceSpecification) (*models.SubstanceSpecification, error) {
	resp, err := c.Update(ctx, "SubstanceSpecification", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSpecification(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSubstanceSpecificationByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstanceSpecification) (*models.SubstanceSpecification, error) {
	resp, err := c.UpdateByID(ctx, "SubstanceSpecification", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSpecification(id, resp)
}

func (c *Client) PatchSubstanceSpecification(ctx context.Context, resource string, params Parameters, entity *models.SubstanceSpecification) ([]*models.SubstanceSpecification, error) {
	resp, err := c.Patch(ctx, "SubstanceSpecification", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSpecifications(resp)
}

func (c *Client) PatchSubstanceSpecificationByID(ctx context.Context, resource, id string, params Parameters, entity *models.SubstanceSpecification) (*models.SubstanceSpecification, error) {
	resp, err := c.PatchByID(ctx, "SubstanceSpecification", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSpecification(id, resp)
}

func (c *Client) DeleteSubstanceSpecification(ctx context.Context, resource string, params Parameters) ([]*models.SubstanceSpecification, error) {
	resp, err := c.Delete(ctx, "SubstanceSpecification", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSpecifications(resp)
}

func (c *Client) DeleteSubstanceSpecificationByID(ctx context.Context, resource, id string, params Parameters) (*models.SubstanceSpecification, error) {
	resp, err := c.DeleteByID(ctx, "SubstanceSpecification", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSubstanceSpecification(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// SupplyDelivery
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSupplyDeliverys(bundle *models.Bundle) ([]*models.SupplyDelivery, error) {
	var entities []*models.SupplyDelivery
	err := EnumBundleResources(bundle, "SupplyDelivery", func(resource ResourceData) error {
		var entity models.SupplyDelivery
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSupplyDeliverys(resp *FhirResponse) ([]*models.SupplyDelivery, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSupplyDeliverys(resp.Bundle)
	case "SupplyDelivery":
		var entity models.SupplyDelivery
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.SupplyDelivery{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSupplyDelivery(id string, resp *FhirResponse) (*models.SupplyDelivery, error) {
	entities, err := fhirRespToSupplyDeliverys(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "SupplyDelivery", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get SupplyDelivery.
func (c *Client) GetSupplyDelivery(ctx context.Context, params Parameters) ([]*models.SupplyDelivery, error) {
	resp, err := c.Get(ctx, "SupplyDelivery", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyDeliverys(resp)
}

// Get SupplyDelivery by ID.
func (c *Client) GetSupplyDeliveryByID(ctx context.Context, id string, params Parameters) (*models.SupplyDelivery, error) {
	resp, err := c.GetByID(ctx, "SupplyDelivery", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyDelivery(id, resp)
}

func (c *Client) CreateSupplyDelivery(ctx context.Context, resource string, params Parameters, entity *models.SupplyDelivery) (*models.SupplyDelivery, error) {
	resp, err := c.Create(ctx, "SupplyDelivery", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyDelivery(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSupplyDelivery(ctx context.Context, resource string, params Parameters, entity *models.SupplyDelivery) (*models.SupplyDelivery, error) {
	resp, err := c.Update(ctx, "SupplyDelivery", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyDelivery(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSupplyDeliveryByID(ctx context.Context, resource, id string, params Parameters, entity *models.SupplyDelivery) (*models.SupplyDelivery, error) {
	resp, err := c.UpdateByID(ctx, "SupplyDelivery", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyDelivery(id, resp)
}

func (c *Client) PatchSupplyDelivery(ctx context.Context, resource string, params Parameters, entity *models.SupplyDelivery) ([]*models.SupplyDelivery, error) {
	resp, err := c.Patch(ctx, "SupplyDelivery", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyDeliverys(resp)
}

func (c *Client) PatchSupplyDeliveryByID(ctx context.Context, resource, id string, params Parameters, entity *models.SupplyDelivery) (*models.SupplyDelivery, error) {
	resp, err := c.PatchByID(ctx, "SupplyDelivery", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyDelivery(id, resp)
}

func (c *Client) DeleteSupplyDelivery(ctx context.Context, resource string, params Parameters) ([]*models.SupplyDelivery, error) {
	resp, err := c.Delete(ctx, "SupplyDelivery", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyDeliverys(resp)
}

func (c *Client) DeleteSupplyDeliveryByID(ctx context.Context, resource, id string, params Parameters) (*models.SupplyDelivery, error) {
	resp, err := c.DeleteByID(ctx, "SupplyDelivery", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyDelivery(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// SupplyRequest
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToSupplyRequests(bundle *models.Bundle) ([]*models.SupplyRequest, error) {
	var entities []*models.SupplyRequest
	err := EnumBundleResources(bundle, "SupplyRequest", func(resource ResourceData) error {
		var entity models.SupplyRequest
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToSupplyRequests(resp *FhirResponse) ([]*models.SupplyRequest, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToSupplyRequests(resp.Bundle)
	case "SupplyRequest":
		var entity models.SupplyRequest
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.SupplyRequest{&entity}, nil
	}
	return nil, nil
}

func fhirRespToSupplyRequest(id string, resp *FhirResponse) (*models.SupplyRequest, error) {
	entities, err := fhirRespToSupplyRequests(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "SupplyRequest", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get SupplyRequest.
func (c *Client) GetSupplyRequest(ctx context.Context, params Parameters) ([]*models.SupplyRequest, error) {
	resp, err := c.Get(ctx, "SupplyRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyRequests(resp)
}

// Get SupplyRequest by ID.
func (c *Client) GetSupplyRequestByID(ctx context.Context, id string, params Parameters) (*models.SupplyRequest, error) {
	resp, err := c.GetByID(ctx, "SupplyRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyRequest(id, resp)
}

func (c *Client) CreateSupplyRequest(ctx context.Context, resource string, params Parameters, entity *models.SupplyRequest) (*models.SupplyRequest, error) {
	resp, err := c.Create(ctx, "SupplyRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSupplyRequest(ctx context.Context, resource string, params Parameters, entity *models.SupplyRequest) (*models.SupplyRequest, error) {
	resp, err := c.Update(ctx, "SupplyRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyRequest(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateSupplyRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.SupplyRequest) (*models.SupplyRequest, error) {
	resp, err := c.UpdateByID(ctx, "SupplyRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyRequest(id, resp)
}

func (c *Client) PatchSupplyRequest(ctx context.Context, resource string, params Parameters, entity *models.SupplyRequest) ([]*models.SupplyRequest, error) {
	resp, err := c.Patch(ctx, "SupplyRequest", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyRequests(resp)
}

func (c *Client) PatchSupplyRequestByID(ctx context.Context, resource, id string, params Parameters, entity *models.SupplyRequest) (*models.SupplyRequest, error) {
	resp, err := c.PatchByID(ctx, "SupplyRequest", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyRequest(id, resp)
}

func (c *Client) DeleteSupplyRequest(ctx context.Context, resource string, params Parameters) ([]*models.SupplyRequest, error) {
	resp, err := c.Delete(ctx, "SupplyRequest", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyRequests(resp)
}

func (c *Client) DeleteSupplyRequestByID(ctx context.Context, resource, id string, params Parameters) (*models.SupplyRequest, error) {
	resp, err := c.DeleteByID(ctx, "SupplyRequest", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToSupplyRequest(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// Task
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToTasks(bundle *models.Bundle) ([]*models.Task, error) {
	var entities []*models.Task
	err := EnumBundleResources(bundle, "Task", func(resource ResourceData) error {
		var entity models.Task
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToTasks(resp *FhirResponse) ([]*models.Task, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToTasks(resp.Bundle)
	case "Task":
		var entity models.Task
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.Task{&entity}, nil
	}
	return nil, nil
}

func fhirRespToTask(id string, resp *FhirResponse) (*models.Task, error) {
	entities, err := fhirRespToTasks(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "Task", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get Task.
func (c *Client) GetTask(ctx context.Context, params Parameters) ([]*models.Task, error) {
	resp, err := c.Get(ctx, "Task", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTasks(resp)
}

// Get Task by ID.
func (c *Client) GetTaskByID(ctx context.Context, id string, params Parameters) (*models.Task, error) {
	resp, err := c.GetByID(ctx, "Task", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTask(id, resp)
}

func (c *Client) CreateTask(ctx context.Context, resource string, params Parameters, entity *models.Task) (*models.Task, error) {
	resp, err := c.Create(ctx, "Task", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTask(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateTask(ctx context.Context, resource string, params Parameters, entity *models.Task) (*models.Task, error) {
	resp, err := c.Update(ctx, "Task", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTask(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateTaskByID(ctx context.Context, resource, id string, params Parameters, entity *models.Task) (*models.Task, error) {
	resp, err := c.UpdateByID(ctx, "Task", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTask(id, resp)
}

func (c *Client) PatchTask(ctx context.Context, resource string, params Parameters, entity *models.Task) ([]*models.Task, error) {
	resp, err := c.Patch(ctx, "Task", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTasks(resp)
}

func (c *Client) PatchTaskByID(ctx context.Context, resource, id string, params Parameters, entity *models.Task) (*models.Task, error) {
	resp, err := c.PatchByID(ctx, "Task", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTask(id, resp)
}

func (c *Client) DeleteTask(ctx context.Context, resource string, params Parameters) ([]*models.Task, error) {
	resp, err := c.Delete(ctx, "Task", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTasks(resp)
}

func (c *Client) DeleteTaskByID(ctx context.Context, resource, id string, params Parameters) (*models.Task, error) {
	resp, err := c.DeleteByID(ctx, "Task", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTask(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// TerminologyCapabilities
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToTerminologyCapabilitiess(bundle *models.Bundle) ([]*models.TerminologyCapabilities, error) {
	var entities []*models.TerminologyCapabilities
	err := EnumBundleResources(bundle, "TerminologyCapabilities", func(resource ResourceData) error {
		var entity models.TerminologyCapabilities
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToTerminologyCapabilitiess(resp *FhirResponse) ([]*models.TerminologyCapabilities, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToTerminologyCapabilitiess(resp.Bundle)
	case "TerminologyCapabilities":
		var entity models.TerminologyCapabilities
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.TerminologyCapabilities{&entity}, nil
	}
	return nil, nil
}

func fhirRespToTerminologyCapabilities(id string, resp *FhirResponse) (*models.TerminologyCapabilities, error) {
	entities, err := fhirRespToTerminologyCapabilitiess(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "TerminologyCapabilities", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get TerminologyCapabilities.
func (c *Client) GetTerminologyCapabilities(ctx context.Context, params Parameters) ([]*models.TerminologyCapabilities, error) {
	resp, err := c.Get(ctx, "TerminologyCapabilities", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTerminologyCapabilitiess(resp)
}

// Get TerminologyCapabilities by ID.
func (c *Client) GetTerminologyCapabilitiesByID(ctx context.Context, id string, params Parameters) (*models.TerminologyCapabilities, error) {
	resp, err := c.GetByID(ctx, "TerminologyCapabilities", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTerminologyCapabilities(id, resp)
}

func (c *Client) CreateTerminologyCapabilities(ctx context.Context, resource string, params Parameters, entity *models.TerminologyCapabilities) (*models.TerminologyCapabilities, error) {
	resp, err := c.Create(ctx, "TerminologyCapabilities", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTerminologyCapabilities(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateTerminologyCapabilities(ctx context.Context, resource string, params Parameters, entity *models.TerminologyCapabilities) (*models.TerminologyCapabilities, error) {
	resp, err := c.Update(ctx, "TerminologyCapabilities", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTerminologyCapabilities(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateTerminologyCapabilitiesByID(ctx context.Context, resource, id string, params Parameters, entity *models.TerminologyCapabilities) (*models.TerminologyCapabilities, error) {
	resp, err := c.UpdateByID(ctx, "TerminologyCapabilities", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTerminologyCapabilities(id, resp)
}

func (c *Client) PatchTerminologyCapabilities(ctx context.Context, resource string, params Parameters, entity *models.TerminologyCapabilities) ([]*models.TerminologyCapabilities, error) {
	resp, err := c.Patch(ctx, "TerminologyCapabilities", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTerminologyCapabilitiess(resp)
}

func (c *Client) PatchTerminologyCapabilitiesByID(ctx context.Context, resource, id string, params Parameters, entity *models.TerminologyCapabilities) (*models.TerminologyCapabilities, error) {
	resp, err := c.PatchByID(ctx, "TerminologyCapabilities", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTerminologyCapabilities(id, resp)
}

func (c *Client) DeleteTerminologyCapabilities(ctx context.Context, resource string, params Parameters) ([]*models.TerminologyCapabilities, error) {
	resp, err := c.Delete(ctx, "TerminologyCapabilities", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTerminologyCapabilitiess(resp)
}

func (c *Client) DeleteTerminologyCapabilitiesByID(ctx context.Context, resource, id string, params Parameters) (*models.TerminologyCapabilities, error) {
	resp, err := c.DeleteByID(ctx, "TerminologyCapabilities", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTerminologyCapabilities(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// TestReport
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToTestReports(bundle *models.Bundle) ([]*models.TestReport, error) {
	var entities []*models.TestReport
	err := EnumBundleResources(bundle, "TestReport", func(resource ResourceData) error {
		var entity models.TestReport
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToTestReports(resp *FhirResponse) ([]*models.TestReport, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToTestReports(resp.Bundle)
	case "TestReport":
		var entity models.TestReport
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.TestReport{&entity}, nil
	}
	return nil, nil
}

func fhirRespToTestReport(id string, resp *FhirResponse) (*models.TestReport, error) {
	entities, err := fhirRespToTestReports(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "TestReport", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get TestReport.
func (c *Client) GetTestReport(ctx context.Context, params Parameters) ([]*models.TestReport, error) {
	resp, err := c.Get(ctx, "TestReport", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestReports(resp)
}

// Get TestReport by ID.
func (c *Client) GetTestReportByID(ctx context.Context, id string, params Parameters) (*models.TestReport, error) {
	resp, err := c.GetByID(ctx, "TestReport", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestReport(id, resp)
}

func (c *Client) CreateTestReport(ctx context.Context, resource string, params Parameters, entity *models.TestReport) (*models.TestReport, error) {
	resp, err := c.Create(ctx, "TestReport", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestReport(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateTestReport(ctx context.Context, resource string, params Parameters, entity *models.TestReport) (*models.TestReport, error) {
	resp, err := c.Update(ctx, "TestReport", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestReport(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateTestReportByID(ctx context.Context, resource, id string, params Parameters, entity *models.TestReport) (*models.TestReport, error) {
	resp, err := c.UpdateByID(ctx, "TestReport", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestReport(id, resp)
}

func (c *Client) PatchTestReport(ctx context.Context, resource string, params Parameters, entity *models.TestReport) ([]*models.TestReport, error) {
	resp, err := c.Patch(ctx, "TestReport", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestReports(resp)
}

func (c *Client) PatchTestReportByID(ctx context.Context, resource, id string, params Parameters, entity *models.TestReport) (*models.TestReport, error) {
	resp, err := c.PatchByID(ctx, "TestReport", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestReport(id, resp)
}

func (c *Client) DeleteTestReport(ctx context.Context, resource string, params Parameters) ([]*models.TestReport, error) {
	resp, err := c.Delete(ctx, "TestReport", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestReports(resp)
}

func (c *Client) DeleteTestReportByID(ctx context.Context, resource, id string, params Parameters) (*models.TestReport, error) {
	resp, err := c.DeleteByID(ctx, "TestReport", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestReport(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// TestScript
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToTestScripts(bundle *models.Bundle) ([]*models.TestScript, error) {
	var entities []*models.TestScript
	err := EnumBundleResources(bundle, "TestScript", func(resource ResourceData) error {
		var entity models.TestScript
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToTestScripts(resp *FhirResponse) ([]*models.TestScript, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToTestScripts(resp.Bundle)
	case "TestScript":
		var entity models.TestScript
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.TestScript{&entity}, nil
	}
	return nil, nil
}

func fhirRespToTestScript(id string, resp *FhirResponse) (*models.TestScript, error) {
	entities, err := fhirRespToTestScripts(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "TestScript", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get TestScript.
func (c *Client) GetTestScript(ctx context.Context, params Parameters) ([]*models.TestScript, error) {
	resp, err := c.Get(ctx, "TestScript", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestScripts(resp)
}

// Get TestScript by ID.
func (c *Client) GetTestScriptByID(ctx context.Context, id string, params Parameters) (*models.TestScript, error) {
	resp, err := c.GetByID(ctx, "TestScript", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestScript(id, resp)
}

func (c *Client) CreateTestScript(ctx context.Context, resource string, params Parameters, entity *models.TestScript) (*models.TestScript, error) {
	resp, err := c.Create(ctx, "TestScript", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestScript(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateTestScript(ctx context.Context, resource string, params Parameters, entity *models.TestScript) (*models.TestScript, error) {
	resp, err := c.Update(ctx, "TestScript", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestScript(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateTestScriptByID(ctx context.Context, resource, id string, params Parameters, entity *models.TestScript) (*models.TestScript, error) {
	resp, err := c.UpdateByID(ctx, "TestScript", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestScript(id, resp)
}

func (c *Client) PatchTestScript(ctx context.Context, resource string, params Parameters, entity *models.TestScript) ([]*models.TestScript, error) {
	resp, err := c.Patch(ctx, "TestScript", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestScripts(resp)
}

func (c *Client) PatchTestScriptByID(ctx context.Context, resource, id string, params Parameters, entity *models.TestScript) (*models.TestScript, error) {
	resp, err := c.PatchByID(ctx, "TestScript", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestScript(id, resp)
}

func (c *Client) DeleteTestScript(ctx context.Context, resource string, params Parameters) ([]*models.TestScript, error) {
	resp, err := c.Delete(ctx, "TestScript", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestScripts(resp)
}

func (c *Client) DeleteTestScriptByID(ctx context.Context, resource, id string, params Parameters) (*models.TestScript, error) {
	resp, err := c.DeleteByID(ctx, "TestScript", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToTestScript(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// ValueSet
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToValueSets(bundle *models.Bundle) ([]*models.ValueSet, error) {
	var entities []*models.ValueSet
	err := EnumBundleResources(bundle, "ValueSet", func(resource ResourceData) error {
		var entity models.ValueSet
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToValueSets(resp *FhirResponse) ([]*models.ValueSet, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToValueSets(resp.Bundle)
	case "ValueSet":
		var entity models.ValueSet
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.ValueSet{&entity}, nil
	}
	return nil, nil
}

func fhirRespToValueSet(id string, resp *FhirResponse) (*models.ValueSet, error) {
	entities, err := fhirRespToValueSets(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "ValueSet", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get ValueSet.
func (c *Client) GetValueSet(ctx context.Context, params Parameters) ([]*models.ValueSet, error) {
	resp, err := c.Get(ctx, "ValueSet", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToValueSets(resp)
}

// Get ValueSet by ID.
func (c *Client) GetValueSetByID(ctx context.Context, id string, params Parameters) (*models.ValueSet, error) {
	resp, err := c.GetByID(ctx, "ValueSet", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToValueSet(id, resp)
}

func (c *Client) CreateValueSet(ctx context.Context, resource string, params Parameters, entity *models.ValueSet) (*models.ValueSet, error) {
	resp, err := c.Create(ctx, "ValueSet", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToValueSet(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateValueSet(ctx context.Context, resource string, params Parameters, entity *models.ValueSet) (*models.ValueSet, error) {
	resp, err := c.Update(ctx, "ValueSet", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToValueSet(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateValueSetByID(ctx context.Context, resource, id string, params Parameters, entity *models.ValueSet) (*models.ValueSet, error) {
	resp, err := c.UpdateByID(ctx, "ValueSet", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToValueSet(id, resp)
}

func (c *Client) PatchValueSet(ctx context.Context, resource string, params Parameters, entity *models.ValueSet) ([]*models.ValueSet, error) {
	resp, err := c.Patch(ctx, "ValueSet", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToValueSets(resp)
}

func (c *Client) PatchValueSetByID(ctx context.Context, resource, id string, params Parameters, entity *models.ValueSet) (*models.ValueSet, error) {
	resp, err := c.PatchByID(ctx, "ValueSet", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToValueSet(id, resp)
}

func (c *Client) DeleteValueSet(ctx context.Context, resource string, params Parameters) ([]*models.ValueSet, error) {
	resp, err := c.Delete(ctx, "ValueSet", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToValueSets(resp)
}

func (c *Client) DeleteValueSetByID(ctx context.Context, resource, id string, params Parameters) (*models.ValueSet, error) {
	resp, err := c.DeleteByID(ctx, "ValueSet", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToValueSet(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// VerificationResult
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToVerificationResults(bundle *models.Bundle) ([]*models.VerificationResult, error) {
	var entities []*models.VerificationResult
	err := EnumBundleResources(bundle, "VerificationResult", func(resource ResourceData) error {
		var entity models.VerificationResult
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToVerificationResults(resp *FhirResponse) ([]*models.VerificationResult, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToVerificationResults(resp.Bundle)
	case "VerificationResult":
		var entity models.VerificationResult
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.VerificationResult{&entity}, nil
	}
	return nil, nil
}

func fhirRespToVerificationResult(id string, resp *FhirResponse) (*models.VerificationResult, error) {
	entities, err := fhirRespToVerificationResults(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "VerificationResult", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get VerificationResult.
func (c *Client) GetVerificationResult(ctx context.Context, params Parameters) ([]*models.VerificationResult, error) {
	resp, err := c.Get(ctx, "VerificationResult", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToVerificationResults(resp)
}

// Get VerificationResult by ID.
func (c *Client) GetVerificationResultByID(ctx context.Context, id string, params Parameters) (*models.VerificationResult, error) {
	resp, err := c.GetByID(ctx, "VerificationResult", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToVerificationResult(id, resp)
}

func (c *Client) CreateVerificationResult(ctx context.Context, resource string, params Parameters, entity *models.VerificationResult) (*models.VerificationResult, error) {
	resp, err := c.Create(ctx, "VerificationResult", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToVerificationResult(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateVerificationResult(ctx context.Context, resource string, params Parameters, entity *models.VerificationResult) (*models.VerificationResult, error) {
	resp, err := c.Update(ctx, "VerificationResult", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToVerificationResult(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateVerificationResultByID(ctx context.Context, resource, id string, params Parameters, entity *models.VerificationResult) (*models.VerificationResult, error) {
	resp, err := c.UpdateByID(ctx, "VerificationResult", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToVerificationResult(id, resp)
}

func (c *Client) PatchVerificationResult(ctx context.Context, resource string, params Parameters, entity *models.VerificationResult) ([]*models.VerificationResult, error) {
	resp, err := c.Patch(ctx, "VerificationResult", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToVerificationResults(resp)
}

func (c *Client) PatchVerificationResultByID(ctx context.Context, resource, id string, params Parameters, entity *models.VerificationResult) (*models.VerificationResult, error) {
	resp, err := c.PatchByID(ctx, "VerificationResult", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToVerificationResult(id, resp)
}

func (c *Client) DeleteVerificationResult(ctx context.Context, resource string, params Parameters) ([]*models.VerificationResult, error) {
	resp, err := c.Delete(ctx, "VerificationResult", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToVerificationResults(resp)
}

func (c *Client) DeleteVerificationResultByID(ctx context.Context, resource, id string, params Parameters) (*models.VerificationResult, error) {
	resp, err := c.DeleteByID(ctx, "VerificationResult", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToVerificationResult(id, resp)
}

// ---------------------------------------------------------------------------------------------------------------------------
// VisionPrescription
// ---------------------------------------------------------------------------------------------------------------------------

func bundleToVisionPrescriptions(bundle *models.Bundle) ([]*models.VisionPrescription, error) {
	var entities []*models.VisionPrescription
	err := EnumBundleResources(bundle, "VisionPrescription", func(resource ResourceData) error {
		var entity models.VisionPrescription
		if err := resource.UnmarshalTo(&entity); err != nil {
			return err
		}
		entities = append(entities, &entity)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func fhirRespToVisionPrescriptions(resp *FhirResponse) ([]*models.VisionPrescription, error) {
	switch resp.ResourceType {
	case "Bundle":
		return bundleToVisionPrescriptions(resp.Bundle)
	case "VisionPrescription":
		var entity models.VisionPrescription
		if err := json.Unmarshal(resp.Body, &entity); err != nil {
			return nil, err
		}
		return []*models.VisionPrescription{&entity}, nil
	}
	return nil, nil
}

func fhirRespToVisionPrescription(id string, resp *FhirResponse) (*models.VisionPrescription, error) {
	entities, err := fhirRespToVisionPrescriptions(resp)
	if err != nil {
		return nil, err
	}
	if err := checkSingular(id, "VisionPrescription", len(entities)); err != nil {
		return nil, err
	}
	return entities[0], nil
}

// Get VisionPrescription.
func (c *Client) GetVisionPrescription(ctx context.Context, params Parameters) ([]*models.VisionPrescription, error) {
	resp, err := c.Get(ctx, "VisionPrescription", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToVisionPrescriptions(resp)
}

// Get VisionPrescription by ID.
func (c *Client) GetVisionPrescriptionByID(ctx context.Context, id string, params Parameters) (*models.VisionPrescription, error) {
	resp, err := c.GetByID(ctx, "VisionPrescription", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToVisionPrescription(id, resp)
}

func (c *Client) CreateVisionPrescription(ctx context.Context, resource string, params Parameters, entity *models.VisionPrescription) (*models.VisionPrescription, error) {
	resp, err := c.Create(ctx, "VisionPrescription", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToVisionPrescription(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateVisionPrescription(ctx context.Context, resource string, params Parameters, entity *models.VisionPrescription) (*models.VisionPrescription, error) {
	resp, err := c.Update(ctx, "VisionPrescription", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToVisionPrescription(StrPtrToStr(entity.ID), resp)
}

func (c *Client) UpdateVisionPrescriptionByID(ctx context.Context, resource, id string, params Parameters, entity *models.VisionPrescription) (*models.VisionPrescription, error) {
	resp, err := c.UpdateByID(ctx, "VisionPrescription", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToVisionPrescription(id, resp)
}

func (c *Client) PatchVisionPrescription(ctx context.Context, resource string, params Parameters, entity *models.VisionPrescription) ([]*models.VisionPrescription, error) {
	resp, err := c.Patch(ctx, "VisionPrescription", params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToVisionPrescriptions(resp)
}

func (c *Client) PatchVisionPrescriptionByID(ctx context.Context, resource, id string, params Parameters, entity *models.VisionPrescription) (*models.VisionPrescription, error) {
	resp, err := c.PatchByID(ctx, "VisionPrescription", id, params, entity)
	if err != nil {
		return nil, err
	}

	return fhirRespToVisionPrescription(id, resp)
}

func (c *Client) DeleteVisionPrescription(ctx context.Context, resource string, params Parameters) ([]*models.VisionPrescription, error) {
	resp, err := c.Delete(ctx, "VisionPrescription", params)
	if err != nil {
		return nil, err
	}

	return fhirRespToVisionPrescriptions(resp)
}

func (c *Client) DeleteVisionPrescriptionByID(ctx context.Context, resource, id string, params Parameters) (*models.VisionPrescription, error) {
	resp, err := c.DeleteByID(ctx, "VisionPrescription", id, params)
	if err != nil {
		return nil, err
	}

	return fhirRespToVisionPrescription(id, resp)
}
