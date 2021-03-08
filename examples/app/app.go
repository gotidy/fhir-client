package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/gotidy/fhir-client"
	"github.com/gotidy/fhir-client/models"
	"github.com/gotidy/ptr"
	"github.com/rs/zerolog"
)

type Application struct {
	log    zerolog.Logger
	client *fhir.Client
}

func NewApplication(log zerolog.Logger, client *fhir.Client) *Application {
	return &Application{log: log, client: client}
}

func (a *Application) Run(ctx context.Context) error {
	john := &models.Patient{
		ID: ptr.String(uuid.NewString()),
		Name: []models.HumanName{
			{
				Given:  []string{"John", "Joshua"},
				Family: ptr.String("Jonson"),
			},
		},
	}

	jane := &models.Patient{
		ID: ptr.String(uuid.NewString()),
		Name: []models.HumanName{
			{
				Given:  []string{"Jane", "Jan"},
				Family: ptr.String("Jack"),
			},
		},
	}

	// Create John
	fmt.Println("### CreatePatient ###")
	patient, err := a.client.CreatePatient(ctx, nil, john)
	if err != nil {
		return err
	}
	fmt.Println(format(patient))

	// Create Jane
	fmt.Println("### CreatePatient ###")
	patient, err = a.client.CreatePatient(ctx, nil, jane)
	if err != nil {
		return err
	}
	fmt.Println(format(patient))

	// Get All
	fmt.Println("### GetPatient ###")
	patients, err := a.client.GetPatient(ctx, nil)
	if err != nil {
		return err
	}
	fmt.Println("Resources count:", len(patients))
	fmt.Println(format(patients))

	// Get Search
	fmt.Println("### GetPatient with params ###")
	params := url.Values{}
	params.Add("id", *(john.ID))
	patients, err = a.client.GetPatient(ctx, params)
	if err != nil {
		return err
	}
	fmt.Println("Resources count:", len(patients))
	fmt.Println(format(patients))

	// Update
	fmt.Println("### UpdatePatient ###")
	john.BirthDate = models.NewDate(1987, 1, 12)
	patient, err = a.client.UpdatePatientByID(ctx, *(john.ID), nil, john)
	if err != nil {
		if e, ok := fhir.AsUnmarshalError(err); ok {
			fmt.Println(string(e.Data))
		}
		return err
	}
	fmt.Println(format(patient))

	// GetByID
	fmt.Println("### GetPatientByID ###")
	patient, err = a.client.GetPatientByID(ctx, *(john.ID), nil)
	if err != nil {
		return err
	}
	fmt.Println("Getting resource with ID:", *(john.ID))
	fmt.Println(format(patient))

	// GetByID wrong ID
	fmt.Println("### GetPatientByID ###")
	patient, err = a.client.GetPatientByID(ctx, "wrong_id", nil)
	fmt.Println(err)
	fmt.Println(format(patient))

	// PatchPatientByID
	fmt.Println("### PatchPatientByID ###")
	john.BirthDate = models.NewDate(1980, 8, 10)
	patient, err = a.client.PatchPatientByID(ctx, *(john.ID), nil, john)
	if err != nil {
		return err
	}
	fmt.Println(format(patient))

	// PatchPatient // not supported
	// fmt.Println("### PatchPatient ###")
	// john.BirthDate = ptr.String("1980-04-30")
	// params = url.Values{}
	// params.Add("_id", *(john.ID))
	// patients, err = a.client.PatchPatient(ctx, params, john)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(format(patients))

	// DeletePatient John
	fmt.Println("### DeletePatient ###")
	john.BirthDate = models.NewDate(1980, 4, 30)
	params = url.Values{}
	params.Add("id", *(john.ID))
	patients, err = a.client.DeletePatient(ctx, params)
	if err != nil {
		return err
	}
	fmt.Println(format(patients))

	// DeletePatientByID Jane
	fmt.Println("### DeletePatientByID ###")
	patient, err = a.client.DeletePatientByID(ctx, *(jane.ID), nil)
	if err != nil {
		return err
	}
	fmt.Println(format(patient))

	// GetByID deleted Patient
	fmt.Println("### GetPatientByID ###")
	patient, err = a.client.GetPatientByID(ctx, *(john.ID), nil)
	fmt.Println(err)
	fmt.Println(format(patient))

	return nil
}
