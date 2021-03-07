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
// 2021-03-07 20:19:38.717755 +0000 UTC

package models

import "encoding/json"

// DeviceRequest is documented here http://hl7.org/fhir/StructureDefinition/DeviceRequest
type DeviceRequest struct {
	ID                    *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string                 `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                 `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference              `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PriorRequest          []Reference              `bson:"priorRequest,omitempty" json:"priorRequest,omitempty"`
	GroupIdentifier       *Identifier              `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status                *RequestStatus           `bson:"status,omitempty" json:"status,omitempty"`
	Intent                RequestIntent            `bson:"intent" json:"intent"`
	Priority              *RequestPriority         `bson:"priority,omitempty" json:"priority,omitempty"`
	Parameter             []DeviceRequestParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Subject               Reference                `bson:"subject" json:"subject"`
	Encounter             *Reference               `bson:"encounter,omitempty" json:"encounter,omitempty"`
	AuthoredOn            *DateTime                `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester             *Reference               `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType         *CodeableConcept         `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Performer             *Reference               `bson:"performer,omitempty" json:"performer,omitempty"`
	ReasonCode            []CodeableConcept        `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference              `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Insurance             []Reference              `bson:"insurance,omitempty" json:"insurance,omitempty"`
	SupportingInfo        []Reference              `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Note                  []Annotation             `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory       []Reference              `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}
type DeviceRequestParameter struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}
type OtherDeviceRequest DeviceRequest

// MarshalJSON marshals the given DeviceRequest as JSON into a byte slice
func (r DeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceRequest: OtherDeviceRequest(r),
		ResourceType:       "DeviceRequest",
	})
}

// UnmarshalDeviceRequest unmarshals a DeviceRequest.
func UnmarshalDeviceRequest(b []byte) (DeviceRequest, error) {
	var deviceRequest DeviceRequest
	if err := json.Unmarshal(b, &deviceRequest); err != nil {
		return deviceRequest, err
	}
	return deviceRequest, nil
}
