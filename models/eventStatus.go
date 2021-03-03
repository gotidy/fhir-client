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
// 2021-03-03 10:55:45.354009 +0000 UTC
// EventStatus is documented here http://hl7.org/fhir/ValueSet/event-status
type EventStatus int

const (
	EventStatusPreparation EventStatus = iota
	EventStatusInProgress
	EventStatusNotDone
	EventStatusOnHold
	EventStatusStopped
	EventStatusCompleted
	EventStatusEnteredInError
	EventStatusUnknown
)

func (code EventStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EventStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "preparation":
		*code = EventStatusPreparation
	case "in-progress":
		*code = EventStatusInProgress
	case "not-done":
		*code = EventStatusNotDone
	case "on-hold":
		*code = EventStatusOnHold
	case "stopped":
		*code = EventStatusStopped
	case "completed":
		*code = EventStatusCompleted
	case "entered-in-error":
		*code = EventStatusEnteredInError
	case "unknown":
		*code = EventStatusUnknown
	default:
		return fmt.Errorf("unknown EventStatus code `%s`", s)
	}
	return nil
}
func (code EventStatus) String() string {
	return code.Code()
}
func (code EventStatus) Code() string {
	switch code {
	case EventStatusPreparation:
		return "preparation"
	case EventStatusInProgress:
		return "in-progress"
	case EventStatusNotDone:
		return "not-done"
	case EventStatusOnHold:
		return "on-hold"
	case EventStatusStopped:
		return "stopped"
	case EventStatusCompleted:
		return "completed"
	case EventStatusEnteredInError:
		return "entered-in-error"
	case EventStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code EventStatus) Display() string {
	switch code {
	case EventStatusPreparation:
		return "Preparation"
	case EventStatusInProgress:
		return "In Progress"
	case EventStatusNotDone:
		return "Not Done"
	case EventStatusOnHold:
		return "On Hold"
	case EventStatusStopped:
		return "Stopped"
	case EventStatusCompleted:
		return "Completed"
	case EventStatusEnteredInError:
		return "Entered in Error"
	case EventStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code EventStatus) Definition() string {
	switch code {
	case EventStatusPreparation:
		return "The core event has not started yet, but some staging activities have begun (e.g. surgical suite preparation).  Preparation stages may be tracked for billing purposes."
	case EventStatusInProgress:
		return "The event is currently occurring."
	case EventStatusNotDone:
		return "The event was terminated prior to any activity beyond preparation.  I.e. The 'main' activity has not yet begun.  The boundary between preparatory and the 'main' activity is context-specific."
	case EventStatusOnHold:
		return "The event has been temporarily stopped but is expected to resume in the future."
	case EventStatusStopped:
		return "The event was terminated prior to the full completion of the intended activity but after at least some of the 'main' activity (beyond preparation) has occurred."
	case EventStatusCompleted:
		return "The event has now concluded."
	case EventStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"stopped\" rather than \"entered-in-error\".)."
	case EventStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this event.  Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which."
	}
	return "<unknown>"
}
