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

import "encoding/json"

// Code generated by go generate; DO NOT EDIT.

// This file was generated by robots at
// 2021-03-03 10:55:44.842639 +0000 UTC
// VerificationResult is documented here http://hl7.org/fhir/StructureDefinition/VerificationResult
type VerificationResult struct {
	ID                *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Target            []Reference                       `bson:"target,omitempty" json:"target,omitempty"`
	TargetLocation    []string                          `bson:"targetLocation,omitempty" json:"targetLocation,omitempty"`
	Need              *CodeableConcept                  `bson:"need,omitempty" json:"need,omitempty"`
	Status            string                            `bson:"status" json:"status"`
	StatusDate        *string                           `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	ValidationType    *CodeableConcept                  `bson:"validationType,omitempty" json:"validationType,omitempty"`
	ValidationProcess []CodeableConcept                 `bson:"validationProcess,omitempty" json:"validationProcess,omitempty"`
	Frequency         *Timing                           `bson:"frequency,omitempty" json:"frequency,omitempty"`
	LastPerformed     *string                           `bson:"lastPerformed,omitempty" json:"lastPerformed,omitempty"`
	NextScheduled     *string                           `bson:"nextScheduled,omitempty" json:"nextScheduled,omitempty"`
	FailureAction     *CodeableConcept                  `bson:"failureAction,omitempty" json:"failureAction,omitempty"`
	PrimarySource     []VerificationResultPrimarySource `bson:"primarySource,omitempty" json:"primarySource,omitempty"`
	Attestation       *VerificationResultAttestation    `bson:"attestation,omitempty" json:"attestation,omitempty"`
	Validator         []VerificationResultValidator     `bson:"validator,omitempty" json:"validator,omitempty"`
}
type VerificationResultPrimarySource struct {
	ID                  *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Who                 *Reference        `bson:"who,omitempty" json:"who,omitempty"`
	Type                []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	CommunicationMethod []CodeableConcept `bson:"communicationMethod,omitempty" json:"communicationMethod,omitempty"`
	ValidationStatus    *CodeableConcept  `bson:"validationStatus,omitempty" json:"validationStatus,omitempty"`
	ValidationDate      *string           `bson:"validationDate,omitempty" json:"validationDate,omitempty"`
	CanPushUpdates      *CodeableConcept  `bson:"canPushUpdates,omitempty" json:"canPushUpdates,omitempty"`
	PushTypeAvailable   []CodeableConcept `bson:"pushTypeAvailable,omitempty" json:"pushTypeAvailable,omitempty"`
}
type VerificationResultAttestation struct {
	ID                        *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Who                       *Reference       `bson:"who,omitempty" json:"who,omitempty"`
	OnBehalfOf                *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	CommunicationMethod       *CodeableConcept `bson:"communicationMethod,omitempty" json:"communicationMethod,omitempty"`
	Date                      *string          `bson:"date,omitempty" json:"date,omitempty"`
	SourceIdentityCertificate *string          `bson:"sourceIdentityCertificate,omitempty" json:"sourceIdentityCertificate,omitempty"`
	ProxyIdentityCertificate  *string          `bson:"proxyIdentityCertificate,omitempty" json:"proxyIdentityCertificate,omitempty"`
	ProxySignature            *Signature       `bson:"proxySignature,omitempty" json:"proxySignature,omitempty"`
	SourceSignature           *Signature       `bson:"sourceSignature,omitempty" json:"sourceSignature,omitempty"`
}
type VerificationResultValidator struct {
	ID                   *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Organization         Reference   `bson:"organization" json:"organization"`
	IdentityCertificate  *string     `bson:"identityCertificate,omitempty" json:"identityCertificate,omitempty"`
	AttestationSignature *Signature  `bson:"attestationSignature,omitempty" json:"attestationSignature,omitempty"`
}
type OtherVerificationResult VerificationResult

// MarshalJSON marshals the given VerificationResult as JSON into a byte slice
func (r VerificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVerificationResult
		ResourceType string `json:"resourceType"`
	}{
		OtherVerificationResult: OtherVerificationResult(r),
		ResourceType:            "VerificationResult",
	})
}

// UnmarshalVerificationResult unmarshals a VerificationResult.
func UnmarshalVerificationResult(b []byte) (VerificationResult, error) {
	var verificationResult VerificationResult
	if err := json.Unmarshal(b, &verificationResult); err != nil {
		return verificationResult, err
	}
	return verificationResult, nil
}
