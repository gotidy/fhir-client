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
// 2021-03-03 10:55:45.470194 +0000 UTC
// MeasureReportType is documented here http://hl7.org/fhir/ValueSet/measure-report-type
type MeasureReportType int

const (
	MeasureReportTypeIndividual MeasureReportType = iota
	MeasureReportTypeSubjectList
	MeasureReportTypeSummary
	MeasureReportTypeDataCollection
)

func (code MeasureReportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MeasureReportType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "individual":
		*code = MeasureReportTypeIndividual
	case "subject-list":
		*code = MeasureReportTypeSubjectList
	case "summary":
		*code = MeasureReportTypeSummary
	case "data-collection":
		*code = MeasureReportTypeDataCollection
	default:
		return fmt.Errorf("unknown MeasureReportType code `%s`", s)
	}
	return nil
}
func (code MeasureReportType) String() string {
	return code.Code()
}
func (code MeasureReportType) Code() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "individual"
	case MeasureReportTypeSubjectList:
		return "subject-list"
	case MeasureReportTypeSummary:
		return "summary"
	case MeasureReportTypeDataCollection:
		return "data-collection"
	}
	return "<unknown>"
}
func (code MeasureReportType) Display() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "Individual"
	case MeasureReportTypeSubjectList:
		return "Subject List"
	case MeasureReportTypeSummary:
		return "Summary"
	case MeasureReportTypeDataCollection:
		return "Data Collection"
	}
	return "<unknown>"
}
func (code MeasureReportType) Definition() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "An individual report that provides information on the performance for a given measure with respect to a single subject."
	case MeasureReportTypeSubjectList:
		return "A subject list report that includes a listing of subjects that satisfied each population criteria in the measure."
	case MeasureReportTypeSummary:
		return "A summary report that returns the number of members in each population criteria for the measure."
	case MeasureReportTypeDataCollection:
		return "A data collection report that contains data-of-interest for the measure."
	}
	return "<unknown>"
}
