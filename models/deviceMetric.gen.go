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
// 2021-03-08 09:41:02.933896 +0000 UTC

package models

import "encoding/json"

// DeviceMetric is documented here http://hl7.org/fhir/StructureDefinition/DeviceMetric
type DeviceMetric struct {
	ID                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              CodeableConcept                `bson:"type" json:"type"`
	Unit              *CodeableConcept               `bson:"unit,omitempty" json:"unit,omitempty"`
	Source            *Reference                     `bson:"source,omitempty" json:"source,omitempty"`
	Parent            *Reference                     `bson:"parent,omitempty" json:"parent,omitempty"`
	OperationalStatus *DeviceMetricOperationalStatus `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	Color             *DeviceMetricColor             `bson:"color,omitempty" json:"color,omitempty"`
	Category          DeviceMetricCategory           `bson:"category" json:"category"`
	MeasurementPeriod *Timing                        `bson:"measurementPeriod,omitempty" json:"measurementPeriod,omitempty"`
	Calibration       []DeviceMetricCalibration      `bson:"calibration,omitempty" json:"calibration,omitempty"`
}
type DeviceMetricCalibration struct {
	ID                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *DeviceMetricCalibrationType  `bson:"type,omitempty" json:"type,omitempty"`
	State             *DeviceMetricCalibrationState `bson:"state,omitempty" json:"state,omitempty"`
	Time              *string                       `bson:"time,omitempty" json:"time,omitempty"`
}
type OtherDeviceMetric DeviceMetric

// MarshalJSON marshals the given DeviceMetric as JSON into a byte slice
func (r DeviceMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceMetric
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceMetric: OtherDeviceMetric(r),
		ResourceType:      "DeviceMetric",
	})
}

// UnmarshalDeviceMetric unmarshals a DeviceMetric.
func UnmarshalDeviceMetric(b []byte) (DeviceMetric, error) {
	var deviceMetric DeviceMetric
	if err := json.Unmarshal(b, &deviceMetric); err != nil {
		return deviceMetric, err
	}
	return deviceMetric, nil
}
