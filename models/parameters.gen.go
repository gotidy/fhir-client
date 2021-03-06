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
// 2021-03-08 09:41:02.788699 +0000 UTC

package models

import "encoding/json"

// Parameters is documented here http://hl7.org/fhir/StructureDefinition/Parameters
type Parameters struct {
	ID            *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *string               `bson:"language,omitempty" json:"language,omitempty"`
	Parameter     []ParametersParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
}
type ParametersParameter struct {
	ID                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                `bson:"name" json:"name"`
	Resource          json.RawMessage       `bson:"resource,omitempty" json:"resource,omitempty"`
	Part              []ParametersParameter `bson:"part,omitempty" json:"part,omitempty"`
}
type OtherParameters Parameters

// MarshalJSON marshals the given Parameters as JSON into a byte slice
func (r Parameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherParameters
		ResourceType string `json:"resourceType"`
	}{
		OtherParameters: OtherParameters(r),
		ResourceType:    "Parameters",
	})
}

// UnmarshalParameters unmarshals a Parameters.
func UnmarshalParameters(b []byte) (Parameters, error) {
	var parameters Parameters
	if err := json.Unmarshal(b, &parameters); err != nil {
		return parameters, err
	}
	return parameters, nil
}
