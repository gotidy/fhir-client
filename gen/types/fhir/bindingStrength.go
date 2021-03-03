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
// 2021-03-03 10:55:39.585986 +0000 UTC
// BindingStrength is documented here http://hl7.org/fhir/ValueSet/binding-strength
type BindingStrength int

const (
	BindingStrengthRequired BindingStrength = iota
	BindingStrengthExtensible
	BindingStrengthPreferred
	BindingStrengthExample
)

func (code BindingStrength) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *BindingStrength) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "required":
		*code = BindingStrengthRequired
	case "extensible":
		*code = BindingStrengthExtensible
	case "preferred":
		*code = BindingStrengthPreferred
	case "example":
		*code = BindingStrengthExample
	default:
		return fmt.Errorf("unknown BindingStrength code `%s`", s)
	}
	return nil
}
func (code BindingStrength) String() string {
	return code.Code()
}
func (code BindingStrength) Code() string {
	switch code {
	case BindingStrengthRequired:
		return "required"
	case BindingStrengthExtensible:
		return "extensible"
	case BindingStrengthPreferred:
		return "preferred"
	case BindingStrengthExample:
		return "example"
	}
	return "<unknown>"
}
func (code BindingStrength) Display() string {
	switch code {
	case BindingStrengthRequired:
		return "Required"
	case BindingStrengthExtensible:
		return "Extensible"
	case BindingStrengthPreferred:
		return "Preferred"
	case BindingStrengthExample:
		return "Example"
	}
	return "<unknown>"
}
func (code BindingStrength) Definition() string {
	switch code {
	case BindingStrengthRequired:
		return "To be conformant, the concept in this element SHALL be from the specified value set."
	case BindingStrengthExtensible:
		return "To be conformant, the concept in this element SHALL be from the specified value set if any of the codes within the value set can apply to the concept being communicated.  If the value set does not cover the concept (based on human review), alternate codings (or, data type allowing, text) may be included instead."
	case BindingStrengthPreferred:
		return "Instances are encouraged to draw from the specified codes for interoperability purposes but are not required to do so to be considered conformant."
	case BindingStrengthExample:
		return "Instances are not expected or even encouraged to draw from the specified value set.  The value set merely provides examples of the types of concepts intended to be included."
	}
	return "<unknown>"
}
