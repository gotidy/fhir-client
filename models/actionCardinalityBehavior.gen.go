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
// 2021-03-03 16:38:34.970702 +0000 UTC

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ActionCardinalityBehavior is documented here http://hl7.org/fhir/ValueSet/action-cardinality-behavior
type ActionCardinalityBehavior int

const (
	ActionCardinalityBehaviorSingle ActionCardinalityBehavior = iota
	ActionCardinalityBehaviorMultiple
)

func (code ActionCardinalityBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionCardinalityBehavior) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "single":
		*code = ActionCardinalityBehaviorSingle
	case "multiple":
		*code = ActionCardinalityBehaviorMultiple
	default:
		return fmt.Errorf("unknown ActionCardinalityBehavior code `%s`", s)
	}
	return nil
}
func (code ActionCardinalityBehavior) String() string {
	return code.Code()
}
func (code ActionCardinalityBehavior) Code() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "single"
	case ActionCardinalityBehaviorMultiple:
		return "multiple"
	}
	return "<unknown>"
}
func (code ActionCardinalityBehavior) Display() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "Single"
	case ActionCardinalityBehaviorMultiple:
		return "Multiple"
	}
	return "<unknown>"
}
func (code ActionCardinalityBehavior) Definition() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "The action may only be selected one time."
	case ActionCardinalityBehaviorMultiple:
		return "The action may be selected multiple times."
	}
	return "<unknown>"
}