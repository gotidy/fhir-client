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
// 2021-03-03 10:55:39.496558 +0000 UTC
// CompositionAttestationMode is documented here http://hl7.org/fhir/ValueSet/composition-attestation-mode
type CompositionAttestationMode int

const (
	CompositionAttestationModePersonal CompositionAttestationMode = iota
	CompositionAttestationModeProfessional
	CompositionAttestationModeLegal
	CompositionAttestationModeOfficial
)

func (code CompositionAttestationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CompositionAttestationMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "personal":
		*code = CompositionAttestationModePersonal
	case "professional":
		*code = CompositionAttestationModeProfessional
	case "legal":
		*code = CompositionAttestationModeLegal
	case "official":
		*code = CompositionAttestationModeOfficial
	default:
		return fmt.Errorf("unknown CompositionAttestationMode code `%s`", s)
	}
	return nil
}
func (code CompositionAttestationMode) String() string {
	return code.Code()
}
func (code CompositionAttestationMode) Code() string {
	switch code {
	case CompositionAttestationModePersonal:
		return "personal"
	case CompositionAttestationModeProfessional:
		return "professional"
	case CompositionAttestationModeLegal:
		return "legal"
	case CompositionAttestationModeOfficial:
		return "official"
	}
	return "<unknown>"
}
func (code CompositionAttestationMode) Display() string {
	switch code {
	case CompositionAttestationModePersonal:
		return "Personal"
	case CompositionAttestationModeProfessional:
		return "Professional"
	case CompositionAttestationModeLegal:
		return "Legal"
	case CompositionAttestationModeOfficial:
		return "Official"
	}
	return "<unknown>"
}
func (code CompositionAttestationMode) Definition() string {
	switch code {
	case CompositionAttestationModePersonal:
		return "The person authenticated the content in their personal capacity."
	case CompositionAttestationModeProfessional:
		return "The person authenticated the content in their professional capacity."
	case CompositionAttestationModeLegal:
		return "The person authenticated the content and accepted legal responsibility for its content."
	case CompositionAttestationModeOfficial:
		return "The organization authenticated the content as consistent with their policies and procedures."
	}
	return "<unknown>"
}
