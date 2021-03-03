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
// 2021-03-03 10:55:39.537099 +0000 UTC
// IdentityAssuranceLevel is documented here http://hl7.org/fhir/ValueSet/identity-assuranceLevel
type IdentityAssuranceLevel int

const (
	IdentityAssuranceLevelLevel1 IdentityAssuranceLevel = iota
	IdentityAssuranceLevelLevel2
	IdentityAssuranceLevelLevel3
	IdentityAssuranceLevelLevel4
)

func (code IdentityAssuranceLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *IdentityAssuranceLevel) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "level1":
		*code = IdentityAssuranceLevelLevel1
	case "level2":
		*code = IdentityAssuranceLevelLevel2
	case "level3":
		*code = IdentityAssuranceLevelLevel3
	case "level4":
		*code = IdentityAssuranceLevelLevel4
	default:
		return fmt.Errorf("unknown IdentityAssuranceLevel code `%s`", s)
	}
	return nil
}
func (code IdentityAssuranceLevel) String() string {
	return code.Code()
}
func (code IdentityAssuranceLevel) Code() string {
	switch code {
	case IdentityAssuranceLevelLevel1:
		return "level1"
	case IdentityAssuranceLevelLevel2:
		return "level2"
	case IdentityAssuranceLevelLevel3:
		return "level3"
	case IdentityAssuranceLevelLevel4:
		return "level4"
	}
	return "<unknown>"
}
func (code IdentityAssuranceLevel) Display() string {
	switch code {
	case IdentityAssuranceLevelLevel1:
		return "Level 1"
	case IdentityAssuranceLevelLevel2:
		return "Level 2"
	case IdentityAssuranceLevelLevel3:
		return "Level 3"
	case IdentityAssuranceLevelLevel4:
		return "Level 4"
	}
	return "<unknown>"
}
func (code IdentityAssuranceLevel) Definition() string {
	switch code {
	case IdentityAssuranceLevelLevel1:
		return "Little or no confidence in the asserted identity's accuracy."
	case IdentityAssuranceLevelLevel2:
		return "Some confidence in the asserted identity's accuracy."
	case IdentityAssuranceLevelLevel3:
		return "High confidence in the asserted identity's accuracy."
	case IdentityAssuranceLevelLevel4:
		return "Very high confidence in the asserted identity's accuracy."
	}
	return "<unknown>"
}
