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
// 2021-03-08 09:41:02.943595 +0000 UTC

package models

// Timing is documented here http://hl7.org/fhir/StructureDefinition/Timing
type Timing struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Event             []DateTime       `bson:"event,omitempty" json:"event,omitempty"`
	Repeat            *TimingRepeat    `bson:"repeat,omitempty" json:"repeat,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}
type TimingRepeat struct {
	ID           *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension    []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	Count        *int         `bson:"count,omitempty" json:"count,omitempty"`
	CountMax     *int         `bson:"countMax,omitempty" json:"countMax,omitempty"`
	Duration     *Decimal     `bson:"duration,omitempty" json:"duration,omitempty"`
	DurationMax  *Decimal     `bson:"durationMax,omitempty" json:"durationMax,omitempty"`
	DurationUnit *string      `bson:"durationUnit,omitempty" json:"durationUnit,omitempty"`
	Frequency    *int         `bson:"frequency,omitempty" json:"frequency,omitempty"`
	FrequencyMax *int         `bson:"frequencyMax,omitempty" json:"frequencyMax,omitempty"`
	Period       *Decimal     `bson:"period,omitempty" json:"period,omitempty"`
	PeriodMax    *Decimal     `bson:"periodMax,omitempty" json:"periodMax,omitempty"`
	PeriodUnit   *string      `bson:"periodUnit,omitempty" json:"periodUnit,omitempty"`
	DayOfWeek    []DaysOfWeek `bson:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty"`
	TimeOfDay    []Time       `bson:"timeOfDay,omitempty" json:"timeOfDay,omitempty"`
	When         []string     `bson:"when,omitempty" json:"when,omitempty"`
	Offset       *int         `bson:"offset,omitempty" json:"offset,omitempty"`
}
