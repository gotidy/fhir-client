// Copyright 2021
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
//
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2021-03-08 09:41:02.543314 +0000 UTC

package models

import "encoding/json"

// NamingSystem is documented here http://hl7.org/fhir/StructureDefinition/NamingSystem
type NamingSystem struct {
	ID                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                 `bson:"name" json:"name"`
	Status            PublicationStatus      `bson:"status" json:"status"`
	Kind              NamingSystemType       `bson:"kind" json:"kind"`
	Date              DateTime               `bson:"date" json:"date"`
	Publisher         *string                `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail        `bson:"contact,omitempty" json:"contact,omitempty"`
	Responsible       *string                `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Type              *CodeableConcept       `bson:"type,omitempty" json:"type,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext         `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept      `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Usage             *string                `bson:"usage,omitempty" json:"usage,omitempty"`
	UniqueId          []NamingSystemUniqueId `bson:"uniqueId" json:"uniqueId"`
}
type NamingSystemUniqueId struct {
	ID                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              NamingSystemIdentifierType `bson:"type" json:"type"`
	Value             string                     `bson:"value" json:"value"`
	Preferred         *bool                      `bson:"preferred,omitempty" json:"preferred,omitempty"`
	Comment           *string                    `bson:"comment,omitempty" json:"comment,omitempty"`
	Period            *Period                    `bson:"period,omitempty" json:"period,omitempty"`
}
type OtherNamingSystem NamingSystem

// MarshalJSON marshals the given NamingSystem as JSON into a byte slice
func (r NamingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNamingSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherNamingSystem: OtherNamingSystem(r),
		ResourceType:      "NamingSystem",
	})
}

// UnmarshalNamingSystem unmarshals a NamingSystem.
func UnmarshalNamingSystem(b []byte) (NamingSystem, error) {
	var namingSystem NamingSystem
	if err := json.Unmarshal(b, &namingSystem); err != nil {
		return namingSystem, err
	}
	return namingSystem, nil
}
