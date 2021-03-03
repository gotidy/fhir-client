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
// 2021-03-03 10:55:39.573357 +0000 UTC
// DeviceMetricOperationalStatus is documented here http://hl7.org/fhir/ValueSet/metric-operational-status
type DeviceMetricOperationalStatus int

const (
	DeviceMetricOperationalStatusOn DeviceMetricOperationalStatus = iota
	DeviceMetricOperationalStatusOff
	DeviceMetricOperationalStatusStandby
	DeviceMetricOperationalStatusEnteredInError
)

func (code DeviceMetricOperationalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceMetricOperationalStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "on":
		*code = DeviceMetricOperationalStatusOn
	case "off":
		*code = DeviceMetricOperationalStatusOff
	case "standby":
		*code = DeviceMetricOperationalStatusStandby
	case "entered-in-error":
		*code = DeviceMetricOperationalStatusEnteredInError
	default:
		return fmt.Errorf("unknown DeviceMetricOperationalStatus code `%s`", s)
	}
	return nil
}
func (code DeviceMetricOperationalStatus) String() string {
	return code.Code()
}
func (code DeviceMetricOperationalStatus) Code() string {
	switch code {
	case DeviceMetricOperationalStatusOn:
		return "on"
	case DeviceMetricOperationalStatusOff:
		return "off"
	case DeviceMetricOperationalStatusStandby:
		return "standby"
	case DeviceMetricOperationalStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code DeviceMetricOperationalStatus) Display() string {
	switch code {
	case DeviceMetricOperationalStatusOn:
		return "On"
	case DeviceMetricOperationalStatusOff:
		return "Off"
	case DeviceMetricOperationalStatusStandby:
		return "Standby"
	case DeviceMetricOperationalStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code DeviceMetricOperationalStatus) Definition() string {
	switch code {
	case DeviceMetricOperationalStatusOn:
		return "The DeviceMetric is operating and will generate DeviceObservations."
	case DeviceMetricOperationalStatusOff:
		return "The DeviceMetric is not operating."
	case DeviceMetricOperationalStatusStandby:
		return "The DeviceMetric is operating, but will not generate any DeviceObservations."
	case DeviceMetricOperationalStatusEnteredInError:
		return "The DeviceMetric was entered in error."
	}
	return "<unknown>"
}
