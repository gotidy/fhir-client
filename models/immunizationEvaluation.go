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
// 2021-03-03 10:55:43.982994 +0000 UTC
// ImmunizationEvaluation is documented here http://hl7.org/fhir/StructureDefinition/ImmunizationEvaluation
type ImmunizationEvaluation struct {
	ID                *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            ImmunizationEvaluationStatusCodes `bson:"status" json:"status"`
	Patient           Reference                         `bson:"patient" json:"patient"`
	Date              *string                           `bson:"date,omitempty" json:"date,omitempty"`
	Authority         *Reference                        `bson:"authority,omitempty" json:"authority,omitempty"`
	TargetDisease     CodeableConcept                   `bson:"targetDisease" json:"targetDisease"`
	ImmunizationEvent Reference                         `bson:"immunizationEvent" json:"immunizationEvent"`
	DoseStatus        CodeableConcept                   `bson:"doseStatus" json:"doseStatus"`
	DoseStatusReason  []CodeableConcept                 `bson:"doseStatusReason,omitempty" json:"doseStatusReason,omitempty"`
	Description       *string                           `bson:"description,omitempty" json:"description,omitempty"`
	Series            *string                           `bson:"series,omitempty" json:"series,omitempty"`
}
type OtherImmunizationEvaluation ImmunizationEvaluation

// MarshalJSON marshals the given ImmunizationEvaluation as JSON into a byte slice
func (r ImmunizationEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunizationEvaluation
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunizationEvaluation: OtherImmunizationEvaluation(r),
		ResourceType:                "ImmunizationEvaluation",
	})
}

// UnmarshalImmunizationEvaluation unmarshals a ImmunizationEvaluation.
func UnmarshalImmunizationEvaluation(b []byte) (ImmunizationEvaluation, error) {
	var immunizationEvaluation ImmunizationEvaluation
	if err := json.Unmarshal(b, &immunizationEvaluation); err != nil {
		return immunizationEvaluation, err
	}
	return immunizationEvaluation, nil
}
