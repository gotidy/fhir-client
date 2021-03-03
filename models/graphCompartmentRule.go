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
// 2021-03-03 10:55:45.128163 +0000 UTC
// GraphCompartmentRule is documented here http://hl7.org/fhir/ValueSet/graph-compartment-rule
type GraphCompartmentRule int

const (
	GraphCompartmentRuleIdentical GraphCompartmentRule = iota
	GraphCompartmentRuleMatching
	GraphCompartmentRuleDifferent
	GraphCompartmentRuleCustom
)

func (code GraphCompartmentRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *GraphCompartmentRule) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "identical":
		*code = GraphCompartmentRuleIdentical
	case "matching":
		*code = GraphCompartmentRuleMatching
	case "different":
		*code = GraphCompartmentRuleDifferent
	case "custom":
		*code = GraphCompartmentRuleCustom
	default:
		return fmt.Errorf("unknown GraphCompartmentRule code `%s`", s)
	}
	return nil
}
func (code GraphCompartmentRule) String() string {
	return code.Code()
}
func (code GraphCompartmentRule) Code() string {
	switch code {
	case GraphCompartmentRuleIdentical:
		return "identical"
	case GraphCompartmentRuleMatching:
		return "matching"
	case GraphCompartmentRuleDifferent:
		return "different"
	case GraphCompartmentRuleCustom:
		return "custom"
	}
	return "<unknown>"
}
func (code GraphCompartmentRule) Display() string {
	switch code {
	case GraphCompartmentRuleIdentical:
		return "Identical"
	case GraphCompartmentRuleMatching:
		return "Matching"
	case GraphCompartmentRuleDifferent:
		return "Different"
	case GraphCompartmentRuleCustom:
		return "Custom"
	}
	return "<unknown>"
}
func (code GraphCompartmentRule) Definition() string {
	switch code {
	case GraphCompartmentRuleIdentical:
		return "The compartment must be identical (the same literal reference)."
	case GraphCompartmentRuleMatching:
		return "The compartment must be the same - the record must be about the same patient, but the reference may be different."
	case GraphCompartmentRuleDifferent:
		return "The compartment must be different."
	case GraphCompartmentRuleCustom:
		return "The compartment rule is defined in the accompanying FHIRPath expression."
	}
	return "<unknown>"
}
