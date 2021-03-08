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
// 2021-03-08 09:41:02.983986 +0000 UTC

package models

// Attachment is documented here http://hl7.org/fhir/StructureDefinition/Attachment
type Attachment struct {
	ID          *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ContentType *string     `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Language    *string     `bson:"language,omitempty" json:"language,omitempty"`
	Data        *string     `bson:"data,omitempty" json:"data,omitempty"`
	URL         *string     `bson:"url,omitempty" json:"url,omitempty"`
	Size        *int        `bson:"size,omitempty" json:"size,omitempty"`
	Hash        *string     `bson:"hash,omitempty" json:"hash,omitempty"`
	Title       *string     `bson:"title,omitempty" json:"title,omitempty"`
	Creation    *DateTime   `bson:"creation,omitempty" json:"creation,omitempty"`
}
