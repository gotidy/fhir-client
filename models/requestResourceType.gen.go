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
// 2021-03-08 09:41:03.264992 +0000 UTC

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// RequestResourceType is documented here http://hl7.org/fhir/ValueSet/request-resource-types
type RequestResourceType int

const (
	RequestResourceTypeAppointment RequestResourceType = iota
	RequestResourceTypeAppointmentResponse
	RequestResourceTypeCarePlan
	RequestResourceTypeClaim
	RequestResourceTypeCommunicationRequest
	RequestResourceTypeContract
	RequestResourceTypeDeviceRequest
	RequestResourceTypeEnrollmentRequest
	RequestResourceTypeImmunizationRecommendation
	RequestResourceTypeMedicationRequest
	RequestResourceTypeNutritionOrder
	RequestResourceTypeServiceRequest
	RequestResourceTypeSupplyRequest
	RequestResourceTypeTask
	RequestResourceTypeVisionPrescription
)

func (code RequestResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *RequestResourceType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "Appointment":
		*code = RequestResourceTypeAppointment
	case "AppointmentResponse":
		*code = RequestResourceTypeAppointmentResponse
	case "CarePlan":
		*code = RequestResourceTypeCarePlan
	case "Claim":
		*code = RequestResourceTypeClaim
	case "CommunicationRequest":
		*code = RequestResourceTypeCommunicationRequest
	case "Contract":
		*code = RequestResourceTypeContract
	case "DeviceRequest":
		*code = RequestResourceTypeDeviceRequest
	case "EnrollmentRequest":
		*code = RequestResourceTypeEnrollmentRequest
	case "ImmunizationRecommendation":
		*code = RequestResourceTypeImmunizationRecommendation
	case "MedicationRequest":
		*code = RequestResourceTypeMedicationRequest
	case "NutritionOrder":
		*code = RequestResourceTypeNutritionOrder
	case "ServiceRequest":
		*code = RequestResourceTypeServiceRequest
	case "SupplyRequest":
		*code = RequestResourceTypeSupplyRequest
	case "Task":
		*code = RequestResourceTypeTask
	case "VisionPrescription":
		*code = RequestResourceTypeVisionPrescription
	default:
		return fmt.Errorf("unknown RequestResourceType code `%s`", s)
	}
	return nil
}
func (code RequestResourceType) String() string {
	return code.Code()
}
func (code RequestResourceType) Code() string {
	switch code {
	case RequestResourceTypeAppointment:
		return "Appointment"
	case RequestResourceTypeAppointmentResponse:
		return "AppointmentResponse"
	case RequestResourceTypeCarePlan:
		return "CarePlan"
	case RequestResourceTypeClaim:
		return "Claim"
	case RequestResourceTypeCommunicationRequest:
		return "CommunicationRequest"
	case RequestResourceTypeContract:
		return "Contract"
	case RequestResourceTypeDeviceRequest:
		return "DeviceRequest"
	case RequestResourceTypeEnrollmentRequest:
		return "EnrollmentRequest"
	case RequestResourceTypeImmunizationRecommendation:
		return "ImmunizationRecommendation"
	case RequestResourceTypeMedicationRequest:
		return "MedicationRequest"
	case RequestResourceTypeNutritionOrder:
		return "NutritionOrder"
	case RequestResourceTypeServiceRequest:
		return "ServiceRequest"
	case RequestResourceTypeSupplyRequest:
		return "SupplyRequest"
	case RequestResourceTypeTask:
		return "Task"
	case RequestResourceTypeVisionPrescription:
		return "VisionPrescription"
	}
	return "<unknown>"
}
func (code RequestResourceType) Display() string {
	switch code {
	case RequestResourceTypeAppointment:
		return "Appointment"
	case RequestResourceTypeAppointmentResponse:
		return "AppointmentResponse"
	case RequestResourceTypeCarePlan:
		return "CarePlan"
	case RequestResourceTypeClaim:
		return "Claim"
	case RequestResourceTypeCommunicationRequest:
		return "CommunicationRequest"
	case RequestResourceTypeContract:
		return "Contract"
	case RequestResourceTypeDeviceRequest:
		return "DeviceRequest"
	case RequestResourceTypeEnrollmentRequest:
		return "EnrollmentRequest"
	case RequestResourceTypeImmunizationRecommendation:
		return "ImmunizationRecommendation"
	case RequestResourceTypeMedicationRequest:
		return "MedicationRequest"
	case RequestResourceTypeNutritionOrder:
		return "NutritionOrder"
	case RequestResourceTypeServiceRequest:
		return "ServiceRequest"
	case RequestResourceTypeSupplyRequest:
		return "SupplyRequest"
	case RequestResourceTypeTask:
		return "Task"
	case RequestResourceTypeVisionPrescription:
		return "VisionPrescription"
	}
	return "<unknown>"
}
func (code RequestResourceType) Definition() string {
	switch code {
	case RequestResourceTypeAppointment:
		return "A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time. This may result in one or more Encounter(s)."
	case RequestResourceTypeAppointmentResponse:
		return "A reply to an appointment request for a patient and/or practitioner(s), such as a confirmation or rejection."
	case RequestResourceTypeCarePlan:
		return "Healthcare plan for patient or group."
	case RequestResourceTypeClaim:
		return "Claim, Pre-determination or Pre-authorization."
	case RequestResourceTypeCommunicationRequest:
		return "A request for information to be sent to a receiver."
	case RequestResourceTypeContract:
		return "Legal Agreement."
	case RequestResourceTypeDeviceRequest:
		return "Medical device request."
	case RequestResourceTypeEnrollmentRequest:
		return "Enrollment request."
	case RequestResourceTypeImmunizationRecommendation:
		return "Guidance or advice relating to an immunization."
	case RequestResourceTypeMedicationRequest:
		return "Ordering of medication for patient or group."
	case RequestResourceTypeNutritionOrder:
		return "Diet, formula or nutritional supplement request."
	case RequestResourceTypeServiceRequest:
		return "A record of a request for service such as diagnostic investigations, treatments, or operations to be performed."
	case RequestResourceTypeSupplyRequest:
		return "Request for a medication, substance or device."
	case RequestResourceTypeTask:
		return "A task to be performed."
	case RequestResourceTypeVisionPrescription:
		return "Prescription for vision correction products for a patient."
	}
	return "<unknown>"
}
