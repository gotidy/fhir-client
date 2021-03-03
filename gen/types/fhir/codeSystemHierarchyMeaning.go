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
// 2021-03-03 10:55:39.491651 +0000 UTC
// CodeSystemHierarchyMeaning is documented here http://hl7.org/fhir/ValueSet/codesystem-hierarchy-meaning
type CodeSystemHierarchyMeaning int

const (
	CodeSystemHierarchyMeaningGroupedBy CodeSystemHierarchyMeaning = iota
	CodeSystemHierarchyMeaningIsA
	CodeSystemHierarchyMeaningPartOf
	CodeSystemHierarchyMeaningClassifiedWith
)

func (code CodeSystemHierarchyMeaning) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CodeSystemHierarchyMeaning) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "grouped-by":
		*code = CodeSystemHierarchyMeaningGroupedBy
	case "is-a":
		*code = CodeSystemHierarchyMeaningIsA
	case "part-of":
		*code = CodeSystemHierarchyMeaningPartOf
	case "classified-with":
		*code = CodeSystemHierarchyMeaningClassifiedWith
	default:
		return fmt.Errorf("unknown CodeSystemHierarchyMeaning code `%s`", s)
	}
	return nil
}
func (code CodeSystemHierarchyMeaning) String() string {
	return code.Code()
}
func (code CodeSystemHierarchyMeaning) Code() string {
	switch code {
	case CodeSystemHierarchyMeaningGroupedBy:
		return "grouped-by"
	case CodeSystemHierarchyMeaningIsA:
		return "is-a"
	case CodeSystemHierarchyMeaningPartOf:
		return "part-of"
	case CodeSystemHierarchyMeaningClassifiedWith:
		return "classified-with"
	}
	return "<unknown>"
}
func (code CodeSystemHierarchyMeaning) Display() string {
	switch code {
	case CodeSystemHierarchyMeaningGroupedBy:
		return "Grouped By"
	case CodeSystemHierarchyMeaningIsA:
		return "Is-A"
	case CodeSystemHierarchyMeaningPartOf:
		return "Part Of"
	case CodeSystemHierarchyMeaningClassifiedWith:
		return "Classified With"
	}
	return "<unknown>"
}
func (code CodeSystemHierarchyMeaning) Definition() string {
	switch code {
	case CodeSystemHierarchyMeaningGroupedBy:
		return "No particular relationship between the concepts can be assumed, except what can be determined by inspection of the definitions of the elements (possible reasons to use this: importing from a source where this is not defined, or where various parts of the hierarchy have different meanings)."
	case CodeSystemHierarchyMeaningIsA:
		return "A hierarchy where the child concepts have an IS-A relationship with the parents - that is, all the properties of the parent are also true for its child concepts. Not that is-a is a property of the concepts, so additional subsumption relationships may be defined using properties or the [subsumes](extension-codesystem-subsumes.html) extension."
	case CodeSystemHierarchyMeaningPartOf:
		return "Child elements list the individual parts of a composite whole (e.g. body site)."
	case CodeSystemHierarchyMeaningClassifiedWith:
		return "Child concepts in the hierarchy may have only one parent, and there is a presumption that the code system is a \"closed world\" meaning all things must be in the hierarchy. This results in concepts such as \"not otherwise classified.\"."
	}
	return "<unknown>"
}
