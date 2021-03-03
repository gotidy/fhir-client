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
// 2021-03-03 10:55:45.41085 +0000 UTC
// MessageSignificanceCategory is documented here http://hl7.org/fhir/ValueSet/message-significance-category
type MessageSignificanceCategory int

const (
	MessageSignificanceCategoryConsequence MessageSignificanceCategory = iota
	MessageSignificanceCategoryCurrency
	MessageSignificanceCategoryNotification
)

func (code MessageSignificanceCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MessageSignificanceCategory) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "consequence":
		*code = MessageSignificanceCategoryConsequence
	case "currency":
		*code = MessageSignificanceCategoryCurrency
	case "notification":
		*code = MessageSignificanceCategoryNotification
	default:
		return fmt.Errorf("unknown MessageSignificanceCategory code `%s`", s)
	}
	return nil
}
func (code MessageSignificanceCategory) String() string {
	return code.Code()
}
func (code MessageSignificanceCategory) Code() string {
	switch code {
	case MessageSignificanceCategoryConsequence:
		return "consequence"
	case MessageSignificanceCategoryCurrency:
		return "currency"
	case MessageSignificanceCategoryNotification:
		return "notification"
	}
	return "<unknown>"
}
func (code MessageSignificanceCategory) Display() string {
	switch code {
	case MessageSignificanceCategoryConsequence:
		return "Consequence"
	case MessageSignificanceCategoryCurrency:
		return "Currency"
	case MessageSignificanceCategoryNotification:
		return "Notification"
	}
	return "<unknown>"
}
func (code MessageSignificanceCategory) Definition() string {
	switch code {
	case MessageSignificanceCategoryConsequence:
		return "The message represents/requests a change that should not be processed more than once; e.g., making a booking for an appointment."
	case MessageSignificanceCategoryCurrency:
		return "The message represents a response to query for current information. Retrospective processing is wrong and/or wasteful."
	case MessageSignificanceCategoryNotification:
		return "The content is not necessarily intended to be current, and it can be reprocessed, though there may be version issues created by processing old notifications."
	}
	return "<unknown>"
}
