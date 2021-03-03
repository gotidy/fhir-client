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

// Code generated by go generate; DO NOT EDIT.

// This file was generated by robots at
// 2021-03-03 10:55:45.102676 +0000 UTC
// HumanName is documented here http://hl7.org/fhir/StructureDefinition/HumanName
type HumanName struct {
	ID        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Use       *NameUse    `bson:"use,omitempty" json:"use,omitempty"`
	Text      *string     `bson:"text,omitempty" json:"text,omitempty"`
	Family    *string     `bson:"family,omitempty" json:"family,omitempty"`
	Given     []string    `bson:"given,omitempty" json:"given,omitempty"`
	Prefix    []string    `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Suffix    []string    `bson:"suffix,omitempty" json:"suffix,omitempty"`
	Period    *Period     `bson:"period,omitempty" json:"period,omitempty"`
}
