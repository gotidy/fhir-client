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
// 2021-03-03 10:55:45.47897 +0000 UTC
// DiscriminatorType is documented here http://hl7.org/fhir/ValueSet/discriminator-type
type DiscriminatorType int

const (
	DiscriminatorTypeValue DiscriminatorType = iota
	DiscriminatorTypeExists
	DiscriminatorTypePattern
	DiscriminatorTypeType
	DiscriminatorTypeProfile
)

func (code DiscriminatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DiscriminatorType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "value":
		*code = DiscriminatorTypeValue
	case "exists":
		*code = DiscriminatorTypeExists
	case "pattern":
		*code = DiscriminatorTypePattern
	case "type":
		*code = DiscriminatorTypeType
	case "profile":
		*code = DiscriminatorTypeProfile
	default:
		return fmt.Errorf("unknown DiscriminatorType code `%s`", s)
	}
	return nil
}
func (code DiscriminatorType) String() string {
	return code.Code()
}
func (code DiscriminatorType) Code() string {
	switch code {
	case DiscriminatorTypeValue:
		return "value"
	case DiscriminatorTypeExists:
		return "exists"
	case DiscriminatorTypePattern:
		return "pattern"
	case DiscriminatorTypeType:
		return "type"
	case DiscriminatorTypeProfile:
		return "profile"
	}
	return "<unknown>"
}
func (code DiscriminatorType) Display() string {
	switch code {
	case DiscriminatorTypeValue:
		return "Value"
	case DiscriminatorTypeExists:
		return "Exists"
	case DiscriminatorTypePattern:
		return "Pattern"
	case DiscriminatorTypeType:
		return "Type"
	case DiscriminatorTypeProfile:
		return "Profile"
	}
	return "<unknown>"
}
func (code DiscriminatorType) Definition() string {
	switch code {
	case DiscriminatorTypeValue:
		return "The slices have different values in the nominated element."
	case DiscriminatorTypeExists:
		return "The slices are differentiated by the presence or absence of the nominated element."
	case DiscriminatorTypePattern:
		return "The slices have different values in the nominated element, as determined by testing them against the applicable ElementDefinition.pattern[x]."
	case DiscriminatorTypeType:
		return "The slices are differentiated by type of the nominated element."
	case DiscriminatorTypeProfile:
		return "The slices are differentiated by conformance of the nominated element to a specified profile. Note that if the path specifies .resolve() then the profile is the target profile on the reference. In this case, validation by the possible profiles is required to differentiate the slices."
	}
	return "<unknown>"
}
