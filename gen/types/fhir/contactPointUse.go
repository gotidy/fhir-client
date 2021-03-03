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
// 2021-03-03 10:55:39.636955 +0000 UTC
// ContactPointUse is documented here http://hl7.org/fhir/ValueSet/contact-point-use
type ContactPointUse int

const (
	ContactPointUseHome ContactPointUse = iota
	ContactPointUseWork
	ContactPointUseTemp
	ContactPointUseOld
	ContactPointUseMobile
)

func (code ContactPointUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ContactPointUse) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "home":
		*code = ContactPointUseHome
	case "work":
		*code = ContactPointUseWork
	case "temp":
		*code = ContactPointUseTemp
	case "old":
		*code = ContactPointUseOld
	case "mobile":
		*code = ContactPointUseMobile
	default:
		return fmt.Errorf("unknown ContactPointUse code `%s`", s)
	}
	return nil
}
func (code ContactPointUse) String() string {
	return code.Code()
}
func (code ContactPointUse) Code() string {
	switch code {
	case ContactPointUseHome:
		return "home"
	case ContactPointUseWork:
		return "work"
	case ContactPointUseTemp:
		return "temp"
	case ContactPointUseOld:
		return "old"
	case ContactPointUseMobile:
		return "mobile"
	}
	return "<unknown>"
}
func (code ContactPointUse) Display() string {
	switch code {
	case ContactPointUseHome:
		return "Home"
	case ContactPointUseWork:
		return "Work"
	case ContactPointUseTemp:
		return "Temp"
	case ContactPointUseOld:
		return "Old"
	case ContactPointUseMobile:
		return "Mobile"
	}
	return "<unknown>"
}
func (code ContactPointUse) Definition() string {
	switch code {
	case ContactPointUseHome:
		return "A communication contact point at a home; attempted contacts for business purposes might intrude privacy and chances are one will contact family or other household members instead of the person one wishes to call. Typically used with urgent cases, or if no other contacts are available."
	case ContactPointUseWork:
		return "An office contact point. First choice for business related contacts during business hours."
	case ContactPointUseTemp:
		return "A temporary contact point. The period can provide more detailed information."
	case ContactPointUseOld:
		return "This contact point is no longer in use (or was never correct, but retained for records)."
	case ContactPointUseMobile:
		return "A telecommunication device that moves and stays with its owner. May have characteristics of all other use codes, suitable for urgent matters, not the first choice for routine business."
	}
	return "<unknown>"
}
