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
// 2021-03-08 09:41:03.07252 +0000 UTC

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TriggerType is documented here http://hl7.org/fhir/ValueSet/trigger-type
type TriggerType int

const (
	TriggerTypeNamedEvent TriggerType = iota
	TriggerTypePeriodic
	TriggerTypeDataChanged
	TriggerTypeDataAdded
	TriggerTypeDataModified
	TriggerTypeDataRemoved
	TriggerTypeDataAccessed
	TriggerTypeDataAccessEnded
)

func (code TriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TriggerType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "named-event":
		*code = TriggerTypeNamedEvent
	case "periodic":
		*code = TriggerTypePeriodic
	case "data-changed":
		*code = TriggerTypeDataChanged
	case "data-added":
		*code = TriggerTypeDataAdded
	case "data-modified":
		*code = TriggerTypeDataModified
	case "data-removed":
		*code = TriggerTypeDataRemoved
	case "data-accessed":
		*code = TriggerTypeDataAccessed
	case "data-access-ended":
		*code = TriggerTypeDataAccessEnded
	default:
		return fmt.Errorf("unknown TriggerType code `%s`", s)
	}
	return nil
}
func (code TriggerType) String() string {
	return code.Code()
}
func (code TriggerType) Code() string {
	switch code {
	case TriggerTypeNamedEvent:
		return "named-event"
	case TriggerTypePeriodic:
		return "periodic"
	case TriggerTypeDataChanged:
		return "data-changed"
	case TriggerTypeDataAdded:
		return "data-added"
	case TriggerTypeDataModified:
		return "data-modified"
	case TriggerTypeDataRemoved:
		return "data-removed"
	case TriggerTypeDataAccessed:
		return "data-accessed"
	case TriggerTypeDataAccessEnded:
		return "data-access-ended"
	}
	return "<unknown>"
}
func (code TriggerType) Display() string {
	switch code {
	case TriggerTypeNamedEvent:
		return "Named Event"
	case TriggerTypePeriodic:
		return "Periodic"
	case TriggerTypeDataChanged:
		return "Data Changed"
	case TriggerTypeDataAdded:
		return "Data Added"
	case TriggerTypeDataModified:
		return "Data Updated"
	case TriggerTypeDataRemoved:
		return "Data Removed"
	case TriggerTypeDataAccessed:
		return "Data Accessed"
	case TriggerTypeDataAccessEnded:
		return "Data Access Ended"
	}
	return "<unknown>"
}
func (code TriggerType) Definition() string {
	switch code {
	case TriggerTypeNamedEvent:
		return "The trigger occurs in response to a specific named event, and no other information about the trigger is specified. Named events are completely pre-coordinated, and the formal semantics of the trigger are not provided."
	case TriggerTypePeriodic:
		return "The trigger occurs at a specific time or periodically as described by a timing or schedule. A periodic event cannot have any data elements, but may have a name assigned as a shorthand for the event."
	case TriggerTypeDataChanged:
		return "The trigger occurs whenever data of a particular type is changed in any way, either added, modified, or removed."
	case TriggerTypeDataAdded:
		return "The trigger occurs whenever data of a particular type is added."
	case TriggerTypeDataModified:
		return "The trigger occurs whenever data of a particular type is modified."
	case TriggerTypeDataRemoved:
		return "The trigger occurs whenever data of a particular type is removed."
	case TriggerTypeDataAccessed:
		return "The trigger occurs whenever data of a particular type is accessed."
	case TriggerTypeDataAccessEnded:
		return "The trigger occurs whenever access to data of a particular type is completed."
	}
	return "<unknown>"
}
