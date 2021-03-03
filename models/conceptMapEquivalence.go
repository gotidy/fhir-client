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
// 2021-03-03 10:55:45.311445 +0000 UTC
// ConceptMapEquivalence is documented here http://hl7.org/fhir/ValueSet/concept-map-equivalence
type ConceptMapEquivalence int

const (
	ConceptMapEquivalenceRelatedto ConceptMapEquivalence = iota
	ConceptMapEquivalenceEquivalent
	ConceptMapEquivalenceEqual
	ConceptMapEquivalenceWider
	ConceptMapEquivalenceSubsumes
	ConceptMapEquivalenceNarrower
	ConceptMapEquivalenceSpecializes
	ConceptMapEquivalenceInexact
	ConceptMapEquivalenceUnmatched
	ConceptMapEquivalenceDisjoint
)

func (code ConceptMapEquivalence) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConceptMapEquivalence) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "relatedto":
		*code = ConceptMapEquivalenceRelatedto
	case "equivalent":
		*code = ConceptMapEquivalenceEquivalent
	case "equal":
		*code = ConceptMapEquivalenceEqual
	case "wider":
		*code = ConceptMapEquivalenceWider
	case "subsumes":
		*code = ConceptMapEquivalenceSubsumes
	case "narrower":
		*code = ConceptMapEquivalenceNarrower
	case "specializes":
		*code = ConceptMapEquivalenceSpecializes
	case "inexact":
		*code = ConceptMapEquivalenceInexact
	case "unmatched":
		*code = ConceptMapEquivalenceUnmatched
	case "disjoint":
		*code = ConceptMapEquivalenceDisjoint
	default:
		return fmt.Errorf("unknown ConceptMapEquivalence code `%s`", s)
	}
	return nil
}
func (code ConceptMapEquivalence) String() string {
	return code.Code()
}
func (code ConceptMapEquivalence) Code() string {
	switch code {
	case ConceptMapEquivalenceRelatedto:
		return "relatedto"
	case ConceptMapEquivalenceEquivalent:
		return "equivalent"
	case ConceptMapEquivalenceEqual:
		return "equal"
	case ConceptMapEquivalenceWider:
		return "wider"
	case ConceptMapEquivalenceSubsumes:
		return "subsumes"
	case ConceptMapEquivalenceNarrower:
		return "narrower"
	case ConceptMapEquivalenceSpecializes:
		return "specializes"
	case ConceptMapEquivalenceInexact:
		return "inexact"
	case ConceptMapEquivalenceUnmatched:
		return "unmatched"
	case ConceptMapEquivalenceDisjoint:
		return "disjoint"
	}
	return "<unknown>"
}
func (code ConceptMapEquivalence) Display() string {
	switch code {
	case ConceptMapEquivalenceRelatedto:
		return "Related To"
	case ConceptMapEquivalenceEquivalent:
		return "Equivalent"
	case ConceptMapEquivalenceEqual:
		return "Equal"
	case ConceptMapEquivalenceWider:
		return "Wider"
	case ConceptMapEquivalenceSubsumes:
		return "Subsumes"
	case ConceptMapEquivalenceNarrower:
		return "Narrower"
	case ConceptMapEquivalenceSpecializes:
		return "Specializes"
	case ConceptMapEquivalenceInexact:
		return "Inexact"
	case ConceptMapEquivalenceUnmatched:
		return "Unmatched"
	case ConceptMapEquivalenceDisjoint:
		return "Disjoint"
	}
	return "<unknown>"
}
func (code ConceptMapEquivalence) Definition() string {
	switch code {
	case ConceptMapEquivalenceRelatedto:
		return "The concepts are related to each other, and have at least some overlap in meaning, but the exact relationship is not known."
	case ConceptMapEquivalenceEquivalent:
		return "The definitions of the concepts mean the same thing (including when structural implications of meaning are considered) (i.e. extensionally identical)."
	case ConceptMapEquivalenceEqual:
		return "The definitions of the concepts are exactly the same (i.e. only grammatical differences) and structural implications of meaning are identical or irrelevant (i.e. intentionally identical)."
	case ConceptMapEquivalenceWider:
		return "The target mapping is wider in meaning than the source concept."
	case ConceptMapEquivalenceSubsumes:
		return "The target mapping subsumes the meaning of the source concept (e.g. the source is-a target)."
	case ConceptMapEquivalenceNarrower:
		return "The target mapping is narrower in meaning than the source concept. The sense in which the mapping is narrower SHALL be described in the comments in this case, and applications should be careful when attempting to use these mappings operationally."
	case ConceptMapEquivalenceSpecializes:
		return "The target mapping specializes the meaning of the source concept (e.g. the target is-a source)."
	case ConceptMapEquivalenceInexact:
		return "The target mapping overlaps with the source concept, but both source and target cover additional meaning, or the definitions are imprecise and it is uncertain whether they have the same boundaries to their meaning. The sense in which the mapping is inexact SHALL be described in the comments in this case, and applications should be careful when attempting to use these mappings operationally."
	case ConceptMapEquivalenceUnmatched:
		return "There is no match for this concept in the target code system."
	case ConceptMapEquivalenceDisjoint:
		return "This is an explicit assertion that there is no mapping between the source and target concept."
	}
	return "<unknown>"
}
