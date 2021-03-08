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
// 2021-03-08 09:41:03.177222 +0000 UTC

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ConsentDataMeaning is documented here http://hl7.org/fhir/ValueSet/consent-data-meaning
type ConsentDataMeaning int

const (
	ConsentDataMeaningInstance ConsentDataMeaning = iota
	ConsentDataMeaningRelated
	ConsentDataMeaningDependents
	ConsentDataMeaningAuthoredby
)

func (code ConsentDataMeaning) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConsentDataMeaning) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "instance":
		*code = ConsentDataMeaningInstance
	case "related":
		*code = ConsentDataMeaningRelated
	case "dependents":
		*code = ConsentDataMeaningDependents
	case "authoredby":
		*code = ConsentDataMeaningAuthoredby
	default:
		return fmt.Errorf("unknown ConsentDataMeaning code `%s`", s)
	}
	return nil
}
func (code ConsentDataMeaning) String() string {
	return code.Code()
}
func (code ConsentDataMeaning) Code() string {
	switch code {
	case ConsentDataMeaningInstance:
		return "instance"
	case ConsentDataMeaningRelated:
		return "related"
	case ConsentDataMeaningDependents:
		return "dependents"
	case ConsentDataMeaningAuthoredby:
		return "authoredby"
	}
	return "<unknown>"
}
func (code ConsentDataMeaning) Display() string {
	switch code {
	case ConsentDataMeaningInstance:
		return "Instance"
	case ConsentDataMeaningRelated:
		return "Related"
	case ConsentDataMeaningDependents:
		return "Dependents"
	case ConsentDataMeaningAuthoredby:
		return "AuthoredBy"
	}
	return "<unknown>"
}
func (code ConsentDataMeaning) Definition() string {
	switch code {
	case ConsentDataMeaningInstance:
		return "The consent applies directly to the instance of the resource."
	case ConsentDataMeaningRelated:
		return "The consent applies directly to the instance of the resource and instances it refers to."
	case ConsentDataMeaningDependents:
		return "The consent applies directly to the instance of the resource and instances that refer to it."
	case ConsentDataMeaningAuthoredby:
		return "The consent applies to instances of resources that are authored by."
	}
	return "<unknown>"
}
