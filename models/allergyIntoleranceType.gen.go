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
// 2021-03-08 09:41:02.992767 +0000 UTC

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AllergyIntoleranceType is documented here http://hl7.org/fhir/ValueSet/allergy-intolerance-type
type AllergyIntoleranceType int

const (
	AllergyIntoleranceTypeAllergy AllergyIntoleranceType = iota
	AllergyIntoleranceTypeIntolerance
)

func (code AllergyIntoleranceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AllergyIntoleranceType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "allergy":
		*code = AllergyIntoleranceTypeAllergy
	case "intolerance":
		*code = AllergyIntoleranceTypeIntolerance
	default:
		return fmt.Errorf("unknown AllergyIntoleranceType code `%s`", s)
	}
	return nil
}
func (code AllergyIntoleranceType) String() string {
	return code.Code()
}
func (code AllergyIntoleranceType) Code() string {
	switch code {
	case AllergyIntoleranceTypeAllergy:
		return "allergy"
	case AllergyIntoleranceTypeIntolerance:
		return "intolerance"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceType) Display() string {
	switch code {
	case AllergyIntoleranceTypeAllergy:
		return "Allergy"
	case AllergyIntoleranceTypeIntolerance:
		return "Intolerance"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceType) Definition() string {
	switch code {
	case AllergyIntoleranceTypeAllergy:
		return "A propensity for hypersensitive reaction(s) to a substance.  These reactions are most typically type I hypersensitivity, plus other \"allergy-like\" reactions, including pseudoallergy."
	case AllergyIntoleranceTypeIntolerance:
		return "A propensity for adverse reactions to a substance that is not judged to be allergic or \"allergy-like\".  These reactions are typically (but not necessarily) non-immune.  They are to some degree idiosyncratic and/or patient-specific (i.e. are not a reaction that is expected to occur with most or all patients given similar circumstances)."
	}
	return "<unknown>"
}
