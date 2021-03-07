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
// 2021-03-07 20:19:38.061281 +0000 UTC

package models

import "encoding/json"

// SubstanceReferenceInformation is documented here http://hl7.org/fhir/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformation struct {
	ID                *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                       `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                                    `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Comment           *string                                       `bson:"comment,omitempty" json:"comment,omitempty"`
	Gene              []SubstanceReferenceInformationGene           `bson:"gene,omitempty" json:"gene,omitempty"`
	GeneElement       []SubstanceReferenceInformationGeneElement    `bson:"geneElement,omitempty" json:"geneElement,omitempty"`
	Classification    []SubstanceReferenceInformationClassification `bson:"classification,omitempty" json:"classification,omitempty"`
	Target            []SubstanceReferenceInformationTarget         `bson:"target,omitempty" json:"target,omitempty"`
}
type SubstanceReferenceInformationGene struct {
	ID                 *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	GeneSequenceOrigin *CodeableConcept `bson:"geneSequenceOrigin,omitempty" json:"geneSequenceOrigin,omitempty"`
	Gene               *CodeableConcept `bson:"gene,omitempty" json:"gene,omitempty"`
	Source             []Reference      `bson:"source,omitempty" json:"source,omitempty"`
}
type SubstanceReferenceInformationGeneElement struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Element           *Identifier      `bson:"element,omitempty" json:"element,omitempty"`
	Source            []Reference      `bson:"source,omitempty" json:"source,omitempty"`
}
type SubstanceReferenceInformationClassification struct {
	ID                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Domain            *CodeableConcept  `bson:"domain,omitempty" json:"domain,omitempty"`
	Classification    *CodeableConcept  `bson:"classification,omitempty" json:"classification,omitempty"`
	Subtype           []CodeableConcept `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Source            []Reference       `bson:"source,omitempty" json:"source,omitempty"`
}
type SubstanceReferenceInformationTarget struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Target            *Identifier      `bson:"target,omitempty" json:"target,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Interaction       *CodeableConcept `bson:"interaction,omitempty" json:"interaction,omitempty"`
	Organism          *CodeableConcept `bson:"organism,omitempty" json:"organism,omitempty"`
	OrganismType      *CodeableConcept `bson:"organismType,omitempty" json:"organismType,omitempty"`
	AmountType        *CodeableConcept `bson:"amountType,omitempty" json:"amountType,omitempty"`
	Source            []Reference      `bson:"source,omitempty" json:"source,omitempty"`
}
type OtherSubstanceReferenceInformation SubstanceReferenceInformation

// MarshalJSON marshals the given SubstanceReferenceInformation as JSON into a byte slice
func (r SubstanceReferenceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceReferenceInformation
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceReferenceInformation: OtherSubstanceReferenceInformation(r),
		ResourceType:                       "SubstanceReferenceInformation",
	})
}

// UnmarshalSubstanceReferenceInformation unmarshals a SubstanceReferenceInformation.
func UnmarshalSubstanceReferenceInformation(b []byte) (SubstanceReferenceInformation, error) {
	var substanceReferenceInformation SubstanceReferenceInformation
	if err := json.Unmarshal(b, &substanceReferenceInformation); err != nil {
		return substanceReferenceInformation, err
	}
	return substanceReferenceInformation, nil
}
