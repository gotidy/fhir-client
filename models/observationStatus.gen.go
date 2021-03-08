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
// 2021-03-08 09:41:03.022352 +0000 UTC

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ObservationStatus is documented here http://hl7.org/fhir/ValueSet/observation-status
type ObservationStatus int

const (
	ObservationStatusRegistered ObservationStatus = iota
	ObservationStatusPreliminary
	ObservationStatusFinal
	ObservationStatusAmended
	ObservationStatusCorrected
	ObservationStatusCancelled
	ObservationStatusEnteredInError
	ObservationStatusUnknown
)

func (code ObservationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ObservationStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "registered":
		*code = ObservationStatusRegistered
	case "preliminary":
		*code = ObservationStatusPreliminary
	case "final":
		*code = ObservationStatusFinal
	case "amended":
		*code = ObservationStatusAmended
	case "corrected":
		*code = ObservationStatusCorrected
	case "cancelled":
		*code = ObservationStatusCancelled
	case "entered-in-error":
		*code = ObservationStatusEnteredInError
	case "unknown":
		*code = ObservationStatusUnknown
	default:
		return fmt.Errorf("unknown ObservationStatus code `%s`", s)
	}
	return nil
}
func (code ObservationStatus) String() string {
	return code.Code()
}
func (code ObservationStatus) Code() string {
	switch code {
	case ObservationStatusRegistered:
		return "registered"
	case ObservationStatusPreliminary:
		return "preliminary"
	case ObservationStatusFinal:
		return "final"
	case ObservationStatusAmended:
		return "amended"
	case ObservationStatusCorrected:
		return "corrected"
	case ObservationStatusCancelled:
		return "cancelled"
	case ObservationStatusEnteredInError:
		return "entered-in-error"
	case ObservationStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code ObservationStatus) Display() string {
	switch code {
	case ObservationStatusRegistered:
		return "Registered"
	case ObservationStatusPreliminary:
		return "Preliminary"
	case ObservationStatusFinal:
		return "Final"
	case ObservationStatusAmended:
		return "Amended"
	case ObservationStatusCorrected:
		return "Corrected"
	case ObservationStatusCancelled:
		return "Cancelled"
	case ObservationStatusEnteredInError:
		return "Entered in Error"
	case ObservationStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code ObservationStatus) Definition() string {
	switch code {
	case ObservationStatusRegistered:
		return "The existence of the observation is registered, but there is no result yet available."
	case ObservationStatusPreliminary:
		return "This is an initial or interim observation: data may be incomplete or unverified."
	case ObservationStatusFinal:
		return "The observation is complete and there are no further actions needed. Additional information such \"released\", \"signed\", etc would be represented using [Provenance](provenance.html) which provides not only the act but also the actors and dates and other related data. These act states would be associated with an observation status of `preliminary` until they are all completed and then a status of `final` would be applied."
	case ObservationStatusAmended:
		return "Subsequent to being Final, the observation has been modified subsequent.  This includes updates/new information and corrections."
	case ObservationStatusCorrected:
		return "Subsequent to being Final, the observation has been modified to correct an error in the test result."
	case ObservationStatusCancelled:
		return "The observation is unavailable because the measurement was not started or not completed (also sometimes called \"aborted\")."
	case ObservationStatusEnteredInError:
		return "The observation has been withdrawn following previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	case ObservationStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which."
	}
	return "<unknown>"
}
