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
// 2021-03-08 09:41:02.980007 +0000 UTC

package models

// Extension is documented here http://hl7.org/fhir/StructureDefinition/Extension
type Extension struct {
	ID        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	URL       string      `bson:"url" json:"url"`

	// An address expressed using postal conventions (as opposed to GPS or other location definition formats).  This data type may be used to convey addresses for use in delivering mail as well as for visiting locations which might not be valid for mail delivery.  There are a variety of postal address formats defined around the world.
	ValueAddress *Address `json:"valueAddress,omitempty"`

	// A duration of time during which an organism (or a process) has existed.
	// ValueAge *Age `json:"valueAge,omitempty"`

	// A  text note which also  contains information about who made the statement and when.
	ValueAnnotation *Annotation `json:"valueAnnotation,omitempty"`

	// For referring to data content defined in other formats.
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueBase64Binary *string `json:"valueBase64Binary,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueCanonical *string `json:"valueCanonical,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueCode *string `json:"valueCode,omitempty"`

	// A concept that may be defined by a formal reference to a terminology or ontology or may be provided by text.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// A reference to a code defined by a terminology system.
	ValueCoding *Coding `json:"valueCoding,omitempty"`

	// Specifies contact information for a person or organization.
	ValueContactDetail *ContactDetail `json:"valueContactDetail,omitempty"`

	// Details for all kinds of technology mediated contact points for a person or organization, including telephone, email, etc.
	ValueContactPoint *ContactPoint `json:"valueContactPoint,omitempty"`

	// A contributor to the content of a knowledge asset, including authors, editors, reviewers, and endorsers.
	// ValueContributor *Contributor `json:"valueContributor,omitempty"`

	// A measured amount (or an amount that can potentially be measured). Note that measured amounts include amounts that are not precisely quantified, including amounts involving arbitrary units and floating currencies.
	// ValueCount *Count `json:"valueCount,omitempty"`

	// Describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueDate *DateTime `json:"valueDate,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueDateTime *DateTime `json:"valueDateTime,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`

	// A length - a value with a unit that is a physical distance.
	// ValueDistance *Distance `json:"valueDistance,omitempty"`

	// Indicates how the medication is/was taken or should be taken by the patient.
	ValueDosage *Dosage `json:"valueDosage,omitempty"`

	// A length of time.
	ValueDuration *Duration `json:"valueDuration,omitempty"`

	// A expression that is evaluated in a specified context and returns a value. The context of use of the expression must specify the context in which the expression is evaluated, and how the result of the expression is used.
	ValueExpression *Expression `json:"valueExpression,omitempty"`

	// A human's name with the ability to identify parts and usage.
	ValueHumanName *HumanName `json:"valueHumanName,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueId *string `json:"valueId,omitempty"`

	// An identifier - identifies some entity uniquely and unambiguously. Typically this is used for business identifiers.
	ValueIdentifier *Identifier `json:"valueIdentifier,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueInstant *string `json:"valueInstant,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueInteger *int `json:"valueInteger,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueMarkdown *string `json:"valueMarkdown,omitempty"`

	// The metadata about a resource. This is content in the resource that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	ValueMeta *Meta `json:"valueMeta,omitempty"`

	// An amount of economic utility in some recognized currency.
	ValueMoney *Money `json:"valueMoney,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueOid *string `json:"valueOid,omitempty"`

	// The parameters to the module. This collection specifies both the input and output parameters. Input parameters are provided by the caller as part of the $evaluate operation. Output parameters are included in the GuidanceResponse.
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`

	// A time period defined by a start and end date and optionally time.
	ValuePeriod *Period `json:"valuePeriod,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValuePositiveInt *int `json:"valuePositiveInt,omitempty"`

	// A measured amount (or an amount that can potentially be measured). Note that measured amounts include amounts that are not precisely quantified, including amounts involving arbitrary units and floating currencies.
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`

	// A set of ordered Quantities defined by a low and high limit.
	ValueRange *Range `json:"valueRange,omitempty"`

	// A relationship of two Quantity values - expressed as a numerator and a denominator.
	ValueRatio *Ratio `json:"valueRatio,omitempty"`

	// A reference from one resource to another.
	ValueReference *Reference `json:"valueReference,omitempty"`

	// Related artifacts such as additional documentation, justification, or bibliographic references.
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact,omitempty"`

	// A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.
	// ValueSampledData *SampledData `json:"valueSampledData,omitempty"`

	// A signature along with supporting context. The signature may be a digital signature that is cryptographic in nature, or some other signature acceptable to the domain. This other signature may be as simple as a graphical image representing a hand-written signature, or a signature ceremony Different signature approaches have different utilities.
	ValueSignature *Signature `json:"valueSignature,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueString *string `json:"valueString,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueTime *string `json:"valueTime,omitempty"`

	// Specifies an event that may occur multiple times. Timing schedules are used to record when things are planned, expected or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
	ValueTiming *Timing `json:"valueTiming,omitempty"`

	// A description of a triggering event. Triggering events can be named events, data events, or periodic, as determined by the type element.
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueUnsignedInt *int `json:"valueUnsignedInt,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueUri *string `json:"valueUri,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueUrl *string `json:"valueUrl,omitempty"`

	// Specifies clinical/business/etc. metadata that can be used to retrieve, index and/or categorize an artifact. This metadata can either be specific to the applicable population (e.g., age category, DRG) or the specific context of care (e.g., venue, care setting, provider of care).
	ValueUsageContext *UsageContext `json:"valueUsageContext,omitempty"`

	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	ValueUuid *string `json:"valueUuid,omitempty"`
}
