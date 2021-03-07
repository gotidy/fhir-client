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
// 2021-03-07 20:19:39.523192 +0000 UTC

package models

import "encoding/json"

// ImmunizationRecommendation is documented here http://hl7.org/fhir/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendation struct {
	ID                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient           Reference                                  `bson:"patient" json:"patient"`
	Date              DateTime                                   `bson:"date" json:"date"`
	Authority         *Reference                                 `bson:"authority,omitempty" json:"authority,omitempty"`
	Recommendation    []ImmunizationRecommendationRecommendation `bson:"recommendation" json:"recommendation"`
}
type ImmunizationRecommendationRecommendation struct {
	ID                           *string                                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension                    []Extension                                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension            []Extension                                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	VaccineCode                  []CodeableConcept                                       `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty"`
	TargetDisease                *CodeableConcept                                        `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
	ContraindicatedVaccineCode   []CodeableConcept                                       `bson:"contraindicatedVaccineCode,omitempty" json:"contraindicatedVaccineCode,omitempty"`
	ForecastStatus               CodeableConcept                                         `bson:"forecastStatus" json:"forecastStatus"`
	ForecastReason               []CodeableConcept                                       `bson:"forecastReason,omitempty" json:"forecastReason,omitempty"`
	DateCriterion                []ImmunizationRecommendationRecommendationDateCriterion `bson:"dateCriterion,omitempty" json:"dateCriterion,omitempty"`
	Description                  *string                                                 `bson:"description,omitempty" json:"description,omitempty"`
	Series                       *string                                                 `bson:"series,omitempty" json:"series,omitempty"`
	SupportingImmunization       []Reference                                             `bson:"supportingImmunization,omitempty" json:"supportingImmunization,omitempty"`
	SupportingPatientInformation []Reference                                             `bson:"supportingPatientInformation,omitempty" json:"supportingPatientInformation,omitempty"`
}
type ImmunizationRecommendationRecommendationDateCriterion struct {
	ID                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Value             DateTime        `bson:"value" json:"value"`
}
type OtherImmunizationRecommendation ImmunizationRecommendation

// MarshalJSON marshals the given ImmunizationRecommendation as JSON into a byte slice
func (r ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunizationRecommendation
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunizationRecommendation: OtherImmunizationRecommendation(r),
		ResourceType:                    "ImmunizationRecommendation",
	})
}

// UnmarshalImmunizationRecommendation unmarshals a ImmunizationRecommendation.
func UnmarshalImmunizationRecommendation(b []byte) (ImmunizationRecommendation, error) {
	var immunizationRecommendation ImmunizationRecommendation
	if err := json.Unmarshal(b, &immunizationRecommendation); err != nil {
		return immunizationRecommendation, err
	}
	return immunizationRecommendation, nil
}
