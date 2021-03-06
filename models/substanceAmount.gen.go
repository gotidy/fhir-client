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
// 2021-03-08 09:41:02.97128 +0000 UTC

package models

// SubstanceAmount is documented here http://hl7.org/fhir/StructureDefinition/SubstanceAmount
type SubstanceAmount struct {
	ID                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	AmountType        *CodeableConcept               `bson:"amountType,omitempty" json:"amountType,omitempty"`
	AmountText        *string                        `bson:"amountText,omitempty" json:"amountText,omitempty"`
	ReferenceRange    *SubstanceAmountReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}
type SubstanceAmountReferenceRange struct {
	ID        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	LowLimit  *Quantity   `bson:"lowLimit,omitempty" json:"lowLimit,omitempty"`
	HighLimit *Quantity   `bson:"highLimit,omitempty" json:"highLimit,omitempty"`
}
