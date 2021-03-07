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
// 2021-03-07 20:19:38.297009 +0000 UTC

package models

import "encoding/json"

// EnrollmentRequest is documented here http://hl7.org/fhir/StructureDefinition/EnrollmentRequest
type EnrollmentRequest struct {
	ID                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                       `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *FinancialResourceStatusCodes `bson:"status,omitempty" json:"status,omitempty"`
	Created           *DateTime                     `bson:"created,omitempty" json:"created,omitempty"`
	Insurer           *Reference                    `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider          *Reference                    `bson:"provider,omitempty" json:"provider,omitempty"`
	Candidate         *Reference                    `bson:"candidate,omitempty" json:"candidate,omitempty"`
	Coverage          *Reference                    `bson:"coverage,omitempty" json:"coverage,omitempty"`
}
type OtherEnrollmentRequest EnrollmentRequest

// MarshalJSON marshals the given EnrollmentRequest as JSON into a byte slice
func (r EnrollmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentRequest: OtherEnrollmentRequest(r),
		ResourceType:           "EnrollmentRequest",
	})
}

// UnmarshalEnrollmentRequest unmarshals a EnrollmentRequest.
func UnmarshalEnrollmentRequest(b []byte) (EnrollmentRequest, error) {
	var enrollmentRequest EnrollmentRequest
	if err := json.Unmarshal(b, &enrollmentRequest); err != nil {
		return enrollmentRequest, err
	}
	return enrollmentRequest, nil
}
