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
// 2021-03-03 10:55:39.542354 +0000 UTC
// DocumentMode is documented here http://hl7.org/fhir/ValueSet/document-mode
type DocumentMode int

const (
	DocumentModeProducer DocumentMode = iota
	DocumentModeConsumer
)

func (code DocumentMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DocumentMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "producer":
		*code = DocumentModeProducer
	case "consumer":
		*code = DocumentModeConsumer
	default:
		return fmt.Errorf("unknown DocumentMode code `%s`", s)
	}
	return nil
}
func (code DocumentMode) String() string {
	return code.Code()
}
func (code DocumentMode) Code() string {
	switch code {
	case DocumentModeProducer:
		return "producer"
	case DocumentModeConsumer:
		return "consumer"
	}
	return "<unknown>"
}
func (code DocumentMode) Display() string {
	switch code {
	case DocumentModeProducer:
		return "Producer"
	case DocumentModeConsumer:
		return "Consumer"
	}
	return "<unknown>"
}
func (code DocumentMode) Definition() string {
	switch code {
	case DocumentModeProducer:
		return "The application produces documents of the specified type."
	case DocumentModeConsumer:
		return "The application consumes documents of the specified type."
	}
	return "<unknown>"
}
