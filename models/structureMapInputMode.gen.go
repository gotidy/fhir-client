// Copyright 2021 Evgeny Safonov
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
// 2021-03-03 16:38:34.78337 +0000 UTC

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StructureMapInputMode is documented here http://hl7.org/fhir/ValueSet/map-input-mode
type StructureMapInputMode int

const (
	StructureMapInputModeSource StructureMapInputMode = iota
	StructureMapInputModeTarget
)

func (code StructureMapInputMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapInputMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "source":
		*code = StructureMapInputModeSource
	case "target":
		*code = StructureMapInputModeTarget
	default:
		return fmt.Errorf("unknown StructureMapInputMode code `%s`", s)
	}
	return nil
}
func (code StructureMapInputMode) String() string {
	return code.Code()
}
func (code StructureMapInputMode) Code() string {
	switch code {
	case StructureMapInputModeSource:
		return "source"
	case StructureMapInputModeTarget:
		return "target"
	}
	return "<unknown>"
}
func (code StructureMapInputMode) Display() string {
	switch code {
	case StructureMapInputModeSource:
		return "Source Instance"
	case StructureMapInputModeTarget:
		return "Target Instance"
	}
	return "<unknown>"
}
func (code StructureMapInputMode) Definition() string {
	switch code {
	case StructureMapInputModeSource:
		return "Names an input instance used a source for mapping."
	case StructureMapInputModeTarget:
		return "Names an instance that is being populated."
	}
	return "<unknown>"
}