// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import "encoding/json"

// Code generated by go generate; DO NOT EDIT.

// This file was generated by robots at
// 2021-03-03 10:55:44.0354 +0000 UTC
// Immunization is documented here http://hl7.org/fhir/StructureDefinition/Immunization
type Immunization struct {
	ID                 *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                       `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status             ImmunizationStatusCodes       `bson:"status" json:"status"`
	StatusReason       *CodeableConcept              `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	VaccineCode        CodeableConcept               `bson:"vaccineCode" json:"vaccineCode"`
	Patient            Reference                     `bson:"patient" json:"patient"`
	Encounter          *Reference                    `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Recorded           *string                       `bson:"recorded,omitempty" json:"recorded,omitempty"`
	PrimarySource      *bool                         `bson:"primarySource,omitempty" json:"primarySource,omitempty"`
	ReportOrigin       *CodeableConcept              `bson:"reportOrigin,omitempty" json:"reportOrigin,omitempty"`
	Location           *Reference                    `bson:"location,omitempty" json:"location,omitempty"`
	Manufacturer       *Reference                    `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	LotNumber          *string                       `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	ExpirationDate     *string                       `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	Site               *CodeableConcept              `bson:"site,omitempty" json:"site,omitempty"`
	Route              *CodeableConcept              `bson:"route,omitempty" json:"route,omitempty"`
	DoseQuantity       *Quantity                     `bson:"doseQuantity,omitempty" json:"doseQuantity,omitempty"`
	Performer          []ImmunizationPerformer       `bson:"performer,omitempty" json:"performer,omitempty"`
	Note               []Annotation                  `bson:"note,omitempty" json:"note,omitempty"`
	ReasonCode         []CodeableConcept             `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference                   `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	IsSubpotent        *bool                         `bson:"isSubpotent,omitempty" json:"isSubpotent,omitempty"`
	SubpotentReason    []CodeableConcept             `bson:"subpotentReason,omitempty" json:"subpotentReason,omitempty"`
	Education          []ImmunizationEducation       `bson:"education,omitempty" json:"education,omitempty"`
	ProgramEligibility []CodeableConcept             `bson:"programEligibility,omitempty" json:"programEligibility,omitempty"`
	FundingSource      *CodeableConcept              `bson:"fundingSource,omitempty" json:"fundingSource,omitempty"`
	Reaction           []ImmunizationReaction        `bson:"reaction,omitempty" json:"reaction,omitempty"`
	ProtocolApplied    []ImmunizationProtocolApplied `bson:"protocolApplied,omitempty" json:"protocolApplied,omitempty"`
}
type ImmunizationPerformer struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}
type ImmunizationEducation struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DocumentType      *string     `bson:"documentType,omitempty" json:"documentType,omitempty"`
	Reference         *string     `bson:"reference,omitempty" json:"reference,omitempty"`
	PublicationDate   *string     `bson:"publicationDate,omitempty" json:"publicationDate,omitempty"`
	PresentationDate  *string     `bson:"presentationDate,omitempty" json:"presentationDate,omitempty"`
}
type ImmunizationReaction struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date              *string     `bson:"date,omitempty" json:"date,omitempty"`
	Detail            *Reference  `bson:"detail,omitempty" json:"detail,omitempty"`
	Reported          *bool       `bson:"reported,omitempty" json:"reported,omitempty"`
}
type ImmunizationProtocolApplied struct {
	ID                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Series            *string           `bson:"series,omitempty" json:"series,omitempty"`
	Authority         *Reference        `bson:"authority,omitempty" json:"authority,omitempty"`
	TargetDisease     []CodeableConcept `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
}
type OtherImmunization Immunization

// MarshalJSON marshals the given Immunization as JSON into a byte slice
func (r Immunization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunization
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunization: OtherImmunization(r),
		ResourceType:      "Immunization",
	})
}

// UnmarshalImmunization unmarshals a Immunization.
func UnmarshalImmunization(b []byte) (Immunization, error) {
	var immunization Immunization
	if err := json.Unmarshal(b, &immunization); err != nil {
		return immunization, err
	}
	return immunization, nil
}
