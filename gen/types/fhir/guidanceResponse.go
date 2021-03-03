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
// 2021-03-03 10:55:38.473081 +0000 UTC
// GuidanceResponse is documented here http://hl7.org/fhir/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
	ID                 *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RequestIdentifier  *Identifier            `bson:"requestIdentifier,omitempty" json:"requestIdentifier,omitempty"`
	Identifier         []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status             GuidanceResponseStatus `bson:"status" json:"status"`
	Subject            *Reference             `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter          *Reference             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	OccurrenceDateTime *string                `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	Performer          *Reference             `bson:"performer,omitempty" json:"performer,omitempty"`
	ReasonCode         []CodeableConcept      `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note               []Annotation           `bson:"note,omitempty" json:"note,omitempty"`
	EvaluationMessage  []Reference            `bson:"evaluationMessage,omitempty" json:"evaluationMessage,omitempty"`
	OutputParameters   *Reference             `bson:"outputParameters,omitempty" json:"outputParameters,omitempty"`
	Result             *Reference             `bson:"result,omitempty" json:"result,omitempty"`
	DataRequirement    []DataRequirement      `bson:"dataRequirement,omitempty" json:"dataRequirement,omitempty"`
}
type OtherGuidanceResponse GuidanceResponse

// MarshalJSON marshals the given GuidanceResponse as JSON into a byte slice
func (r GuidanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGuidanceResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherGuidanceResponse: OtherGuidanceResponse(r),
		ResourceType:          "GuidanceResponse",
	})
}

// UnmarshalGuidanceResponse unmarshals a GuidanceResponse.
func UnmarshalGuidanceResponse(b []byte) (GuidanceResponse, error) {
	var guidanceResponse GuidanceResponse
	if err := json.Unmarshal(b, &guidanceResponse); err != nil {
		return guidanceResponse, err
	}
	return guidanceResponse, nil
}
