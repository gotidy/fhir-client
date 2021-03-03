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
// 2021-03-03 10:55:44.931042 +0000 UTC
// Composition is documented here http://hl7.org/fhir/StructureDefinition/Composition
type Composition struct {
	ID                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            CompositionStatus      `bson:"status" json:"status"`
	Type              CodeableConcept        `bson:"type" json:"type"`
	Category          []CodeableConcept      `bson:"category,omitempty" json:"category,omitempty"`
	Subject           *Reference             `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date              string                 `bson:"date" json:"date"`
	Author            []Reference            `bson:"author" json:"author"`
	Title             string                 `bson:"title" json:"title"`
	Confidentiality   *string                `bson:"confidentiality,omitempty" json:"confidentiality,omitempty"`
	Attester          []CompositionAttester  `bson:"attester,omitempty" json:"attester,omitempty"`
	Custodian         *Reference             `bson:"custodian,omitempty" json:"custodian,omitempty"`
	RelatesTo         []CompositionRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Event             []CompositionEvent     `bson:"event,omitempty" json:"event,omitempty"`
	Section           []CompositionSection   `bson:"section,omitempty" json:"section,omitempty"`
}
type CompositionAttester struct {
	ID                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              CompositionAttestationMode `bson:"mode" json:"mode"`
	Time              *string                    `bson:"time,omitempty" json:"time,omitempty"`
	Party             *Reference                 `bson:"party,omitempty" json:"party,omitempty"`
}
type CompositionRelatesTo struct {
	ID                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              DocumentRelationshipType `bson:"code" json:"code"`
}
type CompositionEvent struct {
	ID                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Detail            []Reference       `bson:"detail,omitempty" json:"detail,omitempty"`
}
type CompositionSection struct {
	ID                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Title             *string              `bson:"title,omitempty" json:"title,omitempty"`
	Code              *CodeableConcept     `bson:"code,omitempty" json:"code,omitempty"`
	Author            []Reference          `bson:"author,omitempty" json:"author,omitempty"`
	Focus             *Reference           `bson:"focus,omitempty" json:"focus,omitempty"`
	Text              *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Mode              *ListMode            `bson:"mode,omitempty" json:"mode,omitempty"`
	OrderedBy         *CodeableConcept     `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Entry             []Reference          `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason       *CodeableConcept     `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
	Section           []CompositionSection `bson:"section,omitempty" json:"section,omitempty"`
}
type OtherComposition Composition

// MarshalJSON marshals the given Composition as JSON into a byte slice
func (r Composition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherComposition
		ResourceType string `json:"resourceType"`
	}{
		OtherComposition: OtherComposition(r),
		ResourceType:     "Composition",
	})
}

// UnmarshalComposition unmarshals a Composition.
func UnmarshalComposition(b []byte) (Composition, error) {
	var composition Composition
	if err := json.Unmarshal(b, &composition); err != nil {
		return composition, err
	}
	return composition, nil
}
