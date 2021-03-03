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
// 2021-03-03 10:55:43.945673 +0000 UTC
// ImplementationGuide is documented here http://hl7.org/fhir/StructureDefinition/ImplementationGuide
type ImplementationGuide struct {
	ID                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                         `bson:"url" json:"url"`
	Version           *string                        `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                         `bson:"name" json:"name"`
	Title             *string                        `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus              `bson:"status" json:"status"`
	Experimental      *bool                          `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                        `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                        `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                        `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                 `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept              `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright         *string                        `bson:"copyright,omitempty" json:"copyright,omitempty"`
	PackageId         string                         `bson:"packageId" json:"packageId"`
	License           *SPDXLicense                   `bson:"license,omitempty" json:"license,omitempty"`
	FhirVersion       []FHIRVersion                  `bson:"fhirVersion" json:"fhirVersion"`
	DependsOn         []ImplementationGuideDependsOn `bson:"dependsOn,omitempty" json:"dependsOn,omitempty"`
	Global            []ImplementationGuideGlobal    `bson:"global,omitempty" json:"global,omitempty"`
	Definition        *ImplementationGuideDefinition `bson:"definition,omitempty" json:"definition,omitempty"`
	Manifest          *ImplementationGuideManifest   `bson:"manifest,omitempty" json:"manifest,omitempty"`
}
type ImplementationGuideDependsOn struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uri               string      `bson:"uri" json:"uri"`
	PackageId         *string     `bson:"packageId,omitempty" json:"packageId,omitempty"`
	Version           *string     `bson:"version,omitempty" json:"version,omitempty"`
}
type ImplementationGuideGlobal struct {
	ID                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ResourceType `bson:"type" json:"type"`
	Profile           string       `bson:"profile" json:"profile"`
}
type ImplementationGuideDefinition struct {
	ID                *string                                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Grouping          []ImplementationGuideDefinitionGrouping  `bson:"grouping,omitempty" json:"grouping,omitempty"`
	Resource          []ImplementationGuideDefinitionResource  `bson:"resource" json:"resource"`
	Page              *ImplementationGuideDefinitionPage       `bson:"page,omitempty" json:"page,omitempty"`
	Parameter         []ImplementationGuideDefinitionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Template          []ImplementationGuideDefinitionTemplate  `bson:"template,omitempty" json:"template,omitempty"`
}
type ImplementationGuideDefinitionGrouping struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
}
type ImplementationGuideDefinitionResource struct {
	ID                *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         Reference     `bson:"reference" json:"reference"`
	FhirVersion       []FHIRVersion `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Name              *string       `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string       `bson:"description,omitempty" json:"description,omitempty"`
	GroupingId        *string       `bson:"groupingId,omitempty" json:"groupingId,omitempty"`
}
type ImplementationGuideDefinitionPage struct {
	ID                *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Title             string                              `bson:"title" json:"title"`
	Generation        GuidePageGeneration                 `bson:"generation" json:"generation"`
	Page              []ImplementationGuideDefinitionPage `bson:"page,omitempty" json:"page,omitempty"`
}
type ImplementationGuideDefinitionParameter struct {
	ID                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              GuideParameterCode `bson:"code" json:"code"`
	Value             string             `bson:"value" json:"value"`
}
type ImplementationGuideDefinitionTemplate struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	Source            string      `bson:"source" json:"source"`
	Scope             *string     `bson:"scope,omitempty" json:"scope,omitempty"`
}
type ImplementationGuideManifest struct {
	ID                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Rendering         *string                               `bson:"rendering,omitempty" json:"rendering,omitempty"`
	Resource          []ImplementationGuideManifestResource `bson:"resource" json:"resource"`
	Page              []ImplementationGuideManifestPage     `bson:"page,omitempty" json:"page,omitempty"`
	Image             []string                              `bson:"image,omitempty" json:"image,omitempty"`
	Other             []string                              `bson:"other,omitempty" json:"other,omitempty"`
}
type ImplementationGuideManifestResource struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         Reference   `bson:"reference" json:"reference"`
	RelativePath      *string     `bson:"relativePath,omitempty" json:"relativePath,omitempty"`
}
type ImplementationGuideManifestPage struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Title             *string     `bson:"title,omitempty" json:"title,omitempty"`
	Anchor            []string    `bson:"anchor,omitempty" json:"anchor,omitempty"`
}
type OtherImplementationGuide ImplementationGuide

// MarshalJSON marshals the given ImplementationGuide as JSON into a byte slice
func (r ImplementationGuide) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImplementationGuide
		ResourceType string `json:"resourceType"`
	}{
		OtherImplementationGuide: OtherImplementationGuide(r),
		ResourceType:             "ImplementationGuide",
	})
}

// UnmarshalImplementationGuide unmarshals a ImplementationGuide.
func UnmarshalImplementationGuide(b []byte) (ImplementationGuide, error) {
	var implementationGuide ImplementationGuide
	if err := json.Unmarshal(b, &implementationGuide); err != nil {
		return implementationGuide, err
	}
	return implementationGuide, nil
}
