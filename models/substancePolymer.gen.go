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
// 2021-03-08 09:41:02.530022 +0000 UTC

package models

import "encoding/json"

// SubstancePolymer is documented here http://hl7.org/fhir/StructureDefinition/SubstancePolymer
type SubstancePolymer struct {
	ID                    *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Class                 *CodeableConcept             `bson:"class,omitempty" json:"class,omitempty"`
	Geometry              *CodeableConcept             `bson:"geometry,omitempty" json:"geometry,omitempty"`
	CopolymerConnectivity []CodeableConcept            `bson:"copolymerConnectivity,omitempty" json:"copolymerConnectivity,omitempty"`
	Modification          []string                     `bson:"modification,omitempty" json:"modification,omitempty"`
	MonomerSet            []SubstancePolymerMonomerSet `bson:"monomerSet,omitempty" json:"monomerSet,omitempty"`
	Repeat                []SubstancePolymerRepeat     `bson:"repeat,omitempty" json:"repeat,omitempty"`
}
type SubstancePolymerMonomerSet struct {
	ID                *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RatioType         *CodeableConcept                             `bson:"ratioType,omitempty" json:"ratioType,omitempty"`
	StartingMaterial  []SubstancePolymerMonomerSetStartingMaterial `bson:"startingMaterial,omitempty" json:"startingMaterial,omitempty"`
}
type SubstancePolymerMonomerSetStartingMaterial struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Material          *CodeableConcept `bson:"material,omitempty" json:"material,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	IsDefining        *bool            `bson:"isDefining,omitempty" json:"isDefining,omitempty"`
	Amount            *SubstanceAmount `bson:"amount,omitempty" json:"amount,omitempty"`
}
type SubstancePolymerRepeat struct {
	ID                      *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	NumberOfUnits           *int                               `bson:"numberOfUnits,omitempty" json:"numberOfUnits,omitempty"`
	AverageMolecularFormula *string                            `bson:"averageMolecularFormula,omitempty" json:"averageMolecularFormula,omitempty"`
	RepeatUnitAmountType    *CodeableConcept                   `bson:"repeatUnitAmountType,omitempty" json:"repeatUnitAmountType,omitempty"`
	RepeatUnit              []SubstancePolymerRepeatRepeatUnit `bson:"repeatUnit,omitempty" json:"repeatUnit,omitempty"`
}
type SubstancePolymerRepeatRepeatUnit struct {
	ID                          *string                                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension                   []Extension                                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension           []Extension                                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OrientationOfPolymerisation *CodeableConcept                                           `bson:"orientationOfPolymerisation,omitempty" json:"orientationOfPolymerisation,omitempty"`
	RepeatUnit                  *string                                                    `bson:"repeatUnit,omitempty" json:"repeatUnit,omitempty"`
	Amount                      *SubstanceAmount                                           `bson:"amount,omitempty" json:"amount,omitempty"`
	DegreeOfPolymerisation      []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation   `bson:"degreeOfPolymerisation,omitempty" json:"degreeOfPolymerisation,omitempty"`
	StructuralRepresentation    []SubstancePolymerRepeatRepeatUnitStructuralRepresentation `bson:"structuralRepresentation,omitempty" json:"structuralRepresentation,omitempty"`
}
type SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Degree            *CodeableConcept `bson:"degree,omitempty" json:"degree,omitempty"`
	Amount            *SubstanceAmount `bson:"amount,omitempty" json:"amount,omitempty"`
}
type SubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Representation    *string          `bson:"representation,omitempty" json:"representation,omitempty"`
	Attachment        *Attachment      `bson:"attachment,omitempty" json:"attachment,omitempty"`
}
type OtherSubstancePolymer SubstancePolymer

// MarshalJSON marshals the given SubstancePolymer as JSON into a byte slice
func (r SubstancePolymer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstancePolymer
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstancePolymer: OtherSubstancePolymer(r),
		ResourceType:          "SubstancePolymer",
	})
}

// UnmarshalSubstancePolymer unmarshals a SubstancePolymer.
func UnmarshalSubstancePolymer(b []byte) (SubstancePolymer, error) {
	var substancePolymer SubstancePolymer
	if err := json.Unmarshal(b, &substancePolymer); err != nil {
		return substancePolymer, err
	}
	return substancePolymer, nil
}
