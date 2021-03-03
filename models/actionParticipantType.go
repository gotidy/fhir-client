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
// 2021-03-03 10:55:45.554774 +0000 UTC
// ActionParticipantType is documented here http://hl7.org/fhir/ValueSet/action-participant-type
type ActionParticipantType int

const (
	ActionParticipantTypePatient ActionParticipantType = iota
	ActionParticipantTypePractitioner
	ActionParticipantTypeRelatedPerson
	ActionParticipantTypeDevice
)

func (code ActionParticipantType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionParticipantType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "patient":
		*code = ActionParticipantTypePatient
	case "practitioner":
		*code = ActionParticipantTypePractitioner
	case "related-person":
		*code = ActionParticipantTypeRelatedPerson
	case "device":
		*code = ActionParticipantTypeDevice
	default:
		return fmt.Errorf("unknown ActionParticipantType code `%s`", s)
	}
	return nil
}
func (code ActionParticipantType) String() string {
	return code.Code()
}
func (code ActionParticipantType) Code() string {
	switch code {
	case ActionParticipantTypePatient:
		return "patient"
	case ActionParticipantTypePractitioner:
		return "practitioner"
	case ActionParticipantTypeRelatedPerson:
		return "related-person"
	case ActionParticipantTypeDevice:
		return "device"
	}
	return "<unknown>"
}
func (code ActionParticipantType) Display() string {
	switch code {
	case ActionParticipantTypePatient:
		return "Patient"
	case ActionParticipantTypePractitioner:
		return "Practitioner"
	case ActionParticipantTypeRelatedPerson:
		return "Related Person"
	case ActionParticipantTypeDevice:
		return "Device"
	}
	return "<unknown>"
}
func (code ActionParticipantType) Definition() string {
	switch code {
	case ActionParticipantTypePatient:
		return "The participant is the patient under evaluation."
	case ActionParticipantTypePractitioner:
		return "The participant is a practitioner involved in the patient's care."
	case ActionParticipantTypeRelatedPerson:
		return "The participant is a person related to the patient."
	case ActionParticipantTypeDevice:
		return "The participant is a system or device used in the care of the patient."
	}
	return "<unknown>"
}
