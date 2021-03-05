package main

import (
	"context"
	"fmt"

	"github.com/gotidy/fhir-client"
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
	// Get
	patients, err := a.client.GetPatient(ctx, nil)
	if err != nil {
		return err
	}
	fmt.Println("### GetPatient ###")
	fmt.Println("Resources count:", len(patients))
	fmt.Println(format(patients))

	// GetByID
	patient, err := a.client.GetPatientByID(ctx, *(patients[0].ID), nil)
	if err != nil {
		return err
	}
	fmt.Println("### GetPatientByID ###")
	fmt.Println("Getting resource with ID:", *(patients[0].ID))
	fmt.Println(format(patient))

	// GetByID wrong ID
	patient, err = a.client.GetPatientByID(ctx, "wrong_id", nil)
	fmt.Println("### GetPatientByID ###")
	fmt.Println(err)
	fmt.Println(format(patient))

	return nil
}
