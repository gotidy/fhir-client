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

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Code generated by go generate; DO NOT EDIT.

// This file was generated by robots at
// 2021-03-03 10:55:39.635425 +0000 UTC
// ObservationDataType is documented here http://hl7.org/fhir/ValueSet/permitted-data-type
type ObservationDataType int

const (
	ObservationDataTypeQuantity ObservationDataType = iota
	ObservationDataTypeCodeableConcept
	ObservationDataTypeString
	ObservationDataTypeBoolean
	ObservationDataTypeInteger
	ObservationDataTypeRange
	ObservationDataTypeRatio
	ObservationDataTypeSampledData
	ObservationDataTypeTime
	ObservationDataTypeDateTime
	ObservationDataTypePeriod
)

func (code ObservationDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ObservationDataType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "Quantity":
		*code = ObservationDataTypeQuantity
	case "CodeableConcept":
		*code = ObservationDataTypeCodeableConcept
	case "string":
		*code = ObservationDataTypeString
	case "boolean":
		*code = ObservationDataTypeBoolean
	case "integer":
		*code = ObservationDataTypeInteger
	case "Range":
		*code = ObservationDataTypeRange
	case "Ratio":
		*code = ObservationDataTypeRatio
	case "SampledData":
		*code = ObservationDataTypeSampledData
	case "time":
		*code = ObservationDataTypeTime
	case "dateTime":
		*code = ObservationDataTypeDateTime
	case "Period":
		*code = ObservationDataTypePeriod
	default:
		return fmt.Errorf("unknown ObservationDataType code `%s`", s)
	}
	return nil
}
func (code ObservationDataType) String() string {
	return code.Code()
}
func (code ObservationDataType) Code() string {
	switch code {
	case ObservationDataTypeQuantity:
		return "Quantity"
	case ObservationDataTypeCodeableConcept:
		return "CodeableConcept"
	case ObservationDataTypeString:
		return "string"
	case ObservationDataTypeBoolean:
		return "boolean"
	case ObservationDataTypeInteger:
		return "integer"
	case ObservationDataTypeRange:
		return "Range"
	case ObservationDataTypeRatio:
		return "Ratio"
	case ObservationDataTypeSampledData:
		return "SampledData"
	case ObservationDataTypeTime:
		return "time"
	case ObservationDataTypeDateTime:
		return "dateTime"
	case ObservationDataTypePeriod:
		return "Period"
	}
	return "<unknown>"
}
func (code ObservationDataType) Display() string {
	switch code {
	case ObservationDataTypeQuantity:
		return "Quantity"
	case ObservationDataTypeCodeableConcept:
		return "CodeableConcept"
	case ObservationDataTypeString:
		return "string"
	case ObservationDataTypeBoolean:
		return "boolean"
	case ObservationDataTypeInteger:
		return "integer"
	case ObservationDataTypeRange:
		return "Range"
	case ObservationDataTypeRatio:
		return "Ratio"
	case ObservationDataTypeSampledData:
		return "SampledData"
	case ObservationDataTypeTime:
		return "time"
	case ObservationDataTypeDateTime:
		return "dateTime"
	case ObservationDataTypePeriod:
		return "Period"
	}
	return "<unknown>"
}
func (code ObservationDataType) Definition() string {
	switch code {
	case ObservationDataTypeQuantity:
		return "A measured amount."
	case ObservationDataTypeCodeableConcept:
		return "A coded concept from a reference terminology and/or text."
	case ObservationDataTypeString:
		return "A sequence of Unicode characters."
	case ObservationDataTypeBoolean:
		return "true or false."
	case ObservationDataTypeInteger:
		return "A signed integer."
	case ObservationDataTypeRange:
		return "A set of values bounded by low and high."
	case ObservationDataTypeRatio:
		return "A ratio of two Quantity values - a numerator and a denominator."
	case ObservationDataTypeSampledData:
		return "A series of measurements taken by a device."
	case ObservationDataTypeTime:
		return "A time during the day, in the format hh:mm:ss."
	case ObservationDataTypeDateTime:
		return "A date, date-time or partial date (e.g. just year or year + month) as used in human communication."
	case ObservationDataTypePeriod:
		return "A time range defined by start and end date/time."
	}
	return "<unknown>"
}
