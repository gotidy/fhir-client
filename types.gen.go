package fhir

import (
	"reflect"

	"github.com/gotidy/fhir-client/models"
)

const (
	AccountResource                           ResourceType = "Account"
	ActivityDefinitionResource                ResourceType = "ActivityDefinition"
	AdverseEventResource                      ResourceType = "AdverseEvent"
	AllergyIntoleranceResource                ResourceType = "AllergyIntolerance"
	AppointmentResource                       ResourceType = "Appointment"
	AppointmentResponseResource               ResourceType = "AppointmentResponse"
	AuditEventResource                        ResourceType = "AuditEvent"
	BasicResource                             ResourceType = "Basic"
	BinaryResource                            ResourceType = "Binary"
	BiologicallyDerivedProductResource        ResourceType = "BiologicallyDerivedProduct"
	BodyStructureResource                     ResourceType = "BodyStructure"
	BundleResource                            ResourceType = "Bundle"
	CapabilityStatementResource               ResourceType = "CapabilityStatement"
	CarePlanResource                          ResourceType = "CarePlan"
	CareTeamResource                          ResourceType = "CareTeam"
	CatalogEntryResource                      ResourceType = "CatalogEntry"
	ChargeItemResource                        ResourceType = "ChargeItem"
	ChargeItemDefinitionResource              ResourceType = "ChargeItemDefinition"
	ClaimResource                             ResourceType = "Claim"
	ClaimResponseResource                     ResourceType = "ClaimResponse"
	ClinicalImpressionResource                ResourceType = "ClinicalImpression"
	CodeSystemResource                        ResourceType = "CodeSystem"
	CommunicationResource                     ResourceType = "Communication"
	CommunicationRequestResource              ResourceType = "CommunicationRequest"
	CompartmentDefinitionResource             ResourceType = "CompartmentDefinition"
	CompositionResource                       ResourceType = "Composition"
	ConceptMapResource                        ResourceType = "ConceptMap"
	ConditionResource                         ResourceType = "Condition"
	ConsentResource                           ResourceType = "Consent"
	ContractResource                          ResourceType = "Contract"
	CoverageResource                          ResourceType = "Coverage"
	CoverageEligibilityRequestResource        ResourceType = "CoverageEligibilityRequest"
	CoverageEligibilityResponseResource       ResourceType = "CoverageEligibilityResponse"
	DetectedIssueResource                     ResourceType = "DetectedIssue"
	DeviceResource                            ResourceType = "Device"
	DeviceDefinitionResource                  ResourceType = "DeviceDefinition"
	DeviceMetricResource                      ResourceType = "DeviceMetric"
	DeviceRequestResource                     ResourceType = "DeviceRequest"
	DeviceUseStatementResource                ResourceType = "DeviceUseStatement"
	DiagnosticReportResource                  ResourceType = "DiagnosticReport"
	DocumentManifestResource                  ResourceType = "DocumentManifest"
	DocumentReferenceResource                 ResourceType = "DocumentReference"
	DomainResourceResource                    ResourceType = "DomainResource"
	EffectEvidenceSynthesisResource           ResourceType = "EffectEvidenceSynthesis"
	EncounterResource                         ResourceType = "Encounter"
	EndpointResource                          ResourceType = "Endpoint"
	EnrollmentRequestResource                 ResourceType = "EnrollmentRequest"
	EnrollmentResponseResource                ResourceType = "EnrollmentResponse"
	EpisodeOfCareResource                     ResourceType = "EpisodeOfCare"
	EventDefinitionResource                   ResourceType = "EventDefinition"
	EvidenceResource                          ResourceType = "Evidence"
	EvidenceVariableResource                  ResourceType = "EvidenceVariable"
	ExampleScenarioResource                   ResourceType = "ExampleScenario"
	ExplanationOfBenefitResource              ResourceType = "ExplanationOfBenefit"
	FamilyMemberHistoryResource               ResourceType = "FamilyMemberHistory"
	FlagResource                              ResourceType = "Flag"
	GoalResource                              ResourceType = "Goal"
	GraphDefinitionResource                   ResourceType = "GraphDefinition"
	GroupResource                             ResourceType = "Group"
	GuidanceResponseResource                  ResourceType = "GuidanceResponse"
	HealthcareServiceResource                 ResourceType = "HealthcareService"
	ImagingStudyResource                      ResourceType = "ImagingStudy"
	ImmunizationResource                      ResourceType = "Immunization"
	ImmunizationEvaluationResource            ResourceType = "ImmunizationEvaluation"
	ImmunizationRecommendationResource        ResourceType = "ImmunizationRecommendation"
	ImplementationGuideResource               ResourceType = "ImplementationGuide"
	InsurancePlanResource                     ResourceType = "InsurancePlan"
	InvoiceResource                           ResourceType = "Invoice"
	LibraryResource                           ResourceType = "Library"
	LinkageResource                           ResourceType = "Linkage"
	ListResource                              ResourceType = "List"
	LocationResource                          ResourceType = "Location"
	MeasureResource                           ResourceType = "Measure"
	MeasureReportResource                     ResourceType = "MeasureReport"
	MediaResource                             ResourceType = "Media"
	MedicationResource                        ResourceType = "Medication"
	MedicationAdministrationResource          ResourceType = "MedicationAdministration"
	MedicationDispenseResource                ResourceType = "MedicationDispense"
	MedicationKnowledgeResource               ResourceType = "MedicationKnowledge"
	MedicationRequestResource                 ResourceType = "MedicationRequest"
	MedicationStatementResource               ResourceType = "MedicationStatement"
	MedicinalProductResource                  ResourceType = "MedicinalProduct"
	MedicinalProductAuthorizationResource     ResourceType = "MedicinalProductAuthorization"
	MedicinalProductContraindicationResource  ResourceType = "MedicinalProductContraindication"
	MedicinalProductIndicationResource        ResourceType = "MedicinalProductIndication"
	MedicinalProductIngredientResource        ResourceType = "MedicinalProductIngredient"
	MedicinalProductInteractionResource       ResourceType = "MedicinalProductInteraction"
	MedicinalProductManufacturedResource      ResourceType = "MedicinalProductManufactured"
	MedicinalProductPackagedResource          ResourceType = "MedicinalProductPackaged"
	MedicinalProductPharmaceuticalResource    ResourceType = "MedicinalProductPharmaceutical"
	MedicinalProductUndesirableEffectResource ResourceType = "MedicinalProductUndesirableEffect"
	MessageDefinitionResource                 ResourceType = "MessageDefinition"
	MessageHeaderResource                     ResourceType = "MessageHeader"
	MolecularSequenceResource                 ResourceType = "MolecularSequence"
	NamingSystemResource                      ResourceType = "NamingSystem"
	NutritionOrderResource                    ResourceType = "NutritionOrder"
	ObservationResource                       ResourceType = "Observation"
	ObservationDefinitionResource             ResourceType = "ObservationDefinition"
	OperationDefinitionResource               ResourceType = "OperationDefinition"
	OperationOutcomeResource                  ResourceType = "OperationOutcome"
	OrganizationResource                      ResourceType = "Organization"
	OrganizationAffiliationResource           ResourceType = "OrganizationAffiliation"
	ParametersResource                        ResourceType = "Parameters"
	PatientResource                           ResourceType = "Patient"
	PaymentNoticeResource                     ResourceType = "PaymentNotice"
	PaymentReconciliationResource             ResourceType = "PaymentReconciliation"
	PersonResource                            ResourceType = "Person"
	PlanDefinitionResource                    ResourceType = "PlanDefinition"
	PractitionerResource                      ResourceType = "Practitioner"
	PractitionerRoleResource                  ResourceType = "PractitionerRole"
	ProcedureResource                         ResourceType = "Procedure"
	ProvenanceResource                        ResourceType = "Provenance"
	QuestionnaireResource                     ResourceType = "Questionnaire"
	QuestionnaireResponseResource             ResourceType = "QuestionnaireResponse"
	RelatedPersonResource                     ResourceType = "RelatedPerson"
	RequestGroupResource                      ResourceType = "RequestGroup"
	ResearchDefinitionResource                ResourceType = "ResearchDefinition"
	ResearchElementDefinitionResource         ResourceType = "ResearchElementDefinition"
	ResearchStudyResource                     ResourceType = "ResearchStudy"
	ResearchSubjectResource                   ResourceType = "ResearchSubject"
	ResourceResource                          ResourceType = "Resource"
	RiskAssessmentResource                    ResourceType = "RiskAssessment"
	RiskEvidenceSynthesisResource             ResourceType = "RiskEvidenceSynthesis"
	ScheduleResource                          ResourceType = "Schedule"
	SearchParameterResource                   ResourceType = "SearchParameter"
	ServiceRequestResource                    ResourceType = "ServiceRequest"
	SlotResource                              ResourceType = "Slot"
	SpecimenResource                          ResourceType = "Specimen"
	SpecimenDefinitionResource                ResourceType = "SpecimenDefinition"
	StructureDefinitionResource               ResourceType = "StructureDefinition"
	StructureMapResource                      ResourceType = "StructureMap"
	SubscriptionResource                      ResourceType = "Subscription"
	SubstanceResource                         ResourceType = "Substance"
	SubstanceNucleicAcidResource              ResourceType = "SubstanceNucleicAcid"
	SubstancePolymerResource                  ResourceType = "SubstancePolymer"
	SubstanceProteinResource                  ResourceType = "SubstanceProtein"
	SubstanceReferenceInformationResource     ResourceType = "SubstanceReferenceInformation"
	SubstanceSourceMaterialResource           ResourceType = "SubstanceSourceMaterial"
	SubstanceSpecificationResource            ResourceType = "SubstanceSpecification"
	SupplyDeliveryResource                    ResourceType = "SupplyDelivery"
	SupplyRequestResource                     ResourceType = "SupplyRequest"
	TaskResource                              ResourceType = "Task"
	TerminologyCapabilitiesResource           ResourceType = "TerminologyCapabilities"
	TestReportResource                        ResourceType = "TestReport"
	TestScriptResource                        ResourceType = "TestScript"
	ValueSetResource                          ResourceType = "ValueSet"
	VerificationResultResource                ResourceType = "VerificationResult"
	VisionPrescriptionResource                ResourceType = "VisionPrescription"
)

var resources = map[ResourceType]struct {
	new func() interface{}
	typ reflect.Type
}{
	AccountResource:                           {new: func() interface{} { return &models.Account{} }, typ: reflect.TypeOf(models.Account{})},
	ActivityDefinitionResource:                {new: func() interface{} { return &models.ActivityDefinition{} }, typ: reflect.TypeOf(models.ActivityDefinition{})},
	AdverseEventResource:                      {new: func() interface{} { return &models.AdverseEvent{} }, typ: reflect.TypeOf(models.AdverseEvent{})},
	AllergyIntoleranceResource:                {new: func() interface{} { return &models.AllergyIntolerance{} }, typ: reflect.TypeOf(models.AllergyIntolerance{})},
	AppointmentResource:                       {new: func() interface{} { return &models.Appointment{} }, typ: reflect.TypeOf(models.Appointment{})},
	AppointmentResponseResource:               {new: func() interface{} { return &models.AppointmentResponse{} }, typ: reflect.TypeOf(models.AppointmentResponse{})},
	AuditEventResource:                        {new: func() interface{} { return &models.AuditEvent{} }, typ: reflect.TypeOf(models.AuditEvent{})},
	BasicResource:                             {new: func() interface{} { return &models.Basic{} }, typ: reflect.TypeOf(models.Basic{})},
	BinaryResource:                            {new: func() interface{} { return &models.Binary{} }, typ: reflect.TypeOf(models.Binary{})},
	BiologicallyDerivedProductResource:        {new: func() interface{} { return &models.BiologicallyDerivedProduct{} }, typ: reflect.TypeOf(models.BiologicallyDerivedProduct{})},
	BodyStructureResource:                     {new: func() interface{} { return &models.BodyStructure{} }, typ: reflect.TypeOf(models.BodyStructure{})},
	BundleResource:                            {new: func() interface{} { return &models.Bundle{} }, typ: reflect.TypeOf(models.Bundle{})},
	CapabilityStatementResource:               {new: func() interface{} { return &models.CapabilityStatement{} }, typ: reflect.TypeOf(models.CapabilityStatement{})},
	CarePlanResource:                          {new: func() interface{} { return &models.CarePlan{} }, typ: reflect.TypeOf(models.CarePlan{})},
	CareTeamResource:                          {new: func() interface{} { return &models.CareTeam{} }, typ: reflect.TypeOf(models.CareTeam{})},
	CatalogEntryResource:                      {new: func() interface{} { return &models.CatalogEntry{} }, typ: reflect.TypeOf(models.CatalogEntry{})},
	ChargeItemResource:                        {new: func() interface{} { return &models.ChargeItem{} }, typ: reflect.TypeOf(models.ChargeItem{})},
	ChargeItemDefinitionResource:              {new: func() interface{} { return &models.ChargeItemDefinition{} }, typ: reflect.TypeOf(models.ChargeItemDefinition{})},
	ClaimResource:                             {new: func() interface{} { return &models.Claim{} }, typ: reflect.TypeOf(models.Claim{})},
	ClaimResponseResource:                     {new: func() interface{} { return &models.ClaimResponse{} }, typ: reflect.TypeOf(models.ClaimResponse{})},
	ClinicalImpressionResource:                {new: func() interface{} { return &models.ClinicalImpression{} }, typ: reflect.TypeOf(models.ClinicalImpression{})},
	CodeSystemResource:                        {new: func() interface{} { return &models.CodeSystem{} }, typ: reflect.TypeOf(models.CodeSystem{})},
	CommunicationResource:                     {new: func() interface{} { return &models.Communication{} }, typ: reflect.TypeOf(models.Communication{})},
	CommunicationRequestResource:              {new: func() interface{} { return &models.CommunicationRequest{} }, typ: reflect.TypeOf(models.CommunicationRequest{})},
	CompartmentDefinitionResource:             {new: func() interface{} { return &models.CompartmentDefinition{} }, typ: reflect.TypeOf(models.CompartmentDefinition{})},
	CompositionResource:                       {new: func() interface{} { return &models.Composition{} }, typ: reflect.TypeOf(models.Composition{})},
	ConceptMapResource:                        {new: func() interface{} { return &models.ConceptMap{} }, typ: reflect.TypeOf(models.ConceptMap{})},
	ConditionResource:                         {new: func() interface{} { return &models.Condition{} }, typ: reflect.TypeOf(models.Condition{})},
	ConsentResource:                           {new: func() interface{} { return &models.Consent{} }, typ: reflect.TypeOf(models.Consent{})},
	ContractResource:                          {new: func() interface{} { return &models.Contract{} }, typ: reflect.TypeOf(models.Contract{})},
	CoverageResource:                          {new: func() interface{} { return &models.Coverage{} }, typ: reflect.TypeOf(models.Coverage{})},
	CoverageEligibilityRequestResource:        {new: func() interface{} { return &models.CoverageEligibilityRequest{} }, typ: reflect.TypeOf(models.CoverageEligibilityRequest{})},
	CoverageEligibilityResponseResource:       {new: func() interface{} { return &models.CoverageEligibilityResponse{} }, typ: reflect.TypeOf(models.CoverageEligibilityResponse{})},
	DetectedIssueResource:                     {new: func() interface{} { return &models.DetectedIssue{} }, typ: reflect.TypeOf(models.DetectedIssue{})},
	DeviceResource:                            {new: func() interface{} { return &models.Device{} }, typ: reflect.TypeOf(models.Device{})},
	DeviceDefinitionResource:                  {new: func() interface{} { return &models.DeviceDefinition{} }, typ: reflect.TypeOf(models.DeviceDefinition{})},
	DeviceMetricResource:                      {new: func() interface{} { return &models.DeviceMetric{} }, typ: reflect.TypeOf(models.DeviceMetric{})},
	DeviceRequestResource:                     {new: func() interface{} { return &models.DeviceRequest{} }, typ: reflect.TypeOf(models.DeviceRequest{})},
	DeviceUseStatementResource:                {new: func() interface{} { return &models.DeviceUseStatement{} }, typ: reflect.TypeOf(models.DeviceUseStatement{})},
	DiagnosticReportResource:                  {new: func() interface{} { return &models.DiagnosticReport{} }, typ: reflect.TypeOf(models.DiagnosticReport{})},
	DocumentManifestResource:                  {new: func() interface{} { return &models.DocumentManifest{} }, typ: reflect.TypeOf(models.DocumentManifest{})},
	DocumentReferenceResource:                 {new: func() interface{} { return &models.DocumentReference{} }, typ: reflect.TypeOf(models.DocumentReference{})},
	DomainResourceResource:                    {new: func() interface{} { return &models.DomainResource{} }, typ: reflect.TypeOf(models.DomainResource{})},
	EffectEvidenceSynthesisResource:           {new: func() interface{} { return &models.EffectEvidenceSynthesis{} }, typ: reflect.TypeOf(models.EffectEvidenceSynthesis{})},
	EncounterResource:                         {new: func() interface{} { return &models.Encounter{} }, typ: reflect.TypeOf(models.Encounter{})},
	EndpointResource:                          {new: func() interface{} { return &models.Endpoint{} }, typ: reflect.TypeOf(models.Endpoint{})},
	EnrollmentRequestResource:                 {new: func() interface{} { return &models.EnrollmentRequest{} }, typ: reflect.TypeOf(models.EnrollmentRequest{})},
	EnrollmentResponseResource:                {new: func() interface{} { return &models.EnrollmentResponse{} }, typ: reflect.TypeOf(models.EnrollmentResponse{})},
	EpisodeOfCareResource:                     {new: func() interface{} { return &models.EpisodeOfCare{} }, typ: reflect.TypeOf(models.EpisodeOfCare{})},
	EventDefinitionResource:                   {new: func() interface{} { return &models.EventDefinition{} }, typ: reflect.TypeOf(models.EventDefinition{})},
	EvidenceResource:                          {new: func() interface{} { return &models.Evidence{} }, typ: reflect.TypeOf(models.Evidence{})},
	EvidenceVariableResource:                  {new: func() interface{} { return &models.EvidenceVariable{} }, typ: reflect.TypeOf(models.EvidenceVariable{})},
	ExampleScenarioResource:                   {new: func() interface{} { return &models.ExampleScenario{} }, typ: reflect.TypeOf(models.ExampleScenario{})},
	ExplanationOfBenefitResource:              {new: func() interface{} { return &models.ExplanationOfBenefit{} }, typ: reflect.TypeOf(models.ExplanationOfBenefit{})},
	FamilyMemberHistoryResource:               {new: func() interface{} { return &models.FamilyMemberHistory{} }, typ: reflect.TypeOf(models.FamilyMemberHistory{})},
	FlagResource:                              {new: func() interface{} { return &models.Flag{} }, typ: reflect.TypeOf(models.Flag{})},
	GoalResource:                              {new: func() interface{} { return &models.Goal{} }, typ: reflect.TypeOf(models.Goal{})},
	GraphDefinitionResource:                   {new: func() interface{} { return &models.GraphDefinition{} }, typ: reflect.TypeOf(models.GraphDefinition{})},
	GroupResource:                             {new: func() interface{} { return &models.Group{} }, typ: reflect.TypeOf(models.Group{})},
	GuidanceResponseResource:                  {new: func() interface{} { return &models.GuidanceResponse{} }, typ: reflect.TypeOf(models.GuidanceResponse{})},
	HealthcareServiceResource:                 {new: func() interface{} { return &models.HealthcareService{} }, typ: reflect.TypeOf(models.HealthcareService{})},
	ImagingStudyResource:                      {new: func() interface{} { return &models.ImagingStudy{} }, typ: reflect.TypeOf(models.ImagingStudy{})},
	ImmunizationResource:                      {new: func() interface{} { return &models.Immunization{} }, typ: reflect.TypeOf(models.Immunization{})},
	ImmunizationEvaluationResource:            {new: func() interface{} { return &models.ImmunizationEvaluation{} }, typ: reflect.TypeOf(models.ImmunizationEvaluation{})},
	ImmunizationRecommendationResource:        {new: func() interface{} { return &models.ImmunizationRecommendation{} }, typ: reflect.TypeOf(models.ImmunizationRecommendation{})},
	ImplementationGuideResource:               {new: func() interface{} { return &models.ImplementationGuide{} }, typ: reflect.TypeOf(models.ImplementationGuide{})},
	InsurancePlanResource:                     {new: func() interface{} { return &models.InsurancePlan{} }, typ: reflect.TypeOf(models.InsurancePlan{})},
	InvoiceResource:                           {new: func() interface{} { return &models.Invoice{} }, typ: reflect.TypeOf(models.Invoice{})},
	LibraryResource:                           {new: func() interface{} { return &models.Library{} }, typ: reflect.TypeOf(models.Library{})},
	LinkageResource:                           {new: func() interface{} { return &models.Linkage{} }, typ: reflect.TypeOf(models.Linkage{})},
	ListResource:                              {new: func() interface{} { return &models.List{} }, typ: reflect.TypeOf(models.List{})},
	LocationResource:                          {new: func() interface{} { return &models.Location{} }, typ: reflect.TypeOf(models.Location{})},
	MeasureResource:                           {new: func() interface{} { return &models.Measure{} }, typ: reflect.TypeOf(models.Measure{})},
	MeasureReportResource:                     {new: func() interface{} { return &models.MeasureReport{} }, typ: reflect.TypeOf(models.MeasureReport{})},
	MediaResource:                             {new: func() interface{} { return &models.Media{} }, typ: reflect.TypeOf(models.Media{})},
	MedicationResource:                        {new: func() interface{} { return &models.Medication{} }, typ: reflect.TypeOf(models.Medication{})},
	MedicationAdministrationResource:          {new: func() interface{} { return &models.MedicationAdministration{} }, typ: reflect.TypeOf(models.MedicationAdministration{})},
	MedicationDispenseResource:                {new: func() interface{} { return &models.MedicationDispense{} }, typ: reflect.TypeOf(models.MedicationDispense{})},
	MedicationKnowledgeResource:               {new: func() interface{} { return &models.MedicationKnowledge{} }, typ: reflect.TypeOf(models.MedicationKnowledge{})},
	MedicationRequestResource:                 {new: func() interface{} { return &models.MedicationRequest{} }, typ: reflect.TypeOf(models.MedicationRequest{})},
	MedicationStatementResource:               {new: func() interface{} { return &models.MedicationStatement{} }, typ: reflect.TypeOf(models.MedicationStatement{})},
	MedicinalProductResource:                  {new: func() interface{} { return &models.MedicinalProduct{} }, typ: reflect.TypeOf(models.MedicinalProduct{})},
	MedicinalProductAuthorizationResource:     {new: func() interface{} { return &models.MedicinalProductAuthorization{} }, typ: reflect.TypeOf(models.MedicinalProductAuthorization{})},
	MedicinalProductContraindicationResource:  {new: func() interface{} { return &models.MedicinalProductContraindication{} }, typ: reflect.TypeOf(models.MedicinalProductContraindication{})},
	MedicinalProductIndicationResource:        {new: func() interface{} { return &models.MedicinalProductIndication{} }, typ: reflect.TypeOf(models.MedicinalProductIndication{})},
	MedicinalProductIngredientResource:        {new: func() interface{} { return &models.MedicinalProductIngredient{} }, typ: reflect.TypeOf(models.MedicinalProductIngredient{})},
	MedicinalProductInteractionResource:       {new: func() interface{} { return &models.MedicinalProductInteraction{} }, typ: reflect.TypeOf(models.MedicinalProductInteraction{})},
	MedicinalProductManufacturedResource:      {new: func() interface{} { return &models.MedicinalProductManufactured{} }, typ: reflect.TypeOf(models.MedicinalProductManufactured{})},
	MedicinalProductPackagedResource:          {new: func() interface{} { return &models.MedicinalProductPackaged{} }, typ: reflect.TypeOf(models.MedicinalProductPackaged{})},
	MedicinalProductPharmaceuticalResource:    {new: func() interface{} { return &models.MedicinalProductPharmaceutical{} }, typ: reflect.TypeOf(models.MedicinalProductPharmaceutical{})},
	MedicinalProductUndesirableEffectResource: {new: func() interface{} { return &models.MedicinalProductUndesirableEffect{} }, typ: reflect.TypeOf(models.MedicinalProductUndesirableEffect{})},
	MessageDefinitionResource:                 {new: func() interface{} { return &models.MessageDefinition{} }, typ: reflect.TypeOf(models.MessageDefinition{})},
	MessageHeaderResource:                     {new: func() interface{} { return &models.MessageHeader{} }, typ: reflect.TypeOf(models.MessageHeader{})},
	MolecularSequenceResource:                 {new: func() interface{} { return &models.MolecularSequence{} }, typ: reflect.TypeOf(models.MolecularSequence{})},
	NamingSystemResource:                      {new: func() interface{} { return &models.NamingSystem{} }, typ: reflect.TypeOf(models.NamingSystem{})},
	NutritionOrderResource:                    {new: func() interface{} { return &models.NutritionOrder{} }, typ: reflect.TypeOf(models.NutritionOrder{})},
	ObservationResource:                       {new: func() interface{} { return &models.Observation{} }, typ: reflect.TypeOf(models.Observation{})},
	ObservationDefinitionResource:             {new: func() interface{} { return &models.ObservationDefinition{} }, typ: reflect.TypeOf(models.ObservationDefinition{})},
	OperationDefinitionResource:               {new: func() interface{} { return &models.OperationDefinition{} }, typ: reflect.TypeOf(models.OperationDefinition{})},
	OperationOutcomeResource:                  {new: func() interface{} { return &models.OperationOutcome{} }, typ: reflect.TypeOf(models.OperationOutcome{})},
	OrganizationResource:                      {new: func() interface{} { return &models.Organization{} }, typ: reflect.TypeOf(models.Organization{})},
	OrganizationAffiliationResource:           {new: func() interface{} { return &models.OrganizationAffiliation{} }, typ: reflect.TypeOf(models.OrganizationAffiliation{})},
	ParametersResource:                        {new: func() interface{} { return &models.Parameters{} }, typ: reflect.TypeOf(models.Parameters{})},
	PatientResource:                           {new: func() interface{} { return &models.Patient{} }, typ: reflect.TypeOf(models.Patient{})},
	PaymentNoticeResource:                     {new: func() interface{} { return &models.PaymentNotice{} }, typ: reflect.TypeOf(models.PaymentNotice{})},
	PaymentReconciliationResource:             {new: func() interface{} { return &models.PaymentReconciliation{} }, typ: reflect.TypeOf(models.PaymentReconciliation{})},
	PersonResource:                            {new: func() interface{} { return &models.Person{} }, typ: reflect.TypeOf(models.Person{})},
	PlanDefinitionResource:                    {new: func() interface{} { return &models.PlanDefinition{} }, typ: reflect.TypeOf(models.PlanDefinition{})},
	PractitionerResource:                      {new: func() interface{} { return &models.Practitioner{} }, typ: reflect.TypeOf(models.Practitioner{})},
	PractitionerRoleResource:                  {new: func() interface{} { return &models.PractitionerRole{} }, typ: reflect.TypeOf(models.PractitionerRole{})},
	ProcedureResource:                         {new: func() interface{} { return &models.Procedure{} }, typ: reflect.TypeOf(models.Procedure{})},
	ProvenanceResource:                        {new: func() interface{} { return &models.Provenance{} }, typ: reflect.TypeOf(models.Provenance{})},
	QuestionnaireResource:                     {new: func() interface{} { return &models.Questionnaire{} }, typ: reflect.TypeOf(models.Questionnaire{})},
	QuestionnaireResponseResource:             {new: func() interface{} { return &models.QuestionnaireResponse{} }, typ: reflect.TypeOf(models.QuestionnaireResponse{})},
	RelatedPersonResource:                     {new: func() interface{} { return &models.RelatedPerson{} }, typ: reflect.TypeOf(models.RelatedPerson{})},
	RequestGroupResource:                      {new: func() interface{} { return &models.RequestGroup{} }, typ: reflect.TypeOf(models.RequestGroup{})},
	ResearchDefinitionResource:                {new: func() interface{} { return &models.ResearchDefinition{} }, typ: reflect.TypeOf(models.ResearchDefinition{})},
	ResearchElementDefinitionResource:         {new: func() interface{} { return &models.ResearchElementDefinition{} }, typ: reflect.TypeOf(models.ResearchElementDefinition{})},
	ResearchStudyResource:                     {new: func() interface{} { return &models.ResearchStudy{} }, typ: reflect.TypeOf(models.ResearchStudy{})},
	ResearchSubjectResource:                   {new: func() interface{} { return &models.ResearchSubject{} }, typ: reflect.TypeOf(models.ResearchSubject{})},
	ResourceResource:                          {new: func() interface{} { return &models.Resource{} }, typ: reflect.TypeOf(models.Resource{})},
	RiskAssessmentResource:                    {new: func() interface{} { return &models.RiskAssessment{} }, typ: reflect.TypeOf(models.RiskAssessment{})},
	RiskEvidenceSynthesisResource:             {new: func() interface{} { return &models.RiskEvidenceSynthesis{} }, typ: reflect.TypeOf(models.RiskEvidenceSynthesis{})},
	ScheduleResource:                          {new: func() interface{} { return &models.Schedule{} }, typ: reflect.TypeOf(models.Schedule{})},
	SearchParameterResource:                   {new: func() interface{} { return &models.SearchParameter{} }, typ: reflect.TypeOf(models.SearchParameter{})},
	ServiceRequestResource:                    {new: func() interface{} { return &models.ServiceRequest{} }, typ: reflect.TypeOf(models.ServiceRequest{})},
	SlotResource:                              {new: func() interface{} { return &models.Slot{} }, typ: reflect.TypeOf(models.Slot{})},
	SpecimenResource:                          {new: func() interface{} { return &models.Specimen{} }, typ: reflect.TypeOf(models.Specimen{})},
	SpecimenDefinitionResource:                {new: func() interface{} { return &models.SpecimenDefinition{} }, typ: reflect.TypeOf(models.SpecimenDefinition{})},
	StructureDefinitionResource:               {new: func() interface{} { return &models.StructureDefinition{} }, typ: reflect.TypeOf(models.StructureDefinition{})},
	StructureMapResource:                      {new: func() interface{} { return &models.StructureMap{} }, typ: reflect.TypeOf(models.StructureMap{})},
	SubscriptionResource:                      {new: func() interface{} { return &models.Subscription{} }, typ: reflect.TypeOf(models.Subscription{})},
	SubstanceResource:                         {new: func() interface{} { return &models.Substance{} }, typ: reflect.TypeOf(models.Substance{})},
	SubstanceNucleicAcidResource:              {new: func() interface{} { return &models.SubstanceNucleicAcid{} }, typ: reflect.TypeOf(models.SubstanceNucleicAcid{})},
	SubstancePolymerResource:                  {new: func() interface{} { return &models.SubstancePolymer{} }, typ: reflect.TypeOf(models.SubstancePolymer{})},
	SubstanceProteinResource:                  {new: func() interface{} { return &models.SubstanceProtein{} }, typ: reflect.TypeOf(models.SubstanceProtein{})},
	SubstanceReferenceInformationResource:     {new: func() interface{} { return &models.SubstanceReferenceInformation{} }, typ: reflect.TypeOf(models.SubstanceReferenceInformation{})},
	SubstanceSourceMaterialResource:           {new: func() interface{} { return &models.SubstanceSourceMaterial{} }, typ: reflect.TypeOf(models.SubstanceSourceMaterial{})},
	SubstanceSpecificationResource:            {new: func() interface{} { return &models.SubstanceSpecification{} }, typ: reflect.TypeOf(models.SubstanceSpecification{})},
	SupplyDeliveryResource:                    {new: func() interface{} { return &models.SupplyDelivery{} }, typ: reflect.TypeOf(models.SupplyDelivery{})},
	SupplyRequestResource:                     {new: func() interface{} { return &models.SupplyRequest{} }, typ: reflect.TypeOf(models.SupplyRequest{})},
	TaskResource:                              {new: func() interface{} { return &models.Task{} }, typ: reflect.TypeOf(models.Task{})},
	TerminologyCapabilitiesResource:           {new: func() interface{} { return &models.TerminologyCapabilities{} }, typ: reflect.TypeOf(models.TerminologyCapabilities{})},
	TestReportResource:                        {new: func() interface{} { return &models.TestReport{} }, typ: reflect.TypeOf(models.TestReport{})},
	TestScriptResource:                        {new: func() interface{} { return &models.TestScript{} }, typ: reflect.TypeOf(models.TestScript{})},
	ValueSetResource:                          {new: func() interface{} { return &models.ValueSet{} }, typ: reflect.TypeOf(models.ValueSet{})},
	VerificationResultResource:                {new: func() interface{} { return &models.VerificationResult{} }, typ: reflect.TypeOf(models.VerificationResult{})},
	VisionPrescriptionResource:                {new: func() interface{} { return &models.VisionPrescription{} }, typ: reflect.TypeOf(models.VisionPrescription{})},
}

// NewResource creates the new resource.
func NewResource(resource ResourceType) interface{} {
	return resources[resource].new()
}

// TypeOf returns the type of resource.
func TypeOf(resource ResourceType) reflect.Type {
	return resources[resource].typ
}
