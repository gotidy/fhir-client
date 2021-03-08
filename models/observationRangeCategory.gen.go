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
// 2021-03-08 09:41:03.186492 +0000 UTC

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ObservationRangeCategory is documented here http://hl7.org/fhir/ValueSet/observation-range-category
type ObservationRangeCategory int

const (
	ObservationRangeCategoryReference ObservationRangeCategory = iota
	ObservationRangeCategoryCritical
	ObservationRangeCategoryAbsolute
)

func (code ObservationRangeCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ObservationRangeCategory) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "reference":
		*code = ObservationRangeCategoryReference
	case "critical":
		*code = ObservationRangeCategoryCritical
	case "absolute":
		*code = ObservationRangeCategoryAbsolute
	default:
		return fmt.Errorf("unknown ObservationRangeCategory code `%s`", s)
	}
	return nil
}
func (code ObservationRangeCategory) String() string {
	return code.Code()
}
func (code ObservationRangeCategory) Code() string {
	switch code {
	case ObservationRangeCategoryReference:
		return "reference"
	case ObservationRangeCategoryCritical:
		return "critical"
	case ObservationRangeCategoryAbsolute:
		return "absolute"
	}
	return "<unknown>"
}
func (code ObservationRangeCategory) Display() string {
	switch code {
	case ObservationRangeCategoryReference:
		return "reference range"
	case ObservationRangeCategoryCritical:
		return "critical range"
	case ObservationRangeCategoryAbsolute:
		return "absolute range"
	}
	return "<unknown>"
}
func (code ObservationRangeCategory) Definition() string {
	switch code {
	case ObservationRangeCategoryReference:
		return "Reference (Normal) Range for Ordinal and Continuous Observations."
	case ObservationRangeCategoryCritical:
		return "Critical Range for Ordinal and Continuous Observations."
	case ObservationRangeCategoryAbsolute:
		return "Absolute Range for Ordinal and Continuous Observations. Results outside this range are not possible."
	}
	return "<unknown>"
}
