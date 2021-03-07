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
// 2021-03-07 20:19:39.705096 +0000 UTC

package models

// Duration is documented here http://hl7.org/fhir/StructureDefinition/Duration
type Duration struct {
	ID         *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension  []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	Value      *string             `bson:"value,omitempty" json:"value,omitempty"`
	Comparator *QuantityComparator `bson:"comparator,omitempty" json:"comparator,omitempty"`
	Unit       *string             `bson:"unit,omitempty" json:"unit,omitempty"`
	System     *string             `bson:"system,omitempty" json:"system,omitempty"`
	Code       *string             `bson:"code,omitempty" json:"code,omitempty"`
}
