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
// 2021-03-08 09:41:03.015326 +0000 UTC

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GuideParameterCode is documented here http://hl7.org/fhir/ValueSet/guide-parameter-code
type GuideParameterCode int

const (
	GuideParameterCodeApply GuideParameterCode = iota
	GuideParameterCodePathResource
	GuideParameterCodePathPages
	GuideParameterCodePathTxCache
	GuideParameterCodeExpansionParameter
	GuideParameterCodeRuleBrokenLinks
	GuideParameterCodeGenerateXml
	GuideParameterCodeGenerateJson
	GuideParameterCodeGenerateTurtle
	GuideParameterCodeHtmlTemplate
)

func (code GuideParameterCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *GuideParameterCode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "apply":
		*code = GuideParameterCodeApply
	case "path-resource":
		*code = GuideParameterCodePathResource
	case "path-pages":
		*code = GuideParameterCodePathPages
	case "path-tx-cache":
		*code = GuideParameterCodePathTxCache
	case "expansion-parameter":
		*code = GuideParameterCodeExpansionParameter
	case "rule-broken-links":
		*code = GuideParameterCodeRuleBrokenLinks
	case "generate-xml":
		*code = GuideParameterCodeGenerateXml
	case "generate-json":
		*code = GuideParameterCodeGenerateJson
	case "generate-turtle":
		*code = GuideParameterCodeGenerateTurtle
	case "html-template":
		*code = GuideParameterCodeHtmlTemplate
	default:
		return fmt.Errorf("unknown GuideParameterCode code `%s`", s)
	}
	return nil
}
func (code GuideParameterCode) String() string {
	return code.Code()
}
func (code GuideParameterCode) Code() string {
	switch code {
	case GuideParameterCodeApply:
		return "apply"
	case GuideParameterCodePathResource:
		return "path-resource"
	case GuideParameterCodePathPages:
		return "path-pages"
	case GuideParameterCodePathTxCache:
		return "path-tx-cache"
	case GuideParameterCodeExpansionParameter:
		return "expansion-parameter"
	case GuideParameterCodeRuleBrokenLinks:
		return "rule-broken-links"
	case GuideParameterCodeGenerateXml:
		return "generate-xml"
	case GuideParameterCodeGenerateJson:
		return "generate-json"
	case GuideParameterCodeGenerateTurtle:
		return "generate-turtle"
	case GuideParameterCodeHtmlTemplate:
		return "html-template"
	}
	return "<unknown>"
}
func (code GuideParameterCode) Display() string {
	switch code {
	case GuideParameterCodeApply:
		return "Apply Metadata Value"
	case GuideParameterCodePathResource:
		return "Resource Path"
	case GuideParameterCodePathPages:
		return "Pages Path"
	case GuideParameterCodePathTxCache:
		return "Terminology Cache Path"
	case GuideParameterCodeExpansionParameter:
		return "Expansion Profile"
	case GuideParameterCodeRuleBrokenLinks:
		return "Broken Links Rule"
	case GuideParameterCodeGenerateXml:
		return "Generate XML"
	case GuideParameterCodeGenerateJson:
		return "Generate JSON"
	case GuideParameterCodeGenerateTurtle:
		return "Generate Turtle"
	case GuideParameterCodeHtmlTemplate:
		return "HTML Template"
	}
	return "<unknown>"
}
func (code GuideParameterCode) Definition() string {
	switch code {
	case GuideParameterCodeApply:
		return "If the value of this string 0..* parameter is one of the metadata fields then all conformance resources will have any specified [Resource].[field] overwritten with the ImplementationGuide.[field], where field is one of: version, date, status, publisher, contact, copyright, experimental, jurisdiction, useContext."
	case GuideParameterCodePathResource:
		return "The value of this string 0..* parameter is a subfolder of the build context's location that is to be scanned to load resources. Scope is (if present) a particular resource type."
	case GuideParameterCodePathPages:
		return "The value of this string 0..1 parameter is a subfolder of the build context's location that contains files that are part of the html content processed by the builder."
	case GuideParameterCodePathTxCache:
		return "The value of this string 0..1 parameter is a subfolder of the build context's location that is used as the terminology cache. If this is not present, the terminology cache is on the local system, not under version control."
	case GuideParameterCodeExpansionParameter:
		return "The value of this string 0..* parameter is a parameter (name=value) when expanding value sets for this implementation guide. This is particularly used to specify the versions of published terminologies such as SNOMED CT."
	case GuideParameterCodeRuleBrokenLinks:
		return "The value of this string 0..1 parameter is either \"warning\" or \"error\" (default = \"error\"). If the value is \"warning\" then IG build tools allow the IG to be considered successfully build even when there is no internal broken links."
	case GuideParameterCodeGenerateXml:
		return "The value of this boolean 0..1 parameter specifies whether the IG publisher creates examples in XML format. If not present, the Publication Tool decides whether to generate XML."
	case GuideParameterCodeGenerateJson:
		return "The value of this boolean 0..1 parameter specifies whether the IG publisher creates examples in JSON format. If not present, the Publication Tool decides whether to generate JSON."
	case GuideParameterCodeGenerateTurtle:
		return "The value of this boolean 0..1 parameter specifies whether the IG publisher creates examples in Turtle format. If not present, the Publication Tool decides whether to generate Turtle."
	case GuideParameterCodeHtmlTemplate:
		return "The value of this string singleton parameter is the name of the file to use as the builder template for each generated page (see templating)."
	}
	return "<unknown>"
}
