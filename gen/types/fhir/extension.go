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
// 2021-03-03 10:55:39.318431 +0000 UTC
// Extension is documented here http://hl7.org/fhir/StructureDefinition/Extension
type Extension struct {
	ID        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Url       string      `bson:"url" json:"url"`
}
