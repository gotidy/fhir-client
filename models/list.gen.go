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
// 2021-03-08 09:41:02.937838 +0000 UTC

package models

import "encoding/json"

// List is documented here http://hl7.org/fhir/StructureDefinition/List
type List struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string          `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative       `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            ListStatus       `bson:"status" json:"status"`
	Mode              ListMode         `bson:"mode" json:"mode"`
	Title             *string          `bson:"title,omitempty" json:"title,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Subject           *Reference       `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date              *DateTime        `bson:"date,omitempty" json:"date,omitempty"`
	Source            *Reference       `bson:"source,omitempty" json:"source,omitempty"`
	OrderedBy         *CodeableConcept `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Note              []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
	Entry             []ListEntry      `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason       *CodeableConcept `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
}
type ListEntry struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Flag              *CodeableConcept `bson:"flag,omitempty" json:"flag,omitempty"`
	Deleted           *bool            `bson:"deleted,omitempty" json:"deleted,omitempty"`
	Date              *DateTime        `bson:"date,omitempty" json:"date,omitempty"`
	Item              Reference        `bson:"item" json:"item"`
}
type OtherList List

// MarshalJSON marshals the given List as JSON into a byte slice
func (r List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherList
		ResourceType string `json:"resourceType"`
	}{
		OtherList:    OtherList(r),
		ResourceType: "List",
	})
}

// UnmarshalList unmarshals a List.
func UnmarshalList(b []byte) (List, error) {
	var list List
	if err := json.Unmarshal(b, &list); err != nil {
		return list, err
	}
	return list, nil
}
