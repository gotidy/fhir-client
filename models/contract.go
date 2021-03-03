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
// 2021-03-03 10:55:44.230478 +0000 UTC
// Contract is documented here http://hl7.org/fhir/StructureDefinition/Contract
type Contract struct {
	ID                    *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Url                   *string                      `bson:"url,omitempty" json:"url,omitempty"`
	Version               *string                      `bson:"version,omitempty" json:"version,omitempty"`
	Status                *ContractResourceStatusCodes `bson:"status,omitempty" json:"status,omitempty"`
	LegalState            *CodeableConcept             `bson:"legalState,omitempty" json:"legalState,omitempty"`
	InstantiatesCanonical *Reference                   `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       *string                      `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	ContentDerivative     *CodeableConcept             `bson:"contentDerivative,omitempty" json:"contentDerivative,omitempty"`
	Issued                *string                      `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies               *Period                      `bson:"applies,omitempty" json:"applies,omitempty"`
	ExpirationType        *CodeableConcept             `bson:"expirationType,omitempty" json:"expirationType,omitempty"`
	Subject               []Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Authority             []Reference                  `bson:"authority,omitempty" json:"authority,omitempty"`
	Domain                []Reference                  `bson:"domain,omitempty" json:"domain,omitempty"`
	Site                  []Reference                  `bson:"site,omitempty" json:"site,omitempty"`
	Name                  *string                      `bson:"name,omitempty" json:"name,omitempty"`
	Title                 *string                      `bson:"title,omitempty" json:"title,omitempty"`
	Subtitle              *string                      `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	Alias                 []string                     `bson:"alias,omitempty" json:"alias,omitempty"`
	Author                *Reference                   `bson:"author,omitempty" json:"author,omitempty"`
	Scope                 *CodeableConcept             `bson:"scope,omitempty" json:"scope,omitempty"`
	Type                  *CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	SubType               []CodeableConcept            `bson:"subType,omitempty" json:"subType,omitempty"`
	ContentDefinition     *ContractContentDefinition   `bson:"contentDefinition,omitempty" json:"contentDefinition,omitempty"`
	Term                  []ContractTerm               `bson:"term,omitempty" json:"term,omitempty"`
	SupportingInfo        []Reference                  `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	RelevantHistory       []Reference                  `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
	Signer                []ContractSigner             `bson:"signer,omitempty" json:"signer,omitempty"`
	Friendly              []ContractFriendly           `bson:"friendly,omitempty" json:"friendly,omitempty"`
	Legal                 []ContractLegal              `bson:"legal,omitempty" json:"legal,omitempty"`
	Rule                  []ContractRule               `bson:"rule,omitempty" json:"rule,omitempty"`
}
type ContractContentDefinition struct {
	ID                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept                        `bson:"type" json:"type"`
	SubType           *CodeableConcept                       `bson:"subType,omitempty" json:"subType,omitempty"`
	Publisher         *Reference                             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	PublicationDate   *string                                `bson:"publicationDate,omitempty" json:"publicationDate,omitempty"`
	PublicationStatus ContractResourcePublicationStatusCodes `bson:"publicationStatus" json:"publicationStatus"`
	Copyright         *string                                `bson:"copyright,omitempty" json:"copyright,omitempty"`
}
type ContractTerm struct {
	ID                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Issued            *string                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies           *Period                     `bson:"applies,omitempty" json:"applies,omitempty"`
	Type              *CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	SubType           *CodeableConcept            `bson:"subType,omitempty" json:"subType,omitempty"`
	Text              *string                     `bson:"text,omitempty" json:"text,omitempty"`
	SecurityLabel     []ContractTermSecurityLabel `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Offer             ContractTermOffer           `bson:"offer" json:"offer"`
	Asset             []ContractTermAsset         `bson:"asset,omitempty" json:"asset,omitempty"`
	Action            []ContractTermAction        `bson:"action,omitempty" json:"action,omitempty"`
	Group             []ContractTerm              `bson:"group,omitempty" json:"group,omitempty"`
}
type ContractTermSecurityLabel struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Number            []int       `bson:"number,omitempty" json:"number,omitempty"`
	Classification    Coding      `bson:"classification" json:"classification"`
	Category          []Coding    `bson:"category,omitempty" json:"category,omitempty"`
	Control           []Coding    `bson:"control,omitempty" json:"control,omitempty"`
}
type ContractTermOffer struct {
	ID                  *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Party               []ContractTermOfferParty  `bson:"party,omitempty" json:"party,omitempty"`
	Topic               *Reference                `bson:"topic,omitempty" json:"topic,omitempty"`
	Type                *CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	Decision            *CodeableConcept          `bson:"decision,omitempty" json:"decision,omitempty"`
	DecisionMode        []CodeableConcept         `bson:"decisionMode,omitempty" json:"decisionMode,omitempty"`
	Answer              []ContractTermOfferAnswer `bson:"answer,omitempty" json:"answer,omitempty"`
	Text                *string                   `bson:"text,omitempty" json:"text,omitempty"`
	LinkId              []string                  `bson:"linkId,omitempty" json:"linkId,omitempty"`
	SecurityLabelNumber []int                     `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
}
type ContractTermOfferParty struct {
	ID                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         []Reference     `bson:"reference" json:"reference"`
	Role              CodeableConcept `bson:"role" json:"role"`
}
type ContractTermOfferAnswer struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type ContractTermAsset struct {
	ID                  *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Scope               *CodeableConcept              `bson:"scope,omitempty" json:"scope,omitempty"`
	Type                []CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	TypeReference       []Reference                   `bson:"typeReference,omitempty" json:"typeReference,omitempty"`
	Subtype             []CodeableConcept             `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Relationship        *Coding                       `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Context             []ContractTermAssetContext    `bson:"context,omitempty" json:"context,omitempty"`
	Condition           *string                       `bson:"condition,omitempty" json:"condition,omitempty"`
	PeriodType          []CodeableConcept             `bson:"periodType,omitempty" json:"periodType,omitempty"`
	Period              []Period                      `bson:"period,omitempty" json:"period,omitempty"`
	UsePeriod           []Period                      `bson:"usePeriod,omitempty" json:"usePeriod,omitempty"`
	Text                *string                       `bson:"text,omitempty" json:"text,omitempty"`
	LinkId              []string                      `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Answer              []ContractTermOfferAnswer     `bson:"answer,omitempty" json:"answer,omitempty"`
	SecurityLabelNumber []int                         `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
	ValuedItem          []ContractTermAssetValuedItem `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
}
type ContractTermAssetContext struct {
	ID                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         *Reference        `bson:"reference,omitempty" json:"reference,omitempty"`
	Code              []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Text              *string           `bson:"text,omitempty" json:"text,omitempty"`
}
type ContractTermAssetValuedItem struct {
	ID                  *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime       *string     `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity            *Quantity   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice           *Money      `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor              *string     `bson:"factor,omitempty" json:"factor,omitempty"`
	Points              *string     `bson:"points,omitempty" json:"points,omitempty"`
	Net                 *Money      `bson:"net,omitempty" json:"net,omitempty"`
	Payment             *string     `bson:"payment,omitempty" json:"payment,omitempty"`
	PaymentDate         *string     `bson:"paymentDate,omitempty" json:"paymentDate,omitempty"`
	Responsible         *Reference  `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Recipient           *Reference  `bson:"recipient,omitempty" json:"recipient,omitempty"`
	LinkId              []string    `bson:"linkId,omitempty" json:"linkId,omitempty"`
	SecurityLabelNumber []int       `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
}
type ContractTermAction struct {
	ID                  *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DoNotPerform        *bool                       `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	Type                CodeableConcept             `bson:"type" json:"type"`
	Subject             []ContractTermActionSubject `bson:"subject,omitempty" json:"subject,omitempty"`
	Intent              CodeableConcept             `bson:"intent" json:"intent"`
	LinkId              []string                    `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Status              CodeableConcept             `bson:"status" json:"status"`
	Context             *Reference                  `bson:"context,omitempty" json:"context,omitempty"`
	ContextLinkId       []string                    `bson:"contextLinkId,omitempty" json:"contextLinkId,omitempty"`
	Requester           []Reference                 `bson:"requester,omitempty" json:"requester,omitempty"`
	RequesterLinkId     []string                    `bson:"requesterLinkId,omitempty" json:"requesterLinkId,omitempty"`
	PerformerType       []CodeableConcept           `bson:"performerType,omitempty" json:"performerType,omitempty"`
	PerformerRole       *CodeableConcept            `bson:"performerRole,omitempty" json:"performerRole,omitempty"`
	Performer           *Reference                  `bson:"performer,omitempty" json:"performer,omitempty"`
	PerformerLinkId     []string                    `bson:"performerLinkId,omitempty" json:"performerLinkId,omitempty"`
	ReasonCode          []CodeableConcept           `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference     []Reference                 `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Reason              []string                    `bson:"reason,omitempty" json:"reason,omitempty"`
	ReasonLinkId        []string                    `bson:"reasonLinkId,omitempty" json:"reasonLinkId,omitempty"`
	Note                []Annotation                `bson:"note,omitempty" json:"note,omitempty"`
	SecurityLabelNumber []int                       `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
}
type ContractTermActionSubject struct {
	ID                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         []Reference      `bson:"reference" json:"reference"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}
type ContractSigner struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              Coding      `bson:"type" json:"type"`
	Party             Reference   `bson:"party" json:"party"`
	Signature         []Signature `bson:"signature" json:"signature"`
}
type ContractFriendly struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type ContractLegal struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type ContractRule struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type OtherContract Contract

// MarshalJSON marshals the given Contract as JSON into a byte slice
func (r Contract) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherContract
		ResourceType string `json:"resourceType"`
	}{
		OtherContract: OtherContract(r),
		ResourceType:  "Contract",
	})
}

// UnmarshalContract unmarshals a Contract.
func UnmarshalContract(b []byte) (Contract, error) {
	var contract Contract
	if err := json.Unmarshal(b, &contract); err != nil {
		return contract, err
	}
	return contract, nil
}
