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
// 2021-03-03 10:55:38.504217 +0000 UTC
// BiologicallyDerivedProduct is documented here http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProduct struct {
	ID                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                              `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ProductCategory   *BiologicallyDerivedProductCategory     `bson:"productCategory,omitempty" json:"productCategory,omitempty"`
	ProductCode       *CodeableConcept                        `bson:"productCode,omitempty" json:"productCode,omitempty"`
	Status            *BiologicallyDerivedProductStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Request           []Reference                             `bson:"request,omitempty" json:"request,omitempty"`
	Quantity          *int                                    `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Parent            []Reference                             `bson:"parent,omitempty" json:"parent,omitempty"`
	Collection        *BiologicallyDerivedProductCollection   `bson:"collection,omitempty" json:"collection,omitempty"`
	Processing        []BiologicallyDerivedProductProcessing  `bson:"processing,omitempty" json:"processing,omitempty"`
	Manipulation      *BiologicallyDerivedProductManipulation `bson:"manipulation,omitempty" json:"manipulation,omitempty"`
	Storage           []BiologicallyDerivedProductStorage     `bson:"storage,omitempty" json:"storage,omitempty"`
}
type BiologicallyDerivedProductCollection struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Collector         *Reference  `bson:"collector,omitempty" json:"collector,omitempty"`
	Source            *Reference  `bson:"source,omitempty" json:"source,omitempty"`
}
type BiologicallyDerivedProductProcessing struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
	Procedure         *CodeableConcept `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Additive          *Reference       `bson:"additive,omitempty" json:"additive,omitempty"`
}
type BiologicallyDerivedProductManipulation struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
}
type BiologicallyDerivedProductStorage struct {
	ID                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Temperature       *string                                 `bson:"temperature,omitempty" json:"temperature,omitempty"`
	Scale             *BiologicallyDerivedProductStorageScale `bson:"scale,omitempty" json:"scale,omitempty"`
	Duration          *Period                                 `bson:"duration,omitempty" json:"duration,omitempty"`
}
type OtherBiologicallyDerivedProduct BiologicallyDerivedProduct

// MarshalJSON marshals the given BiologicallyDerivedProduct as JSON into a byte slice
func (r BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBiologicallyDerivedProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherBiologicallyDerivedProduct: OtherBiologicallyDerivedProduct(r),
		ResourceType:                    "BiologicallyDerivedProduct",
	})
}

// UnmarshalBiologicallyDerivedProduct unmarshals a BiologicallyDerivedProduct.
func UnmarshalBiologicallyDerivedProduct(b []byte) (BiologicallyDerivedProduct, error) {
	var biologicallyDerivedProduct BiologicallyDerivedProduct
	if err := json.Unmarshal(b, &biologicallyDerivedProduct); err != nil {
		return biologicallyDerivedProduct, err
	}
	return biologicallyDerivedProduct, nil
}
