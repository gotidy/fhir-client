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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// CapabilityStatement is documented here http://hl7.org/fhir/StructureDefinition/CapabilityStatement
type CapabilityStatement struct {
	Id                  *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                            `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                         `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                 *string                            `bson:"url,omitempty" json:"url,omitempty"`
	Version             *string                            `bson:"version,omitempty" json:"version,omitempty"`
	Name                *string                            `bson:"name,omitempty" json:"name,omitempty"`
	Title               *string                            `bson:"title,omitempty" json:"title,omitempty"`
	Status              PublicationStatus                  `bson:"status" json:"status"`
	Experimental        *bool                              `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                string                             `bson:"date" json:"date"`
	Publisher           *string                            `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact             []ContactDetail                    `bson:"contact,omitempty" json:"contact,omitempty"`
	Description         *string                            `bson:"description,omitempty" json:"description,omitempty"`
	UseContext          []UsageContext                     `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept                  `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose             *string                            `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright           *string                            `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Kind                CapabilityStatementKind            `bson:"kind" json:"kind"`
	Instantiates        []string                           `bson:"instantiates,omitempty" json:"instantiates,omitempty"`
	Imports             []string                           `bson:"imports,omitempty" json:"imports,omitempty"`
	Software            *CapabilityStatementSoftware       `bson:"software,omitempty" json:"software,omitempty"`
	Implementation      *CapabilityStatementImplementation `bson:"implementation,omitempty" json:"implementation,omitempty"`
	FhirVersion         FHIRVersion                        `bson:"fhirVersion" json:"fhirVersion"`
	Format              []string                           `bson:"format" json:"format"`
	PatchFormat         []string                           `bson:"patchFormat,omitempty" json:"patchFormat,omitempty"`
	ImplementationGuide []string                           `bson:"implementationGuide,omitempty" json:"implementationGuide,omitempty"`
	Rest                []CapabilityStatementRest          `bson:"rest,omitempty" json:"rest,omitempty"`
	Messaging           []CapabilityStatementMessaging     `bson:"messaging,omitempty" json:"messaging,omitempty"`
	Document            []CapabilityStatementDocument      `bson:"document,omitempty" json:"document,omitempty"`
}
type CapabilityStatementSoftware struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Version           *string     `bson:"version,omitempty" json:"version,omitempty"`
	ReleaseDate       *string     `bson:"releaseDate,omitempty" json:"releaseDate,omitempty"`
}
type CapabilityStatementImplementation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       string      `bson:"description" json:"description"`
	Url               *string     `bson:"url,omitempty" json:"url,omitempty"`
	Custodian         *Reference  `bson:"custodian,omitempty" json:"custodian,omitempty"`
}
type CapabilityStatementRest struct {
	Id                *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              RestfulCapabilityMode                        `bson:"mode" json:"mode"`
	Documentation     *string                                      `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Security          *CapabilityStatementRestSecurity             `bson:"security,omitempty" json:"security,omitempty"`
	Resource          []CapabilityStatementRestResource            `bson:"resource,omitempty" json:"resource,omitempty"`
	Interaction       []CapabilityStatementRestInteraction         `bson:"interaction,omitempty" json:"interaction,omitempty"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	Operation         []CapabilityStatementRestResourceOperation   `bson:"operation,omitempty" json:"operation,omitempty"`
	Compartment       []string                                     `bson:"compartment,omitempty" json:"compartment,omitempty"`
}
type CapabilityStatementRestSecurity struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Cors              *bool             `bson:"cors,omitempty" json:"cors,omitempty"`
	Service           []CodeableConcept `bson:"service,omitempty" json:"service,omitempty"`
	Description       *string           `bson:"description,omitempty" json:"description,omitempty"`
}
type CapabilityStatementRestResource struct {
	Id                *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ResourceType                                 `bson:"type" json:"type"`
	Profile           *string                                      `bson:"profile,omitempty" json:"profile,omitempty"`
	SupportedProfile  []string                                     `bson:"supportedProfile,omitempty" json:"supportedProfile,omitempty"`
	Documentation     *string                                      `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Interaction       []CapabilityStatementRestResourceInteraction `bson:"interaction,omitempty" json:"interaction,omitempty"`
	Versioning        *ResourceVersionPolicy                       `bson:"versioning,omitempty" json:"versioning,omitempty"`
	ReadHistory       *bool                                        `bson:"readHistory,omitempty" json:"readHistory,omitempty"`
	UpdateCreate      *bool                                        `bson:"updateCreate,omitempty" json:"updateCreate,omitempty"`
	ConditionalCreate *bool                                        `bson:"conditionalCreate,omitempty" json:"conditionalCreate,omitempty"`
	ConditionalRead   *ConditionalReadStatus                       `bson:"conditionalRead,omitempty" json:"conditionalRead,omitempty"`
	ConditionalUpdate *bool                                        `bson:"conditionalUpdate,omitempty" json:"conditionalUpdate,omitempty"`
	ConditionalDelete *ConditionalDeleteStatus                     `bson:"conditionalDelete,omitempty" json:"conditionalDelete,omitempty"`
	ReferencePolicy   []ReferenceHandlingPolicy                    `bson:"referencePolicy,omitempty" json:"referencePolicy,omitempty"`
	SearchInclude     []string                                     `bson:"searchInclude,omitempty" json:"searchInclude,omitempty"`
	SearchRevInclude  []string                                     `bson:"searchRevInclude,omitempty" json:"searchRevInclude,omitempty"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	Operation         []CapabilityStatementRestResourceOperation   `bson:"operation,omitempty" json:"operation,omitempty"`
}
type CapabilityStatementRestResourceInteraction struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              TypeRestfulInteraction `bson:"code" json:"code"`
	Documentation     *string                `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type CapabilityStatementRestResourceSearchParam struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string          `bson:"name" json:"name"`
	Definition        *string         `bson:"definition,omitempty" json:"definition,omitempty"`
	Type              SearchParamType `bson:"type" json:"type"`
	Documentation     *string         `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type CapabilityStatementRestResourceOperation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Definition        string      `bson:"definition" json:"definition"`
	Documentation     *string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type CapabilityStatementRestInteraction struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              SystemRestfulInteraction `bson:"code" json:"code"`
	Documentation     *string                  `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type CapabilityStatementMessaging struct {
	Id                *string                                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Endpoint          []CapabilityStatementMessagingEndpoint         `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	ReliableCache     *int                                           `bson:"reliableCache,omitempty" json:"reliableCache,omitempty"`
	Documentation     *string                                        `bson:"documentation,omitempty" json:"documentation,omitempty"`
	SupportedMessage  []CapabilityStatementMessagingSupportedMessage `bson:"supportedMessage,omitempty" json:"supportedMessage,omitempty"`
}
type CapabilityStatementMessagingEndpoint struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Protocol          Coding      `bson:"protocol" json:"protocol"`
	Address           string      `bson:"address" json:"address"`
}
type CapabilityStatementMessagingSupportedMessage struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              EventCapabilityMode `bson:"mode" json:"mode"`
	Definition        string              `bson:"definition" json:"definition"`
}
type CapabilityStatementDocument struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              DocumentMode `bson:"mode" json:"mode"`
	Documentation     *string      `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Profile           string       `bson:"profile" json:"profile"`
}
type OtherCapabilityStatement CapabilityStatement

// MarshalJSON marshals the given CapabilityStatement as JSON into a byte slice.
func (r CapabilityStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCapabilityStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherCapabilityStatement: OtherCapabilityStatement(r),
		ResourceType:             "CapabilityStatement",
	})
}

// UnmarshalCapabilityStatement unmarshals a CapabilityStatement.
func UnmarshalCapabilityStatement(b []byte) (CapabilityStatement, error) {
	var capabilityStatement CapabilityStatement
	if err := json.Unmarshal(b, &capabilityStatement); err != nil {
		return capabilityStatement, err
	}
	return capabilityStatement, nil
}
