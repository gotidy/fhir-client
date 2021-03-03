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
// 2021-03-03 10:55:45.132084 +0000 UTC
// ImmunizationEvaluationStatusCodes is documented here http://hl7.org/fhir/ValueSet/immunization-evaluation-status
type ImmunizationEvaluationStatusCodes int

const (
	ImmunizationEvaluationStatusCodesInProgress ImmunizationEvaluationStatusCodes = iota
	ImmunizationEvaluationStatusCodesNotDone
	ImmunizationEvaluationStatusCodesOnHold
	ImmunizationEvaluationStatusCodesCompleted
	ImmunizationEvaluationStatusCodesEnteredInError
	ImmunizationEvaluationStatusCodesStopped
	ImmunizationEvaluationStatusCodesUnknown
)

func (code ImmunizationEvaluationStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ImmunizationEvaluationStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "in-progress":
		*code = ImmunizationEvaluationStatusCodesInProgress
	case "not-done":
		*code = ImmunizationEvaluationStatusCodesNotDone
	case "on-hold":
		*code = ImmunizationEvaluationStatusCodesOnHold
	case "completed":
		*code = ImmunizationEvaluationStatusCodesCompleted
	case "entered-in-error":
		*code = ImmunizationEvaluationStatusCodesEnteredInError
	case "stopped":
		*code = ImmunizationEvaluationStatusCodesStopped
	case "unknown":
		*code = ImmunizationEvaluationStatusCodesUnknown
	default:
		return fmt.Errorf("unknown ImmunizationEvaluationStatusCodes code `%s`", s)
	}
	return nil
}
func (code ImmunizationEvaluationStatusCodes) String() string {
	return code.Code()
}
func (code ImmunizationEvaluationStatusCodes) Code() string {
	switch code {
	case ImmunizationEvaluationStatusCodesInProgress:
		return "in-progress"
	case ImmunizationEvaluationStatusCodesNotDone:
		return "not-done"
	case ImmunizationEvaluationStatusCodesOnHold:
		return "on-hold"
	case ImmunizationEvaluationStatusCodesCompleted:
		return "completed"
	case ImmunizationEvaluationStatusCodesEnteredInError:
		return "entered-in-error"
	case ImmunizationEvaluationStatusCodesStopped:
		return "stopped"
	case ImmunizationEvaluationStatusCodesUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code ImmunizationEvaluationStatusCodes) Display() string {
	switch code {
	case ImmunizationEvaluationStatusCodesInProgress:
		return "In Progress"
	case ImmunizationEvaluationStatusCodesNotDone:
		return "Not Done"
	case ImmunizationEvaluationStatusCodesOnHold:
		return "On Hold"
	case ImmunizationEvaluationStatusCodesCompleted:
		return "Completed"
	case ImmunizationEvaluationStatusCodesEnteredInError:
		return "Entered in Error"
	case ImmunizationEvaluationStatusCodesStopped:
		return "Stopped"
	case ImmunizationEvaluationStatusCodesUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code ImmunizationEvaluationStatusCodes) Definition() string {
	switch code {
	case ImmunizationEvaluationStatusCodesInProgress:
		return "The administration has started but has not yet completed."
	case ImmunizationEvaluationStatusCodesNotDone:
		return "The administration was terminated prior to any impact on the subject (though preparatory actions may have been taken)"
	case ImmunizationEvaluationStatusCodesOnHold:
		return "Actions implied by the administration have been temporarily halted, but are expected to continue later. May also be called 'suspended'."
	case ImmunizationEvaluationStatusCodesCompleted:
		return "All actions that are implied by the administration have occurred."
	case ImmunizationEvaluationStatusCodesEnteredInError:
		return "The administration was entered in error and therefore nullified."
	case ImmunizationEvaluationStatusCodesStopped:
		return "Actions implied by the administration have been permanently halted, before all of them occurred."
	case ImmunizationEvaluationStatusCodesUnknown:
		return "The authoring system does not know which of the status values currently applies for this request. Note: This concept is not to be used for 'other' - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
