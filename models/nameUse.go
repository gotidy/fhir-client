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
// 2021-03-03 10:55:45.351201 +0000 UTC
// NameUse is documented here http://hl7.org/fhir/ValueSet/name-use
type NameUse int

const (
	NameUseUsual NameUse = iota
	NameUseOfficial
	NameUseTemp
	NameUseNickname
	NameUseAnonymous
	NameUseOld
	NameUseMaiden
)

func (code NameUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *NameUse) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "usual":
		*code = NameUseUsual
	case "official":
		*code = NameUseOfficial
	case "temp":
		*code = NameUseTemp
	case "nickname":
		*code = NameUseNickname
	case "anonymous":
		*code = NameUseAnonymous
	case "old":
		*code = NameUseOld
	case "maiden":
		*code = NameUseMaiden
	default:
		return fmt.Errorf("unknown NameUse code `%s`", s)
	}
	return nil
}
func (code NameUse) String() string {
	return code.Code()
}
func (code NameUse) Code() string {
	switch code {
	case NameUseUsual:
		return "usual"
	case NameUseOfficial:
		return "official"
	case NameUseTemp:
		return "temp"
	case NameUseNickname:
		return "nickname"
	case NameUseAnonymous:
		return "anonymous"
	case NameUseOld:
		return "old"
	case NameUseMaiden:
		return "maiden"
	}
	return "<unknown>"
}
func (code NameUse) Display() string {
	switch code {
	case NameUseUsual:
		return "Usual"
	case NameUseOfficial:
		return "Official"
	case NameUseTemp:
		return "Temp"
	case NameUseNickname:
		return "Nickname"
	case NameUseAnonymous:
		return "Anonymous"
	case NameUseOld:
		return "Old"
	case NameUseMaiden:
		return "Name changed for Marriage"
	}
	return "<unknown>"
}
func (code NameUse) Definition() string {
	switch code {
	case NameUseUsual:
		return "Known as/conventional/the one you normally use."
	case NameUseOfficial:
		return "The formal name as registered in an official (government) registry, but which name might not be commonly used. May be called \"legal name\"."
	case NameUseTemp:
		return "A temporary name. Name.period can provide more detailed information. This may also be used for temporary names assigned at birth or in emergency situations."
	case NameUseNickname:
		return "A name that is used to address the person in an informal manner, but is not part of their formal or usual name."
	case NameUseAnonymous:
		return "Anonymous assigned name, alias, or pseudonym (used to protect a person's identity for privacy reasons)."
	case NameUseOld:
		return "This name is no longer in use (or was never correct, but retained for records)."
	case NameUseMaiden:
		return "A name used prior to changing name because of marriage. This name use is for use by applications that collect and store names that were used prior to a marriage. Marriage naming customs vary greatly around the world, and are constantly changing. This term is not gender specific. The use of this term does not imply any particular history for a person's name."
	}
	return "<unknown>"
}
