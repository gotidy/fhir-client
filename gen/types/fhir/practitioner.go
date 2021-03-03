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
// 2021-03-03 10:55:38.953853 +0000 UTC
// Practitioner is documented here http://hl7.org/fhir/StructureDefinition/Practitioner
type Practitioner struct {
	ID                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                       `bson:"active,omitempty" json:"active,omitempty"`
	Name              []HumanName                 `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint              `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           []Address                   `bson:"address,omitempty" json:"address,omitempty"`
	Gender            *AdministrativeGender       `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate         *string                     `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Photo             []Attachment                `bson:"photo,omitempty" json:"photo,omitempty"`
	Qualification     []PractitionerQualification `bson:"qualification,omitempty" json:"qualification,omitempty"`
	Communication     []CodeableConcept           `bson:"communication,omitempty" json:"communication,omitempty"`
}
type PractitionerQualification struct {
	ID                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Period            *Period         `bson:"period,omitempty" json:"period,omitempty"`
	Issuer            *Reference      `bson:"issuer,omitempty" json:"issuer,omitempty"`
}
type OtherPractitioner Practitioner

// MarshalJSON marshals the given Practitioner as JSON into a byte slice
func (r Practitioner) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitioner
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitioner: OtherPractitioner(r),
		ResourceType:      "Practitioner",
	})
}

// UnmarshalPractitioner unmarshals a Practitioner.
func UnmarshalPractitioner(b []byte) (Practitioner, error) {
	var practitioner Practitioner
	if err := json.Unmarshal(b, &practitioner); err != nil {
		return practitioner, err
	}
	return practitioner, nil
}
